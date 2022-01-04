package delivery

import (
	"enigmacamp.com/completetesting/manager"
	"enigmacamp.com/completetesting/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type useCaseManagerMock struct {
	mock.Mock
}

func (uc *useCaseManagerMock) StudentUseCase() usecase.IStudentUseCase {
	args := uc.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(usecase.IStudentUseCase)
}

type DeliveryInitTestSuite struct {
	suite.Suite
	routerTest         *gin.Engine
	useCaseManagerTest manager.UseCaseManager
}

func (suite *DeliveryInitTestSuite) SetupTest() {
	suite.routerTest = gin.Default()
	suite.useCaseManagerTest = new(useCaseManagerMock)
}
func (suite *DeliveryInitTestSuite) TestDeliveryInit_CreateServer_Success() {
	suite.useCaseManagerTest.(*useCaseManagerMock).On("StudentUseCase").Return(usecase.NewStudentUseCase(nil))
	err := NewServer(suite.routerTest, suite.useCaseManagerTest)
	assert.Nil(suite.T(), err)
}
func (suite *DeliveryInitTestSuite) TestDeliveryInit_CreateServer_Failed() {
	suite.useCaseManagerTest.(*useCaseManagerMock).On("StudentUseCase").Return(nil)
	err := NewServer(suite.routerTest, suite.useCaseManagerTest)
	assert.NotNil(suite.T(), err)
}

func TestDeliveryInitTestSuite(t *testing.T) {
	suite.Run(t, new(DeliveryInitTestSuite))
}

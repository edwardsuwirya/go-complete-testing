package delivery

import (
	"enigmacamp.com/completetesting/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type mockUseCaseManager struct {
}

func (uc *mockUseCaseManager) StudentUseCase() usecase.IStudentUseCase {
	return nil
}

type DeliveryInitTestSuite struct {
	suite.Suite
}

func (suite *DeliveryInitTestSuite) TestNewServer() {
	mockRoutes := NewServer(new(mockUseCaseManager))
	assert.NotNil(suite.T(), mockRoutes.RouterEngine)
	assert.NotNil(suite.T(), mockRoutes.routers)
	assert.NotNil(suite.T(), mockRoutes.publicRoute)
}
func TestDeliveryInitTestSuite(t *testing.T) {
	suite.Run(t, new(DeliveryInitTestSuite))
}

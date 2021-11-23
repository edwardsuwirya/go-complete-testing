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
	server := NewServer(new(mockUseCaseManager))
	assert.NotNil(suite.T(), server.RouterEngine)
	assert.NotNil(suite.T(), server.routers)
	assert.NotNil(suite.T(), server.publicRoute)
	assert.Equal(suite.T(), "/api", server.publicRoute.BasePath())
}
func TestDeliveryInitTestSuite(t *testing.T) {
	suite.Run(t, new(DeliveryInitTestSuite))
}

package delivery

import (
	"enigmacamp.com/completetesting/usecase"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type mockInfra struct {
}

func (i *mockInfra) SqlDb() *sqlx.DB {
	mockdb, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockdb, "sqlmock")
	return sqlxDB
}

type mockUseCaseManager struct {
}

func (uc *mockUseCaseManager) StudentUseCase() usecase.IStudentUseCase {
	return nil
}

type DeliveryInitTestSuite struct {
	suite.Suite
}

func (suite *DeliveryInitTestSuite) TestNewServer() {
	mockRoutes := NewServer(new(mockInfra), new(mockUseCaseManager))
	assert.NotNil(suite.T(), mockRoutes.routerEngine)
	assert.NotNil(suite.T(), mockRoutes.routers)
	assert.NotNil(suite.T(), mockRoutes.publicRoute)
}
func (suite *DeliveryInitTestSuite) TestStartEngine() {
	mockRoutes := NewServer(new(mockInfra), new(mockUseCaseManager))
	engine := mockRoutes.StartEngine()
	assert.NotNil(suite.T(), engine)
}

func TestDeliveryInitTestSuite(t *testing.T) {
	suite.Run(t, new(DeliveryInitTestSuite))
}

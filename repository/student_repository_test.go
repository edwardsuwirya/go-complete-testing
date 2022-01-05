package repository

import (
	"enigmacamp.com/completetesting/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

var dummyStudents = []model.Student{
	{
		Id:       1,
		Name:     "Dummy name 1",
		Gender:   "M",
		Age:      1,
		JoinDate: time.Time{},
		IdCard:   "dummy id card 1",
		Senior:   false,
	}, {
		Id:       2,
		Name:     "Dummy name 2",
		Gender:   "F",
		Age:      2,
		JoinDate: time.Time{},
		IdCard:   "dummy id card 2",
		Senior:   true,
	},
}

type StudentRepositoryTestSuite struct {
	suite.Suite
	mockResource *sqlx.DB
	mock         sqlmock.Sqlmock
}

func (suite *StudentRepositoryTestSuite) SetupTest() {
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(mockdb, "sqlmock")
	suite.mockResource = sqlxDB
	suite.mock = mock
}

func (suite *StudentRepositoryTestSuite) TearDownTest() {
	suite.mockResource.Close()
}

func (suite *StudentRepositoryTestSuite) TestStudentRepository_GetAll() {
	rows := sqlmock.NewRows([]string{"id", "name", "gender", "age", "join_date", "id_card", "senior"})
	for _, d := range dummyStudents {
		rows.AddRow(d.Id, d.Name, d.Gender, d.Age, d.JoinDate, d.IdCard, d.Senior)
	}
	suite.mock.ExpectQuery("SELECT (.+) FROM M_STUDENT").
		WillReturnRows(rows)

	repo := NewStudentRepository(suite.mockResource)
	all, err := repo.GetAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(all))
	assert.Equal(suite.T(), 1, all[0].Id)
}
func TestStudentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StudentRepositoryTestSuite))
}

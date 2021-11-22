package delivery

import (
	"bytes"
	"encoding/json"
	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
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

type MockResponse struct {
	Message model.Student
}
type MockErrorResponse struct {
	Message string
}

type studentUseCaseMock struct {
	mock.Mock
}

func (s *studentUseCaseMock) NewRegistration(student model.Student) (*model.Student, error) {
	args := s.Called(student)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), nil
}

func (s *studentUseCaseMock) FindStudentInfoById(idCard string) (*model.Student, error) {
	args := s.Called(idCard)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), nil
}

type StudentApiTestSuite struct {
	suite.Suite
	useCaseTest     usecase.IStudentUseCase
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *StudentApiTestSuite) SetupTest() {
	suite.useCaseTest = new(studentUseCaseMock)
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/api")
}

func (suite *StudentApiTestSuite) TestStudentApi_CreateStudent_Success() {
	dummyStudent := dummyStudents[1]
	suite.useCaseTest.(*studentUseCaseMock).On("NewRegistration", dummyStudent).Return(&dummyStudent, nil)

	studentApi := NewStudentApi(suite.useCaseTest).(*StudentApi)
	studentApi.InitRouter(suite.routerGroupTest)
	handler := studentApi.createStudent
	suite.routerTest.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyStudent)
	request, _ := http.NewRequest(http.MethodPost, "/api/student", bytes.NewBuffer(reqBody))
	request.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 200)

	//expectedRespBody, _ := json.Marshal(gin.H{
	//	"message": dummyStudent,
	//})
	//assert.Equal(suite.T(), expectedRespBody, rr.Body.Bytes())
	a := rr.Body.String()
	actualStudent := new(MockResponse)
	json.Unmarshal([]byte(a), actualStudent)
	assert.Equal(suite.T(), dummyStudent.Name, actualStudent.Message.Name)
}
func (suite *StudentApiTestSuite) TestStudentApi_CreateStudent_FailedBinding() {
	suite.useCaseTest.(*studentUseCaseMock).On("NewRegistration", nil).Return(nil, errors.New("failed"))
	studentApi := NewStudentApi(suite.useCaseTest).(*StudentApi)
	studentApi.InitRouter(suite.routerGroupTest)
	handler := studentApi.createStudent
	suite.routerTest.POST("", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/api/student", nil)
	request.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 400)
}

func (suite *StudentApiTestSuite) TestStudentApi_CreateStudent_FailedUseCase() {
	dummyStudent := dummyStudents[1]
	suite.useCaseTest.(*studentUseCaseMock).On("NewRegistration", dummyStudent).Return(nil, errors.New("failed"))
	studentApi := NewStudentApi(suite.useCaseTest).(*StudentApi)
	studentApi.InitRouter(suite.routerGroupTest)
	handler := studentApi.createStudent
	suite.routerTest.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyStudent)
	request, _ := http.NewRequest(http.MethodPost, "/api/student", bytes.NewBuffer(reqBody))
	request.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 500)
	a := rr.Body.String()
	actualError := new(MockErrorResponse)
	json.Unmarshal([]byte(a), actualError)
	assert.Equal(suite.T(), "failed", actualError.Message)
}

func (suite *StudentApiTestSuite) TestStudentApi_GetById_Success() {
	dummyStudent := dummyStudents[0]
	suite.useCaseTest.(*studentUseCaseMock).On("FindStudentInfoById", "2").Return(&dummyStudent, nil)

	studentApi := NewStudentApi(suite.useCaseTest).(*StudentApi)
	studentApi.InitRouter(suite.routerGroupTest)
	handler := studentApi.getStudentById
	suite.routerTest.GET("/:idcard", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/student/2", nil)
	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 200)

	a := rr.Body.String()
	actualStudent := new(MockResponse)
	json.Unmarshal([]byte(a), actualStudent)
	assert.Equal(suite.T(), dummyStudent.Name, actualStudent.Message.Name)
}

func TestStudentApiTestSuite(t *testing.T) {
	suite.Run(t, new(StudentApiTestSuite))
}

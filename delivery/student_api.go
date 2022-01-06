package delivery

import (
	appresponse "enigmacamp.com/completetesting/delivery/app_response"
	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/usecase"
	"enigmacamp.com/completetesting/util/app_status"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentApi struct {
	BaseApi
	useCase     usecase.IStudentUseCase
	publicRoute *gin.RouterGroup
}

func NewStudentApi(publicRoute *gin.RouterGroup, usecase usecase.IStudentUseCase) (*StudentApi, error) {
	if publicRoute == nil || usecase == nil {
		return nil, errors.New("Empty Router or UseCase")
	}
	studentApi := StudentApi{
		useCase:     usecase,
		publicRoute: publicRoute,
	}
	studentApi.InitRouter()
	return &studentApi, nil
}
func (api *StudentApi) InitRouter() {
	studentRoute := api.publicRoute.Group("/student")
	studentRoute.GET("", api.getAllStudent)
	studentRoute.GET("/:idcard", api.getStudentById)
	studentRoute.POST("", api.createStudent)
}

func (api *StudentApi) getAllStudent(c *gin.Context) {
	students, err := api.useCase.GetStudentList()
	resp := appresponse.NewJsonResponse(c)
	if err != nil {
		errMsg := appresponse.NewErrorMessage(http.StatusInternalServerError, app_status.GeneralError, "Failed Get Student List")
		api.errLogging(c, err, errMsg)
		resp.SendError(errMsg)
		return
	}
	resp.SendData(appresponse.NewResponseMessage(app_status.Success, "Student List", students))
}
func (api *StudentApi) getStudentById(c *gin.Context) {
	name := c.Param("idcard")
	student, err := api.useCase.FindStudentInfoById(name)
	resp := appresponse.NewJsonResponse(c)
	if err != nil {
		errMsg := appresponse.NewErrorMessage(http.StatusBadRequest, app_status.GeneralError, "Failed Get Student By ID")
		api.errLogging(c, err, errMsg)
		resp.SendError(errMsg)
		return
	}
	resp.SendData(appresponse.NewResponseMessage(app_status.Success, "Student By Id", student))
}
func (api *StudentApi) createStudent(c *gin.Context) {
	var student model.Student
	resp := appresponse.NewJsonResponse(c)
	err := c.BindJSON(&student)
	if err != nil {
		errMsg := appresponse.NewErrorMessage(http.StatusBadRequest, app_status.ErrorLackInfo, app_status.StatusText(app_status.ErrorLackInfo))
		api.errLogging(c, err, errMsg)
		resp.SendError(errMsg)
		return
	}
	registeredStudent, err := api.useCase.NewRegistration(student)
	if err != nil {
		errMsg := appresponse.NewErrorMessage(http.StatusInternalServerError, app_status.GeneralError, "Failed Create Student")
		api.errLogging(c, err, errMsg)
		resp.SendError(errMsg)
		return
	}
	resp.SendData(appresponse.NewResponseMessage(app_status.Success, "Student Registration", registeredStudent))
}

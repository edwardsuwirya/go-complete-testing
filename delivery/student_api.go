package delivery

import (
	appresponse "enigmacamp.com/completetesting/delivery/app_response"
	"enigmacamp.com/completetesting/util/app_status"
	"enigmacamp.com/completetesting/util/logger"
	"errors"
	"net/http"

	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/usecase"
	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	useCase     usecase.IStudentUseCase
	publicRoute *gin.RouterGroup
	logger      *logger.AppLogger
}

func NewStudentApi(publicRoute *gin.RouterGroup, usecase usecase.IStudentUseCase, logger *logger.AppLogger) (*StudentApi, error) {
	if publicRoute == nil || usecase == nil {
		return nil, errors.New("Empty Router or UseCase")
	}
	studentApi := StudentApi{
		useCase:     usecase,
		publicRoute: publicRoute,
		logger:      logger,
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
	resp := appresponse.NewJsonResponse(c, api.logger)
	if err != nil {
		resp.SendError(http.StatusInternalServerError, appresponse.NewErrorMessage(app_status.GeneralError, "Failed Get Student List"), err)
		return
	}
	resp.SendData(appresponse.NewResponseMessage(app_status.Success, "Student List", students))
}
func (api *StudentApi) getStudentById(c *gin.Context) {
	name := c.Param("idcard")
	student, err := api.useCase.FindStudentInfoById(name)
	resp := appresponse.NewJsonResponse(c, api.logger)
	if err != nil {
		resp.SendError(http.StatusBadRequest, appresponse.NewErrorMessage(app_status.GeneralError, "Failed Get Student By ID"), err)
		return
	}
	resp.SendData(appresponse.NewResponseMessage(app_status.Success, "Student By Id", student))
}
func (api *StudentApi) createStudent(c *gin.Context) {
	var student model.Student
	resp := appresponse.NewJsonResponse(c, api.logger)
	err := c.BindJSON(&student)
	if err != nil {
		resp.SendError(http.StatusBadRequest, appresponse.NewErrorMessage(app_status.ErrorLackInfo, app_status.StatusText(app_status.ErrorLackInfo)), err)
		return
	}
	registeredStudent, err := api.useCase.NewRegistration(student)
	if err != nil {
		resp.SendError(http.StatusInternalServerError, appresponse.NewErrorMessage(app_status.GeneralError, "Failed Create Student"), err)
		return
	}
	resp.SendData(appresponse.NewResponseMessage(app_status.Success, "Student Registration", registeredStudent))
}

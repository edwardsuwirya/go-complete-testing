package delivery

import (
	appresponse "enigmacamp.com/completetesting/httputil"
	"errors"
	"net/http"

	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/usecase"
	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	usecase     usecase.IStudentUseCase
	publicRoute *gin.RouterGroup
}

func NewStudentApi(publicRoute *gin.RouterGroup, usecase usecase.IStudentUseCase) (*StudentApi, error) {
	if publicRoute == nil || usecase == nil {
		return nil, errors.New("Empty Router or UseCase")
	}
	studentApi := StudentApi{
		usecase:     usecase,
		publicRoute: publicRoute,
	}
	studentApi.InitRouter()
	return &studentApi, nil
}
func (api *StudentApi) InitRouter() {
	studentRoute := api.publicRoute.Group("/student")
	studentRoute.GET("/:idcard", api.getStudentById)
	studentRoute.POST("", api.createStudent)
}

// getStudentById godoc
// @Summary      Show student
// @Description  get student by ID Card
// @Tags         students
// @Produce      json
// @Param        idcard   path      int  true  "Student ID Card"
// @Success      200  {object}  model.Student
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /student/{idcard} [get]
func (api *StudentApi) getStudentById(c *gin.Context) {
	name := c.Param("idcard")
	student, err := api.usecase.FindStudentInfoById(name)
	resp := appresponse.NewJsonResponse(c)
	if err != nil {
		resp.SendError(appresponse.NewErrorMessage(http.StatusBadRequest, "X00", err.Error()))
		return
	}
	resp.SendData(appresponse.NewResponseMessage("00", "", student))
}
func (api *StudentApi) createStudent(c *gin.Context) {
	var student model.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "Data is required"})
		return
	}
	registeredStudent, err := api.usecase.NewRegistration(student)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": registeredStudent,
	})
}

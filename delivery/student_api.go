package delivery

import (
	"log"
	"net/http"

	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/usecase"
	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	router  *gin.RouterGroup
	usecase usecase.IStudentUseCase
}

func NewStudentApi(router *gin.RouterGroup, usecase usecase.IStudentUseCase) *StudentApi {
	userRoute := router.Group("/student")
	studentApi := StudentApi{
		router:  userRoute,
		usecase: usecase,
	}
	studentApi.initRouter()
	return &studentApi
}
func (api *StudentApi) initRouter() {
	api.router.GET("/:idcard", api.getStudentById)
	api.router.POST("", api.createStudent)
}

func (api *StudentApi) getStudentById(c *gin.Context) {
	name := c.Param("idcard")
	student, err := api.usecase.FindStudentInfoById(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": student,
	})
}
func (api *StudentApi) createStudent(c *gin.Context) {
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		log.Println(err)
		return
	}
	registeredStudent, err := api.usecase.NewRegistration(student)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": registeredStudent,
	})
}

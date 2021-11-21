package delivery

import (
	"log"
	"net/http"

	"enigmacamp.com/completetesting/db"
	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/repository"
	"enigmacamp.com/completetesting/usecase"
	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	router  *gin.RouterGroup
	usecase usecase.IStudentUseCase
}

func NewStudentApi(router *gin.RouterGroup, resource *db.Resource) *StudentApi {
	userRoute := router.Group("/student")
	studentRepo := repository.NewStudentRepository(resource)
	studentApi := StudentApi{
		router:  userRoute,
		usecase: usecase.NewStudentUseCase(studentRepo),
	}
	studentApi.initRouter()
	return &studentApi
}
func (api *StudentApi) initRouter() {
	api.router.GET("/:name", api.getStudentById)
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

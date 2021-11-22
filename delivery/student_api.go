package delivery

import (
	"net/http"

	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/usecase"
	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	usecase usecase.IStudentUseCase
}

func NewStudentApi(usecase usecase.IStudentUseCase) IDelivery {
	studentApi := StudentApi{
		usecase: usecase,
	}
	return &studentApi
}
func (api *StudentApi) InitRouter(publicRoute *gin.RouterGroup) {
	userRoute := publicRoute.Group("/student")
	userRoute.GET("/:idcard", api.getStudentById)
	userRoute.POST("", api.createStudent)
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
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
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

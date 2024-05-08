package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nathaliapavan/edu-api/shared"
)

func routeHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

type Student struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

var Students = []Student{
	{ID: shared.GetUuid(), Name: "Peter Parker", Age: 18},
}

func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func routeGetStudent(c *gin.Context) {
	idByParam := c.Params.ByName("id")
	uuid, err := shared.GetUuidByString(idByParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error generating uuid",
		})
		return
	}

	studentIndex := -1
	var studentFound Student
	for i, studentElement := range Students {
		if studentElement.ID == uuid {
			studentIndex = i
			studentFound = studentElement
			break
		}
	}

	if studentIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}

	c.JSON(http.StatusOK, studentFound)
}

func routePostStudent(c *gin.Context) {
	var student Student
	err := c.Bind(&student)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "unable to get payload",
		})
		return
	}
	student.ID = shared.GetUuid()
	Students = append(Students, student)
	c.JSON(http.StatusCreated, student)
}

func routePutStudent(c *gin.Context) {
	var studentToUpdate Student
	var matchStudent Student

	err := c.BindJSON(&studentToUpdate)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "unable to get payload",
		})
		return
	}

	idByParam := c.Params.ByName("id")
	uuid, err := shared.GetUuidByString(idByParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error generating uuid",
		})
		return
	}

	studentIndex := -1
	for i, studentElement := range Students {
		if studentElement.ID == uuid {
			studentIndex = i
			matchStudent = studentElement
			break
		}
	}

	if studentIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}

	matchStudent.Name = studentToUpdate.Name
	matchStudent.Age = studentToUpdate.Age
	Students[studentIndex] = matchStudent

	c.JSON(http.StatusOK, gin.H{
		"message": "student updated successfully",
	})
}

func routeDeleteStudent(c *gin.Context) {
	idByParam := c.Params.ByName("id")
	uuid, err := shared.GetUuidByString(idByParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error generating uuid",
		})
		return
	}

	studentIndex := -1
	for i, studentElement := range Students {
		if studentElement.ID == uuid {
			studentIndex = i
			break
		}
	}

	if studentIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}

	Students = append(Students[:studentIndex], Students[studentIndex+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": "student removed successfully",
	})
}

func main() {
	gin.ForceConsoleColor()
	service := gin.Default()
	getRoutes(service)
	service.Run()
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/health", routeHealth)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	groupStudents.GET("/:id", routeGetStudent)
	groupStudents.POST("/", routePostStudent)
	groupStudents.PUT("/:id", routePutStudent)
	groupStudents.DELETE("/:id", routeDeleteStudent)

	return c
}

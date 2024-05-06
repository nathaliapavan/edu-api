package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func routeHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Students = []Student{
	{ID: 1, Name: "Peter Parker", Age: 18},
}

func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func routeGetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "id not found",
		})
		return
	}

	var student Student
	for _, studentElement := range Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
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
	student.ID = len(Students) + 1
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

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "id not found",
		})
		return
	}

	studentIndex := -1
	for i, studentElement := range Students {
		if studentElement.ID == id {
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id not found",
		})
	}

	studentIndex := -1
	for i, studentElement := range Students {
		if studentElement.ID == id {
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

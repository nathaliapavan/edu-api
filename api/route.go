package api

import (
	infra_controller "github.com/nathaliapavan/edu-api/api/controller/infra"
	student_controller "github.com/nathaliapavan/edu-api/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/health", infra_controller.Health)

	groupStudent := s.Engine.Group("/students")
	groupStudent.GET("/", student_controller.List)
	groupStudent.GET("/:id", student_controller.Details)
	groupStudent.POST("/", student_controller.Create)
	groupStudent.PUT("/:id", student_controller.Update)
	groupStudent.DELETE("/:id", student_controller.Delete)
}

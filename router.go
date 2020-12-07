package main

import (
	"neko_server_go"
	"neko_server_go/core"
	"neko_server_go/enum"
	"neko_server_go_example/handler"
)

var Router = neko_server_go.Router{
	"/enrollment/view":      core.CreateMethodsHandler(enum.HttpMethodsGet, handler.EnrollmentView),
	"/enrollment/add":       core.CreateMethodsHandler(enum.HttpMethodsPost, handler.EnrollmentAdd),
	"/enrollment/statistic": core.CreateMethodsHandler(enum.HttpMethodsGet, handler.EnrollmentStatistic),
	"/course/add":           core.CreateMethodsHandler(enum.HttpMethodsPost, handler.CourseAdd),
	"/course":               core.CreateMethodsHandler(enum.HttpMethodsGet, handler.CourseAddTemplate),
	"/":                     core.CreateMethodsHandler(enum.HttpMethodsPost, handler.Index),
}

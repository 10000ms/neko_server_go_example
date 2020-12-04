package main

import (
	"neko_server_go"
	"neko_server_go_example/handler"
)

var Router = neko_server_go.Router{
	"/enrollment/view":      handler.EnrollmentView,
	"/enrollment/add":       handler.EnrollmentAdd,
	"/enrollment/statistic": handler.EnrollmentStatistic,
	"/course/add":           handler.CourseAdd,
	"/course":               handler.CourseAddTemplate,
	"/":                     handler.Index,
}

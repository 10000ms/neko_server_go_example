package handler

import (
	"fmt"
	"neko_server_go"
	"neko_server_go/core"
	"neko_server_go/utils"
	"neko_server_go_example/db"
	"strconv"
)

func EnrollmentView(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	l := core.TemplateLoader{
		Path: c.App.Setting["Path"].(string),
	}
	render := core.Render{
		TemplateLoader: l,
	}
	courses, err := db.QueryAllCourse(c, c.App.Db["Default"])
	a := map[string]interface{}{
		"courses": courses,
	}
	t := render.Render("/template/student/enrollment.html", a)
	_, err = fmt.Fprintf(w, t) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}

func EnrollmentAdd(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	courseId := c.Request.URL.Query()["course_id"]
	age := c.Request.URL.Query()["age"][0]
	email := c.Request.URL.Query()["email"][0]
	name := c.Request.URL.Query()["name"][0]

	// TODO：支持事务
	SID, err := db.InsertStudent(c, c.App.Db["Default"], name, email, age)
	if err != nil {
		_, _ = fmt.Fprintf(w, "error")
		return
	}
	courseIds := make([]int64, 0)
	for i := 0; i < len(courseId); i++ {
		i, _ := strconv.ParseInt(courseId[i], 10, 64)
		courseIds = append(courseIds, i)
	}
	err = db.Enrollment(c, c.App.Db["Default"], SID, courseIds)
	msg := ""
	if err != nil {
		msg = c.LogRequestID() + " Insert Error"
	} else {
		msg = c.LogRequestID() + " Insert Success"
	}

	_, err = fmt.Fprintf(w, msg) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}

}

func EnrollmentStatistic(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	l := core.TemplateLoader{
		Path: c.App.Setting["Path"].(string),
	}
	render := core.Render{
		TemplateLoader: l,
	}
	statistic, err := db.EnrollmentStatistic(c, c.App.Db["Default"])
	a := map[string]interface{}{
		"statistic": statistic,
	}
	t := render.Render("/template/enrollment/statistic.html", a)
	_, err = fmt.Fprintf(w, t) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}

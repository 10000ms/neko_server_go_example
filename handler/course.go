package handler

import (
	"fmt"
	"neko_server_go"
	"neko_server_go/core"
	"neko_server_go/utils"
	"neko_server_go_example/db"
)

func CourseAddTemplate(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	l := core.TemplateLoader{
		Path: c.App.Setting["Path"].(string),
	}
	t := l.GetSource("/template/course/add.html")
	utils.LogInfo("in CourseAddTemplate")
	_, err := fmt.Fprintf(w, t) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}

func CourseAdd(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	courseName := c.Request.URL.Query()["name"][0]

	err := db.InsertCourse(c, c.App.Db["Default"], courseName)
	var msg string
	if err != nil {
		msg = c.LogRequestID() + " Insert Error : " + courseName
	} else {
		msg = c.LogRequestID() + " Insert Success : " + courseName
	}

	_, err = fmt.Fprintf(w, msg) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}

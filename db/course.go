package db

import (
	"database/sql"
	"neko_server_go"
	"neko_server_go/utils"
)

type Course struct {
	ID int64
	Name string
}


func InsertCourse(c *neko_server_go.Context, db *sql.DB, courseName string) error {
	insertSQL := `
	     INSERT INTO course (name)
         VALUES (?);
    `
	_, err := db.Exec(insertSQL, courseName)
	if err != nil {
		utils.LogError(c.LogRequestID(), "insert course error: ", err)
		return err
	} else {
		utils.LogInfo(c.LogRequestID(), "insert course success!")
		return nil
	}
}

func QueryAllCourse(c *neko_server_go.Context, db *sql.DB) ([]*Course, error) {
	querySQL := `
 		SELECT 
			id, name
		FROM
			course
	`
	rows, err := db.Query(querySQL)
	if err != nil {
		utils.LogError(c.LogRequestID(), "query course error: ", err)
		return nil, err
	}
	courses := make([]*Course, 0)
	for rows.Next() {
		course := Course{}
		err = rows.Scan(&course.ID, &course.Name)
		if err != nil {
			utils.LogError(c.LogRequestID(), "query course error: ", err)
		} else {
			courses = append(courses, &course)
		}
	}
	utils.LogDebug("QueryAllCourse: ", courses)
	return courses, nil
}

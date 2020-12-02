package db

import (
	"database/sql"
	"neko_server_go"
	"neko_server_go/utils"
)

type EnrollmentS struct {
	Name string
	Count int64
	AvgAge float64
}

func Enrollment(c *neko_server_go.Context, db *sql.DB, studentID int64, courseID []int64) error {
	for i := 0; i < len(courseID); i++ {
		insertSQL := `
			INSERT INTO student_course (student_id, course_id)
			VALUES (?,?);
		`
		_, err := db.Exec(insertSQL, studentID, courseID[i])
		if err != nil {
			utils.LogError(c.LogRequestID(), "insert Enrollment error: ", err)
			return err
		} else {
			utils.LogInfo(c.LogRequestID(), "insert Enrollment success:")
		}
	}

	return nil
}


func EnrollmentStatistic(c *neko_server_go.Context, db *sql.DB) ([]*EnrollmentS, error) {
	JoinSQL := `
		SELECT course.name, count(student.id), avg(student.age)
		FROM course
		INNER JOIN student_course ON course.id = student_course.course_id
		INNER JOIN student ON student.id = student_course.student_id
		GROUP BY course.name
	`
	rows, err := db.Query(JoinSQL)
	if err != nil {
		utils.LogError(c.LogRequestID(), "query EnrollmentStatistic error: ", err)
		return nil, err
	}
	es := make([]*EnrollmentS, 0)
	for rows.Next() {
		s := EnrollmentS{}
		err = rows.Scan(&s.Name, &s.Count, &s.AvgAge)
		if err != nil {
			utils.LogError(c.LogRequestID(), "query EnrollmentStatistic error: ", err)
		} else {
			es = append(es, &s)
		}
	}
	utils.LogDebug("QueryAllCourse: ", es)
	return es, nil


}

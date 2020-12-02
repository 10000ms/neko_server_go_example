package db

import (
	"database/sql"
	"neko_server_go"
	"neko_server_go/utils"
)

func InsertStudent(c *neko_server_go.Context, db *sql.DB, name string, age string, email string) (int64, error) {
	insertSQL := `
	     INSERT INTO student (name, email, age)
         VALUES (?,?,?);
    `
	r, err := db.Exec(insertSQL, name, age, email)
	if err != nil {
		utils.LogError(c.LogRequestID(), "insert course error: ", err)
		return 0, err
	} else {
		lId, _ := r.LastInsertId()
		utils.LogInfo(c.LogRequestID(), "insert course success: ", lId)
		return lId, nil
	}
}

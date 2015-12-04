package alarm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Service struct {
	Real_address               string
	Esb_address                string
	Name                       string
	Description                string
	Register_user_mail_address string
	Last_modify_date           time.Time
	Environment                string
	Major_version              int
	Minor_version              int
}

func LoadServiceRegisterUserEmail(appname string) (string, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/esb_management?charset=utf8")
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	rows, err := db.Query("select register_user_mail_address from services where name=?", appname)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	var email string
	for rows.Next() {
		row_err := rows.Scan(&email)
		if row_err != nil {
			log.Println(err)
		}
	}

	return email, err
}

package dto

import (
	"database/sql"

	// importing the SQL driver
	_ "github.com/go-sql-driver/mysql"
)

// DTO structure
type DTO struct {
	id     int
	status string
}

func InitDTO(dataSourceName string) (*sql.DB, error) {
	if len(dataSourceName) == 0 {
		dataSourceName = "root@/loa"
	}
	db, err := sql.Open("mysql", dataSourceName)
	checkErr(err)

	return db, err

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

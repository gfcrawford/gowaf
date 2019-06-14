package dto

import (
	"database/sql"
	"fmt"

	// importing the SQL driver
	_ "github.com/go-sql-driver/mysql"
)

// DTO structure
type DTO struct {
	ID     int
	Status string
	Error  error
	DB     *sql.DB
}

// Init function initialises the base DTO struct
func (d *DTO) Init(dataSourceName string) {
	if len(dataSourceName) == 0 {
		dataSourceName = "root@/loa"
	}
	db, err := sql.Open("mysql", dataSourceName)
	d.Status = "Initialised"
	d.Error = err
	d.DB = db
}

// Destroy function closes the db connection in the DTO struct
func (d *DTO) Destroy() {
	d.DB.Close()
}

// Count all the objects in a table
func (d *DTO) Count(whereClause ...string) int {
	rows, err := d.DB.Query("SELECT count(*) FROM player")
	if err == nil {
		for rows.Next() {
            	var count int
            	err = rows.Scan(&count)
            	fmt.Println("Count:",count)
        	}
	}	
	return 99
}

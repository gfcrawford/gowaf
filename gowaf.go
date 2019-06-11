package main

import (
	"fmt"

	"github.com/gfcrawford/gowaf/include/dto"
)

// User struct
type User struct {
	dto.DTO
	username string
	password string
}

// Fields returns an array of the struct fields
func (u *User) Fields() [2]string {
	var a [2]string
	a[0] = "username"
	a[1] = "password"
	return a
}

func main() {
	fmt.Println("gowaf")
	d := new(dto.DTO)
	fmt.Println(d.Status)
	d.Init("Fred:Flintstone@/loa")
	rows, err := d.DB.Query("SELECT * FROM player")
	if err == nil {
		for rows.Next() {
			var id int
			var name string
			var familyname string
			var craft int
			err = rows.Scan(&id, &name, &familyname, &craft)

			fmt.Println(id)
			fmt.Println(name)
			fmt.Println(familyname)
			fmt.Println(craft)

		}
	}
	fmt.Println(d.Count())
}

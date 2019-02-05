package main

import (
	"fmt"

	"github.com/gfcrawford/gowaf/include/dto"
)

func main() {
	fmt.Println("gowaf")
	db, error := dto.InitDTO("root@/loa")
	if error == nil {
		// rows, err := db.Query("SELECT * FROM player")
		// if err == nil {
		// 	for rows.Next() {
		// 		var id int
		// 		var name string
		// 		var familyname string
		// 		var craft int
		// 		err = rows.Scan(&id, &name, &familyname, &craft)

		// 		fmt.Println(id)
		// 		fmt.Println(name)
		// 		fmt.Println(familyname)
		// 		fmt.Println(craft)

		// 	}
		// }

		db.Close()
	}
}

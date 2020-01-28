package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//Connect to the "company_db" database
	db, err :=
		sql.Open("postgres", "postgresql://root@localhost:26257/company_db?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	//Inserting a Row in to DB
	if _, err := db.Exec(
		`INSERT INTO tbl_employee (full_name, department, designation, created_at, update_at) 
		VALUES ('Julian Salgado', 'IT', 'Developer', NOW(), NOW());`); err != nil {
		log.Fatal(err)

	}

	//Select statement
	rows, err := db.Query("select employee_id, full_name FROM tbl_employee;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employeeID int64
		var fullName string
		if err := rows.Scan(&employeeID, &fullName); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("EMployee ID: %d \t Employee Name: %s\n",
			employeeID, fullName)
	}

	fmt.Println("Mensaje para que no haya error, graciaaas")
}

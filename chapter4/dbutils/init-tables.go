package dbutils

import (
	"database/sql"
	"fmt"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}

	_, statementError := statement.Exec()

	if statementError != nil {
		log.Println(statementError)
	}

	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()

	statement, _ = dbDriver.Prepare(station)
	statement.Exec()

	fmt.Println("Successfully build")
}

package db

import (
	"database/sql"
	 "fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	dbent "internal/db/entities"
)

// ModelQ is an interface that ensure implement Queries on the db
type ModelQ interface {
	QueryDB(n int) []dbent.Author
}

// DBModel manages a db and its manager
type DBModel struct {
	DB      *sql.DB
	Manager DBManager
}

func (m DBModel) QueryDB(n int) []dbent.Author {
	registers := make([]dbent.Author, 0)
  strQuery := fmt.Sprintf(m.Manager.QueryStr(),n)
	rows, _ := m.DB.Query(strQuery)

	for rows.Next() {
		var id, name, date string
		if err := rows.Scan(&id, &name, &date); err != nil {
			// Check for a scan error.
			log.Fatal(err)
		}
		timecu, _ := time.Parse(time.RFC3339, date)

		iid, _ := strconv.Atoi(id)

		newreg := dbent.Author{
			Id:    int64(iid),
			Name:  name,
			Dated: timecu,
		}
		registers = append(registers, newreg)
	}
	return registers
}

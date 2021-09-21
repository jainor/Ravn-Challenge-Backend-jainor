package main

import (
	"database/sql"
	"fmt"

	idb "internal/db"
	mq "internal/messagequeue"
)

func main() {

	dbmanager := idb.GetManager()
	psqlInfo := dbmanager.ConnectionStr()

	fmt.Println(psqlInfo)
	db, err := sql.Open(dbmanager.Name(), psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	//err = db.Ping()

	if err != nil {
		panic(err)
	}

	model := idb.DBModel{DB: db, Manager: dbmanager}

	mq.ResponseMessage(model)
}

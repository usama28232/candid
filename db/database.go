package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/usama28232/candid/shared"

	_ "github.com/lib/pq"
)

var (
	db    *sql.DB = nil
	mutex sync.Mutex
)

func Init() (string, error) {

	dbHost, _ := shared.GetConfigByKey(shared.DB_HOST)
	dbPort, _ := shared.GetConfigByKey(shared.DB_PORT)
	dbName, _ := shared.GetConfigByKey(shared.DB_NAME)
	dbUser, _ := shared.GetConfigByKey(shared.DB_USER)
	dbPw, _ := shared.GetConfigByKey(shared.DB_PW)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost.(string), int(dbPort.(float64)), dbUser.(string), dbPw.(string), dbName.(string))
	var err error
	// open database
	if db != nil {
		return psqlconn, nil
	}
	mutex.Lock()
	defer mutex.Unlock()
	db, err = sql.Open("postgres", psqlconn)
	checkError(err)

	return psqlconn, err
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func getConnection() (*sql.DB, error) {
	return db, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Execute(query string, args ...any) error {
	conn, connErr := getConnection()
	if connErr == nil {
		var err error
		if len(args) > 0 {
			_, err = conn.Exec(query, args...)

		} else {
			_, err = conn.Exec(query)
		}
		return err
	}
	return connErr
}

func Query(query string, args ...any) (*([][]any), error) {
	conn, connErr := getConnection()
	if connErr != nil {
		return nil, connErr
	}
	var rows *sql.Rows
	var err error
	if len(args) > 0 {
		rows, err = conn.Query(query, args...)

	} else {
		rows, err = conn.Query(query)
	}
	if err == nil {
		defer rows.Close()
		cols, _ := rows.Columns()
		data := [][]any{}
		for rows.Next() {
			columns, columnPointers := generateColumnPointers(cols)
			err := rows.Scan(columnPointers...)
			if err == nil {
				data = append(data, columns)
			}
		}

		return &data, nil
	} else {
		fmt.Println("Error in Query", err)
	}
	// db.Close()
	return nil, err
}

func generateColumnPointers(cols []string) ([]any, []any) {
	columnCount := len(cols)
	columns := make([]any, columnCount)
	columnPointers := make([]any, columnCount)
	for i := 0; i < columnCount; i++ {
		columnPointers[i] = &columns[i]
	}
	return columns, columnPointers
}

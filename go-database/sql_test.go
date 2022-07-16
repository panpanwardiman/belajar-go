package godatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES(2, 'Wardiman')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully insert new customer.")
}

func TestSelectQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT * FROM")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}

	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	id := "2"
	name := "Wardiman"

	query := "SELECT * FROM customer where id = ? and name = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, id, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Name: ", name)
	} else {
		fmt.Println("Failed")
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sql := "INSERT INTO customer (name) VALUE (?)"
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		name := "arjun " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, name)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully insert with ID ", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}

package db

import (
	"awesomeProject2/parser"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type Expression struct {
	ExpressionId int
	Expression   string
	Status       string
	Result       float64
	TimeSpent    int // in milliseconds
}

type Operation struct {
	Operation string
	Time      int
}

const dbPath = "db/db.db"

func AddExpressionToDB(exp Expression) error {

	exp.Expression = parser.ToPostfix(exp.Expression)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into equasions (expressionId, expression, status, result, timeSpent) VALUES ((SELECT max(expressionId) from equasions) + 1, ?, ?, ?, ?)",
		exp.Expression, exp.Status, exp.Result, exp.TimeSpent)
	return err
}

func GetResult(id int) (float64, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT result from equasions where expressionId = ?", id)

	var result float64
	for rows.Next() {
		err = rows.Scan(&result)
		return result, err
	}
	return 0, errors.New("no result found")
}

func SetResult(expression string, result float64) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE equasions set result = ?, status = 'done' where expression = ?", result, expression)
	return err
}

func GetExpressions() ([]Expression, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * from equasions")

	var result []Expression
	for rows.Next() {
		p := Expression{}
		err := rows.Scan(&p.ExpressionId, &p.Expression, &p.Status, &p.Result, &p.TimeSpent)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}

func AddOperation(oper Operation) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into operations (operation, time) VALUES (?, ?)",
		oper.Operation, oper.Time)
	return err
}

func GetOperations() ([]Operation, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * from operations	")

	var result []Operation
	for rows.Next() {
		p := Operation{}
		err := rows.Scan(&p.Operation, &p.Time)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}

func GetQueue() ([]Expression, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * from equasions where status = 'sent'")

	var result []Expression
	if rows == nil {
		return make([]Expression, 0), nil
	}
	for rows.Next() {
		p := Expression{}
		err := rows.Scan(&p.ExpressionId, &p.Expression, &p.Status, &p.Result, &p.TimeSpent)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}

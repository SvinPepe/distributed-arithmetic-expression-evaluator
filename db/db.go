package db

import (
	"database/sql"
	"errors"
)

type Expression struct {
	expressionId int
	expression   string
	status       string
	result       float64
	timeSpent    int // in milliseconds
}

func AddExpressionToDB(exp Expression) error {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into equasions (expressionId, expression, status, result, timeSpent) VALUES (?, ?, ?, ?, ?)",
		exp.expressionId, exp.expression, exp.status, exp.result, exp.timeSpent)
	return err
}

func GetResult(id int) (float64, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT from equasions result where expressionId = ?", id)

	var result float64
	for rows.Next() {
		err = rows.Scan(&result)
		return result, err
	}
	return 0, errors.New("no result found")
}

func GetExpressions() ([]Expression, error) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * from equasions")

	var result []Expression
	for rows.Next() {
		p := Expression{}
		err := rows.Scan(&p.expressionId, &p.expression, &p.status, &p.result, &p.timeSpent)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}

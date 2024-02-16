package main

import (
	. "awesomeProject2/db"
)

func AddExpression(exp Expression) {
	err := AddExpressionToDB(exp)
	if err != nil {
		return // TODO отправить код 500 поидее
	}

}

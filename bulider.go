package sql

import "github.com/WebXense/sql/stm"

// Eq equals to "=" in SQL expressions
func Eq(field string, value interface{}) *stm.Statement {
	return newStatement(stm.OP_EQUAL, field, value)
}

// Neq equals to "<>" in SQL expressions
func Neq(field string, value interface{}) *stm.Statement {
	return newStatement(stm.OP_NOT_EQUAL, field, value)
}

// Gt equals to ">" in SQL expressions
func Gt(field string, value interface{}) *stm.Statement {
	return newStatement(stm.OP_GREATER, field, value)
}

// Gte equals to ">=" in SQL expressions
func Gte(field string, value interface{}) *stm.Statement {
	return newStatement(stm.OP_GREATER_EQUAL, field, value)
}

// Lt equals to "<" in SQL expressions
func Lt(field string, value interface{}) *stm.Statement {
	return newStatement(stm.OP_LESS, field, value)
}

// Lte equals to "<=" in SQL expressions
func Lte(field string, value interface{}) *stm.Statement {
	return newStatement(stm.OP_LESS_EQUAL, field, value)
}

// In equals to "IN" in SQL expressions
func In(field string, value ...interface{}) *stm.Statement {
	return newStatement(stm.OP_IN, field, value)
}

// Nin equals to "NOT IN" in SQL expressions
func Nin(field string, value ...interface{}) *stm.Statement {
	return newStatement(stm.OP_NOT_IN, field, value)
}

// Lk equals to "LIKE" in SQL expressions
func Lk(field string, value interface{}) *stm.Statement {
	return newStatement(stm.OP_LIKE, field, value)
}

// Nlk equals to "NOT LIKE" in SQL expressions
func Nlk(field string, value interface{}) *stm.Statement {
	return newStatement(stm.OP_NOT_LIKE, field, value)
}

// Null equals to "IS NULL" in SQL expressions
func Null(field string) *stm.Statement {
	return newStatement(stm.OP_NULL, field, "IS NULL")
}

// NotNull equals to "IS NOT NULL" in SQL expressions
func NotNull(field string) *stm.Statement {
	return newStatement(stm.OP_NOT_NULL, field, "IS NOT NULL")
}

// And equals to "AND" in SQL expressions
func And(statements ...*stm.Statement) *stm.Statement {
	return newStatement(stm.OP_AND, "", "", statements...)
}

// Or equals to "OR" in SQL expressions
func Or(statements ...*stm.Statement) *stm.Statement {
	return newStatement(stm.OP_OR, "", "", statements...)
}

func newStatement(op string, field string, value interface{}, children ...*stm.Statement) *stm.Statement {
	return &stm.Statement{
		Operator: op,
		Field:    field,
		Value:    value,
		Children: children,
	}
}

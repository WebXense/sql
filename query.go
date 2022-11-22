package sql

import (
	"github.com/WebXense/sql/stm"
	"gorm.io/gorm"
)

// FindOne finds the first matching record
func FindOne[T any](tx *gorm.DB, statement *stm.Statement) (*T, error) {
	tx = consumeStatement(tx, statement)
	result := new(T)
	err := tx.First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAll finds all the matching records
func FindAll[T any](tx *gorm.DB, statement *stm.Statement, p *Pagination, s *Sort) ([]T, error) {
	tx = consumeStatement(tx, statement)
	tx = consumePagination(tx, p)
	tx = consumeSort(tx, s)
	results := make([]T, 0)
	err := tx.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Count counts the number of matching records
func Count[T any](tx *gorm.DB, statement *stm.Statement) (int64, error) {
	tx = consumeStatement(tx, statement)
	var count int64
	err := tx.Model(new(T)).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func consumeStatement(tx *gorm.DB, statement *stm.Statement) *gorm.DB {
	if statement == nil {
		return tx
	}
	s, v := statement.Build()
	return tx.Where(s, v...)
}

func consumePagination(tx *gorm.DB, p *Pagination) *gorm.DB {
	if p != nil {
		if p.Page > 0 {
			tx = tx.Offset((p.Page - 1) * p.Size)
		}
		if p.Size > 0 {
			tx = tx.Limit(p.Size)
		}
	}
	return tx
}

func consumeSort(tx *gorm.DB, s *Sort) *gorm.DB {
	if s != nil && s.SortBy != "" {
		var order string
		if s.Asc {
			order = "ASC"
		} else {
			order = "DESC"
		}
		tx = tx.Order(s.SortBy + " " + order)
	}
	return tx
}

package models

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

const (
	CreatedAtOrder = "created_at DESC"
	IDOrder        = "id DESC"
	IDOrderAsc     = "id ASC"
	SortAsc        = "sort ASC"
)

const (
	MaxSize     = 100
	DefaultSize = 10
)

func paginate(page, size *int, db *gorm.DB) *gorm.DB {
	if *page == 0 {
		*page = 1
	}

	switch {
	case *size >= MaxSize:
		*size = MaxSize
	case *size <= 0:
		*size = DefaultSize
	}

	offset := (*page - 1) * *size
	return db.Offset(offset).Limit(*size)
}

func GDelete[T any](db *gorm.DB, t *T, query string, args ...any) error {
	return db.Model(t).Where(query, args...).Delete(&t).Error
}

func GFirst[T any](db *gorm.DB, t *T) (*T, error) {
	var result T
	err := db.Where(t).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return &result, nil
}

func GWhereFirst[T any](db *gorm.DB, t *T, query string, args ...any) (*T, error) {
	var result T
	err := db.Model(t).Where(query, args...).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return &result, nil
}

func GWhereLast[T any](db *gorm.DB, t *T, query string, args ...any) (*T, error) {
	var result T
	err := db.Model(t).Where(query, args...).Last(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return &result, nil
}

func quoteColumns(columns string) string {
	subColumns := strings.Split(columns, ",")
	quotedColumns := make([]string, 0, len(subColumns))
	for i := 0; i != len(subColumns); i++ {
		quotedColumns = append(quotedColumns, "`"+strings.TrimSpace(subColumns[i])+"`")
	}
	return strings.Join(quotedColumns, ", ")
}

func GWhereFirstSelect[T any](db *gorm.DB, t *T, columns string, query string, args ...any) (*T, error) {
	var result T
	q := db.Model(t)

	if columns != "" && columns != "*" {
		columns = quoteColumns(columns)
		q = q.Select(columns)
	}
	err := q.Where(query, args...).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return &result, nil
}

func GWhereLastSelect[T any](db *gorm.DB, t *T, columns string, query string, args ...any) (*T, error) {
	var result T
	q := db.Model(t)

	if columns != "" && columns != "*" {
		columns = quoteColumns(columns)
		q = q.Select(columns)
	}
	err := q.Where(query, args...).Last(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return &result, nil
}

func GWhereExist[T any](db *gorm.DB, t *T, query string, args ...any) (bool, error) {
	var count int64
	err := db.Model(t).Where(query, args...).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

// GWhereAllSelect cannot use for tables without "ID" column.
func GWhereAllSelect[T any](db *gorm.DB, columns string, query string, args ...any) ([]T, error) {
	var t T

	return GWhereAllSelectOrder(db, &t, columns, "", query, args...)
}

func GWhereAllSelectOrder[T any](db *gorm.DB, t *T, columns string, order string, query string, args ...any) ([]T, error) {
	var results []T
	q := db.Model(t)

	if order == "" {
		order = IDOrder
	}

	if columns != "" && columns != "*" {
		columns = quoteColumns(columns)
		q = q.Select(columns)
	}
	err := q.Order(order).Where(query, args...).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GPaginate cannot use for tables without "ID" column.
func GPaginate[T any](db *gorm.DB, params *ListPageInput, query string, args ...any) ([]*T, int64, int64, error) {
	var t T
	return GPaginateOrder(db, &t, params, "", query, args...)
}

func GPaginateOrder[T any](db *gorm.DB, t *T, params *ListPageInput, order, query string, args ...any) ([]*T, int64, int64, error) {
	var results []*T
	var count int64

	if order == "" {
		order = CreatedAtOrder
	}
	page := params.Page
	size := params.Size
	q := db.Where(query, args...)

	err := paginate(&page, &size, q).Order(order).Find(&results).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, 0, err
	}

	err = db.Model(&t).Where(query, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, 0, err
	}

	pages := int64(math.Ceil(float64(count) / float64(size)))

	return results, count, pages, nil
}

// GPaginatorSelect cannot use for tables without "ID" column.
func GPaginatorSelect[T any](db *gorm.DB, params *ListPageInput, columns string, query string, args ...any) ([]T, int64, error) {
	var t T
	return GPaginatorSelectOrder(db, &t, params, columns, "", query, args...)
}

func GPaginatorSelectOrder[T any](db *gorm.DB, t *T, params *ListPageInput, columns string, order string, query string, args ...any) ([]T, int64, error) {
	var results []T
	var total int64
	page := params.Page
	size := params.Size

	if order == "" {
		order = IDOrder
	}

	q := db.Model(t).Where(query, args...)

	if columns != "" && columns != "*" {
		columns = quoteColumns(columns)
		q = q.Select(columns)
	}

	err := paginate(&page, &size, q).Order(order).Find(&results).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Model(t).Where(query, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return results, total, nil
}

func GWhereCount[T any](db *gorm.DB, t *T, query string, args ...any) (int64, error) {
	var count int64
	err := db.Model(t).Where(query, args...).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GWhereAllPluckIds cannot use for tables without "ID" column.
func GWhereAllPluckIds[T any](db *gorm.DB, column string, query string, args ...any) ([]uint, error) {
	var t T

	return GWhereAllOrderPluckIds(db, &t, column, "", query, args...)
}

func GWhereAllOrderPluckIds[T any](db *gorm.DB, t *T, column string, order string, query string, args ...any) ([]uint, error) {
	var ids []uint
	if order == "" {
		order = IDOrder
	}
	err := db.Model(t).Order(order).Where(query, args...).Pluck(column, &ids).Error
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func GUpdatesWhere[T any](db *gorm.DB, t *T, values map[string]any, query string, args ...any) (int64, error) {
	result := db.Model(t).Where(query, args...).Updates(values)
	return result.RowsAffected, result.Error
}

func GWhereSum[T any](db *gorm.DB, t *T, column string, query string, args ...any) (int64, error) {
	var result int64
	err := db.Model(&t).
		Where(query, args...).
		Pluck(fmt.Sprintf("IFNULL(SUM(%s), 0) AS result", column), &result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func GWhereJoinGroup[T any, E any](db *gorm.DB, t *T, _ *E, table string, joinTable string, equalColumn string,
	joinEqualColumn string, column string, group string, query string, args ...any) ([]E, error) {
	var results []E
	q := db.Model(t)

	if column != "" && column != "*" {
		q = q.Select(column)
	}

	err := q.Joins(fmt.Sprintf("INNER JOIN %s on %s.%s = %s.%s", joinTable, table, equalColumn, joinTable, joinEqualColumn)).
		Where(query, args...).
		Group(group).
		Scan(&results).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return results, err
	}

	return results, nil
}

func GWhereJoinOrderGroup[T any, E any](db *gorm.DB, t *T, _ *E, table string, joinTable string, equalColumn string,
	joinEqualColumn string, column string, order string, group string, query string, args ...interface{}) ([]E, error) {
	var results []E
	err := db.Model(t).
		Select(column).
		Joins(fmt.Sprintf("INNER JOIN %s on %s.%s = %s.%s", joinTable, table, equalColumn, joinTable, joinEqualColumn)).
		Where(query, args...).
		Order(order).
		Group(group).
		Scan(&results).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return results, err
	}

	return results, nil
}

func GWhereSelectJoin[T any, E any](db *gorm.DB, t *T, _ *E, table string, joinTable string, equalColumn string,
	joinEqualColumn string, column string, query string, args ...any) ([]E, error) {
	var results []E
	q := db.Model(t)

	if column != "" && column != "*" {
		q = q.Select(column)
	}

	err := q.Joins(fmt.Sprintf("INNER JOIN %s on %s.%s = %s.%s", joinTable, table, equalColumn, joinTable, joinEqualColumn)).
		Where(query, args...).
		Scan(&results).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return results, err
	}

	return results, nil
}

func GWhereSelectOrderJoin[T any, E any](db *gorm.DB, t *T, _ *E, table string, joinTable string, equalColumn string,
	joinEqualColumn string, column string, order string, query string, args ...interface{}) ([]E, error) {
	var results []E
	err := db.Model(t).
		Select(column).
		Joins(fmt.Sprintf("INNER JOIN %s on %s.%s = %s.%s", joinTable, table, equalColumn, joinTable, joinEqualColumn)).
		Where(query, args...).
		Order(order).
		Scan(&results).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return results, err
	}

	return results, nil
}

func GWhereSelectGroup[T any, E any](db *gorm.DB, t *T, _ *E, column string, group string, query string, args ...interface{}) ([]E, error) {
	var result []E
	q := db.Model(t)

	if column != "" && column != "*" {
		//column = quoteColumns(column)
		q = q.Select(column)
	}

	err := q.Where(query, args...).Group(group).Scan(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

type uintString struct {
	Uint   uint   `json:"uint" gorm:"column:uint"`
	String string `json:"string" gorm:"column:string"`
}

type uintUint struct {
	UintA uint `json:"uintA" gorm:"column:uintA"`
	UintB uint `json:"uintB" gorm:"column:uintB"`
}

func GUintStringMap[T any](db *gorm.DB, t *T, uintColumn string, strColumn string, query string, args ...any) (map[uint]string, error) {
	var values []uintString
	err := db.Model(t).Select(fmt.Sprintf("%s AS uint, %s AS string", uintColumn, strColumn)).Where(query, args...).Scan(&values).Error
	if err != nil {
		return nil, err
	}

	result := make(map[uint]string, len(values))
	for i := 0; i != len(values); i++ {
		result[values[i].Uint] = values[i].String
	}
	return result, nil
}

func GUintUintMap[T any](db *gorm.DB, t *T, uintAColumn string, uintBColumn string, query string, args ...any) (map[uint]uint, error) {
	var values []uintUint
	err := db.Model(t).Select(fmt.Sprintf("%s AS uintA, %s AS uintB", uintAColumn, uintBColumn)).Where(query, args...).Scan(&values).Error
	if err != nil {
		return nil, err
	}

	result := make(map[uint]uint, len(values))
	for i := 0; i != len(values); i++ {
		result[values[i].UintA] = values[i].UintB
	}

	return result, nil
}

func GUintUintsMap[T any](db *gorm.DB, t *T, uintColumn string, uintsColumn string, query string, args ...any) (map[uint][]uint, error) {
	var values []uintUint
	err := db.Model(t).Select(fmt.Sprintf("%s AS uintA, %s AS uintB", uintColumn, uintsColumn)).Where(query, args...).Scan(&values).Error
	if err != nil {
		return nil, err
	}

	result := make(map[uint][]uint, len(values))
	for i := 0; i != len(values); i++ {
		result[values[i].UintA] = append(result[values[i].UintA], values[i].UintB)
	}

	return result, nil
}

type uintStruct struct {
	Uint uint `json:"uint" gorm:"column:uint"`
}

func GUintSet[T any](db *gorm.DB, t *T, uintColumn string, query string, args ...any) (map[uint]struct{}, error) {
	var values []uintStruct
	err := db.Model(t).Select(fmt.Sprintf("%s as uint", uintColumn)).Where(query, args...).Scan(&values).Error
	if err != nil {
		return nil, err
	}

	result := make(map[uint]struct{}, len(values))
	for i := 0; i != len(values); i++ {
		result[values[i].Uint] = struct{}{}
	}

	return result, nil
}

func GStringUintMap[T any](db *gorm.DB, t *T, strColumn string, uintColumn string, query string, args ...any) (map[string]uint, error) {
	var values []uintString
	err := db.Model(t).Select(fmt.Sprintf("%s AS string, %s AS uint", strColumn, uintColumn)).Where(query, args...).Scan(&values).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string]uint, len(values))
	for i := 0; i != len(values); i++ {
		result[values[i].String] = values[i].Uint
	}

	return result, nil
}

func GInsert[T any](db *gorm.DB, t *T) error {
	return db.Model(t).Create(t).Error
}

func GBatchInsert[T any](db *gorm.DB, t *T, data any) error {
	return db.Model(t).CreateInBatches(data, 1000).Error
}

func GUpdateColumn[T any](db *gorm.DB, t *T, key string, value any, conditions string, args ...any) error {
	err := db.
		Model(&t).
		Where(conditions, args...).
		Update(key, value).
		Error
	if err != nil {
		return err
	}

	return nil
}

// GMaxNewSort cannot use for tables without "sort" column.
func GMaxNewSort[T any](db *gorm.DB, query string, args ...any) (int, error) {
	return GMaxSelectNewSort[T](db, "sort", query, args...)
}

func GMaxSelectNewSort[T any](db *gorm.DB, column string, query string, args ...any) (int, error) {
	var result []int
	var t T

	err := db.
		Model(&t).
		Select(fmt.Sprintf("COALESCE(MAX(%s),0) + 1 as g_select_new_sort", column)).
		Where(query, args...).
		Pluck("g_select_new_sort", &result).Error
	if err != nil {
		return 0, err
	}
	return result[0], nil
}

func GFirstOrCreate[T any](db *gorm.DB, conditions *T) (*T, error) {
	var t T
	err := db.FirstOrCreate(&t, conditions).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GSave[T any](db *gorm.DB, t *T, data interface{}, condition string, args ...interface{}) error {
	return db.Model(t).Where(condition, args...).Save(data).Error
}

func GWhereCountDistinct[T any](db *gorm.DB, t *T, column string, query string, args ...interface{}) (count int64, err error) {
	err = db.Model(t).Select(fmt.Sprintf("COUNT(DISTINCT %s)", column)).Where(query, args...).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GWhereCountGroupDistinct[T any](db *gorm.DB, t *T, group string, distinct string, query string, args ...interface{}) (result map[uint]uint, err error) {
	var values []uintUint
	err = db.Model(t).Select(fmt.Sprintf("%s AS uintA, COUNT(DISTINCT %s) AS uintB", group, distinct)).
		Where(query, args...).
		Group(group).
		Scan(&values).
		Error
	if err != nil {
		return nil, err
	}
	result = make(map[uint]uint, len(values))
	for i := 0; i < len(values); i++ {
		result[values[i].UintA] = values[i].UintB
	}
	return result, nil
}

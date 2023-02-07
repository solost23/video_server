package models

import "gorm.io/gorm"

// 抛弃每个model携带相同的数据库sql， 采用泛型统一sql
func GInsert[T any](db *gorm.DB, t *T) error {
	return db.Model(t).Create(t).Error
}

func GWhereDelete[T any](db *gorm.DB, t *T, conditions string, args ...interface{}) error {
	return db.Model(t).Where(conditions, args...).Error
}

func GWhereUpdates[T any](db *gorm.DB, t *T, values map[string]interface{}, conditions string, args ...interface{}) error {
	return db.Model(t).Where(conditions, args...).Updates(values).Error
}

func GWhereOne[T any](db *gorm.DB, t *T, query string, args ...interface{}) (result *T, err error) {
	err = db.Model(t).Where(query, args...).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GWhereAll[T any](db *gorm.DB, t *T, query string, args ...interface{}) (results []*T, err error) {
	err = db.Model(t).Where(query, args...).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GWherePageListOrder[T any](db *gorm.DB, t *T, order string, params *ListPageInput, conditions string, args ...interface{}) (results []*T, count int64, err error) {
	if order == "" {
		order = "created_at DESC"
	}
	offset := (params.Page - 1) * params.Size

	err = db.Model(t).Where(conditions, args...).Offset(offset).Limit(params.Size).Order(order).Find(&results).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Model(t).Where(conditions, args...).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return results, count, nil
}

func GWhereAllOrderPluckIds[T any](db *gorm.DB, t *T, column string, order string, query string, args ...any) ([]uint, error) {
	var ids []uint
	if order == "" {
		order = "id DESC"
	}
	err := db.Model(t).Order(order).Where(query, args...).Pluck(column, &ids).Error
	if err != nil {
		return nil, err
	}

	return ids, nil
}

package db

import "gorm.io/gorm"

type ORM interface {
	Close() error
	Error() error
	Model(value interface{}) ORM

	Create(args interface{}) error
	Update(args interface{}) error
	Delete(model interface{}, args ...interface{}) error
	Find(object interface{}) error
	Where(query interface{}, args ...interface{}) ORM
	Order(value interface{}) ORM

	Joins(assoc string) ORM
	Raw(query string, args ...interface{}) ORM
	Scan(object interface{}) error
	GetGormInstance() *gorm.DB
	Preload(assoc string) ORM

	Count(count *int64) error
}

type ORMAssociation interface {
	Replace(values ...interface{}) error
	Find(out interface{}, conds ...interface{}) error
	Clear() error
}

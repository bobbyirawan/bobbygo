package db

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	mysqldb struct {
		db  *gorm.DB
		err error
	}

	MySqlOption struct {
		ConnectionString                     string
		MaxLifeTimeConnection                time.Duration
		MaxIdleConnection, MaxOpenConnection int
		Logger                               log.Logger
	}
)

func (d *mysqldb) Error() error {
	return d.err
}

func (d *mysqldb) Close() error {
	sql, err := d.db.DB()

	if err != nil {
		return err
	}

	if err := sql.Close(); err != nil {
		return err
	}
	return nil
}

func (d *mysqldb) Begin() ORM {
	var (
		db  = d.db.Begin()
		err = db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) Commit() error {
	return d.db.Commit().Error
}

func (d *mysqldb) Rollback() error {
	return d.db.Rollback().Error
}

func (d *mysqldb) Offset(offset int64) ORM {
	var (
		db  = d.db.Offset(int(offset))
		err = d.db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) Limit(limit int64) ORM {
	var (
		db  = d.db.Limit(int(limit))
		err = d.db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) First(object interface{}) error {
	var (
		res = d.db.First(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *mysqldb) FirstWithLocking(object interface{}) error {
	var (
		res = d.db.Clauses(clause.Locking{Strength: "UPDATE"}).First(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *mysqldb) Last(object interface{}) error {
	var (
		res = d.db.Last(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *mysqldb) LastWithLocking(object interface{}) error {
	var (
		res = d.db.Clauses(clause.Locking{Strength: "UPDATE"}).Last(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *mysqldb) Find(object interface{}) error {
	var (
		res = d.db.Find(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *mysqldb) Take(object interface{}) error {
	var (
		res = d.db.Take(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *mysqldb) TakeWithLocking(object interface{}) error {
	var (
		res = d.db.Clauses(clause.Locking{Strength: "UPDATE"}).Take(object)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *mysqldb) Model(value interface{}) ORM {
	var (
		db  = d.db.Model(value)
		err = db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) OmitAssoc() ORM {
	var (
		db  = d.db.Omit(clause.Associations)
		err = db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) Select(query interface{}, args ...interface{}) ORM {
	var (
		db  = d.db.Select(query, args...)
		err = db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) Where(query interface{}, args ...interface{}) ORM {
	var (
		db  = d.db.Where(query, args...)
		err = db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) Order(value interface{}) ORM {
	var (
		db  = d.db.Order(value)
		err = d.db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) Create(args interface{}) error {
	return d.db.Create(args).Error
}

func (d *mysqldb) Update(args interface{}) error {
	return d.db.Updates(args).Error
}

func (d *mysqldb) UpdateColumns(args interface{}) error {
	return d.db.UpdateColumns(args).Error
}

func (d *mysqldb) Delete(model interface{}, args ...interface{}) error {
	return d.db.Delete(model, args...).Error
}

func (d *mysqldb) WithContext(ctx context.Context) ORM {
	var (
		db = d.db.WithContext(ctx)
	)

	return &mysqldb{db: db, err: nil}
}

func (d *mysqldb) Raw(query string, args ...interface{}) ORM {
	var (
		db  = d.db.Raw(query, args...)
		err = db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) Exec(query string, args ...interface{}) ORM {
	var (
		db  = d.db.Exec(query, args...)
		err = db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) Scan(object interface{}) error {
	var (
		db = d.db.Scan(object)
	)

	return db.Error
}

func (d *mysqldb) Preload(assoc string) ORM {
	var (
		db  = d.db.Preload(assoc)
		err = db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) Joins(assoc string) ORM {
	var (
		db  = d.db.Joins(assoc)
		err = db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) GetGormInstance() *gorm.DB {
	return d.db
}

func (d *mysqldb) Count(count *int64) error {
	var (
		res = d.db.Count(count)
	)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *mysqldb) Association(column string) ORMAssociation {
	return d.db.Association(column)
}

func (d *mysqldb) Or(query interface{}, args ...interface{}) ORM {
	var (
		db  = d.db.Or(query, args...)
		err = db.Error
	)

	return &mysqldb{db, err}
}

func (d *mysqldb) Save(data interface{}) error {
	return d.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(data).Error
}

func NewMySql(option *MySqlOption) (ORM, error) {
	var (
		opts = &gorm.Config{}
	)

	// if option.Logger != nil {
	// 	opts.Logger = logger.New(option.Logger, logger.Config{
	// 		SlowThreshold:             time.Second,
	// 		LogLevel:                  logger.Info,
	// 		Colorful:                  false,
	// 		IgnoreRecordNotFoundError: false,
	// 	})
	// }

	db, err := gorm.Open(mysql.Open(option.ConnectionString), opts)

	if err != nil {
		return nil, err
	}

	sql, err := db.DB()

	if err != nil {
		return nil, err
	}

	sql.SetConnMaxLifetime(option.MaxLifeTimeConnection)
	sql.SetMaxOpenConns(option.MaxOpenConnection)
	sql.SetMaxIdleConns(option.MaxIdleConnection)

	return &mysqldb{db: db}, nil
}

package db

// import (
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type (
// 	postgresdb struct {
// 		db  *gorm.DB
// 		err error
// 	}

// 	PostgreSqlOption struct {
// 		ConnectionString                     string
// 		MaxLifeTimeConnection                time.Duration
// 		MaxIdleConnection, MaxOpenConnection int
// 		Logger                               log.Logger
// 	}
// )

// func (d *postgresdb) Error() error {
// 	return d.err
// }

// func (d *postgresdb) Close() error {
// 	sql, err := d.db.DB()

// 	if err != nil {
// 		return err
// 	}

// 	if err := sql.Close(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (d *postgresdb) Find(object interface{}) error {
// 	var (
// 		res = d.db.Find(object)
// 	)

// 	if res.Error != nil {
// 		return res.Error
// 	}

// 	return nil
// }

// func (d *postgresdb) Create(args interface{}) error {
// 	return d.db.Create(args).Error
// }

// func (d *postgresdb) Update(args interface{}) error {
// 	return d.db.Updates(args).Error
// }

// func (d *postgresdb) Delete(model interface{}, args ...interface{}) error {
// 	return d.db.Delete(model, args...).Error
// }

// func (d *postgresdb) Where(query interface{}, args ...interface{}) ORM {
// 	var (
// 		db  = d.db.Where(query, args...)
// 		err = db.Error
// 	)
// 	return &postgresdb{db, err}
// }

// func (d *postgresdb) GetGormInstance() *gorm.DB {
// 	return d.db
// }

// func NewPostgreSql(option *PostgreSqlOption) (ORM, error) {
// 	var (
// 		opts = &gorm.Config{}
// 	)

// 	// if option.Logger != nil {
// 	// 	opts.Logger = logger.New(option.Logger, logger.Config{
// 	// 		SlowThreshold:             time.Second,
// 	// 		LogLevel:                  logger.Info,
// 	// 		Colorful:                  false,
// 	// 		IgnoreRecordNotFoundError: false,
// 	// 	})
// 	// }

// 	db, err := gorm.Open(postgres.Open(option.ConnectionString), opts)

// 	if err != nil {
// 		return nil, err
// 	}

// 	sql, err := db.DB()

// 	if err != nil {
// 		return nil, err
// 	}

// 	sql.SetConnMaxLifetime(option.MaxLifeTimeConnection)
// 	sql.SetMaxOpenConns(option.MaxOpenConnection)
// 	sql.SetMaxIdleConns(option.MaxIdleConnection)

// 	return &postgresdb{db: db}, nil
// }

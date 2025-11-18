package database

import (
	"context"
	"fmt"
	"time"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type PostgresInfra struct {
	DB *gorm.DB
}

func NewPostgresInfra(ctx context.Context) (*PostgresInfra, error) {

	dbWriteDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
		config.GlobalEnv.DBWriteUser,
		config.GlobalEnv.DBWritePass,
		config.GlobalEnv.DBWriteHost,
		config.GlobalEnv.DBWritePort,
		config.GlobalEnv.DBWriteName,
		config.GlobalEnv.DBWriteSchema,
	)

	dbReadDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
		config.GlobalEnv.DBReadUser,
		config.GlobalEnv.DBReadPass,
		config.GlobalEnv.DBReadHost,
		config.GlobalEnv.DBReadPort,
		config.GlobalEnv.DBReadName,
		config.GlobalEnv.DBReadSchema,
	)

	db, err := gorm.Open(postgres.Open(dbWriteDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{postgres.Open(dbWriteDSN)},
		Replicas: []gorm.Dialector{postgres.Open(dbReadDSN)},
		Policy:   dbresolver.RandomPolicy{},
	}))
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresInfra{
		DB: db,
	}, nil
}

func (p *PostgresInfra) Shutdown(ctx context.Context) error {
	if p.DB != nil {
		sqlDB, _ := p.DB.DB()
		sqlDB.Close()
	}

	return nil
}


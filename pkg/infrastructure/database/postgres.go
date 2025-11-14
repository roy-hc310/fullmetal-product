package database

import (
	"context"
	"fmt"
	"time"

	"github.com/roy-hc310/fullmetal-product/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresInfra struct {
	DBWrite *gorm.DB
	DBRead  *gorm.DB
}

func NewPostgresInfra(ctx context.Context) (*PostgresInfra, error) {

	dbWriteURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
		config.GlobalEnv.DBWriteUser,
		config.GlobalEnv.DBWritePass,
		config.GlobalEnv.DBWriteHost,
		config.GlobalEnv.DBWritePort,
		config.GlobalEnv.DBWriteName,
		config.GlobalEnv.DBWriteSchema,
	)

	dbWrite, err := gorm.Open(postgres.Open(dbWriteURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDBWrite, err := dbWrite.DB()
	if err != nil {
		return nil, err
	}
	sqlDBWrite.SetMaxOpenConns(10)
	sqlDBWrite.SetMaxIdleConns(5)
	sqlDBWrite.SetConnMaxLifetime(time.Hour)
	sqlDBWrite.SetConnMaxIdleTime(30 * time.Minute)

	err = sqlDBWrite.Ping()
	if err != nil {
		return nil, err
	}

	//------------------ Read DB ------------------//

	dbReadURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
		config.GlobalEnv.DBReadUser,
		config.GlobalEnv.DBReadPass,
		config.GlobalEnv.DBReadHost,
		config.GlobalEnv.DBReadPort,
		config.GlobalEnv.DBReadName,
		config.GlobalEnv.DBReadSchema,
	)

	dbRead, err := gorm.Open(postgres.Open(dbReadURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDBRead, err := dbRead.DB()
	if err != nil {
		return nil, err
	}
	sqlDBRead.SetMaxOpenConns(10)
	sqlDBRead.SetMaxIdleConns(5)
	sqlDBRead.SetConnMaxLifetime(time.Hour)
	sqlDBRead.SetConnMaxIdleTime(30 * time.Minute)

	err = sqlDBRead.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresInfra{
		DBWrite: dbWrite,
		DBRead:  dbRead,
	}, nil
}

func (p *PostgresInfra) Shutdown(ctx context.Context) error {
	if p.DBWrite != nil {
		sqlDB, _ := p.DBWrite.DB()
		sqlDB.Close()
	}
	if p.DBRead != nil {
		sqlDB, _ := p.DBRead.DB()
		sqlDB.Close()
	}

	return nil
}

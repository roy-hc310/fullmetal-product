package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
)

type PostgresInfra struct {
	DBWrite *pgxpool.Pool
	DBRead  *pgxpool.Pool
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

	dbWriteConfig, err := pgxpool.ParseConfig(dbWriteURL)
	if err != nil {
		return nil, err
	}
	dbWriteConfig.MaxConns = 10
	dbWriteConfig.MinConns = 0
	dbWriteConfig.MaxConnLifetime = time.Hour
	dbWriteConfig.MaxConnIdleTime = time.Minute * 30
	dbWriteConfig.HealthCheckPeriod = time.Minute * 5

	dbWrite, err := pgxpool.NewWithConfig(context.Background(), dbWriteConfig)
	if err != nil {
		return nil, err
	}

	if err := dbWrite.Ping(ctx); err != nil {
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

	dbReadConfig, err := pgxpool.ParseConfig(dbReadURL)
	if err != nil {
		return nil, err
	}
	dbReadConfig.MaxConns = 10
	dbReadConfig.MinConns = 0
	dbReadConfig.MaxConnLifetime = time.Hour
	dbReadConfig.MaxConnIdleTime = time.Minute * 30
	dbReadConfig.HealthCheckPeriod = time.Minute * 5

	dbRead, err := pgxpool.NewWithConfig(context.Background(), dbReadConfig)
	if err != nil {
		return nil, err
	}

	if err := dbRead.Ping(ctx); err != nil {
		return nil, err
	}

	return &PostgresInfra{
		DBWrite: dbWrite,
		DBRead:  dbRead,
	}, nil
}

func (p *PostgresInfra) Shutdown(ctx context.Context) error {
	if p.DBWrite != nil {
		p.DBWrite.Close()
	}
	if p.DBRead != nil {
		p.DBRead.Close()
	}

	return nil
}

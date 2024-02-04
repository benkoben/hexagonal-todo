package postgres

import (
	"fmt"
	"time"
    "context"

	sq "github.com/Masterminds/squirrel"
    "github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// default values
var (
	defaultHost                  = "localhost"
	defaultPort           uint16 = 5432
	defaultAquireTimeout               = time.Second * 30
	defaultMaxConnections        = 5
    defaultSSLMode               = "disable"
)

type PostgresClientOptions struct {
	// Username to authenticate with
	Username string
	// Password for username to authenticate with
	Password string
	// Hosts where postgres server is running. Defaults to localhost
	Host string
	// Port on which postgres is running on
	Port uint16
	// Name of the database
	Database string
	// Max wait time when all connetions are busy
	AcquireTimeout time.Duration
	// Max simultaneous connections to use, defaults to 5, must be at least 2
	MaxConnections int
    // ssl mode enable or disable
    SSLMode string
}

// Wrap Connpool and squirrel statement builder
type DB struct {
	*pgxpool.Pool
	QueryBuilder *sq.StatementBuilderType
}

// Construct DB
func NewDB(ctx context.Context, opt *PostgresClientOptions) (*DB, error) {
	if opt.Host == "" {
		opt.Host = defaultHost
	}

	if opt.Port == 0 {
		opt.Port = defaultPort
	}

	if opt.MaxConnections == 0 {
		opt.MaxConnections = defaultMaxConnections
	}

	if opt.AcquireTimeout == 0 {
		opt.AcquireTimeout = defaultAquireTimeout
	}

	if opt.Username == "" {
		return &DB{}, fmt.Errorf("username cannot be empty")
	}

	if opt.Password == "" {
		return &DB{}, fmt.Errorf("password cannot be empty")
	}

	if opt.Database == "" {
		return &DB{}, fmt.Errorf("database name cannot be empty")
	}

    if opt.SSLMode == "" {
        opt.SSLMode = defaultSSLMode
    }

    // postgres://postgres:123456@127.0.0.1:5432/dummy
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		opt.Username,
        opt.Password,
        opt.Host,
        opt.Port,
        opt.Database,
        opt.SSLMode,
	)

    fmt.Println(url)
    
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	db, err := pgxpool.New(ctx, url)

	if err != nil {
		return nil, fmt.Errorf("could not create database connpool: %v", err)
	}
    
    err = db.Ping(ctx)
    if err != nil {
        return nil, fmt.Errorf("unsuccessful ping to database")
    }

	return &DB{
        db,
		&psql,
	}, nil
}

// ErrorCode returns the error code of the given error
func (db *DB) ErrorCode(err error) string {
	pgErr := err.(*pgconn.PgError)
	return pgErr.Code
}

package postgres

import (
	"fmt"
	"time"
    "context"

	"crypto/tls"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

// default values
var (
	defaultHost                  = "localhost"
	defaultPort           uint16 = 5432
	defaultAquireTimeout               = time.Second * 30
	defaultMaxConnections        = 5
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
	// tls.Config. TLS is disabled if nil
	TLSConfig *tls.Config
	// Name of the database
	Database string
	// Max wait time when all connetions are busy
	AcquireTimeout time.Duration
	// Max simultaneous connections to use, defaults to 5, must be at least 2
	MaxConnections int
}

// Wrap Connpool and squirrel statement builder
type DB struct {
	*pgx.ConnPool
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

	cfg := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Port:      opt.Port,
			Database:  opt.Database,
			User:      opt.Username,
			Password:  opt.Password,
			TLSConfig: opt.TLSConfig,
		},
		MaxConnections: opt.MaxConnections,
		AcquireTimeout: opt.AcquireTimeout,
	}
    
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	connpool, err := pgx.NewConnPool(cfg)

	if err != nil {
		return nil, fmt.Errorf("could not create database connpool: %v", err)
	}

    // Acquire an explicit connection to database 
    conn, err := connpool.Acquire()
    if err != nil {
        return nil, fmt.Errorf("could not aquire database connection")
    }
    
    // Ping data base with aquired connection to verifiy reachability
    err = conn.Ping(ctx)
    if err != nil {
        return nil, fmt.Errorf("unsuccessful ping to database")
    }

	return &DB{
        ConnPool: connpool,
		QueryBuilder: &psql,
	}, nil
}

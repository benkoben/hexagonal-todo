package postgres

import (
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

// default values
var (
    defaultHost = "localhost"
    defaultPort = "5432"
    defaultTimeout = time.Second * 30
    defaultMaxConnections = 5
)

type PostgresClientOptions struct {
    // Username to authenticate with
	Username string 
    // Password for username to authenticate with
	Password string 
    // Hosts where postgres server is running. Defaults to localhost
    Host string
    // tls.Config. TLS is disabled if nil
    TLSConfig *tls.Config 
    // Name of the database 
    Database string
    // Max wait time when all connetions are busy
    Timeout time.Duration
    // Max simultaneous connections to use, defaults to 5, must be at least 2
    MaxConnections
}

type DB struct {
    *pgx.ConnPool
    QueryBuilder *sq.StatementBuilderType
}

func NewDB(opt *PostgresClientOptions) (*DB, error) {
   if opt.Host == "" {
        opt.Host = defaultHost
   } 

   if opt.Port == "" {
        opt.Port = defaultPort
   }

   if opt.MaxConnections == 0 {
        opt.MaxConnections = defaultMaxConnections
   }

   if opt.Timeout == 0 {
        opt.Timeout = defaultTimeout
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

   cfg := pgx.ConnConfig{
        Port: opt.Port,
        Database: opt.Database,
        Username: opt.Username,
        Password: opt.Password,
        TLSConfig: opt.TLSconfig,    
   }

   connpool, error := pgx.NewConnPool(pgx.ConnPoolConfig{
        cfg,
        AcquireTimeout: opt.Timeout,
   }) 

   return &{
       
   }
}

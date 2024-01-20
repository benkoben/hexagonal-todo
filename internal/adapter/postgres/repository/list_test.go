package repository

// TODO: Find out how to test the behaviour of NewDB. Which should return a DB struct
//      In different scenarios such as. Only defaults are used, optionals are used. Required fields are missing.

import (
	"testing"
	"time"
    "context"

	"github.com/benkoben/hexagonal-todo/internal/adapter/postgres"

    "github.com/jackc/pgx"
    "github.com/jackc/pgconn"

    "github.com/google/go-cmp/cmp"
)

type MockDB interface{
    Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type FakePostgres struct {}

func TestNewDB(t *testing.T) {
    
    // I want to test if the defaults work
    connTests := []struct{
       options postgres.PostgresClientOptions
       want MockDB
       wantErr error
    }{
        {
            options: postgres.PostgresClientOptions{
                Username: "pgadmin",
                Password: "Syp9393",
                Database: "postgres",
            },
            want: pgx.ConnPool{},
            wantErr: nil,
        },
    }
    
    for _, test := range connTests {
        ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(test.options.Timeout))
        defer cancel()

        got, gotErr := postgres.NewDB(ctx, &test.options)
        
        if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("NewDB(%q) = unexpected results, (-want, +got)\n%s\n", test.options, diff)
		}

        if gotErr != nil && test.wantErr == nil {
           t.Errorf("NewDB(%q) unexpected error", test.options) 
        }
    }

}

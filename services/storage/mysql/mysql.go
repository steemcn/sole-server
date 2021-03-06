package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // is needed for mysql driver registeration
	"github.com/jmoiron/sqlx"
	"github.com/solefaucet/sole-server/services/storage"
)

// Storage implements Storage interface for data storage
type Storage struct {
	db *sqlx.DB
}

var _ storage.Storage = Storage{}

// New returns a Storage with data source name
func New(dsn string) Storage {
	return Storage{
		db: sqlx.MustConnect("mysql", dsn).Unsafe(),
	}
}

// SetMaxOpenConns alias sql.DB.SetMaxOpenConns
func (s *Storage) SetMaxOpenConns(n int) {
	s.db.SetMaxOpenConns(n)
}

// SetMaxIdleConns alias sql.DB.SetMaxIdleConns
func (s *Storage) SetMaxIdleConns(n int) {
	s.db.SetMaxIdleConns(n)
}

// mysql error codes
const (
	errcodeDuplicate = 1062
)

func (s Storage) selects(dest interface{}, rawSQL string, args ...interface{}) error {
	if err := s.db.Select(dest, rawSQL, args...); err != nil {
		return fmt.Errorf("query %v error: %v", rawSQL, err)
	}

	return nil
}

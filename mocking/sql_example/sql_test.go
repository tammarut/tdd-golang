package sql_example_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/tammarut/tdd-golang/mocking/sql_example"
)

func TestOpenDB(t *testing.T) {
	mockError := errors.New("Fake Error")
	subtests := []struct {
		name        string
		u, p, a, db string
		sqlOpener   func(string, string) (*sql.DB, error)
		expectedErr error
	}{
		{
			name: "happy case",
			u:    "u",
			p:    "p",
			a:    "ad",
			db:   "mydb",
			sqlOpener: func(driver, source string) (db *sql.DB, err error) {
				if source != "u:p@ad/mydb" {
					return nil, errors.New("Wrong connection string")
				}
				return nil, nil
			},
		},
		{
			name: "failed case",
			sqlOpener: func(driver, source string) (db *sql.DB, err error) {
				return nil, mockError
			},
			expectedErr: mockError,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			_, err := sql_example.OpenDB(subtest.u, subtest.p, subtest.a, subtest.db, subtest.sqlOpener)
			if !errors.Is(err, subtest.expectedErr) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.expectedErr, err)
			}
		})
	}
}

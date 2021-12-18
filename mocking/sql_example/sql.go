package sql_example

import (
	"database/sql"
	"fmt"
)

type sqlOpener func(string, string) (*sql.DB, error)

/*
➜ I need to mock a function (Higher Order Functions)
✅ Good:
    - Clear and proximal to function under test
    - Stateless
❌ Bad:
    - Parameter list can get ugly (using a type for the function can help)
    - Think of dependency graph
*/
func OpenDB(user, password, address, db string, opener sqlOpener) (*sql.DB, error) {
	connection := fmt.Sprintf("%s:%s@%s/%s", user, password, address, db)

	return opener("mysql", connection)
}

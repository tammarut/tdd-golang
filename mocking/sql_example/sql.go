package sql_example

import (
	"database/sql"
	"fmt"
)

// type sqlOpener func(string, string) (*sql.DB, error)

/*
➜ I need to mock a function (Higher Order Functions)
✅ Good:
    - Clear and proximal to function under test
    - Stateless
❌ Bad:
    - Parameter list can get ugly (using a type for the function can help)
    - Think of dependency graph
*/
// func OpenDB(user, password, address, db string, opener sqlOpener) (*sql.DB, error) {
// 	connection := fmt.Sprintf("%s:%s@%s/%s", user, password, address, db)

// 	return opener("mysql", connection)
// }

/*
➜ I need to mock a function (Monkey Patching)
✅ Good:
	- Don't need to modify function signature
	- Don't need to grow that dependency graph into other packages
❌ Bad:
	- Allergic to parallelism (stateful)
	- Have to make variable public (if testing from file_test package)
*/
var SQLOpen = sql.Open

func OpenDB(user, password, address, db string) (*sql.DB, error) {
	connection := fmt.Sprintf("%s:%s@%s/%s", user, password, address, db)
	return SQLOpen("mysql", connection)
}

// Package driver provides a constant driver name for SQL usage, and imports the needed PostgreSQL driver.
package driver

import _ "github.com/jackc/pgx/v4/stdlib"

// Driver is the name of the PostgreSQL database driver used by Jaunty.
const Driver = "pgx"

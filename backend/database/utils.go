package database

import "strings"

func NoRowsFound(err error) bool {
	return strings.Contains(err.Error(), "no rows in result set")
}

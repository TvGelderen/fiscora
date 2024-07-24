package database

import "strings"

const DefaultFetchLimit = 25

func NoRowsFound(err error) bool {
	return strings.Contains(err.Error(), "no rows in result set")
}

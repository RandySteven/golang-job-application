package utils

import "fmt"

func SelectAllFrom(table string) string {
	return fmt.Sprintf("SELECT * FROM %s", table)
}

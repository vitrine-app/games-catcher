package main

import "database/sql"

func validateString(nullableString sql.NullString) string {
	if nullableString.Valid {
		return nullableString.String
	}
	return ""
}

func validateNumber(nullableNumber sql.NullInt64) int64 {
	if nullableNumber.Valid {
		return nullableNumber.Int64
	}
	return -1
}

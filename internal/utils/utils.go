package utils

import "database/sql"

func NullToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func NullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

package util

import (
	"database/sql"
	"strconv"
)

func ConvertNullString(nullStr sql.NullString) string {
	if nullStr.Valid {
		return nullStr.String
	}
	return ""
}

// Helper function to convert nullable time to Unix timestamp
func ConvertNullTime(nullTime sql.NullTime) int64 {
	if nullTime.Valid {
		return nullTime.Time.Unix()
	}
	return 0
}

// Helper function to convert nullable float (as string) to float32
func ConvertNullFloat(nullStr sql.NullString) float32 {
	if nullStr.Valid {
		// Assuming nullStr.String holds a valid float value
		val, err := strconv.ParseFloat(nullStr.String, 32)
		if err == nil {
			return float32(val)
		}
	}
	return 0.0
}

// Helper function to convert nullable int32 to int32
func ConvertNullInt32(nullInt sql.NullInt32) int32 {
	if nullInt.Valid {
		return nullInt.Int32
	}
	return 0
}

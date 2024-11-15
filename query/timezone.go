package query

import (
	"context"
	"strings"
	"time"
)

var (
	defaultFormat = "2006-01-02 15:04"
)

// GetTimezoneArea ...
// Bali, Lombok, Manado is GMT+8
func GetTimezoneArea(areaCode string) int32 {
	areaCode = strings.ToUpper(strings.TrimSpace(areaCode))
	switch areaCode {
	case "DPS", "AMI", "MDC":
		return int32(8)
	default:
		return int32(7)
	}
}

func GetTimezone(areaCode string) string {
	areaCode = strings.ToUpper(strings.TrimSpace(areaCode))
	//use Asia/Jayapura to WIT
	switch areaCode {
	case "DPS", "AMI", "MDC":
		return "Asia/Makassar"
	default:
		return "Asia/Jakarta"
	}
}

func TimeLoadLoc(t time.Time, tz string) time.Time {
	loc, _ := time.LoadLocation(tz)
	return t.In(loc)
}

func GetGenerateDateRangeWithTz(ctx context.Context, mustArr []map[string]interface{}, startTime, endTime string, field string) ([]map[string]interface{}, error) {
	if endTime != "" && startTime != "" {
		tz := "Asia/Jakarta"

		st, err := ParseStartDateTimeTz(startTime, tz)
		if err != nil {
			return nil, err
		}

		et, err := ParseEndDateTimeTz(endTime, tz)
		if err != nil {
			return nil, err
		}

		mustArr = GenerateRange(mustArr, map[string]interface{}{
			field: map[string]interface{}{
				"gte": st,
				"lte": et,
			},
		})
	} else if startTime != "" {
		tz := "Asia/Jakarta"

		st, err := ParseStartDateTz(startTime, tz)
		if err != nil {
			return nil, err
		}

		mustArr = GenerateRange(mustArr, map[string]interface{}{
			field: map[string]interface{}{
				"gte": st,
			},
		})
	} else if endTime != "" {
		tz := "Asia/Jakarta"

		et, err := ParseEndDateTz(endTime, tz)
		if err != nil {
			return nil, err
		}

		mustArr = GenerateRange(mustArr, map[string]interface{}{
			field: map[string]interface{}{
				"lte": et,
			},
		})
	}
	return mustArr, nil
}

package query

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func GenerateMatch(must []map[string]interface{}, query map[string]interface{}) []map[string]interface{} {
	match := map[string]interface{}{
		"match": query,
	}
	return append(must, match)
}

// Equal to SELECT * FROM X WHERE X IN
func GenerateTerms(must []map[string]interface{}, terms map[string]interface{}) []map[string]interface{} {
	return append(must, map[string]interface{}{
		"terms": terms,
	})
}

func GenerateWildcard(must []map[string]interface{}, query map[string]interface{}) []map[string]interface{} {
	match := map[string]interface{}{
		"wildcard": query,
	}
	return append(must, match)
}

func SplitIDs(input string) []string {
	return strings.Split(input, ",")
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}


func GenerateValueWildcard(text *string) *string {
	if text == nil {
		return nil
	}
	t := fmt.Sprintf("*%s*", StringValue(text))
	return &t
}

func GenerateMultiMatch(must []map[string]interface{}, query string, fields []string) []map[string]interface{} {
	match := map[string]interface{}{
		"multi_match": map[string]interface{}{
			"query":  query,
			"fields": fields,
		},
	}
	return append(must, match)
}

func GenerateRange(must []map[string]interface{}, query map[string]interface{}) []map[string]interface{} {
	rangeQuery := map[string]interface{}{
		"range": query,
	}
	return append(must, rangeQuery)
}

func GenerateShould(must []map[string]interface{}, mustShould []map[string]interface{}) []map[string]interface{} {
	booleanQ := ESBoolShould{Should: mustShould}
	boolean := map[string]interface{}{
		"bool": booleanQ,
	}
	return append(must, boolean)
}

func GenerateUpdateScript(field string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		"source": fmt.Sprintf("ctx._source.%s = params.value", field),
		"lang":   "painless",
		"params": map[string]interface{}{
			"value": value,
		},
	}
}

func ParseStartDate(d string) (string, error) {
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return "", err
	}
	return BeginningOfDay(t).Format("2006-01-02T15:04:05.999Z"), nil
}

func ParseEndDate(d string) (string, error) {
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return "", err
	}
	return EndOfDay(t).Format("2006-01-02T15:04:05.999Z"), nil
}

func SetTimeOnDate(t time.Time, hour, min, sec, nSec int) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, hour, min, sec, nSec, t.Location())
}

func BeginningOfDay(t time.Time) time.Time {
	return SetTimeOnDate(t, 0, 0, 0, 0)
}

func EndOfDay(t time.Time) time.Time {
	return SetTimeOnDate(t, 23, 59, 59, 999999999)
}


// LogErrFormat ...
func LogErrFormat(funcName, reqId string, err error) string {
	return fmt.Sprintf("%s [request_id]: %s [error]: %+v", funcName, reqId, err)
}

// LogReqFormat ...
func LogReqFormat(funcName, reqId, request string) string {
	return fmt.Sprintf("%s [request_id]: %s [request]: %s", funcName, reqId, request)
}

// LogResFormat ...
func LogResFormat(funcName, reqId, response string) string {
	return fmt.Sprintf("%s [request_id]: %s [response]: %s", funcName, reqId, response)
}

// GenerateSort ...
func GenerateSort(sortField map[string]string) []map[string]interface{} {
	var sortArr []map[string]interface{}

	for k, v := range sortField {
		sort := map[string]interface{}{
			k: map[string]interface{}{
				"order": v,
			},
		}
		sortArr = append(sortArr, sort)
	}
	return sortArr
}

func GenerateSumAggregate(aggField string, field string) map[string]interface{} {
	return map[string]interface{}{
		aggField: map[string]interface{}{
			"sum": map[string]interface{}{
				"field": field,
			},
		},
	}
}

func GenerateHideMenuFilter(hideMenu string) map[string]interface{} {
	mustNot := map[string]interface{}{
		"term": map[string]interface{}{
			"hide_menu": hideMenu,
		},
	}
	return mustNot
}

func GetContextString(ctx context.Context, ctxKey interface{}) string {
	val := ctx.Value(ctxKey)
	if val != nil {
		str, ok := val.(string)
		if ok {
			return str
		}
	}
	return ""
}



func ParseStartDateTz(d string, tz string) (string, error) {
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return "", err
	}

	return BeginningOfDayTz(t, tz).UTC().Format("2006-01-02T15:04:05.999Z"), nil
}

func ParseEndDateTz(d string, tz string) (string, error) {
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return "", err
	}
	return EndOfDayTz(t, tz).UTC().Format("2006-01-02T15:04:05.999Z"), nil
}

func ParseStartDateTimeTz(d string, tz string) (string, error) {
	t, err := time.Parse(defaultFormat, d)
	if err != nil {
		return "", err
	}

	return BeginningOfDayTz(t, tz).UTC().Format("2006-01-02T15:04:05.999Z"), nil
}

func ParseEndDateTimeTz(d string, tz string) (string, error) {
	t, err := time.Parse(defaultFormat, d)
	if err != nil {
		return "", err
	}
	return EndOfDayTz(t, tz).UTC().Format("2006-01-02T15:04:05.999Z"), nil
}

func BeginningOfDayTz(t time.Time, tz string) time.Time {
	return SetTimeOnDateTz(t, 0, 0, 0, 0, tz)
}

func EndOfDayTz(t time.Time, tz string) time.Time {
	return SetTimeOnDateTz(t, 23, 59, 59, 999999999, tz)
}

func SetTimeOnDateTz(t time.Time, hour, min, sec, nSec int, tz string) time.Time {
	year, month, day := t.Date()
	loc, _ := time.LoadLocation(tz)
	return time.Date(year, month, day, hour, min, sec, nSec, loc)
}

func ConvertTimeToRFC3339(waktu time.Time) string {
	return waktu.Format(time.RFC3339)
}





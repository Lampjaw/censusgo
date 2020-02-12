// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import (
	"fmt"
	"reflect"
	"strings"
)

const queryTagName = "queryProp"

type iCensusOperator interface {
	getKeyValueStringFormat() string
	getPropertySpacer() string
	getTermSpacer() string
}

type censusOperator struct {
	iCensusOperator
}

func (o *censusOperator) BaseString() string {
	queryArgs := make([]string, 0)

	t := reflect.ValueOf(o)

	for i := 0; i < t.NumField(); i++ {
		if tag, ok := t.Type().Field(i).Tag.Lookup(queryTagName); ok {
			fieldValue := t.Field(i).Interface()

			if o.isValueNilOrDefault(fieldValue) || o.isValueTagDefault(fieldValue, tag) {
				continue
			}

			propName := strings.Split(tag, ",")[:1]
			propValue := o.getStringValue(fieldValue)

			queryArgs = append(queryArgs, fmt.Sprintf(o.getKeyValueStringFormat(), propName, propValue))
		}
	}

	return strings.Join(queryArgs, o.getPropertySpacer())
}

func (o *censusOperator) isValueNilOrDefault(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case string:
		return v == ""
	case []interface{}:
		return len(v) == 0
	case bool:
		return v == false
	}

	return false
}

func (o *censusOperator) isValueTagDefault(fieldValue interface{}, tag string) bool {
	tagArgs := strings.Split(tag, ",")

	var defaultValue string
	if defaultSet, _ := fmt.Sscanf(strings.Join(tagArgs[1:], ","), "default=%s", &defaultValue); defaultSet > 0 {
		if fieldValue.(string) == defaultValue {
			return true
		}
	}

	return false
}

func (o *censusOperator) getStringValue(value interface{}) string {
	valType := reflect.TypeOf(value).Kind()

	if valType != reflect.String && valType == reflect.Array {
		values := value.([]interface{})

		var sValues []string
		for _, v := range values {
			sValues = append(sValues, v.(string))
		}

		return strings.Join(sValues, o.getTermSpacer())
	}

	return value.(string)
}

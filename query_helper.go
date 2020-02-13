// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import (
	"fmt"
	"reflect"
	"strings"
)

const queryTagName = "queryProp"

type censusOperator interface {
	getKeyValueStringFormat() string
	getPropertySpacer() string
	getTermSpacer() string
}

func operatorToString(op censusOperator) string {
	queryArgs := make([]string, 0)

	v := reflect.ValueOf(op)
	ind := reflect.Indirect(v)
	s := ind.Type()

	for i := 0; i < s.NumField(); i++ {
		if tag, ok := s.Field(i).Tag.Lookup(queryTagName); ok {
			fieldValue := ind.Field(i)
			fieldType := fieldValue.Type()

			if isValueNilOrDefault(fieldValue, fieldType) || isValueTagDefault(fieldValue, tag) {
				continue
			}

			propName := strings.Split(tag, ",")[0]
			propValue := getPropertyValueAsString(fieldValue, op)

			queryArgs = append(queryArgs, fmt.Sprintf(op.getKeyValueStringFormat(), propName, propValue))
		}
	}

	return strings.Join(queryArgs, op.getPropertySpacer())
}

func isValueNilOrDefault(value reflect.Value, valType reflect.Type) bool {
	vi := value.Interface()

	switch reflect.TypeOf(vi).Kind() {
	case reflect.String:
		return value.String() == ""
	case reflect.Slice:
		return value.Len() == 0
	case reflect.Bool:
		return value.Bool() == false
	}

	return false
}

func isValueTagDefault(value reflect.Value, tag string) bool {
	vi := value.Interface()
	tagArgs := strings.Split(tag, ",")

	var defaultValue string
	defaultSet, _ := fmt.Sscanf(strings.Join(tagArgs[1:], ","), "default=%s", &defaultValue)

	if defaultSet > 0 {
		if fmt.Sprintf("%v", vi) == defaultValue {
			return true
		}
	}

	return false
}

func getPropertyValueAsString(value reflect.Value, op censusOperator) string {
	vi := value.Interface()
	rt := reflect.ValueOf(vi).Kind()

	if rt == reflect.Slice {
		var sValues []string
		for i := 0; i < value.Len(); i++ {
			sValues = append(sValues, fmt.Sprintf("%v", value.Index(i)))
		}

		return strings.Join(sValues, op.getTermSpacer())
	}

	return fmt.Sprintf("%v", value)
}

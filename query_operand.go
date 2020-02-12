// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import (
	"fmt"
	"time"
)

type operatorType int

const (
	equals operatorType = iota
	notEquals
	isLessThan
	isLessThanOrEquals
	isGreaterThan
	isGreaterThanOrEquals
	startsWith
	contains
)

type queryOperand struct {
	comparator interface{}
	operator   operatorType
}

func newQueryOperand() *queryOperand {
	return &queryOperand{}
}

func (o *queryOperand) Equals(value interface{}) {
	o.comparator = value
	o.operator = equals
}

func (o *queryOperand) NotEquals(value interface{}) {
	o.comparator = value
	o.operator = notEquals
}

func (o *queryOperand) IsLessThan(value interface{}) {
	o.comparator = value
	o.operator = isLessThan
}

func (o *queryOperand) IsLessThanOrEquals(value interface{}) {
	o.comparator = value
	o.operator = isLessThanOrEquals
}

func (o *queryOperand) IsGreaterThan(value interface{}) {
	o.comparator = value
	o.operator = isGreaterThan
}

func (o *queryOperand) IsGreaterThanOrEquals(value interface{}) {
	o.comparator = value
	o.operator = isGreaterThanOrEquals
}

func (o *queryOperand) StartsWith(value interface{}) {
	o.comparator = value
	o.operator = startsWith
}

func (o *queryOperand) Contains(value interface{}) {
	o.comparator = value
	o.operator = contains
}

func (o *queryOperand) String() string {
	var mod = ""

	switch o.operator {
	case notEquals:
		mod = "!"
	case isLessThan:
		mod = "<"
	case isLessThanOrEquals:
		mod = "["
	case isGreaterThan:
		mod = ">"
	case isGreaterThanOrEquals:
		mod = "]"
	case startsWith:
		mod = "^"
	case contains:
		mod = "*"
	}

	return fmt.Sprintf("=%s%s", mod, o.getComparatorString())
}

func (o *queryOperand) getComparatorString() string {
	if t, ok := o.comparator.(time.Time); ok {
		return t.Format("2006-01-02 15:04:05")
	}
	return fmt.Sprintf("%s", o.comparator)
}

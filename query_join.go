// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import "fmt"

type queryJoin struct {
	censusOperator
	join       []*queryJoin
	collection string
	List       bool             `queryProp:"list"`
	Outer      bool             `queryProp:"outer,default=true"`
	Show       []string         `queryProp:"show"`
	Hide       []string         `queryProp:"hide"`
	Terms      []*queryArgument `queryProp:"terms"`
	On         string           `queryProp:"on"`
	To         string           `queryProp:"to"`
	InjectAt   string           `queryProp:"inject_at"`
}

func newQueryJoin(collection string) *queryJoin {
	return &queryJoin{
		join:       make([]*queryJoin, 0),
		collection: collection,
		List:       false,
		Outer:      true,
		Show:       make([]string, 0),
		Hide:       make([]string, 0),
		Terms:      make([]*queryArgument, 0),
		On:         "",
		To:         "",
		InjectAt:   "",
	}
}

func (j *queryJoin) IsList(isList bool) *queryJoin {
	j.List = isList
	return j
}

func (j *queryJoin) IsOuterJoin(isOuter bool) *queryJoin {
	j.Outer = isOuter
	return j
}

func (j *queryJoin) ShowFields(fields ...string) *queryJoin {
	j.Show = fields
	return j
}

func (j *queryJoin) HideFields(fields ...string) *queryJoin {
	j.Hide = fields
	return j
}

func (j *queryJoin) OnField(field string) *queryJoin {
	j.On = field
	return j
}

func (j *queryJoin) ToField(field string) *queryJoin {
	j.To = field
	return j
}

func (j *queryJoin) WithInjectAt(field string) *queryJoin {
	j.InjectAt = field
	return j
}

func (j *queryJoin) Where(field string) *queryOperand {
	arg := newQueryArgument(field)

	j.Terms = append(j.Terms, arg)

	return arg.operand
}

func (j *queryJoin) JoinCollection(collection string) *queryJoin {
	newJoin := newQueryJoin(collection)
	j.join = append(j.join, newJoin)
	return newJoin
}

func (j *queryJoin) String() string {
	baseString := j.baseString(j)

	if len(baseString) > 0 {
		baseString = "^" + baseString
	}

	subJoinString := ""
	for _, subJoin := range j.join {
		subJoinString += fmt.Sprintf("(%s)", subJoin.String())
	}

	return j.collection + baseString + subJoinString
}

func (j *queryJoin) getKeyValueStringFormat() string {
	return "%s:%s"
}

func (j *queryJoin) getPropertySpacer() string {
	return "^"
}

func (j *queryJoin) getTermSpacer() string {
	return "'"
}

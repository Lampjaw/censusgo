// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import "fmt"

type queryJoin struct {
	censusOperator
	list       bool             `queryProp:"list"`
	outer      bool             `queryProp:"outer,default=true"`
	show       []string         `queryProp:"show"`
	hide       []string         `queryProp:"hide"`
	terms      []*queryArgument `queryProp:"terms"`
	on         string           `queryProp:"on"`
	to         string           `queryProp:"to"`
	injectAt   string           `queryProp:"inject_at"`
	join       []*queryJoin
	collection string
}

func newQueryJoin(collection string) *queryJoin {
	return &queryJoin{
		list:       false,
		outer:      true,
		show:       make([]string, 0),
		hide:       make([]string, 0),
		terms:      make([]*queryArgument, 0),
		on:         "",
		to:         "",
		injectAt:   "",
		join:       make([]*queryJoin, 0),
		collection: collection,
	}
}

func (j *queryJoin) IsList(isList bool) {
	j.list = isList
}

func (j *queryJoin) IsOuterJoin(isOuter bool) {
	j.outer = isOuter
}

func (j *queryJoin) ShowFields(fields ...string) {
	j.show = fields
}

func (j *queryJoin) HideFields(fields ...string) {
	j.hide = fields
}

func (j *queryJoin) OnField(field string) {
	j.on = field
}

func (j *queryJoin) ToField(field string) {
	j.to = field
}

func (j *queryJoin) WithInjectAt(field string) {
	j.injectAt = field
}

func (j *queryJoin) Where(field string) *queryOperand {
	arg := newQueryArgument(field)

	j.terms = append(j.terms, arg)

	return arg.operand
}

func (j *queryJoin) JoinCollection(collection string) *queryJoin {
	newJoin := newQueryJoin(collection)
	j.join = append(j.join, newJoin)
	return newJoin
}

func (j *queryJoin) String() string {
	baseString := j.BaseString()

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

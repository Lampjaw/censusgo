// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import (
	"strings"
)

type query struct {
	censusOperator
	Collection      string
	censusClient    *CensusClient
	terms           []*queryArgument
	exactMatchFirst bool         `queryProp:"exactMatchFirst"`
	timing          bool         `queryProp:"timing"`
	includeNull     bool         `queryProp:"includeNull"`
	caseSensitive   bool         `queryProp:"case,default=true"`
	retry           bool         `queryProp:"retry,default=true"`
	limit           int          `queryProp:"limit,default=-1"`
	limitPerDB      int          `queryProp:"limitPerDB,default=-1"`
	start           int          `queryProp:"start,default=-1"`
	show            []string     `queryProp:"show"`
	hide            []string     `queryProp:"hide"`
	sort            []string     `queryProp:"sort"`
	has             []string     `queryProp:"has"`
	resolve         []string     `queryProp:"resolve"`
	join            []*queryJoin `queryProp:"join"`
	tree            []*queryTree `queryProp:"tree"`
	distinct        string       `queryProp:"distinct"`
	language        string       `queryProp:"lang"`
}

func newQuery(collection string, censusClient *CensusClient) *query {
	return &query{
		Collection:      collection,
		censusClient:    censusClient,
		terms:           make([]*queryArgument, 0),
		exactMatchFirst: false,
		timing:          false,
		includeNull:     false,
		caseSensitive:   true,
		retry:           true,
		limit:           -1,
		limitPerDB:      -1,
		start:           -1,
		show:            make([]string, 0),
		hide:            make([]string, 0),
		sort:            make([]string, 0),
		has:             make([]string, 0),
		resolve:         make([]string, 0),
		join:            make([]*queryJoin, 0),
		tree:            make([]*queryTree, 0),
		distinct:        "",
		language:        "",
	}
}

func (q *query) JoinCollection(collection string) *queryJoin {
	newJoin := newQueryJoin(collection)
	q.join = append(q.join, newJoin)
	return newJoin
}

func (q *query) TreeField(field string) *queryTree {
	newTree := newQueryTree(field)
	q.tree = append(q.tree, newTree)
	return newTree
}

func (q *query) Where(field string) *queryOperand {
	newArg := newQueryArgument(field)
	q.terms = append(q.terms, newArg)
	return newArg.operand
}

func (q *query) ShowFields(fields ...string) *query {
	q.show = append(q.show, fields...)
	return q
}

func (q *query) HideFields(fields ...string) *query {
	q.hide = append(q.hide, fields...)
	return q
}

func (q *query) SetLimit(limit int) *query {
	q.limit = limit
	return q
}

func (q *query) SetStart(start int) *query {
	q.start = start
	return q
}

func (q *query) AddResolve(resolves ...string) *query {
	q.resolve = append(q.resolve, resolves...)
	return q
}

func (q *query) SetLanguage(language CensusLanguage) *query {
	switch language {
	case LangEnglish:
		return q.SetLanguageString("en")
	case LangGerman:
		return q.SetLanguageString("de")
	case LangSpanish:
		return q.SetLanguageString("es")
	case LangFrench:
		return q.SetLanguageString("fr")
	case LangItalian:
		return q.SetLanguageString("it")
	case LangTurkish:
		return q.SetLanguageString("tr")
	}

	return q
}

func (q *query) SetLanguageString(language string) *query {
	q.language = language
	return q
}

func (q *query) GetResults() ([]interface{}, error) {
	return q.censusClient.executeQuery(q)
}

func (q *query) GetResultsBatch() ([]interface{}, error) {
	return q.censusClient.executeQueryBatch(q)
}

func (q *query) GetUrl() string {
	return q.censusClient.createRequestURL(q)
}

func (q *query) String() string {
	baseString := q.BaseString()

	var terms []string
	for _, t := range q.terms {
		terms = append(terms, t.String())
	}

	sTerms := strings.Join(terms, q.getPropertySpacer())

	if len(baseString) > 0 {
		baseString = "?" + baseString

		if len(sTerms) > 0 {
			sTerms = "&" + sTerms
		}
	} else if len(sTerms) > 0 {
		sTerms = "?" + sTerms
	}

	return q.Collection + "/" + baseString + sTerms
}

func (q *query) getKeyValueStringFormat() string {
	return "c:%s=%s"
}

func (q *query) getPropertySpacer() string {
	return "&"
}

func (q *query) getTermSpacer() string {
	return ","
}

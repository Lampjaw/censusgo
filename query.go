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
	ExactMatchFirst bool         `queryProp:"exactMatchFirst"`
	Timing          bool         `queryProp:"timing"`
	IncludeNull     bool         `queryProp:"includeNull"`
	CaseSensitive   bool         `queryProp:"case,default=true"`
	Retry           bool         `queryProp:"retry,default=true"`
	Limit           int          `queryProp:"limit,default=-1"`
	LimitPerDB      int          `queryProp:"limitPerDB,default=-1"`
	Start           int          `queryProp:"start,default=-1"`
	Show            []string     `queryProp:"show"`
	Hide            []string     `queryProp:"hide"`
	Sort            []string     `queryProp:"sort"`
	Has             []string     `queryProp:"has"`
	Resolve         []string     `queryProp:"resolve"`
	Join            []*queryJoin `queryProp:"join"`
	Tree            []*queryTree `queryProp:"tree"`
	Distinct        string       `queryProp:"distinct"`
	Language        string       `queryProp:"lang"`
}

func newQuery(collection string, censusClient *CensusClient) *query {
	return &query{
		Collection:      collection,
		censusClient:    censusClient,
		terms:           make([]*queryArgument, 0),
		ExactMatchFirst: false,
		Timing:          false,
		IncludeNull:     false,
		CaseSensitive:   true,
		Retry:           true,
		Limit:           -1,
		LimitPerDB:      -1,
		Start:           -1,
		Show:            make([]string, 0),
		Hide:            make([]string, 0),
		Sort:            make([]string, 0),
		Has:             make([]string, 0),
		Resolve:         make([]string, 0),
		Join:            make([]*queryJoin, 0),
		Tree:            make([]*queryTree, 0),
		Distinct:        "",
		Language:        "",
	}
}

func (q *query) JoinCollection(collection string) *queryJoin {
	newJoin := newQueryJoin(collection)
	q.Join = append(q.Join, newJoin)
	return newJoin
}

func (q *query) TreeField(field string) *queryTree {
	newTree := newQueryTree(field)
	q.Tree = append(q.Tree, newTree)
	return newTree
}

func (q *query) Where(field string) *queryOperand {
	newArg := newQueryArgument(field)
	q.terms = append(q.terms, newArg)
	return newArg.operand
}

func (q *query) ShowFields(fields ...string) *query {
	q.Show = append(q.Show, fields...)
	return q
}

func (q *query) HideFields(fields ...string) *query {
	q.Hide = append(q.Hide, fields...)
	return q
}

func (q *query) SetLimit(limit int) *query {
	q.Limit = limit
	return q
}

func (q *query) SetStart(start int) *query {
	q.Start = start
	return q
}

func (q *query) AddResolve(resolves ...string) *query {
	q.Resolve = append(q.Resolve, resolves...)
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
	q.Language = language
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
	baseString := q.baseString(q)

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

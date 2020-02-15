// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

import (
	"encoding/json"
	"strings"
)

// Query is the expression root for a census query
type Query struct {
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

// NewQuery creates a new Query object
func NewQuery(collection string, censusClient *CensusClient) *Query {
	return &Query{
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

// JoinCollection joins the query with another collection
func (q *Query) JoinCollection(collection string) *queryJoin {
	newJoin := newQueryJoin(collection)
	q.Join = append(q.Join, newJoin)
	return newJoin
}

// TreeField creates a tree with a specific field
func (q *Query) TreeField(field string) *queryTree {
	newTree := newQueryTree(field)
	q.Tree = append(q.Tree, newTree)
	return newTree
}

// Where begins an argument expression with a field designation
func (q *Query) Where(field string) *queryOperand {
	newArg := newQueryArgument(field)
	q.terms = append(q.terms, newArg)
	return newArg.operand
}

// ShowFields lists the specific fields to include from each record
func (q *Query) ShowFields(fields ...string) *Query {
	q.Show = append(q.Show, fields...)
	return q
}

// HideFields lists the specific fields to exclude from each record
func (q *Query) HideFields(fields ...string) *Query {
	q.Hide = append(q.Hide, fields...)
	return q
}

// SetLimit sets the maximum number of records to return
func (q *Query) SetLimit(limit int) *Query {
	q.Limit = limit
	return q
}

// SetStart sets the record to begin the query from if the collection isn't stored in a cluster
func (q *Query) SetStart(start int) *Query {
	q.Start = start
	return q
}

// AddResolve adds resolution effects to the query
func (q *Query) AddResolve(resolves ...string) *Query {
	q.Resolve = append(q.Resolve, resolves...)
	return q
}

// SetLanguage sets the localization string set to only return a specific language
func (q *Query) SetLanguage(language CensusLanguage) *Query {
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

// SetLanguageString sets the localization string set to only return a specific language
func (q *Query) SetLanguageString(language string) *Query {
	q.Language = language
	return q
}

// GetResults returns the records from the census API as stated in the query expression
func (q *Query) GetResults() ([]byte, error) {
	res, err := q.censusClient.executeQuery(q)
	if err != nil {
		return nil, err
	}
	return json.Marshal(res)
}

// GetResultsBatch returns ALL records from the census API as stated in the query expression in batches
func (q *Query) GetResultsBatch() ([]byte, error) {
	res, err := q.censusClient.executeQueryBatch(q)
	if err != nil {
		return nil, err
	}
	return json.Marshal(res)
}

// GetURL returns the URL equivalent of the query expression
func (q *Query) GetURL() string {
	return q.censusClient.createRequestURL(q)
}

func (q *Query) String() string {
	baseString := operatorToString(q)

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

func (q *Query) getKeyValueStringFormat() string {
	return "c:%s=%s"
}

func (q *Query) getPropertySpacer() string {
	return "&"
}

func (q *Query) getTermSpacer() string {
	return ","
}

// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

// VERSION for semver
const VERSION = "1.1.0"

// QueryBuilder is a factory used to generate census queries that share the same census service ID and collection namespace
type QueryBuilder struct {
	CensusClient *CensusClient
}

// NewQueryBuilder returns a new QueryBuilder object
func NewQueryBuilder(serviceID string, collectionNamespace string) *QueryBuilder {
	return &QueryBuilder{
		CensusClient: NewCensusClient(serviceID, collectionNamespace),
	}
}

// NewQuery creates a new census query expression
func (b *QueryBuilder) NewQuery(collection string) *Query {
	return NewQuery(collection, b.CensusClient)
}

// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

const VERSION = "1.0.0"

type QueryBuilder struct {
	CensusClient *CensusClient
}

func NewQueryBuilder(serviceID string, collectionNamespace string) *QueryBuilder {
	return &QueryBuilder{
		CensusClient: NewCensusClient(serviceID, collectionNamespace),
	}
}

func (b *QueryBuilder) NewQuery(collection string) *query {
	return newQuery(collection, b.CensusClient)
}

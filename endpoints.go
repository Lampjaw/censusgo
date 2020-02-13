// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

var (
	endpointCensus = "http://census.daybreakgames.com/"

	endpointServiceID       = func(sID string) string { return endpointCensus + "s:" + sID + "/" }
	endpointCollection      = func(sID, ns, op, col string) string { return endpointServiceID(sID) + op + "/" + ns + "/" + col }
	endpointCollectionGet   = func(sID, ns, col string) string { return endpointCollection(sID, ns, "get", col) }
	endpointCollectionCount = func(sID, ns, col string) string { return endpointCollection(sID, ns, "count", col) }
)

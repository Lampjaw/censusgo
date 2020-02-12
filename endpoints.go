// Censusgo - DaybreakGames Census bindings for Go
// Available at https://github.com/lampjaw/censusgo

package censusgo

var (
	EndpointCensus = "http://census.daybreakgames.com/"

	EndpointServiceID       = func(sID string) string { return EndpointCensus + "s:" + sID + "/" }
	EndpointCollection      = func(sID, ns, op, col string) string { return EndpointServiceID(sID) + op + "/" + ns + "/" + col }
	EndpointCollectionGet   = func(sID, ns, col string) string { return EndpointCollection(sID, ns, "get", col) }
	EndpointCollectionCount = func(sID, ns, col string) string { return EndpointCollection(sID, ns, "count", col) }
)

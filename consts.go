package sameas

const (
	defaultSrvUrl = "http://sameas.org/"
	queryWord     = "?q="
	queryURL      = "?uri="
)

type RawQueryType string

const (
	Rdf  RawQueryType = "rdf"
	N3   RawQueryType = "n3"
	Html RawQueryType = "html"
	Json RawQueryType = "json"
	Text RawQueryType = "text"
)

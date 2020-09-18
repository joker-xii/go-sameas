package sameas

import (
	"net/http"
)

type QueryClient struct {
	ServiceURL string
	Client     *http.Client
}

type Result struct {
	NumDuplicates int64    `json:"numDuplicates,string"`
	Duplicates    []string `json:"duplicates"`
	URI           string   `json:"uri"`
}

func NewClient() *QueryClient {
	return &QueryClient{
		ServiceURL: defaultSrvUrl,
		Client:     http.DefaultClient,
	}
}

func NewCustomClient(serviceUrl string, client *http.Client) *QueryClient {
	if serviceUrl == "" {
		serviceUrl = defaultSrvUrl
	}
	return &QueryClient{
		ServiceURL: serviceUrl,
		Client:     client,
	}
}

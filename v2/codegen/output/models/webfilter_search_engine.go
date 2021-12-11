package models

const WebfilterSearchEnginePath = "webfilter/search-engine/"

type WebfilterSearchEngine struct {
	Charset       *string `json:"charset,omitempty"`
	Hostname      *string `json:"hostname,omitempty"`
	Name          *string `json:"name,omitempty"`
	Query         *string `json:"query,omitempty"`
	Safesearch    *string `json:"safesearch,omitempty"`
	SafesearchStr *string `json:"safesearch-str,omitempty"`
	Url           *string `json:"url,omitempty"`
}

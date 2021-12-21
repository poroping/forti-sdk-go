package models

const CmdbBasePath = "/api/v2/cmdb/"

type CmdbRequest struct {
	HTTPMethod string
	Payload    []byte
	Params     CmdbRequestParams
	Path       string
	Mkey       *string
	NoVdom     bool
}

type CmdbRequestParams struct {
	AllowAppend       *bool    `json:"allow_append,omitempty"`
	Action            string   `json:"action,omitempty"`
	Filter            []string `json:"filter,omitempty"`
	Format            []string `json:"format,omitempty"`
	PlainTextPassword *bool    `json:"plain-text-password,omitempty"`
	Vdom              string   `json:"vdom,omitempty"`
}

type CmdbResponse struct {
	HTTPMethod      string      `json:"http_method,omitempty"`
	Size            *int64      `json:"size,omitempty"`
	MatchedCount    *int64      `json:"matched_count,omitempty"`
	NextIdx         *int64      `json:"next_idx,omitempty"`
	Revision        *string     `json:"revision,omitempty"`
	RevisionChanged *bool       `json:"revision_changed,omitempty"`
	OldRevision     *string     `json:"old_revision,omitempty"`
	Results         interface{} `json:"results,omitempty"`
	Vdom            *string     `json:"vdom,omitempty"`
	Path            *string     `json:"path,omitempty"`
	ChildPath       *string     `json:"child_path,omitempty"`
	Name            *string     `json:"name,omitempty"`
	Mkey            interface{} `json:"mkey,omitempty"`
	ChildMkey       interface{} `json:"child_mkey,omitempty"`
	Status          string      `json:"status,omitempty"`
	HTTPStatus      int64       `json:"http_status,omitempty"`
	Serial          *string     `json:"serial,omitempty"`
	Version         *string     `json:"version,omitempty"`
	Build           *int64      `json:"build,omitempty"`
	FullPath        *string
}

type CmdbError500 struct {
	HTTPMethod      string `json:"http_method,omitempty"`
	Revision        string `json:"revision,omitempty"`
	RevisionChanged bool   `json:"revision_changed,omitempty"`
	CLIError        string `json:"cli_error,omitempty"`
	Error           int64  `json:"error,omitempty"`
	Status          string `json:"status,omitempty"`
	HTTPStatus      int64  `json:"http_status,omitempty"`
	Vdom            string `json:"vdom,omitempty"`
	Path            string `json:"path,omitempty"`
	Name            string `json:"name,omitempty"`
	Serial          string `json:"serial,omitempty"`
	Version         string `json:"version,omitempty"`
	Build           int64  `json:"build,omitempty"`
}
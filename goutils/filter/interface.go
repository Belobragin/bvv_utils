package filter

import (
	"net/url"
)

type FilterI interface {
	ParseQueries(url.Values) error
	ProduceFilterQuery(isAdmin bool, startArgs ...string) (string, []any, error)
}

type ProjectFilterI interface {
	// ParseQueries parses the request query to use for database search
	ParseQueries(url.Values) error
	// DistillatekQuery validates the robustness of the query and excludes banned fields
	DistillateQuery() (map[string]interface{}, error)
}

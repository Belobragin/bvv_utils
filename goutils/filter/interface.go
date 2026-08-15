package filter

import (
	"net/url"
)

type FilterI interface {
	ParseQueries(url.Values) error
	ProduceFilterQuery(isAdmin bool, startArgs ...string) (string, []any, error)
}

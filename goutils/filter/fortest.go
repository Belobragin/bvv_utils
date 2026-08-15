package filter

import (
	"net/url"
	"time"

	"github.com/bvv_utils/goutils/util"
	"github.com/google/uuid"
)

type ParseQueryTest struct {
	Description string
	InU         url.Values
	OutE        error
}

type ProduceFilterQuery struct {
	Description string
	InFilter    FilterI
	InIsAdmin   bool
	InStartArgs []string
	OutQuery    string
	OutArgs     []any
	OutE        error
}

var (
	V0                     = setStandardMockValue()
	MockGUUID              = "b3b05871-cc1c-4969-b9ea-7fdd950cf4f8"
	MockGUUID_uuid, _      = uuid.Parse(MockGUUID)
	MockCategoryGUUIDFake  = "b3b05871-cc1c-4969-b9ea-7fdd950cf4f" // no last digit
	MockPage               = 2
	MockDataTimeS          = "2023-08-15T11:55:26.371"
	MockDataTime, _        = time.Parse(util.TimeFormatContext, MockDataTimeS)
	MockFakeFilterKey      = "fake_key"
	MockSortOrderCorrect   = "asc"
	MockSortOrderIncorrect = "qsc"
)

func setStandardMockValue() url.Values {
	var v0 = url.Values{}
	v0.Set("created_at", MockDataTimeS)
	v0.Set("updated_at", MockDataTimeS)
	return v0
}

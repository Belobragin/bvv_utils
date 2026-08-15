package filter

import (
	"errors"
	"maps"
	"net/url"
	"testing"

	"github.com/belobragin/bvv_utils/goutils/mistake"
)

func Test_StandardFilterParse(t *testing.T) {
	var (
		v1 = url.Values{}
	)
	maps.Copy(v1, V0)
	v1.Set("sort_order", MockSortOrderIncorrect)

	var (
		e  error
		ff = new(FilterStandard)

		arrayTests = []ParseQueryTest{
			{"positive_test", V0, nil},
			{"negative-1", v1, mistake.ErrInvalidSortOrder},
		}
	)
	for _, tt := range arrayTests {
		t.Run(tt.Description, func(t *testing.T) {
			if e = ff.StandardFilterParse(tt.InU); !errors.Is(e, tt.OutE) {
				t.Errorf("expected %v, get %v", tt.OutE, e)
			}
		})
	}
}

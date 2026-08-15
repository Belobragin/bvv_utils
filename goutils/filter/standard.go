package filter

import (
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/belobragin/bvv_utils/goutils/util"
)

var InputStandardFilterFields = []string{"sort_order", "sort_by", "page"}

type FilterStandard struct {
	// nullable: false
	// filter by time of creation created_at.
	// in:query
	// example: "2023-08-28T14:11:51.213+0300"
	CreatedAt string `json:"created_at"`
	// nullable: false
	// filter by time of update updated_at.
	// in:query
	// example: "2023-08-28T14:11:51.213+0300"
	UpdatedAt string `json:"updated_at"`
	// nullable: true
	// Sort filter results with a next parameter.
	// in:query
	// enum: catalog_uuid, address_type, point_id, delivery_finished, created_at, updated_at
	// default: created_at
	// example: created_at
	SortBy string `json:"sort_by"`
	// nullable: true
	// Sort direction.
	// in:query
	// enum: asc,desc
	// default:asc
	// example: asc
	SortOrder string `json:"sort_order"`
	// nullable: true
	// Page number.
	// in:query
	// default: 1
	// example: 1
	Page int `json:"page"`
}

func (f *FilterStandard) GetPage() int {
	return f.Page
}

func (f *FilterStandard) GetCreatedAt() string {
	return f.CreatedAt
}

func (f *FilterStandard) GetUpdatedAt() string {
	return f.UpdatedAt
}

func (f *FilterStandard) GetSortOrder() string {
	if f.SortOrder == "" {
		return "DESC"
	}
	return f.SortOrder
}

func (f *FilterStandard) GetSortBy() string {
	return f.SortBy
}

func (f *FilterStandard) StandardFilterParse(values url.Values) error {
	f.CreatedAt = values.Get(util.FilterKeyCreatedTimeStamp)
	f.UpdatedAt = values.Get(util.FilterKeyUpdatedTimeStamp)
	f.SortOrder = values.Get(util.FilterSortOrder)
	f.SortBy = values.Get(util.FilterSortBy)
	page := values.Get(util.FilterKeyPage)
	if page != "" {
		page, err := strconv.Atoi(page)
		if err != nil {
			return mistake.NewAddErr(err, mistake.ErrInvalidParameterFilter)
		}
		f.Page = page
	}
	if f.Page == 0 {
		f.Page = 1
	}
	if err := f.validateStandardFilter(); err != nil {
		return err
	}
	return nil
}

func (f *FilterStandard) validateStandardFilter() error {
	var (
		s string
	)

	s = f.GetCreatedAt()
	if _, e := time.Parse(util.TimeFormatContext, s); e != nil && s != "" {
		return mistake.NewAddErr(e, mistake.ErrInvalidParameterFilter)
	}

	s = f.GetUpdatedAt()
	if _, e := time.Parse(util.TimeFormatContext, s); e != nil && s != "" {
		return mistake.NewAddErr(e, mistake.ErrInvalidParameterFilter)
	}
	if f.Page == 0 {
		f.Page = 1
	}

	if s = f.GetSortOrder(); s != "" && !slices.Contains(util.PermittedSortOrder, strings.ToLower(s)) {
		return mistake.ErrInvalidSortOrder
	}
	return nil
}

func (f *FilterStandard) ProduceStandardFilterQuery(
	interQuery string, argNumP *int, args []any, wherePrefixSet bool,
) ([]any, string, error) {
	if argNumP == nil {
		return nil, "", mistake.ErrNillFilterData
	}
	argNum := *argNumP

	sortOrder := f.GetSortOrder()
	var finalQuery string
	if orderBy := f.GetSortBy(); orderBy == "" {
		finalQuery = ` limit $1 offset $2`
	} else {
		finalQuery = fmt.Sprintf(` order by %s %s NULLS LAST limit $1 offset $2`, orderBy, sortOrder)
	}

	var (
		nextArgNum string
		keyID      interface{}
	)
	if keyID = f.GetCreatedAt(); keyID != "" {
		argNum++
		nextArgNum = fmt.Sprintf("$%d", argNum)
		if wherePrefixSet {
			interQuery = interQuery + fmt.Sprintf(` and %s = %s`, util.FilterKeyCreatedTimeStamp, nextArgNum)
		} else {
			interQuery = interQuery + fmt.Sprintf(` where %s = %s`, util.FilterKeyCreatedTimeStamp, nextArgNum)
			wherePrefixSet = true
		}
		args = append(args, keyID)
	}
	if keyID = f.GetUpdatedAt(); keyID != "" {
		argNum++
		nextArgNum = fmt.Sprintf("$%d", argNum)
		if wherePrefixSet {
			interQuery = interQuery + fmt.Sprintf(` and %s = %s`, util.FilterKeyUpdatedTimeStamp, nextArgNum)
		} else {
			interQuery = interQuery + fmt.Sprintf(` where %s = %s`, util.FilterKeyUpdatedTimeStamp, nextArgNum)
			wherePrefixSet = true
		}
		args = append(args, keyID)
	}
	return args, interQuery + finalQuery, nil
}

func (f *FilterStandard) MakeInitalArgs() []any {
	return []any{
		util.OutputLimitPerPage, // limit
		(f.GetPage() - 1) * 30,  // offset
	}
}

type FilterStandardBrief struct {
	// nullable: true
	// Sort filter results with a next parameter.
	// in:query
	// enum: catalog_uuid, address_type, point_id, delivery_finished, created_at, updated_at
	// default: created_at
	// example: created_at
	SortBy string `json:"sort_by"`
	// nullable: true
	// Sort direction.
	// in:query
	// enum: asc,desc
	// default:asc
	// example: asc
	SortOrder string `json:"sort_order"`
	// nullable: true
	// Page number.
	// in:query
	// default: 1
	// example: 1
	Page int `json:"page"`
}

func (f *FilterStandardBrief) GetPage() int {
	return f.Page
}

func (f *FilterStandardBrief) GetSortOrder() string {
	if f.SortOrder == "" {
		return "DESC"
	}
	return f.SortOrder
}

func (f *FilterStandardBrief) GetSortBy() string {
	return f.SortBy
}

func (f *FilterStandardBrief) StandardFilterParse(values url.Values) error {
	f.SortOrder = values.Get(util.FilterSortOrder)
	f.SortBy = values.Get(util.FilterSortBy)
	page := values.Get(util.FilterKeyPage)
	if page != "" {
		page, err := strconv.Atoi(page)
		if err != nil {
			return mistake.NewAddErr(err, mistake.ErrInvalidParameterFilter)
		}
		f.Page = page
	}
	if f.Page == 0 {
		f.Page = 1
	}
	if err := f.validateStandardFilterBrief(); err != nil {
		return err
	}
	return nil
}

func (f *FilterStandardBrief) ProduceStandardFilterQuery(
	interQuery string, argNumP *int, args []any, wherePrefixSet bool,
) ([]any, string, error) {
	if argNumP == nil {
		return nil, "", mistake.ErrNillFilterData
	}
	sortOrder := f.GetSortOrder()
	var finalQuery string
	if orderBy := f.GetSortBy(); orderBy == "" {
		finalQuery = ` limit $1 offset $2`
	} else {
		finalQuery = fmt.Sprintf(` order by %s %s NULLS LAST limit $1 offset $2`, orderBy, sortOrder)
	}
	return args, interQuery + finalQuery, nil
}

func (f *FilterStandardBrief) MakeInitalArgs() []any {
	return []any{
		util.OutputLimitPerPage, // limit
		(f.GetPage() - 1) * 30,  // offset
	}
}

func (f *FilterStandardBrief) validateStandardFilterBrief() error {

	if f.Page == 0 {
		f.Page = 1
	}

	if s := f.GetSortOrder(); s != "" && !slices.Contains(util.PermittedSortOrder, strings.ToLower(s)) {
		return mistake.ErrInvalidSortOrder
	}
	return nil
}

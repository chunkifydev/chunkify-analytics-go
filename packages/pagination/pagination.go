// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/chunkifydev/chunkify-analytics-go/internal/apijson"
	"github.com/chunkifydev/chunkify-analytics-go/internal/requestconfig"
	"github.com/chunkifydev/chunkify-analytics-go/option"
	"github.com/chunkifydev/chunkify-analytics-go/packages/param"
	"github.com/chunkifydev/chunkify-analytics-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type EventsPage[T any] struct {
	Events     []T    `json:"events"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r EventsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *EventsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EventsPage[T]) GetNextPage() (res *EventsPage[T], err error) {
	if len(r.Events) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EventsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EventsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EventsPageAutoPager[T any] struct {
	page *EventsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewEventsPageAutoPager[T any](page *EventsPage[T], err error) *EventsPageAutoPager[T] {
	return &EventsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EventsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Events) == 0 {
		return false
	}
	if r.idx >= len(r.page.Events) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Events) == 0 {
			return false
		}
	}
	r.cur = r.page.Events[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EventsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EventsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EventsPageAutoPager[T]) Index() int {
	return r.run
}

type VideosPage[T any] struct {
	Videos     []T    `json:"videos"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Videos      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r VideosPage[T]) RawJSON() string { return r.JSON.raw }
func (r *VideosPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *VideosPage[T]) GetNextPage() (res *VideosPage[T], err error) {
	if len(r.Videos) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *VideosPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &VideosPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type VideosPageAutoPager[T any] struct {
	page *VideosPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewVideosPageAutoPager[T any](page *VideosPage[T], err error) *VideosPageAutoPager[T] {
	return &VideosPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *VideosPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Videos) == 0 {
		return false
	}
	if r.idx >= len(r.page.Videos) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Videos) == 0 {
			return false
		}
	}
	r.cur = r.page.Videos[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *VideosPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *VideosPageAutoPager[T]) Err() error {
	return r.err
}

func (r *VideosPageAutoPager[T]) Index() int {
	return r.run
}

type ViewsPage[T any] struct {
	Views      []T    `json:"views"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Views       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ViewsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *ViewsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ViewsPage[T]) GetNextPage() (res *ViewsPage[T], err error) {
	if len(r.Views) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ViewsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ViewsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ViewsPageAutoPager[T any] struct {
	page *ViewsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewViewsPageAutoPager[T any](page *ViewsPage[T], err error) *ViewsPageAutoPager[T] {
	return &ViewsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ViewsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Views) == 0 {
		return false
	}
	if r.idx >= len(r.page.Views) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Views) == 0 {
			return false
		}
	}
	r.cur = r.page.Views[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ViewsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ViewsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ViewsPageAutoPager[T]) Index() int {
	return r.run
}

type ErrorsPage[T any] struct {
	Errors     []T    `json:"errors"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ErrorsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *ErrorsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ErrorsPage[T]) GetNextPage() (res *ErrorsPage[T], err error) {
	if len(r.Errors) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ErrorsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ErrorsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ErrorsPageAutoPager[T any] struct {
	page *ErrorsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewErrorsPageAutoPager[T any](page *ErrorsPage[T], err error) *ErrorsPageAutoPager[T] {
	return &ErrorsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ErrorsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Errors) == 0 {
		return false
	}
	if r.idx >= len(r.page.Errors) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Errors) == 0 {
			return false
		}
	}
	r.cur = r.page.Errors[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ErrorsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ErrorsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ErrorsPageAutoPager[T]) Index() int {
	return r.run
}

type DimensionValuesPage[T any] struct {
	Values     []T    `json:"values"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Values      respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r DimensionValuesPage[T]) RawJSON() string { return r.JSON.raw }
func (r *DimensionValuesPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DimensionValuesPage[T]) GetNextPage() (res *DimensionValuesPage[T], err error) {
	if len(r.Values) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DimensionValuesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DimensionValuesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DimensionValuesPageAutoPager[T any] struct {
	page *DimensionValuesPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewDimensionValuesPageAutoPager[T any](page *DimensionValuesPage[T], err error) *DimensionValuesPageAutoPager[T] {
	return &DimensionValuesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DimensionValuesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Values) == 0 {
		return false
	}
	if r.idx >= len(r.page.Values) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Values) == 0 {
			return false
		}
	}
	r.cur = r.page.Values[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DimensionValuesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *DimensionValuesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *DimensionValuesPageAutoPager[T]) Index() int {
	return r.run
}

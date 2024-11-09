// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/jx"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

// Ref: #
type CategoriesResponse struct {
	Data    []Category `json:"data"`
	Page    int        `json:"page"`
	Pages   int        `json:"pages"`
	PerPage int        `json:"per_page"`
}

// GetData returns the value of Data.
func (s *CategoriesResponse) GetData() []Category {
	return s.Data
}

// GetPage returns the value of Page.
func (s *CategoriesResponse) GetPage() int {
	return s.Page
}

// GetPages returns the value of Pages.
func (s *CategoriesResponse) GetPages() int {
	return s.Pages
}

// GetPerPage returns the value of PerPage.
func (s *CategoriesResponse) GetPerPage() int {
	return s.PerPage
}

// SetData sets the value of Data.
func (s *CategoriesResponse) SetData(val []Category) {
	s.Data = val
}

// SetPage sets the value of Page.
func (s *CategoriesResponse) SetPage(val int) {
	s.Page = val
}

// SetPages sets the value of Pages.
func (s *CategoriesResponse) SetPages(val int) {
	s.Pages = val
}

// SetPerPage sets the value of PerPage.
func (s *CategoriesResponse) SetPerPage(val int) {
	s.PerPage = val
}

func (*CategoriesResponse) v1AdminCategoriesGetRes() {}

// Ref: #
type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// GetID returns the value of ID.
func (s *Category) GetID() int {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *Category) GetTitle() string {
	return s.Title
}

// SetID sets the value of ID.
func (s *Category) SetID(val int) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *Category) SetTitle(val string) {
	s.Title = val
}

func (*Category) v1AdminCategoriesPostRes() {}
func (*Category) v1AdminCategoriesPutRes()  {}

// Ref: #
type CategoryInput struct {
	Title string `json:"title"`
}

// GetTitle returns the value of Title.
func (s *CategoryInput) GetTitle() string {
	return s.Title
}

// SetTitle sets the value of Title.
func (s *CategoryInput) SetTitle(val string) {
	s.Title = val
}

// Ref: #
type Error struct {
	// Error message.
	Message string       `json:"message"`
	Details ErrorDetails `json:"details"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// GetDetails returns the value of Details.
func (s *Error) GetDetails() ErrorDetails {
	return s.Details
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// SetDetails sets the value of Details.
func (s *Error) SetDetails(val ErrorDetails) {
	s.Details = val
}

type ErrorDetails map[string]jx.Raw

func (s *ErrorDetails) init() ErrorDetails {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// V1AdminCategoriesDeleteNoContent is response for V1AdminCategoriesDelete operation.
type V1AdminCategoriesDeleteNoContent struct{}

func (*V1AdminCategoriesDeleteNoContent) v1AdminCategoriesDeleteRes() {}

type V1AdminCategoriesDeleteNotFound Error

func (*V1AdminCategoriesDeleteNotFound) v1AdminCategoriesDeleteRes() {}

type V1AdminCategoriesDeleteUnauthorized Error

func (*V1AdminCategoriesDeleteUnauthorized) v1AdminCategoriesDeleteRes() {}

type V1AdminCategoriesGetBadRequest Error

func (*V1AdminCategoriesGetBadRequest) v1AdminCategoriesGetRes() {}

type V1AdminCategoriesGetNotFound Error

func (*V1AdminCategoriesGetNotFound) v1AdminCategoriesGetRes() {}

type V1AdminCategoriesGetUnauthorized Error

func (*V1AdminCategoriesGetUnauthorized) v1AdminCategoriesGetRes() {}

type V1AdminCategoriesPostBadRequest Error

func (*V1AdminCategoriesPostBadRequest) v1AdminCategoriesPostRes() {}

type V1AdminCategoriesPostConflict Error

func (*V1AdminCategoriesPostConflict) v1AdminCategoriesPostRes() {}

type V1AdminCategoriesPostUnauthorized Error

func (*V1AdminCategoriesPostUnauthorized) v1AdminCategoriesPostRes() {}

type V1AdminCategoriesPutBadRequest Error

func (*V1AdminCategoriesPutBadRequest) v1AdminCategoriesPutRes() {}

type V1AdminCategoriesPutConflict Error

func (*V1AdminCategoriesPutConflict) v1AdminCategoriesPutRes() {}

type V1AdminCategoriesPutNotFound Error

func (*V1AdminCategoriesPutNotFound) v1AdminCategoriesPutRes() {}

type V1AdminCategoriesPutUnauthorized Error

func (*V1AdminCategoriesPutUnauthorized) v1AdminCategoriesPutRes() {}

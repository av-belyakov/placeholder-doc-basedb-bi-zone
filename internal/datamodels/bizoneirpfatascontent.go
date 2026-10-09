package datamodels

import (
	"errors"
	"fmt"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func NewBiZoneIRPSContent() *BiZoneIRPSContent {
	return &BiZoneIRPSContent{}
}

func (c *BiZoneIRPSContent) Get() *BiZoneIRPSContent {
	return c
}

// GetContent для поля 'content'
func (c *BiZoneIRPSContent) GetContent() string {
	return c.Content
}

// SetContent для поля 'content'
func (c *BiZoneIRPSContent) SetContent(v string) error {
	c.Content = v

	return nil
}

// SetAnyContent для поля 'content'
func (c *BiZoneIRPSContent) SetAnyContent(a any) error {
	return c.SetContent(fmt.Sprint(a))
}

// GetNocase для поля 'nocase'
func (c *BiZoneIRPSContent) GetNocase() bool {
	return c.Nocase
}

// SetNocase для поля 'nocase'
func (c *BiZoneIRPSContent) SetNocase(v bool) error {
	c.Nocase = v

	return nil
}

// SetAnyNocase для поля 'nocase'
func (c *BiZoneIRPSContent) SetAnyNocase(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'nocase'")
	}

	return c.SetNocase(v)
}

// GetHTTPURI для поля 'http_uri'
func (c *BiZoneIRPSContent) GetHTTPURI() bool {
	return c.HTTPURI
}

// SetHTTPURI для поля 'http_uri'
func (c *BiZoneIRPSContent) SetHTTPURI(v bool) error {
	c.Nocase = v

	return nil
}

// SetAnyHTTPURI для поля 'http_uri'
func (c *BiZoneIRPSContent) SetAnyHTTPURI(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_uri'")
	}

	return c.SetHTTPURI(v)
}

// GetRawbytes для поля 'rawbytes'
func (c *BiZoneIRPSContent) GetRawbytes() bool {
	return c.Rawbytes
}

// SetRawbytes для поля 'rawbytes'
func (c *BiZoneIRPSContent) SetRawbytes(v bool) error {
	c.Rawbytes = v

	return nil
}

// SetAnyRawbytes для поля 'rawbytes'
func (c *BiZoneIRPSContent) SetAnyRawbytes(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'rawbytes'")
	}

	return c.SetRawbytes(v)
}

// GetHTTPCookie для поля 'http_cookie'
func (c *BiZoneIRPSContent) GetHTTPCookie() bool {
	return c.HTTPCookie
}

// SetHTTPCookie для поля 'http_cookie'
func (c *BiZoneIRPSContent) SetHTTPCookie(v bool) error {
	c.HTTPCookie = v

	return nil
}

// SetAnyHTTPCookie для поля 'http_cookie'
func (c *BiZoneIRPSContent) SetAnyHTTPCookie(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_cookie'")
	}

	return c.SetHTTPCookie(v)
}

// GetHTTPHeader для поля 'http_header'
func (c *BiZoneIRPSContent) GetHTTPHeader() bool {
	return c.HTTPHeader
}

// SetHTTPHeader для поля 'http_header'
func (c *BiZoneIRPSContent) SetHTTPHeader(v bool) error {
	c.HTTPHeader = v

	return nil
}

// SetAnyHTTPHeader для поля 'http_header'
func (c *BiZoneIRPSContent) SetAnyHTTPHeader(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_header'")
	}

	return c.SetHTTPHeader(v)
}

// GetHTTPMethod для поля 'http_method'
func (c *BiZoneIRPSContent) GetHTTPMethod() bool {
	return c.HTTPMethod
}

// SetHTTPMethod для поля 'http_method'
func (c *BiZoneIRPSContent) SetHTTPMethod(v bool) error {
	c.HTTPMethod = v

	return nil
}

// SetAnyHTTPMethod для поля 'http_method'
func (c *BiZoneIRPSContent) SetAnyHTTPMethod(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_method'")
	}

	return c.SetHTTPMethod(v)
}

// GetHTTPRawURI для поля 'http_raw_uri'
func (c *BiZoneIRPSContent) GetHTTPRawURI() bool {
	return c.HTTPRawURI
}

// SetHTTPRawURI для поля 'http_raw_uri'
func (c *BiZoneIRPSContent) SetHTTPRawURI(v bool) error {
	c.HTTPRawURI = v

	return nil
}

// SetAnyHTTPRawURI для поля 'http_raw_uri'
func (c *BiZoneIRPSContent) SetAnyHTTPRawURI(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_raw_uri'")
	}

	return c.SetHTTPRawURI(v)
}

// GetHTTPStatMsg для поля 'http_stat_msg'
func (c *BiZoneIRPSContent) GetHTTPStatMsg() bool {
	return c.HTTPStatMsg
}

// SetHTTPStatMsg для поля 'http_stat_msg'
func (c *BiZoneIRPSContent) SetHTTPStatMsg(v bool) error {
	c.HTTPStatMsg = v

	return nil
}

// SetAnyHTTPStatMsg для поля 'http_stat_msg'
func (c *BiZoneIRPSContent) SetAnyHTTPStatMsg(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_stat_msg'")
	}

	return c.SetHTTPStatMsg(v)
}

// GetHTTPStatCode для поля 'http_stat_code'
func (c *BiZoneIRPSContent) GetHTTPStatCode() bool {
	return c.HTTPStatCode
}

// SetHTTPStatCode для поля 'http_stat_code'
func (c *BiZoneIRPSContent) SetHTTPStatCode(v bool) error {
	c.HTTPStatCode = v

	return nil
}

// SetAnyHTTPStatCode для поля 'http_stat_code'
func (c *BiZoneIRPSContent) SetAnyHTTPStatCode(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_stat_code'")
	}

	return c.SetHTTPStatCode(v)
}

// GetHTTPRawCookie для поля 'http_raw_cookie'
func (c *BiZoneIRPSContent) GetHTTPRawCookie() bool {
	return c.HTTPRawCookie
}

// SetHTTPRawCookie для поля 'http_raw_cookie'
func (c *BiZoneIRPSContent) SetHTTPRawCookie(v bool) error {
	c.HTTPRawCookie = v

	return nil
}

// SetAnyHTTPRawCookie для поля 'http_raw_cookie'
func (c *BiZoneIRPSContent) SetAnyHTTPRawCookie(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_raw_cookie'")
	}

	return c.SetHTTPRawCookie(v)
}

// GetHTTPRawHeader для поля 'http_raw_header'
func (c *BiZoneIRPSContent) GetHTTPRawHeader() bool {
	return c.HTTPRawHeader
}

// SetHTTPRawHeader для поля 'http_raw_header'
func (c *BiZoneIRPSContent) SetHTTPRawHeader(v bool) error {
	c.HTTPRawHeader = v

	return nil
}

// SetAnyHTTPRawHeader для поля 'http_raw_header'
func (c *BiZoneIRPSContent) SetAnyHTTPRawHeader(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_raw_header'")
	}

	return c.SetHTTPRawHeader(v)
}

// GetHTTPClientBody для поля 'http_client_body'
func (c *BiZoneIRPSContent) GetHTTPClientBody() bool {
	return c.HTTPClientBody
}

// SetHTTPClientBody для поля 'http_client_body'
func (c *BiZoneIRPSContent) SetHTTPClientBody(v bool) error {
	c.HTTPClientBody = v

	return nil
}

// SetAnyHTTPClientBody для поля 'http_client_body'
func (c *BiZoneIRPSContent) SetAnyHTTPClientBody(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_client_body'")
	}

	return c.SetHTTPClientBody(v)
}

// GetDistance для поля 'distance'
func (c *BiZoneIRPSContent) GetDistance() *string {
	return c.Distance
}

// SetDistance для поля 'distance'
func (c *BiZoneIRPSContent) SetDistance(v *string) error {
	c.Distance = v

	return nil
}

// SetAnyDistance для поля 'distance'
func (c *BiZoneIRPSContent) SetAnyDistance(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'distance'")
	}

	return c.SetDistance(&v)
}

// GetFastPattern для поля 'fast_pattern'
func (c *BiZoneIRPSContent) GetFastPattern() *string {
	return c.FastPattern
}

// SetFastPattern для поля 'fast_pattern'
func (c *BiZoneIRPSContent) SetFastPattern(v *string) error {
	c.FastPattern = v

	return nil
}

// SetAnyFastPattern для поля 'fast_pattern'
func (c *BiZoneIRPSContent) SetAnyFastPattern(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'fast_pattern'")
	}

	return c.SetFastPattern(&v)
}

// GetDepth для поля 'depth'
func (c *BiZoneIRPSContent) GetDepth() *string {
	return c.Depth
}

// SetDepth для поля 'depth'
func (c *BiZoneIRPSContent) SetDepth(v *string) error {
	c.Depth = v

	return nil
}

// SetAnyDepth для поля 'depth'
func (c *BiZoneIRPSContent) SetAnyDepth(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'depth'")
	}

	return c.SetDepth(&v)
}

// GetOffset для поля 'offset'
func (c *BiZoneIRPSContent) GetOffset() *string {
	return c.Offset
}

// SetOffset для поля 'offset'
func (c *BiZoneIRPSContent) SetOffset(v *string) error {
	c.Offset = v

	return nil
}

// SetAnyOffset для поля 'offset'
func (c *BiZoneIRPSContent) SetAnyOffset(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'offset'")
	}

	return c.SetOffset(&v)
}

// GetWithin для поля 'within'
func (c *BiZoneIRPSContent) GetWithin() *string {
	return c.Within
}

// SetWithin для поля 'within'
func (c *BiZoneIRPSContent) SetWithin(v *string) error {
	c.Within = v

	return nil
}

// SetAnyWithin для поля 'within'
func (c *BiZoneIRPSContent) SetAnyWithin(a any) error {
	v, ok := a.(string)
	if !ok {
		return errors.New("type conversion error for field 'within'")
	}

	return c.SetWithin(&v)
}

// ToStringBeautiful форматированный вывод
func (c *BiZoneIRPSContent) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := supportingfunctions.GetWhitespace(num)

	fmt.Fprintf(&str, "%s'content': '%s'\n", ws, c.Content)
	fmt.Fprintf(&str, "%s'nocase': '%t'\n", ws, c.Nocase)
	fmt.Fprintf(&str, "%s'http_uri': '%t'\n", ws, c.HTTPURI)
	fmt.Fprintf(&str, "%s'rawbytes': '%t'\n", ws, c.Rawbytes)
	fmt.Fprintf(&str, "%s'http_cookie': '%t'\n", ws, c.HTTPCookie)
	fmt.Fprintf(&str, "%s'http_header': '%t'\n", ws, c.HTTPHeader)
	fmt.Fprintf(&str, "%s'http_method': '%t'\n", ws, c.HTTPMethod)
	fmt.Fprintf(&str, "%s'http_raw_uri': '%t'\n", ws, c.HTTPRawURI)
	fmt.Fprintf(&str, "%s'http_stat_msg': '%t'\n", ws, c.HTTPStatMsg)
	fmt.Fprintf(&str, "%s'http_stat_code': '%t'\n", ws, c.HTTPStatCode)
	fmt.Fprintf(&str, "%s'http_raw_cookie': '%t'\n", ws, c.HTTPRawCookie)
	fmt.Fprintf(&str, "%s'http_raw_header': '%t'\n", ws, c.HTTPRawHeader)
	fmt.Fprintf(&str, "%s'http_client_body': '%t'\n", ws, c.HTTPClientBody)
	fmt.Fprintf(&str, "%s'depth': '%s'\n", ws, *c.Depth)
	fmt.Fprintf(&str, "%s'offset': '%s'\n", ws, *c.Offset)
	fmt.Fprintf(&str, "%s'within': '%s'\n", ws, *c.Within)
	fmt.Fprintf(&str, "%s'distance': '%s'\n", ws, *c.Distance)
	fmt.Fprintf(&str, "%s'fast_pattern': '%s'\n", ws, *c.FastPattern)

	return str.String()
}

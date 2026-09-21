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

func (d *BiZoneIRPSContent) Get() *BiZoneIRPSContent {
	return d
}

// GetContent для поля content
func (d *BiZoneIRPSContent) GetContent() string {
	return d.Content
}

// SetContent для поля content
func (d *BiZoneIRPSContent) SetContent(v string) error {
	d.Content = v

	return nil
}

// SetAnyContent для поля content
func (d *BiZoneIRPSContent) SetAnyContent(a any) error {
	return d.SetContent(fmt.Sprint(a))
}

// GetNocase для поля nocase
func (d *BiZoneIRPSContent) GetNocase() bool {
	return d.Nocase
}

// SetNocase для поля nocase
func (d *BiZoneIRPSContent) SetNocase(v bool) error {
	d.Nocase = v

	return nil
}

// SetAnyNocase для поля nocase
func (d *BiZoneIRPSContent) SetAnyNocase(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'nocase'")
	}

	return d.SetNocase(v)
}

// GetHTTPURI для поля http_uri
func (d *BiZoneIRPSContent) GetHTTPURI() bool {
	return d.HTTPURI
}

// SetHTTPURI для поля http_uri
func (d *BiZoneIRPSContent) SetHTTPURI(v bool) error {
	d.Nocase = v

	return nil
}

// SetAnyHTTPURI для поля http_uri
func (d *BiZoneIRPSContent) SetAnyHTTPURI(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_uri'")
	}

	return d.SetHTTPURI(v)
}

// GetRawbytes для поля rawbytes
func (d *BiZoneIRPSContent) GetRawbytes() bool {
	return d.Rawbytes
}

// SetRawbytes для поля rawbytes
func (d *BiZoneIRPSContent) SetRawbytes(v bool) error {
	d.Rawbytes = v

	return nil
}

// SetAnyRawbytes для поля rawbytes
func (d *BiZoneIRPSContent) SetAnyRawbytes(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'rawbytes'")
	}

	return d.SetRawbytes(v)
}

// GetHTTPCookie для поля http_cookie
func (d *BiZoneIRPSContent) GetHTTPCookie() bool {
	return d.HTTPCookie
}

// SetHTTPCookie для поля http_cookie
func (d *BiZoneIRPSContent) SetHTTPCookie(v bool) error {
	d.HTTPCookie = v

	return nil
}

// SetAnyHTTPCookie для поля http_cookie
func (d *BiZoneIRPSContent) SetAnyHTTPCookie(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_cookie'")
	}

	return d.SetHTTPCookie(v)
}

// GetHTTPHeader для поля http_header
func (d *BiZoneIRPSContent) GetHTTPHeader() bool {
	return d.HTTPHeader
}

// SetHTTPHeader для поля http_header
func (d *BiZoneIRPSContent) SetHTTPHeader(v bool) error {
	d.HTTPHeader = v

	return nil
}

// SetAnyHTTPHeader для поля http_header
func (d *BiZoneIRPSContent) SetAnyHTTPHeader(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_header'")
	}

	return d.SetHTTPHeader(v)
}

// GetHTTPMethod для поля http_method
func (d *BiZoneIRPSContent) GetHTTPMethod() bool {
	return d.HTTPMethod
}

// SetHTTPMethod для поля http_method
func (d *BiZoneIRPSContent) SetHTTPMethod(v bool) error {
	d.HTTPMethod = v

	return nil
}

// SetAnyHTTPMethod для поля http_method
func (d *BiZoneIRPSContent) SetAnyHTTPMethod(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_method'")
	}

	return d.SetHTTPMethod(v)
}

// GetHTTPRawURI для поля http_raw_uri
func (d *BiZoneIRPSContent) GetHTTPRawURI() bool {
	return d.HTTPRawURI
}

// SetHTTPRawURI для поля http_raw_uri
func (d *BiZoneIRPSContent) SetHTTPRawURI(v bool) error {
	d.HTTPRawURI = v

	return nil
}

// SetAnyHTTPRawURI для поля http_raw_uri
func (d *BiZoneIRPSContent) SetAnyHTTPRawURI(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_raw_uri'")
	}

	return d.SetHTTPRawURI(v)
}

// GetHTTPStatMsg для поля http_stat_msg
func (d *BiZoneIRPSContent) GetHTTPStatMsg() bool {
	return d.HTTPStatMsg
}

// SetHTTPStatMsg для поля http_stat_msg
func (d *BiZoneIRPSContent) SetHTTPStatMsg(v bool) error {
	d.HTTPStatMsg = v

	return nil
}

// SetAnyHTTPStatMsg для поля http_stat_msg
func (d *BiZoneIRPSContent) SetAnyHTTPStatMsg(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_stat_msg'")
	}

	return d.SetHTTPStatMsg(v)
}

// GetHTTPStatCode для поля http_stat_code
func (d *BiZoneIRPSContent) GetHTTPStatCode() bool {
	return d.HTTPStatCode
}

// SetHTTPStatCode для поля http_stat_code
func (d *BiZoneIRPSContent) SetHTTPStatCode(v bool) error {
	d.HTTPStatCode = v

	return nil
}

// SetAnyHTTPStatCode для поля http_stat_code
func (d *BiZoneIRPSContent) SetAnyHTTPStatCode(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_stat_code'")
	}

	return d.SetHTTPStatCode(v)
}

// GetHTTPRawCookie для поля http_raw_cookie
func (d *BiZoneIRPSContent) GetHTTPRawCookie() bool {
	return d.HTTPRawCookie
}

// SetHTTPRawCookie для поля http_raw_cookie
func (d *BiZoneIRPSContent) SetHTTPRawCookie(v bool) error {
	d.HTTPRawCookie = v

	return nil
}

// SetAnyHTTPRawCookie для поля http_raw_cookie
func (d *BiZoneIRPSContent) SetAnyHTTPRawCookie(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_raw_cookie'")
	}

	return d.SetHTTPRawCookie(v)
}

// GetHTTPRawHeader для поля http_raw_header
func (d *BiZoneIRPSContent) GetHTTPRawHeader() bool {
	return d.HTTPRawHeader
}

// SetHTTPRawHeader для поля http_raw_header
func (d *BiZoneIRPSContent) SetHTTPRawHeader(v bool) error {
	d.HTTPRawHeader = v

	return nil
}

// SetAnyHTTPRawHeader для поля http_raw_header
func (d *BiZoneIRPSContent) SetAnyHTTPRawHeader(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_raw_header'")
	}

	return d.SetHTTPRawHeader(v)
}

// GetHTTPClientBody для поля http_client_body
func (d *BiZoneIRPSContent) GetHTTPClientBody() bool {
	return d.HTTPClientBody
}

// SetHTTPClientBody для поля http_client_body
func (d *BiZoneIRPSContent) SetHTTPClientBody(v bool) error {
	d.HTTPClientBody = v

	return nil
}

// SetAnyHTTPClientBody для поля http_client_body
func (d *BiZoneIRPSContent) SetAnyHTTPClientBody(a any) error {
	v, ok := a.(bool)
	if !ok {
		return errors.New("type conversion error for field 'http_client_body'")
	}

	return d.SetHTTPClientBody(v)
}

// GetDistance для поля distance
func (d *BiZoneIRPSContent) GetDistance() *string {
	return d.Distance
}

// SetDistance для поля distance
func (d *BiZoneIRPSContent) SetDistance(v *string) error {
	d.Distance = v

	return nil
}

// SetAnyDistance для поля distance
func (d *BiZoneIRPSContent) SetAnyDistance(a any) error {
	v, ok := a.(*string)
	if !ok {
		return errors.New("type conversion error for field 'distance'")
	}

	return d.SetDistance(v)
}

// GetFastPattern для поля fast_pattern
func (d *BiZoneIRPSContent) GetFastPattern() *bool {
	return d.FastPattern
}

// SetFastPattern для поля fast_pattern
func (d *BiZoneIRPSContent) SetFastPattern(v *bool) error {
	d.FastPattern = v

	return nil
}

// SetAnyFastPattern для поля fast_pattern
func (d *BiZoneIRPSContent) SetAnyFastPattern(a any) error {
	v, ok := a.(*bool)
	if !ok {
		return errors.New("type conversion error for field 'fast_pattern'")
	}

	return d.SetFastPattern(v)
}

// GetDepth для поля depth
func (d *BiZoneIRPSContent) GetDepth() *int {
	return d.Depth
}

// SetDepth для поля depth
func (d *BiZoneIRPSContent) SetDepth(v *int) error {
	d.Depth = v

	return nil
}

// SetAnyDepth для поля depth
func (d *BiZoneIRPSContent) SetAnyDepth(a any) error {
	v, ok := a.(*int)
	if !ok {
		return errors.New("type conversion error for field 'depth'")
	}

	return d.SetDepth(v)
}

// GetOffset для поля offset
func (d *BiZoneIRPSContent) GetOffset() *int {
	return d.Offset
}

// SetOffset для поля offset
func (d *BiZoneIRPSContent) SetOffset(v *int) error {
	d.Offset = v

	return nil
}

// SetAnyOffset для поля offset
func (d *BiZoneIRPSContent) SetAnyOffset(a any) error {
	v, ok := a.(*int)
	if !ok {
		return errors.New("type conversion error for field 'offset'")
	}

	return d.SetOffset(v)
}

// GetWithin для поля within
func (d *BiZoneIRPSContent) GetWithin() *int {
	return d.Within
}

// SetWithin для поля within
func (d *BiZoneIRPSContent) SetWithin(v *int) error {
	d.Within = v

	return nil
}

// SetAnyWithin для поля within
func (d *BiZoneIRPSContent) SetAnyWithin(a any) error {
	v, ok := a.(*int)
	if !ok {
		return errors.New("type conversion error for field 'within'")
	}

	return d.SetWithin(v)
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
	fmt.Fprintf(&str, "%s'distance': '%s'\n", ws, *c.Distance)
	fmt.Fprintf(&str, "%s'fast_pattern': '%t'\n", ws, *c.FastPattern)
	fmt.Fprintf(&str, "%s'depth': '%d'\n", ws, *c.Depth)
	fmt.Fprintf(&str, "%s'offset': '%d'\n", ws, *c.Offset)
	fmt.Fprintf(&str, "%s'within': '%d'\n", ws, *c.Within)

	return str.String()
}

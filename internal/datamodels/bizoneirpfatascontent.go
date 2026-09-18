package datamodels

import (
	"errors"
	"fmt"
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

// SetAnySRuleBody для поля nocase
func (d *BiZoneIRPSContent) SetAnySRuleBody(a any) error {
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

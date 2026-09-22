package handlers

// NewListBiZoneHandlerSContents начальный обработчик событий полей 'data.data_security.s_content.*'
func NewListBiZoneHandlerSContents(sc *SupportingStructureForSContentType) map[string][]func(any) error {
	return map[string][]func(any) error{
		// --- depth ---
		"data.data_security.s_content.depth": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.depth",
					a,
					sc.GetSContentTmp().SetAnyDepth,
				)
			},
		},
		// --- nocase ---
		"data.data_security.s_content.nocase": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.nocase",
					a,
					sc.GetSContentTmp().SetAnyNocase,
				)
			},
		},
		// --- content ---
		"data.data_security.s_content.content": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.content",
					a,
					sc.GetSContentTmp().SetAnyContent,
				)
			},
		},
		// --- offset ---
		"data.data_security.s_content.offset": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.offset",
					a,
					sc.GetSContentTmp().SetAnyOffset,
				)
			},
		},
		// --- within ---
		"data.data_security.s_content.within": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.within",
					a,
					sc.GetSContentTmp().SetAnyWithin,
				)
			},
		},
		// --- http_uri ---
		"data.data_security.s_content.http_uri": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_uri",
					a,
					sc.GetSContentTmp().SetAnyHTTPURI,
				)
			},
		},
		// --- rawbytes ---
		"data.data_security.s_content.rawbytes": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.rawbytes",
					a,
					sc.GetSContentTmp().SetAnyRawbytes,
				)
			},
		},
		// --- distance ---
		"data.data_security.s_content.distance": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.distance",
					a,
					sc.GetSContentTmp().SetAnyDistance,
				)
			},
		},
		// --- http_cookie ---
		"data.data_security.s_content.http_cookie": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_cookie",
					a,
					sc.GetSContentTmp().SetAnyHTTPCookie,
				)
			},
		},
		// --- http_header ---
		"data.data_security.s_content.http_header": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_header",
					a,
					sc.GetSContentTmp().SetAnyHTTPHeader,
				)
			},
		},
		// --- http_method ---
		"data.data_security.s_content.http_method": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_method",
					a,
					sc.GetSContentTmp().SetAnyHTTPMethod,
				)
			},
		},
		// --- fast_pattern ---
		"data.data_security.s_content.fast_pattern": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.fast_pattern",
					a,
					sc.GetSContentTmp().SetAnyFastPattern,
				)
			},
		},
		// --- http_raw_uri ---
		"data.data_security.s_content.http_raw_uri": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_raw_uri",
					a,
					sc.GetSContentTmp().SetAnyHTTPRawURI,
				)
			},
		},
		// --- http_stat_msg ---
		"data.data_security.s_content.http_stat_msg": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_stat_msg",
					a,
					sc.GetSContentTmp().SetAnyHTTPStatMsg,
				)
			},
		},
		// --- http_stat_code ---
		"data.data_security.s_content.http_stat_code": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_stat_code",
					a,
					sc.GetSContentTmp().SetAnyHTTPStatCode,
				)
			},
		},
		// --- http_raw_cookie ---
		"data.data_security.s_content.http_raw_cookie": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_raw_cookie",
					a,
					sc.GetSContentTmp().SetAnyHTTPRawCookie,
				)
			},
		},
		// --- http_raw_header ---
		"data.data_security.s_content.http_raw_header": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_raw_header",
					a,
					sc.GetSContentTmp().SetAnyHTTPRawHeader,
				)
			},
		},
		// --- http_client_body ---
		"data.data_security.s_content.http_client_body": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_client_body",
					a,
					sc.GetSContentTmp().SetAnyHTTPClientBody,
				)
			},
		},
	}
}

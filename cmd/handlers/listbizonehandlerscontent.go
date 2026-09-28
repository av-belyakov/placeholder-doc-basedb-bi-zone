package handlers

// NewListBiZoneHandlerSContents начальный обработчик событий полей 'data.data_security.s_content.*'
func NewListBiZoneHandlerSContents(sc *SupportingStructureForSContentType) map[string][]func(any) error {
	return map[string][]func(any) error{
		// --- depth ---
		"data.data_security.s_content.depth:": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.depth:",
					a,
					sc.sContentTmp.SetAnyDepth,
				)
			},
		},
		// --- nocase ---
		"data.data_security.s_content.nocase": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.nocase",
					a,
					sc.sContentTmp.SetAnyNocase,
				)
			},
		},
		// --- content ---
		"data.data_security.s_content.content": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.content",
					a,
					sc.sContentTmp.SetAnyContent,
				)
			},
		},
		// --- offset ---
		"data.data_security.s_content.offset:": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.offset:",
					a,
					sc.sContentTmp.SetAnyOffset,
				)
			},
		},
		// --- within ---
		"data.data_security.s_content.within:": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.within:",
					a,
					sc.sContentTmp.SetAnyWithin,
				)
			},
		},
		// --- http_uri ---
		"data.data_security.s_content.http_uri": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_uri",
					a,
					sc.sContentTmp.SetAnyHTTPURI,
				)
			},
		},
		// --- rawbytes ---
		"data.data_security.s_content.rawbytes": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.rawbytes",
					a,
					sc.sContentTmp.SetAnyRawbytes,
				)
			},
		},
		// --- distance ---
		"data.data_security.s_content.distance:": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.distance:",
					a,
					sc.sContentTmp.SetAnyDistance,
				)
			},
		},
		// --- http_cookie ---
		"data.data_security.s_content.http_cookie": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_cookie",
					a,
					sc.sContentTmp.SetAnyHTTPCookie,
				)
			},
		},
		// --- http_header ---
		"data.data_security.s_content.http_header": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_header",
					a,
					sc.sContentTmp.SetAnyHTTPHeader,
				)
			},
		},
		// --- http_method ---
		"data.data_security.s_content.http_method": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_method",
					a,
					sc.sContentTmp.SetAnyHTTPMethod,
				)
			},
		},
		// --- fast_pattern ---
		"data.data_security.s_content.fast_pattern": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.fast_pattern",
					a,
					sc.sContentTmp.SetAnyFastPattern,
				)
			},
		},
		// --- http_raw_uri ---
		"data.data_security.s_content.http_raw_uri": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_raw_uri",
					a,
					sc.sContentTmp.SetAnyHTTPRawURI,
				)
			},
		},
		// --- http_stat_msg ---
		"data.data_security.s_content.http_stat_msg": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_stat_msg",
					a,
					sc.sContentTmp.SetAnyHTTPStatMsg,
				)
			},
		},
		// --- http_stat_code ---
		"data.data_security.s_content.http_stat_code": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_stat_code",
					a,
					sc.sContentTmp.SetAnyHTTPStatCode,
				)
			},
		},
		// --- http_raw_cookie ---
		"data.data_security.s_content.http_raw_cookie": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_raw_cookie",
					a,
					sc.sContentTmp.SetAnyHTTPRawCookie,
				)
			},
		},
		// --- http_raw_header ---
		"data.data_security.s_content.http_raw_header": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_raw_header",
					a,
					sc.sContentTmp.SetAnyHTTPRawHeader,
				)
			},
		},
		// --- http_client_body ---
		"data.data_security.s_content.http_client_body": {
			func(a any) error {
				return sc.HandlerValue(
					"data.data_security.s_content.http_client_body",
					a,
					sc.sContentTmp.SetAnyHTTPClientBody,
				)
			},
		},
	}
}

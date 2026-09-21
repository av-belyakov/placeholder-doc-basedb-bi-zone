package handlers

// NewListBiZoneHandlerSContent начальный обработчик событий полей 'data.data_security.s_content.*'
func NewListBiZoneHandlerSContent(ssc *SupportiveSContentType) map[string][]func(any) error {
	return map[string][]func(any) error{
		// --- depth ---
		"data.data_security.s_content.depth": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.depth",
					a,
					ssc.GetSContentTmp().SetAnyDepth,
				)
			},
		},
		// --- nocase ---
		"data.data_security.s_content.nocase": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.nocase",
					a,
					ssc.GetSContentTmp().SetAnyNocase,
				)
			},
		},
		// --- content ---
		"data.data_security.s_content.content": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.content",
					a,
					ssc.GetSContentTmp().SetAnyContent,
				)
			},
		},
		// --- offset ---
		"data.data_security.s_content.offset": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.offset",
					a,
					ssc.GetSContentTmp().SetAnyOffset,
				)
			},
		},
		// --- within ---
		"data.data_security.s_content.within": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.within",
					a,
					ssc.GetSContentTmp().SetAnyWithin,
				)
			},
		},
		// --- http_uri ---
		"data.data_security.s_content.http_uri": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_uri",
					a,
					ssc.GetSContentTmp().SetAnyHTTPURI,
				)
			},
		},
		// --- rawbytes ---
		"data.data_security.s_content.rawbytes": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.rawbytes",
					a,
					ssc.GetSContentTmp().SetAnyRawbytes,
				)
			},
		},
		// --- distance ---
		"data.data_security.s_content.distance": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.distance",
					a,
					ssc.GetSContentTmp().SetAnyDistance,
				)
			},
		},
		// --- http_cookie ---
		"data.data_security.s_content.http_cookie": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_cookie",
					a,
					ssc.GetSContentTmp().SetAnyHTTPCookie,
				)
			},
		},
		// --- http_header ---
		"data.data_security.s_content.http_header": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_header",
					a,
					ssc.GetSContentTmp().SetAnyHTTPHeader,
				)
			},
		},
		// --- http_method ---
		"data.data_security.s_content.http_method": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_method",
					a,
					ssc.GetSContentTmp().SetAnyHTTPMethod,
				)
			},
		},
		// --- fast_pattern ---
		"data.data_security.s_content.fast_pattern": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.fast_pattern",
					a,
					ssc.GetSContentTmp().SetAnyFastPattern,
				)
			},
		},
		// --- http_raw_uri ---
		"data.data_security.s_content.http_raw_uri": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_raw_uri",
					a,
					ssc.GetSContentTmp().SetAnyHTTPRawURI,
				)
			},
		},
		// --- http_stat_msg ---
		"data.data_security.s_content.http_stat_msg": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_stat_msg",
					a,
					ssc.GetSContentTmp().SetAnyHTTPStatMsg,
				)
			},
		},
		// --- http_stat_code ---
		"data.data_security.s_content.http_stat_code": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_stat_code",
					a,
					ssc.GetSContentTmp().SetAnyHTTPStatCode,
				)
			},
		},
		// --- http_raw_cookie ---
		"data.data_security.s_content.http_raw_cookie": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_raw_cookie",
					a,
					ssc.GetSContentTmp().SetAnyHTTPRawCookie,
				)
			},
		},
		// --- http_raw_header ---
		"data.data_security.s_content.http_raw_header": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_raw_header",
					a,
					ssc.GetSContentTmp().SetAnyHTTPRawHeader,
				)
			},
		},
		// --- http_client_body ---
		"data.data_security.s_content.http_client_body": {
			func(a any) error {
				return ssc.HandlerValue(
					"data.data_security.s_content.http_client_body",
					a,
					ssc.GetSContentTmp().SetAnyHTTPClientBody,
				)
			},
		},
	}
}

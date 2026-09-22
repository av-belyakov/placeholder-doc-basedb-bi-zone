package handlers

// NewListBiZoneHandlerDataSecurity начальный обработчик событий полей 'data.data_security.*'
func NewListBiZoneHandlerDataSecurity(sds *SupportingStructureForDataSecurityType) map[string][]func(any) error {
	return map[string][]func(any) error{
		// --- i_sid ---
		"data.data_security.i_sid": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.i_sid",
					a,
					sds.GetDataSecurityTmp().SetAnyISid,
				)
			},
		},
		// --- s_msg ---
		"data.data_security.s_msg": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.s_msg",
					a,
					sds.GetDataSecurityTmp().SetAnySMsg,
				)
			},
		},
		// --- i_rev ---
		"data.data_security.i_rev": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.i_rev",
					a,
					sds.GetDataSecurityTmp().SetAnyIRev,
				)
			},
		},
		// --- i_b_rec ---
		"data.data_security.i_b_rec": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.i_b_rec",
					a,
					sds.GetDataSecurityTmp().SetAnyIBRec,
				)
			},
		},
		// --- s_rule_body ---
		"data.data_security.s_rule_body": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.s_rule_body",
					a,
					sds.GetDataSecurityTmp().SetAnySRuleBody,
				)
			},
		},
		// --- s_classtype ---
		"data.data_security.s_classtype": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.s_classtype",
					a,
					sds.GetDataSecurityTmp().SetAnySClasstype,
				)
			},
		},
		// --- s_add_date ---
		"data.data_security.s_add_date": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.s_add_date",
					a,
					sds.GetDataSecurityTmp().SetAnySAddDate,
				)
			},
		},
		// --- i_priority ---
		"data.data_security.i_priority": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.i_priority",
					a,
					sds.GetDataSecurityTmp().SetAnyIPriority,
				)
			},
		},
		// --- s_source_name ---
		"data.data_security.s_source_name": {
			func(a any) error {
				return sds.HandlerValue(
					"data.data_security.s_source_name",
					a,
					sds.GetDataSecurityTmp().SetAnySSourceName,
				)
			},
		},
	}
}

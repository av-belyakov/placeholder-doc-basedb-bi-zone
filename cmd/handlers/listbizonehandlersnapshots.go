package handlers

// NewListBiZoneHandlerSnapshots обработчик для значений типа 'data.snapshots.*' основного объекта
func NewListBiZoneHandlerSnapshots(s *SupportingStructureForSnapshotsType) map[string][]func(any) error {
	return map[string][]func(any) error{
		//--- os ---
		"snapshots.os": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.os",
					a,
					s.snapshotTmp.SetAnyOS,
				)
			},
		},
		//--- fqdn ---
		"snapshots.fqdn": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.fqdn",
					a,
					s.snapshotTmp.SetAnyFqdn,
				)
			},
		},
		//--- title ---
		"snapshots.title": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.title",
					a,
					s.snapshotTmp.SetAnyTitle,
				)
			},
		},
		//--- domain ---
		"snapshots.domain": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.domain",
					a,
					s.snapshotTmp.SetAnyDomain,
				)
			},
		},
		//--- cmdb_id ---
		"snapshots.cmdb_id": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.cmdb_id",
					a,
					s.snapshotTmp.SetAnyCMDBID,
				)
			},
		},
		//--- os_type ---
		"snapshots.os_type": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.os_type",
					a,
					s.snapshotTmp.SetAnyOSType,
				)
			},
		},
		//--- hostname ---
		"snapshots.hostname": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.hostname",
					a,
					s.snapshotTmp.SetAnyHostname)
			},
		},
		//--- severity ---
		"snapshots.severity": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.severity",
					a,
					s.snapshotTmp.SetAnySeverity)
			},
		},
		//--- user_cmdb_name ---
		"snapshots.user_cmdb_name": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.user_cmdb_name",
					a,
					s.snapshotTmp.SetAnyUserCMDBName,
				)
			},
		},
		//--- user_cmdb_id ---
		"snapshots.user_cmdb_id": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.user_cmdb_id",
					a,
					s.snapshotTmp.SetAnyUserCmdbId,
				)
			},
		},
		//ниже работа со срезам содержащими простые типы
		//--- ip_addresses ---
		"snapshots.ip_addresses": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.ip_addresses",
					a,
					s.snapshotTmp.SetAnyIPAddresse,
				)
			}},
		//--- mac_addresses ---
		"snapshots.mac_addresses": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.mac_addresses",
					a,
					s.snapshotTmp.SetAnyMACAddresse,
				)
			},
		},
	}
}

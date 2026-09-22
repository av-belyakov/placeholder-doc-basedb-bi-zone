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
					s.GetSnapshotTmp().SetAnyOS,
				)
			},
		},
		//--- fqdn ---
		"snapshots.fqdn": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.fqdn",
					a,
					s.GetSnapshotTmp().SetAnyFqdn,
				)
			},
		},
		//--- title ---
		"snapshots.title": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.title",
					a,
					s.GetSnapshotTmp().SetAnyTitle,
				)
			},
		},
		//--- domain ---
		"snapshots.domain": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.domain",
					a,
					s.GetSnapshotTmp().SetAnyDomain,
				)
			},
		},
		//--- cmdb_id ---
		"snapshots.cmdb_id": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.cmdb_id",
					a,
					s.GetSnapshotTmp().SetAnyCMDBID,
				)
			},
		},
		//--- os_type ---
		"snapshots.os_type": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.os_type",
					a,
					s.GetSnapshotTmp().SetAnyOSType,
				)
			},
		},
		//--- hostname ---
		"snapshots.hostname": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.hostname",
					a,
					s.GetSnapshotTmp().SetAnyHostname)
			},
		},
		//--- severity ---
		"snapshots.severity": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.severity",
					a,
					s.GetSnapshotTmp().SetAnySeverity)
			},
		},
		//--- user_cmdb_name ---
		"snapshots.user_cmdb_name": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.user_cmdb_name",
					a,
					s.GetSnapshotTmp().SetAnyUserCMDBName,
				)
			},
		},
		//--- user_cmdb_id ---
		"snapshots.user_cmdb_id": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.user_cmdb_id",
					a,
					s.GetSnapshotTmp().SetAnyUserCmdbId,
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
					s.GetSnapshotTmp().SetAnyIPAddresse,
				)
			}},
		//--- mac_addresses ---
		"snapshots.mac_addresses": {
			func(a any) error {
				return s.HandlerValue(
					"snapshots.mac_addresses",
					a,
					s.GetSnapshotTmp().SetAnyMACAddresse,
				)
			},
		},
	}
}

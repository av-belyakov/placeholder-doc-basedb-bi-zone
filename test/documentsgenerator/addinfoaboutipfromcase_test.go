package documentsgenerator

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func TestGetIpAddressFromCaseType(t *testing.T) {
	description := `<p>{<br>\n  \"desc\": \"bz-test\",<br>\n  \"tags\": [<br>\n\"bz-test\"<br>\n  ],<br>
\n  \"agent\": 8030069,<br>\n  \"createdBy\": \"admin\",<br>\n  \"job_title\": \"bz-test\",<br>
\n  \"source_ip\": \"185.71.64.18\",<br>\n  \"target_ip\": \"185.71.64.18\",<br>
\n  \"siem2_link\": \"https://siem.siem-2.cloud.gcm/events/SIEM-9\",<br>\n  \"raw_payload\": {<br>
\n\"event\": {<br>\n  \"end\": 1791202081000,<br>\n  \"ext\": [<br>
\n{<br>\n  \"ip\": \"185.71.64.18\",<br>\n  \"city\": null,<br>\n  \"geoIso\": \"RU\",<br>
\n  \"geoCountry\": \"\\u0420\\u043e\\u0441\\u0441\\u0438\\u044f\"<br>\n}<br>\n  ],<br>
\n  \"home\": [<br>\n{<br>\n  \"ip\": \"251.1.161.50\"<br>\n}<br>\n  ],<br>\n  \"tags\": [<br>
\n\"bz-test\"<br>\n  ],<br>\n  \"ttps\": [<br>\n{<br>\n  \"name\": \"Steal or Forge Kerberos Tickets\",<br>
\n  \"tactic\": {<br>\n\"name\": \"Credential Access\",<br>\n\"externalId\": \"TA0006\"<br>\n  },<br>
\n  \"externalId\": \"T1558\"<br>\n}<br>\n  ],<br>\n  \"rules\": [<br>\n{<br>\n  \"id\": \"2\",<br>
\n  \"addedAt\": 1791202081000,<br>\n  \"message\": \"Zeek Weird.log\"<br>\n}<br>\n  ],<br>
\n  \"start\": 1791202081000,<br>\n  \"tasks\": [],<br>\n  \"title\": \"bz-test\",<br>\n  \"object\": {<br>
\n\"name\": \"\\u0413\\u0423\\u0421\\u041f\",<br>\n\"group\": {<br>
\n  \"name\": \"\\u0413\\u0426\\u041c (\\u0433.\\u041c\\u043e\\u0441\\u043a\\u0432\\u0430)\"<br>\n},<br>
\n\"tenant\": {<br>\n  \"name\": \"\\u0413\\u0423\\u0421\\u041f\"<br>\n},<br>\n\"sensors\": [<br>
\n  {<br>\n\"id\": \"8030069\",<br>\n\"name\": \"\\u0413\\u0423\\u0421\\u041f, \\u041c\\u043e\\u0441\\u043a\\u0432\\u0430\"<br>
\n  }<br>\n]<br>\n  },<br>\n  \"status\": {<br>\n\"name\": \"new\"<br>\n  },<br>
\n  \"eventUrl\": \"https://siem.siem-2.cloud.gcm/events/SIEM-9\",<br>\n  \"lastTime\": 1791202081000,<br>
\n  \"firstTime\": 1791202081000,<br>\n  \"assignedTo\": null,<br>\n  \"attributes\": [],<br>
\n  \"nccciClass\": {<br>\n\"name\": \"\\u0421\\u043e\\u0431\\u044b\\u0442\\u0438\\u0435 \\u043d\\u0435 \\u0441\\u0432\\u044f\\u0437\\u0430\\u043d\\u043e \\u0441 \\u043a\\u043e\\u043c\\u043f\\u044c\\u044e\\u0442\\u0435\\u0440\\u043d\\u043e\\u0439 \\u0430\\u0442\\u0430\\u043a\\u043e\\u0439\"<br>\n  },<br>\n  \"withImpact\": null,<br>\n  \"attachments\": [],<br>\n  \"description\": \"bz-test\",<br>\n  \"recommendation\": {<br>\n\"name\": \"\\u041f\\u0440\\u043e\\u0432\\u0435\\u0441\\u0442\\u0438 \\u0430\\u043d\\u0430\\u043b\\u0438\\u0437 \\u0441\\u0435\\u0442\\u0435\\u0432\\u043e\\u0433\\u043e \\u0442\\u0440\\u0430\\u0444\\u0438\\u043a\\u0430\"<br>\n  },<br>\n  \"resolutionStatusId\": null<br>\n},<br>\n\"eventId\": 9,<br>\n\"outboxId\": 9,<br>\n\"publicId\": \"SIEM-9\",<br>\n\"callbacks\": [<br>\n  {<br>\n\"url\": \"https://siem.siem-2.cloud.gcm/api/v2/events/callbacks/cb_6f53a4ec1b23488db2ec4b92ba71a738\",<br>\n\"targetIrpId\": 2<br>\n  }<br>\n],<br>\n\"createdBy\": {<br>\n  \"username\": \"admin\"<br>\n},<br>\n\"createMode\": \"withAlerts\",<br>\n\"targetIrpIds\": [<br>\n  2<br>\n],<br>\n\"irpCreateMode\": \"case\"<br>\n  },<br>\n  \"severity_id\": 2,<br>\n  \"ext_endpoint\": \"185.71.64.18\",<br>\n  \"siem2_irp_id\": 2,<br>\n  \"data_security\": [<br>\n{<br>\n  \"gid\": 1,<br>\n  \"msg\": \"Zeek Weird.log\",<br>\n  \"rev\": 1,<br>\n  \"sid\": 2,<br>\n  \"classtype\": null,<br>\n  \"rule_body\": null<br>\n}<br>\n  ],<br>\n  \"home_endpoint\": \"185.71.64.18\",<br>\n  \"last_seen_time\": \"2026-10-05T12:08:01Z\",<br>\n  \"first_seen_time\": \"2026-10-05T12:08:01Z\",<br>\n  \"detection_pattern\": [<br>\n2<br>\n  ],<br>\n  \"unmapped_event_card\": \"SIEM-9\",<br>\n  \"unmapped_agent_array\": [<br>\n8030069<br>\n  ],<br>\n  \"unmapped_sensor_name\": \"\\u0413\\u0423\\u0421\\u041f, \\u041c\\u043e\\u0441\\u043a\\u0432\\u0430\",<br>\n  \"metadata_product_name\": \"SIEM2\",<br>\n  \"unmapped_sensor_name_array\": [<br>\n\"\\u0413\\u0423\\u0421\\u041f, \\u041c\\u043e\\u0441\\u043a\\u0432\\u0430\"<br>
\n  ],<br>\n  \"unmapped_ext_endpoint_array\": [<br>\n\"185.71.64.18\"<br>\n  ],<br>
\n  \"unmapped_home_endpoint_array\": [<br>\n\"185.71.64.18\"<br>\n  ]<br>\n}</p>`

	reg := regexp.MustCompile(`\b(?:(?:25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.){3}(?:25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\b`)
	ipaddresses := reg.FindAllStringSubmatch(description, -1)
	ips := make([]string, 0, len(ipaddresses))

	for _, ip := range ipaddresses {
		ips = append(ips, ip[0])
	}

	ips = supportingfunctions.SliceUniqueValues(ips)

	for k, ip := range ips {
		fmt.Printf("%d. %s\n", k, ip)
	}
}

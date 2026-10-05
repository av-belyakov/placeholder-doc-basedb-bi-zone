package documentsgenerator_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitAretString(t *testing.T) {
	result := []string(nil)
	alertsStr := []string{
		"alert tcp $EXTERNAL_NET $HTTP_PORTS -> $HOME_NET any (msg:\"ETPRO WEB_CLIENT Microsoft Internet Explorer JPEG Rendering Buffer Overflow\"; flow:to_client,established; content:\"image/jpeg\"; nocase; content:\"|0D 0A 0D 0A FF D8 FF|\"; content:\"|FF DA 00 0C 03|\"; distance:0; pcre:!\"/^..(\\x02.\\x03|\\x04.\\x05)/sR\"; reference:cve,CVE-2005-1988; reference:bugtraq,14282; classtype:attempted-user; sid:2800778; rev:3; metadata:affected_product Web_Browsers, affected_product Web_Browser_Plugins, attack_target Client_Endpoint, created_at 2010_09_26, deployment Perimeter, confidence High, signature_severity Major, tag Web_Client_Attacks, updated_at 2010_09_30;)",
		"alert tcp $SMTP_SERVERS [25,587] -> $EXTERNAL_NET any (msg:\"Potential SMTP Brute-Force attempt (T1110)\"; content: \"535 \"; detection_filter:track by_src, count 10, seconds 3600; content: \"POST\"; classtype:unsuccessful-user; sid:90000086; rev:2; metadata: author RCM_NVS, created 2024_08_14;)",
		"alert tcp any any -> $HOME_NET any (msg:\"FOX-SRT - Exploit - Possible Apache Log4J RCE Request Observed (CVE-2021-44228)\"; flow:established, to_server; content:\"${jndi:ldap://\"; fast_pattern:only; flowbits:set, fox.apachelog4j.rce; threshold:type limit, track by_dst, count 1, seconds 3600; classtype:web-application-attack; priority:3; reference:url, www.lunasec.io/docs/blog/log4j-zero-day/; metadata:CVE 2021-44228; metadata:created_at 2021-12-10; metadata:ids suricata; sid:21003726; rev:1;)\n",
	}

	rgx, err := regexp.Compile(`content:\s?"([^"]*)";`)
	assert.NoError(t, err)

	for _, alertStr := range alertsStr {
		matchFound := rgx.FindAllString(alertStr, -1)
		result = append(result, matchFound...)
	}

	fmt.Println("Найдено совпадений:")
	for k, v := range result {
		str := strings.TrimSuffix(strings.TrimPrefix(v, "content:\""), "\";")
		str = strings.TrimSuffix(strings.TrimPrefix(v, "content: \""), "\";")
		fmt.Printf("%d. %s\n", k, strings.ReplaceAll(str, "\"", ""))
	}
}

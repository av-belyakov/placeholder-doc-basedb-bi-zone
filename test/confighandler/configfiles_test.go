package confighandler_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
)

func TestConfigFileHandler(t *testing.T) {
	const (
		KAFKA_SSLPASSWD         = "687df-343rfs-1233f-aasq1"
		DATABASESTORAGE_PASSWD  = "f990-gggr-02jfg-fww2"
		DATABASEWRITELOG_PASSWD = "cjis8w-dff0w0-fy2y3"
	)

	var (
		listTesting []TestOptions
		testOptions TestOptions

		cfg *confighandler.Config
		err error
	)

	unsetAllEnviromentEnvAny()

	os.Setenv("GO_PHDOCBASEDBBZ_MAIN", "test")
	os.Setenv("GO_PHDOCBASEDBBZ_KSSLPASSWORD", KAFKA_SSLPASSWD)
	os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD", DATABASESTORAGE_PASSWD)
	os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD", DATABASEWRITELOG_PASSWD)

	// --- Общие настройки (из config.yml) ---
	testOptions = TestOptions{
		name: "Общие настройки (чтение файла 'config.yml')",
		function: func() {
			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.NetworkHost},
					expectedParameters: TestTypeElements{valueString: "192.168.9.45"},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetCommon().Zabbix.NetworkPort},
					expectedParameters: TestTypeElements{valueInt: 10051},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.ZabbixHost},
					expectedParameters: TestTypeElements{valueString: "test-uchet-db.cloud.gcm"},
				},
				// для отслеживания ошибок
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[0].EventType},
					expectedParameters: TestTypeElements{valueString: "error"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[0].ZabbixKey},
					expectedParameters: TestTypeElements{valueString: "placeholder_db_bizone.error"},
				},
				{
					inputParameters:    TestTypeElements{valueBool: cfg.GetCommon().Zabbix.EventTypes[0].IsTransmit},
					expectedParameters: TestTypeElements{valueBool: true},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetCommon().Zabbix.EventTypes[0].Handshake.TimeInterval},
					expectedParameters: TestTypeElements{valueInt: 0},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[0].Handshake.Message},
					expectedParameters: TestTypeElements{valueString: ""},
				},
				// для информационных сообщений о выполненной работе
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[1].EventType},
					expectedParameters: TestTypeElements{valueString: "info"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[1].ZabbixKey},
					expectedParameters: TestTypeElements{valueString: "placeholder_db_bizone.info"},
				},
				{
					inputParameters:    TestTypeElements{valueBool: cfg.GetCommon().Zabbix.EventTypes[1].IsTransmit},
					expectedParameters: TestTypeElements{valueBool: true},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetCommon().Zabbix.EventTypes[1].Handshake.TimeInterval},
					expectedParameters: TestTypeElements{valueInt: 0},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[1].Handshake.Message},
					expectedParameters: TestTypeElements{valueString: "I'm still alive"},
				},
				// для регулярного отстукивания что модуль еще работает
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[2].EventType},
					expectedParameters: TestTypeElements{valueString: "handshake"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[2].ZabbixKey},
					expectedParameters: TestTypeElements{valueString: "placeholder_db_bizone.handshake"},
				},
				{
					inputParameters:    TestTypeElements{valueBool: cfg.GetCommon().Zabbix.EventTypes[2].IsTransmit},
					expectedParameters: TestTypeElements{valueBool: true},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetCommon().Zabbix.EventTypes[2].Handshake.TimeInterval},
					expectedParameters: TestTypeElements{valueInt: 1},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().Zabbix.EventTypes[2].Handshake.Message},
					expectedParameters: TestTypeElements{valueString: "0"},
				},
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	testOptions = TestOptions{
		name: "Общие настройки каждого файла (чтение файла config_test.yml)",
		function: func() {
			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetCommon().RegionalObject},
					expectedParameters: TestTypeElements{valueString: "gcm-test"},
				},
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки NATS ---
	testOptions = TestOptions{
		name: "Настройки NATS (чтение файла config_test.yml)",
		function: func() {
			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetNATS().Host},
					expectedParameters: TestTypeElements{valueString: "192.168.9.208"},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetNATS().Port},
					expectedParameters: TestTypeElements{valueInt: 4222},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetNATS().CacheTTL},
					expectedParameters: TestTypeElements{valueInt: 3600},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetNATS().Subscriptions["some"]},
					expectedParameters: TestTypeElements{valueString: "somesubscription"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetNATS().EnrichingQueries["get_geoip_info"]},
					expectedParameters: TestTypeElements{valueString: "object.geoip-request.test"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetNATS().EnrichingQueries["get_sensor_info"]},
					expectedParameters: TestTypeElements{valueString: "object.sensor-info-request.test"},
				},
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки Kafka ---
	testOptions = TestOptions{
		name: "Настройки Kafka (чтение файла config_test.yml)",
		function: func() {
			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().Host},
					expectedParameters: TestTypeElements{valueString: "192.168.9.73"},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetKafka().Port},
					expectedParameters: TestTypeElements{valueInt: 30932},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetKafka().CacheTTL},
					expectedParameters: TestTypeElements{valueInt: 3600},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().Topics["alerts"]},
					expectedParameters: TestTypeElements{valueString: "object.topicalerttype.test"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().Topics["soar-alerts"]},
					expectedParameters: TestTypeElements{valueString: "object.topicsoaralertstype.test"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().AuthType},
					expectedParameters: TestTypeElements{valueString: "sasl-ssl"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SASLMechanism},
					expectedParameters: TestTypeElements{valueString: "SCRAM-SHA-512"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLUsername},
					expectedParameters: TestTypeElements{valueString: "siem2"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLCaFile},
					expectedParameters: TestTypeElements{valueString: "/secrets/truststore.jks"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLCertFile},
					expectedParameters: TestTypeElements{valueString: "/secrets/ca.crt"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLKeyFile},
					expectedParameters: TestTypeElements{valueString: "/secrets/ca.crt"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLPassword},
					expectedParameters: TestTypeElements{valueString: KAFKA_SSLPASSWD},
				},
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки DATABASESTORAGE ---
	testOptions = TestOptions{
		name: "Настройки DATABASESTORAGE (чтение файла config_test.yml)",
		function: func() {
			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().Host},
					expectedParameters: TestTypeElements{valueString: "192.168.9.208"},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetStorageDB().Port},
					expectedParameters: TestTypeElements{valueInt: 9200},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().User},
					expectedParameters: TestTypeElements{valueString: "placeholder-docbasedb-bizone"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().NameDB},
					expectedParameters: TestTypeElements{valueString: ""},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().Storage["alerts"]},
					expectedParameters: TestTypeElements{valueString: "module_placeholderdb_bizone_alerts.test"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().Storage["soar-alerts"]},
					expectedParameters: TestTypeElements{valueString: "module_placeholderdb_bizone_soar-alerts.test"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().Passwd},
					expectedParameters: TestTypeElements{valueString: DATABASESTORAGE_PASSWD},
				},
			}
		}}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки DATABASEWRITELOG ---
	testOptions = TestOptions{
		name: "Настройки DATABASEWRITELOG (чтение файла config_test.yml)",
		function: func() {
			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().Host},
					expectedParameters: TestTypeElements{valueString: "192.168.9.208"},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetLogDB().Port},
					expectedParameters: TestTypeElements{valueInt: 9200},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().NameDB},
					expectedParameters: TestTypeElements{valueString: ""},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().StorageNameDB},
					expectedParameters: TestTypeElements{valueString: "db-placeholder_docbasedb-bizone_log"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().User},
					expectedParameters: TestTypeElements{valueString: "user-placeholder-docbasedb-bizone_log"},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().Passwd},
					expectedParameters: TestTypeElements{valueString: DATABASEWRITELOG_PASSWD},
				},
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки DEBUGSERVER ---
	testOptions = TestOptions{
		name: "Настройки DEBUGSERVER (чтение файла config_test.yml)",
		function: func() {
			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueBool: cfg.GetDebugServer().Enable},
					expectedParameters: TestTypeElements{valueBool: true},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetDebugServer().Host},
					expectedParameters: TestTypeElements{valueString: "localhost"},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetDebugServer().Port},
					expectedParameters: TestTypeElements{valueInt: 6565},
				},
			}
		}}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки NATS (через переменные окружения) ---
	testOptions = TestOptions{
		name: "Настройки NATS (через переменные окружения)",
		function: func() {
			const (
				HOST          = "nats.cloud.gcm.test.test"
				PORT          = 4545
				CACHE_TTL     = 10
				SUBSCRIPTIONS = "send:sender.alert;receive:receiver.alert"
				ENRICHINGQUER = "get_geoip_info:object.geoip-request;get_sensor_info:object.sensor-info-request"
			)

			os.Setenv("GO_PHDOCBASEDBBZ_NHOST", HOST)
			os.Setenv("GO_PHDOCBASEDBBZ_NPORT", strconv.Itoa(PORT))
			os.Setenv("GO_PHDOCBASEDBBZ_NCACHETTL", strconv.Itoa(CACHE_TTL))
			os.Setenv("GO_PHDOCBASEDBBZ_NSUBSCRIPTIONS", SUBSCRIPTIONS)
			os.Setenv("GO_PHDOCBASEDBBZ_NENRICHINGQUER", ENRICHINGQUER)

			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetNATS().Host},
					expectedParameters: TestTypeElements{valueString: HOST},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetNATS().Port},
					expectedParameters: TestTypeElements{valueInt: PORT},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetNATS().CacheTTL},
					expectedParameters: TestTypeElements{valueInt: CACHE_TTL},
				},
			}

			for k, v := range []string{SUBSCRIPTIONS, ENRICHINGQUER} {
				if k == 0 {
					if !strings.Contains(v, ";") {
						if tmp := strings.Split(v, ":"); len(tmp) == 2 {
							testOptions.items = append(testOptions.items, TestParametrs{
								inputParameters:    TestTypeElements{valueString: cfg.GetNATS().Subscriptions[tmp[0]]},
								expectedParameters: TestTypeElements{valueString: tmp[1]},
							})
						}
					} else {
						for sl := range strings.SplitSeq(v, ";") {
							if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
								testOptions.items = append(testOptions.items, TestParametrs{
									inputParameters:    TestTypeElements{valueString: cfg.GetNATS().Subscriptions[tmp[0]]},
									expectedParameters: TestTypeElements{valueString: tmp[1]},
								})
							}
						}
					}

					continue
				}

				if !strings.Contains(v, ";") {
					if tmp := strings.Split(v, ":"); len(tmp) == 2 {
						testOptions.items = append(testOptions.items, TestParametrs{
							inputParameters:    TestTypeElements{valueString: cfg.GetNATS().EnrichingQueries[tmp[0]]},
							expectedParameters: TestTypeElements{valueString: tmp[1]},
						})
					}
				} else {
					for sl := range strings.SplitSeq(v, ";") {
						if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
							testOptions.items = append(testOptions.items, TestParametrs{
								inputParameters:    TestTypeElements{valueString: cfg.GetNATS().EnrichingQueries[tmp[0]]},
								expectedParameters: TestTypeElements{valueString: tmp[1]},
							})
						}
					}
				}
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки Kafka (через переменные окружения) ---
	testOptions = TestOptions{
		name: "Настройки Kafka (через переменные окружения)",
		function: func() {
			const (
				HOST           = "45.6.36.1"
				PORT           = 1180
				CACHE_TTL      = 35
				TOPICS         = "topiconw:phdocbasedbbz;topictwo:phdocbaseddmz2"
				AUTH_TYPE      = "none"
				SASL_MECHANISM = "PLAIN"
				SSL_USERNAME   = "any-login"
				SSL_PASSWORD   = "pass-here!@#"
				SSL_CA_FILE    = "/secrets/anycafile.jks"
				SSL_CERT_FILE  = "/secrets/anycertfile.jks"
				SSL_KEY_FILE   = "/secrets/anykeyfile.jks"
			)

			os.Setenv("GO_PHDOCBASEDBBZ_KHOST", HOST)
			os.Setenv("GO_PHDOCBASEDBBZ_KPORT", strconv.Itoa(PORT))
			os.Setenv("GO_PHDOCBASEDBBZ_KCACHETTL", strconv.Itoa(CACHE_TTL))
			os.Setenv("GO_PHDOCBASEDBBZ_KTOPICS", TOPICS)
			os.Setenv("GO_PHDOCBASEDBBZ_KAUTHTYPE", AUTH_TYPE)
			os.Setenv("GO_PHDOCBASEDBBZ_KSASLMECHANISM", SASL_MECHANISM)
			os.Setenv("GO_PHDOCBASEDBBZ_KSSLUSERNAME", SSL_USERNAME)
			os.Setenv("GO_PHDOCBASEDBBZ_KSSLPASSWORD", SSL_PASSWORD)
			os.Setenv("GO_PHDOCBASEDBBZ_KSSLCAFILE", SSL_CA_FILE)
			os.Setenv("GO_PHDOCBASEDBBZ_KCERTFILE", SSL_CERT_FILE)
			os.Setenv("GO_PHDOCBASEDBBZ_KKEYFILE", SSL_KEY_FILE)

			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().Host},
					expectedParameters: TestTypeElements{valueString: HOST},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetKafka().Port},
					expectedParameters: TestTypeElements{valueInt: PORT},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetKafka().CacheTTL},
					expectedParameters: TestTypeElements{valueInt: CACHE_TTL},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().AuthType},
					expectedParameters: TestTypeElements{valueString: AUTH_TYPE},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SASLMechanism},
					expectedParameters: TestTypeElements{valueString: SASL_MECHANISM},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLUsername},
					expectedParameters: TestTypeElements{valueString: SSL_USERNAME},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLPassword},
					expectedParameters: TestTypeElements{valueString: SSL_PASSWORD},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLCaFile},
					expectedParameters: TestTypeElements{valueString: SSL_CA_FILE},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLCertFile},
					expectedParameters: TestTypeElements{valueString: SSL_CERT_FILE},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetKafka().SSLKeyFile},
					expectedParameters: TestTypeElements{valueString: SSL_KEY_FILE},
				},
			}

			if !strings.Contains(TOPICS, ";") {
				if tmp := strings.Split(TOPICS, ":"); len(tmp) == 2 {
					testOptions.items = append(testOptions.items, TestParametrs{
						inputParameters:    TestTypeElements{valueString: cfg.GetKafka().Topics[tmp[0]]},
						expectedParameters: TestTypeElements{valueString: tmp[1]},
					})
				}
			} else {
				for sl := range strings.SplitSeq(TOPICS, ";") {
					if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
						testOptions.items = append(testOptions.items, TestParametrs{
							inputParameters:    TestTypeElements{valueString: cfg.GetKafka().Topics[tmp[0]]},
							expectedParameters: TestTypeElements{valueString: tmp[1]},
						})
					}
				}
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки DATABASESTOREGE (через переменные окружения) ---
	testOptions = TestOptions{
		name: "Настройки DATABASESTOREGE (через переменные окружения)",
		function: func() {
			const (
				HOST   = "112.63.23.59"
				PORT   = 8074
				NAME   = "phdocbasedb"
				USER   = "anybduser"
				PASSWD = "AnyBD@User"
				TABLES = "table.one:table-storage-one;table.two:my-table-storage-two"
			)

			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEHOST", HOST)
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEPORT", strconv.Itoa(PORT))
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGENAME", NAME)
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEUSER", USER)
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD", PASSWD)
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGETABLES", TABLES)

			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().Host},
					expectedParameters: TestTypeElements{valueString: HOST},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetStorageDB().Port},
					expectedParameters: TestTypeElements{valueInt: PORT},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().NameDB},
					expectedParameters: TestTypeElements{valueString: NAME},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().User},
					expectedParameters: TestTypeElements{valueString: USER},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().Passwd},
					expectedParameters: TestTypeElements{valueString: PASSWD},
				},
			}

			if !strings.Contains(TABLES, ";") {
				if tmp := strings.Split(TABLES, ":"); len(tmp) == 2 {
					testOptions.items = append(testOptions.items, TestParametrs{
						inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().Storage[tmp[0]]},
						expectedParameters: TestTypeElements{valueString: tmp[1]},
					})
				}
			} else {
				for sl := range strings.SplitSeq(TABLES, ";") {
					if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
						testOptions.items = append(testOptions.items, TestParametrs{
							inputParameters:    TestTypeElements{valueString: cfg.GetStorageDB().Storage[tmp[0]]},
							expectedParameters: TestTypeElements{valueString: tmp[1]},
						})
					}
				}
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	// --- Настройки DATABASEWRITELOG (через переменные окружения) ---
	testOptions = TestOptions{
		name: "Настройки DATABASEWRITELOG (через переменные окружения)",
		function: func() {
			const (
				HOST        = "112.100.100.159"
				PORT        = 8099
				NAME        = "phdocbasedblog"
				USER        = "anybduserlog"
				PASSWD      = "AnyBD@UserLoGG"
				STORAGENAME = "storage_logs"
			)

			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGHOST", HOST)
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGPORT", strconv.Itoa(PORT))
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGNAME", NAME)
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGUSER", USER)
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD", PASSWD)
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME", STORAGENAME)

			cfg, err = confighandler.New(constants.Root_Dir)
			testOptions.err = err
			testOptions.items = []TestParametrs{
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().Host},
					expectedParameters: TestTypeElements{valueString: HOST},
				},
				{
					inputParameters:    TestTypeElements{valueInt: cfg.GetLogDB().Port},
					expectedParameters: TestTypeElements{valueInt: PORT},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().NameDB},
					expectedParameters: TestTypeElements{valueString: NAME},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().User},
					expectedParameters: TestTypeElements{valueString: USER},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().Passwd},
					expectedParameters: TestTypeElements{valueString: PASSWD},
				},
				{
					inputParameters:    TestTypeElements{valueString: cfg.GetLogDB().StorageNameDB},
					expectedParameters: TestTypeElements{valueString: STORAGENAME},
				},
			}
		},
	}
	testOptions.function()
	listTesting = append(listTesting, testOptions)

	//--------------------------
	// --- Выполнение тестов ---
	for testNum, tt := range listTesting {
		t.Run(fmt.Sprintf("Тест %d. %s", testNum+1, tt.name), func(t *testing.T) {
			assert.NoError(t, tt.err)

			for _, v := range tt.items {
				/*fmt.Printf(
				`%d.
				expectedParameters.valueInt:'%d' = '%d':valueInt.inputParameters
				expectedParameters.valueString:'%s' = '%s':valueString.inputParameters`,
				k+1,
				v.expectedParameters.valueInt, v.inputParameters.valueInt,
				v.expectedParameters.valueString, v.inputParameters.valueString)*/

				assert.Equal(t, v.expectedParameters.valueInt, v.inputParameters.valueInt)
				assert.Equal(t, v.expectedParameters.valueString, v.inputParameters.valueString)
			}
		})
	}

	//t.Run("", func(t *testing.T) {})

	t.Cleanup(func() {
		unsetAllEnviromentEnvAny()
	})
}

func unsetAllEnviromentEnvAny() {
	os.Unsetenv("GO_PHDOCBASEDBBZ_MAIN")

	//настройка наименования регионального объекта
	os.Unsetenv("GO_PHDOCBASEDBBZ_REGIONALOBJECT")

	//настройки NATS
	os.Unsetenv("GO_PHDOCBASEDBBZ_NHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_NPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_NCACHETTL")
	os.Unsetenv("GO_PHDOCBASEDBBZ_NSUBLISTENER")
	os.Unsetenv("GO_PHDOCBASEDBBZ_NENRICHINGQUER")

	//настройки Kafka
	os.Unsetenv("GO_PHDOCBASEDBBZ_KHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KTOPICS")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KCACHETTL")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KLOGIN")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KPASSWD")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KCERTPATH")
	os.Unsetenv("GO_PHDOCBASEDBBZ_KTRUSTSTOREPATH")

	// Настройки доступа к БД в которую будут записыватся alert и case
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGENAME")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEUSER")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGETABLES")

	//настройки доступа к БД в которую будут записыватся логи
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGNAME")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGUSER")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME")
}

type TestTypeElements struct {
	valueString string
	valueInt    int
	valueBool   bool
}

type TestParametrs struct {
	inputParameters    TestTypeElements
	expectedParameters TestTypeElements
}

type TestOptions struct {
	items    []TestParametrs
	function func()
	err      error
	name     string
}

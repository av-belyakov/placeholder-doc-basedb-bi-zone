package confighandler_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
)

const Root_Dir = "placeholder-doc-basedb-bi-zone"

var (
	conf *confighandler.Config

	err error
)

func TestMain(m *testing.M) {
	cleanEnv()

	//загружаем ключи и пароли
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	os.Setenv("GO_PHDOCBASEDBBZ_MAIN", "test")

	conf, err = confighandler.New(Root_Dir)
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(m.Run())
}

func TestConfigHandler(t *testing.T) {
	t.Run("Тест чтения конфигурационного файла", func(t *testing.T) {
		t.Run("Тест 1. Проверка настройки NATS из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.GetNATS().Host, "192.168.9.208")
			assert.Equal(t, conf.GetNATS().Port, 4222)
			assert.Equal(t, conf.GetNATS().CacheTTL, 3600)

			listRequests := map[string]string{
				"get_geoip_info":  "object.geoip-request.test",
				"get_sensor_info": "object.sensor-info-request.test",
			}
			for k, v := range listRequests {
				s, ok := conf.GetNATS().EnrichingQueries[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}

			listSubscriptions := map[string]string{
				"some": "somesubscription",
			}
			for k, v := range listSubscriptions {
				s, ok := conf.GetNATS().Subscriptions[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}

		})

		t.Run("Тест 2. Проверка настройки имени регионального объекта, полученного из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.Common.RegionalObject, "gcm-test")
		})

		t.Run("Тест 3. Проверка настройки KAFKA из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.GetKafka().Host, "localhost")
			assert.Equal(t, conf.GetKafka().Port, 9092)
			assert.Equal(t, conf.GetKafka().CacheTTL, 3600)

			listTopics := map[string]string{
				"alerts":      "object.topicalerttype.test",
				"soar-alerts": "object.topicsoaralertstype.test",
			}
			for k, v := range listTopics {
				s, ok := conf.GetKafka().Topics[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 4. Проверка настройки DATABASESTORAGE из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.GetStorageDB().Host, "192.168.9.208")
			assert.Equal(t, conf.GetStorageDB().Port, 9200)
			assert.Equal(t, conf.GetStorageDB().User, "placeholder-docbasedb-bizone")
			assert.Equal(t, conf.GetStorageDB().Passwd, os.Getenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD"))
			assert.Equal(t, conf.GetStorageDB().NameDB, "")

			list := map[string]string{
				"alerts":      "module_placeholderdb_bizone_alerts.test",
				"soar-alerts": "module_placeholderdb_bizone_soar-alerts.test",
			}
			for k, v := range list {
				s, ok := conf.GetStorageDB().Storage[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 5. Проверка настройки DATABASEWRITELOG из файла config_test.yml", func(t *testing.T) {
			assert.Equal(t, conf.GetLogDB().Host, "192.168.9.208")
			assert.Equal(t, conf.GetLogDB().Port, 9200)
			assert.Equal(t, conf.GetLogDB().User, "placeholder-docbasedb-bizone")
			assert.Equal(t, conf.GetLogDB().Passwd, os.Getenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD"))
			assert.Equal(t, conf.GetLogDB().NameDB, "")
			assert.Equal(t, conf.GetLogDB().StorageNameDB, "placeholder_docbasedb-bizone")
		})

		t.Run("Тест 6. Проверка настроек сервера отладки", func(t *testing.T) {
			assert.True(t, conf.DebugServer.Enable)
			assert.Equal(t, conf.DebugServer.Host, "localhost")
			assert.Equal(t, conf.DebugServer.Port, int(6565))
		})
	})

	t.Run("Тест чтения переменных окружения", func(t *testing.T) {
		t.Run("Тест 1. Проверка настройки NATS", func(t *testing.T) {
			os.Setenv("GO_PHDOCBASEDBBZ_NHOST", "127.0.0.1")
			os.Setenv("GO_PHDOCBASEDBBZ_NPORT", "4242")
			os.Setenv("GO_PHDOCBASEDBBZ_NCACHETTL", "650")
			os.Setenv("GO_PHDOCBASEDBBZ_NSUBLISTENER", "alert:object.env.alerttype.test;case:object.env.casetype.test")
			os.Setenv("GO_PHDOCBASEDBBZ_NENRICHINGQUER", "get_geoip_info:object.geoipreq;get_sensor_info:object.sensorinforeq")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetNATS().Host, "127.0.0.1")
			assert.Equal(t, conf.GetNATS().Port, 4242)
			assert.Equal(t, conf.GetNATS().CacheTTL, 650)

			listEnrichingQueries := map[string]string{
				"get_geoip_info":  "object.geoipreq",
				"get_sensor_info": "object.sensorinforeq",
			}
			for k, v := range listEnrichingQueries {
				s, ok := conf.GetNATS().EnrichingQueries[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}

			listSubscriptions := map[string]string{
				"alert": "object.env.alerttype.test",
				"case":  "object.env.casetype.test",
			}
			for k, v := range listSubscriptions {
				s, ok := conf.GetNATS().Subscriptions[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 2. Проверка настройки Kafka", func(t *testing.T) {
			os.Setenv("GO_PHDOCBASEDBBZ_KHOST", "127.0.0.1")
			os.Setenv("GO_PHDOCBASEDBBZ_KPORT", "8977")
			os.Setenv("GO_PHDOCBASEDBBZ_KTOPICS", "alerts:new_topic_for_alerts;soma-ealerts:some_new_alert_topic")
			os.Setenv("GO_PHDOCBASEDBBZ_KCACHETTL", "4800")
			os.Setenv("GO_PHDOCBASEDBBZ_KLOGIN")
			os.Setenv("GO_PHDOCBASEDBBZ_KPASSWD")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetKafka().Host, "127.0.0.1")
			assert.Equal(t, conf.GetKafka().Port, 8977)
			assert.Equal(t, conf.GetKafka().CacheTTL, 4800)

			listTopics := map[string]string{
				"alerts":       "new_topic_for_alerts",
				"soma-ealerts": "some_new_alert_topic",
			}
			for k, v := range listTopics {
				s, ok := conf.GetKafka().Topics[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 3. Проверка настройки DATABASESTORAGE", func(t *testing.T) {
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEHOST", "examle.database.cm")
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEPORT", "9559")
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGENAME", "any_name")
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEUSER", "any_user")
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD", "my_new_passwd")
			os.Setenv("GO_PHDOCBASEDBBZ_DBSTORAGEN", "alert:test.env.module_placeholderdb_alert;case:test.env.module_placeholderdb_case")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetStorageDB().Host, "examle.database.cm")
			assert.Equal(t, conf.GetStorageDB().Port, 9559)
			assert.Equal(t, conf.GetStorageDB().NameDB, "any_name")
			assert.Equal(t, conf.GetStorageDB().User, "any_user")
			assert.Equal(t, conf.GetStorageDB().Passwd, os.Getenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD"))

			list := map[string]string{
				"alert": "test.env.module_placeholderdb_alert",
				"case":  "test.env.module_placeholderdb_case",
			}
			for k, v := range list {
				s, ok := conf.GetStorageDB().Storage[k]
				assert.True(t, ok)
				assert.Equal(t, s, v)
			}
		})

		t.Run("Тест 4. Проверка настройки DATABASEWRITELOG", func(t *testing.T) {
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGHOST", "domaniname.database.cm")
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGPORT", "8989")
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGUSER", "somebody_user")
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGNAME", "any_name_db")
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD", "your_passwd")
			os.Setenv("GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME", "log_storage")

			conf, err := confighandler.New(Root_Dir)
			assert.NoError(t, err)

			assert.Equal(t, conf.GetLogDB().Host, "domaniname.database.cm")
			assert.Equal(t, conf.GetLogDB().Port, 8989)
			assert.Equal(t, conf.GetLogDB().User, "somebody_user")
			assert.Equal(t, conf.GetLogDB().Passwd, os.Getenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD"))
			assert.Equal(t, conf.GetLogDB().NameDB, "any_name_db")
			assert.Equal(t, conf.GetLogDB().StorageNameDB, "log_storage")
		})
	})

	t.Cleanup(func() {
		cleanEnv()
	})
}

func cleanEnv() {
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

	// Настройки доступа к БД в которую будут записыватся alert и case
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEN")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGENAME")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEUSER")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD")

	//настройки доступа к БД в которую будут записыватся логи
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGHOST")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGPORT")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGNAME")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGUSER")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGPASSWD")
	os.Unsetenv("GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME")

}

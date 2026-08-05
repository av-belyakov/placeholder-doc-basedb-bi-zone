// Пакет confighandler формирует конфигурационные настройки приложения
package confighandler

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/supportingfunctions"
)

func New(rootDir string) (*Config, error) {
	cfg := &Config{}

	var (
		validate *validator.Validate
		envList  map[string]string = map[string]string{
			"GO_PHDOCBASEDBBZ_MAIN":           "",
			"GO_PHDOCBASEDBBZ_REGIONALOBJECT": "",

			//Подключение к NATS
			"GO_PHDOCBASEDBBZ_NHOST":          "",
			"GO_PHDOCBASEDBBZ_NPORT":          "",
			"GO_PHDOCBASEDBBZ_NCACHETTL":      "",
			"GO_PHDOCBASEDBBZ_NSUBSCRIPTIONS": "",
			"GO_PHDOCBASEDBBZ_NENRICHINGQUER": "",

			//Подключение к Kafka
			"GO_PHDOCBASEDBBZ_KHOST":          "",
			"GO_PHDOCBASEDBBZ_KPORT":          "",
			"GO_PHDOCBASEDBBZ_KTOPICS":        "",
			"GO_PHDOCBASEDBBZ_KCACHETTL":      "",
			"GO_PHDOCBASEDBBZ_KAUTHTYPE":      "",
			"GO_PHDOCBASEDBBZ_KSASLMECHANISM": "",
			"GO_PHDOCBASEDBBZ_KSSLUSERNAME":   "",
			"GO_PHDOCBASEDBBZ_KSSLPASSWORD":   "",
			"GO_PHDOCBASEDBBZ_KSSLCAFILE":     "",
			"GO_PHDOCBASEDBBZ_KCERTFILE":      "",
			"GO_PHDOCBASEDBBZ_KKEYFILE":       "",

			//Настройки доступа к БД в которую будут записыватся полученные объекты
			"GO_PHDOCBASEDBBZ_DBSTORAGEHOST":   "",
			"GO_PHDOCBASEDBBZ_DBSTORAGEPORT":   "",
			"GO_PHDOCBASEDBBZ_DBSTORAGENAME":   "",
			"GO_PHDOCBASEDBBZ_DBSTORAGEUSER":   "",
			"GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD": "",
			"GO_PHDOCBASEDBBZ_DBSTORAGETABLES": "",

			//Настройки доступа к БД в которую будут записыватся логи
			"GO_PHDOCBASEDBBZ_DBWLOGHOST":        "",
			"GO_PHDOCBASEDBBZ_DBWLOGPORT":        "",
			"GO_PHDOCBASEDBBZ_DBWLOGNAME":        "",
			"GO_PHDOCBASEDBBZ_DBWLOGUSER":        "",
			"GO_PHDOCBASEDBBZ_DBWLOGPASSWD":      "",
			"GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME": "",
		}
	)

	getFileName := func(sf, confPath string, lfs []fs.DirEntry) (string, error) {
		for _, v := range lfs {
			if v.Name() == sf && !v.IsDir() {
				return filepath.Join(confPath, v.Name()), nil
			}
		}

		return "", fmt.Errorf("file '%s' is not found", sf)
	}

	setCommonSettings := func(fn string) error {
		viper.SetConfigFile(fn)
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		ls := Logs{}
		if ok := viper.IsSet("LOGGING"); ok {
			if err := viper.GetViper().Unmarshal(&ls); err != nil {
				return err
			}

			cfg.Common.Logs = ls.Logging
		}

		z := ZabbixSet{}
		if ok := viper.IsSet("ZABBIX"); ok {
			if err := viper.GetViper().Unmarshal(&z); err != nil {
				return err
			}

			np := 10051
			if z.Zabbix.NetworkPort != 0 && z.Zabbix.NetworkPort < 65536 {
				np = z.Zabbix.NetworkPort
			}

			cfg.Common.Zabbix = ZabbixOptions{
				NetworkPort: np,
				NetworkHost: z.Zabbix.NetworkHost,
				ZabbixHost:  z.Zabbix.ZabbixHost,
				EventTypes:  z.Zabbix.EventTypes,
			}
		}

		return nil
	}

	setSpecial := func(fn string) error {
		viper.SetConfigFile(fn)
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		//Настройка наименования регионального объекта
		if viper.IsSet("COMMONINFO.regional_object") {
			cfg.Common.RegionalObject = viper.GetString("COMMONINFO.regional_object")
		}

		//Настройки для модуля подключения к NATS
		if viper.IsSet("NATS.host") {
			cfg.NATS.Host = viper.GetString("NATS.host")
		}
		if viper.IsSet("NATS.port") {
			cfg.NATS.Port = viper.GetInt("NATS.port")
		}
		if viper.IsSet("NATS.cache_ttl") {
			cfg.NATS.CacheTTL = viper.GetInt("NATS.cache_ttl")
		}
		if viper.IsSet("NATS.subscriptions") {
			cfg.NATS.Subscriptions = viper.GetStringMapString("NATS.subscriptions")
		}
		if viper.IsSet("NATS.enriching_queries") {
			cfg.NATS.EnrichingQueries = viper.GetStringMapString("NATS.enriching_queries")
		}

		//Настройки для модуля подключения к Kafka
		if viper.IsSet("KAFKA.host") {
			cfg.Kafka.Host = viper.GetString("KAFKA.host")
		}
		if viper.IsSet("KAFKA.port") {
			cfg.Kafka.Port = viper.GetInt("KAFKA.port")
		}
		if viper.IsSet("KAFKA.cache_ttl") {
			cfg.Kafka.CacheTTL = viper.GetInt("KAFKA.cache_ttl")
		}
		if viper.IsSet("KAFKA.topics") {
			cfg.Kafka.Topics = viper.GetStringMapString("KAFKA.topics")
		}
		if viper.IsSet("KAFKA.auth_type") {
			cfg.Kafka.AuthType = viper.GetString("KAFKA.auth_type")
		}
		if viper.IsSet("KAFKA.sasl_mechanism") {
			cfg.Kafka.SASLMechanism = viper.GetString("KAFKA.sasl_mechanism")
		}
		if viper.IsSet("KAFKA.ssl_username") {
			cfg.Kafka.SSLUsername = viper.GetString("KAFKA.ssl_username")
		}
		if viper.IsSet("KAFKA.ssl_ca_file") {
			cfg.Kafka.SSLCaFile = viper.GetString("KAFKA.ssl_ca_file")
		}
		if viper.IsSet("KAFKA.ssl_cert_file") {
			cfg.Kafka.SSLCertFile = viper.GetString("KAFKA.ssl_cert_file")
		}
		if viper.IsSet("KAFKA.ssl_key_file") {
			cfg.Kafka.SSLKeyFile = viper.GetString("KAFKA.ssl_key_file")
		}

		// Настройки доступа к БД в которую будет записыватся основная информация
		if viper.IsSet("DATABASESTORAGE.host") {
			cfg.StorageDB.Host = viper.GetString("DATABASESTORAGE.host")
		}
		if viper.IsSet("DATABASESTORAGE.port") {
			cfg.StorageDB.Port = viper.GetInt("DATABASESTORAGE.port")
		}
		if viper.IsSet("DATABASESTORAGE.user") {
			cfg.StorageDB.User = viper.GetString("DATABASESTORAGE.user")
		}
		if viper.IsSet("DATABASESTORAGE.namedb") {
			cfg.StorageDB.NameDB = viper.GetString("DATABASESTORAGE.namedb")
		}
		if viper.IsSet("DATABASESTORAGE.storage_name_db") {
			cfg.StorageDB.Storage = viper.GetStringMapString("DATABASESTORAGE.storage_name_db")
		}

		// Настройки доступа к БД в которую будут записыватся логи
		if viper.IsSet("DATABASEWRITELOG.host") {
			cfg.LogDB.Host = viper.GetString("DATABASEWRITELOG.host")
		}
		if viper.IsSet("DATABASEWRITELOG.port") {
			cfg.LogDB.Port = viper.GetInt("DATABASEWRITELOG.port")
		}
		if viper.IsSet("DATABASEWRITELOG.user") {
			cfg.LogDB.User = viper.GetString("DATABASEWRITELOG.user")
		}
		if viper.IsSet("DATABASEWRITELOG.namedb") {
			cfg.LogDB.NameDB = viper.GetString("DATABASEWRITELOG.namedb")
		}
		if viper.IsSet("DATABASEWRITELOG.storage_name_db") {
			cfg.LogDB.StorageNameDB = viper.GetString("DATABASEWRITELOG.storage_name_db")
		}

		// Настройки для отладочного сервера
		if viper.IsSet("DEBUGSERVER.enable") {
			cfg.DebugServer.Enable = viper.GetBool("DEBUGSERVER.enable")
		}
		if viper.IsSet("DEBUGSERVER.host") {
			cfg.DebugServer.Host = viper.GetString("DEBUGSERVER.host")
		}
		if viper.IsSet("DEBUGSERVER.port") {
			cfg.DebugServer.Port = viper.GetInt("DEBUGSERVER.port")
		}

		return nil
	}

	validate = validator.New(validator.WithRequiredStructEnabled())

	for v := range envList {
		if env, ok := os.LookupEnv(v); ok {
			envList[v] = env
		}
	}

	rootPath, err := supportingfunctions.GetPathRoot(rootDir)
	if err != nil {
		return cfg, err
	}

	confPath := filepath.Join(rootPath, "config")
	list, err := os.ReadDir(confPath)
	if err != nil {
		return cfg, err
	}

	fileNameCommon, err := getFileName("config.yml", confPath, list)
	if err != nil {
		return cfg, err
	}

	//читаем общий конфигурационный файл
	if err := setCommonSettings(fileNameCommon); err != nil {
		return cfg, err
	}

	var fn string
	switch envList["GO_PHDOCBASEDBBZ_MAIN"] {
	case "development":
		fn, err = getFileName("config_dev.yml", confPath, list)
		if err != nil {
			return cfg, err
		}
	case "test":
		fn, err = getFileName("config_test.yml", confPath, list)
		if err != nil {
			return cfg, err
		}
	default:
		fn, err = getFileName("config_prod.yml", confPath, list)
		if err != nil {
			return cfg, err
		}
	}

	if err := setSpecial(fn); err != nil {
		return cfg, err
	}

	//Настройка наименования регионального объекта
	if envList["GO_PHDOCBASEDBBZ_REGIONALOBJECT"] != "" {
		cfg.Common.RegionalObject = envList["GO_PHDOCBASEDBBZ_REGIONALOBJECT"]
	}

	//Настройки для модуля подключения к NATS
	if envList["GO_PHDOCBASEDBBZ_NHOST"] != "" {
		cfg.NATS.Host = envList["GO_PHDOCBASEDBBZ_NHOST"]
	}
	if envList["GO_PHDOCBASEDBBZ_NPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_NPORT"]); err == nil {
			cfg.NATS.Port = p
		}
	}
	if envList["GO_PHDOCBASEDBBZ_NCACHETTL"] != "" {
		if ttl, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_NCACHETTL"]); err == nil {
			cfg.NATS.CacheTTL = ttl
		}
	}
	if envList["GO_PHDOCBASEDBBZ_NSUBSCRIPTIONS"] != "" {
		sublistener := envList["GO_PHDOCBASEDBBZ_NSUBSCRIPTIONS"]
		if !strings.Contains(sublistener, ";") {
			if tmp := strings.Split(sublistener, ":"); len(tmp) == 2 {
				cfg.NATS.Subscriptions[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(sublistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					cfg.NATS.Subscriptions[tmp[0]] = tmp[1]
				}
			}
		}
	}
	if envList["GO_PHDOCBASEDBBZ_NENRICHINGQUER"] != "" {
		reqlistener := envList["GO_PHDOCBASEDBBZ_NENRICHINGQUER"]
		if !strings.Contains(reqlistener, ";") {
			if tmp := strings.Split(reqlistener, ":"); len(tmp) == 2 {
				cfg.NATS.EnrichingQueries[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(reqlistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					cfg.NATS.EnrichingQueries[tmp[0]] = tmp[1]
				}
			}
		}
	}

	//Настройки для модуля подключения к Kafka
	if envList["GO_PHDOCBASEDBBZ_KHOST"] != "" {
		cfg.Kafka.Host = envList["GO_PHDOCBASEDBBZ_KHOST"]
	}
	if envList["GO_PHDOCBASEDBBZ_KPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_KPORT"]); err == nil {
			cfg.Kafka.Port = p
		}
	}
	if envList["GO_PHDOCBASEDBBZ_KCACHETTL"] != "" {
		if ttl, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_KCACHETTL"]); err == nil {
			cfg.Kafka.CacheTTL = ttl
		}
	}
	if envList["GO_PHDOCBASEDBBZ_KTOPICS"] != "" {
		sublistener := envList["GO_PHDOCBASEDBBZ_KTOPICS"]
		if !strings.Contains(sublistener, ";") {
			if tmp := strings.Split(sublistener, ":"); len(tmp) == 2 {
				cfg.Kafka.Topics[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(sublistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					cfg.Kafka.Topics[tmp[0]] = tmp[1]
				}
			}
		}
	}
	if envList["GO_PHDOCBASEDBBZ_KAUTHTYPE"] != "" {
		cfg.Kafka.AuthType = envList["GO_PHDOCBASEDBBZ_KAUTHTYPE"]
	}
	if envList["GO_PHDOCBASEDBBZ_KSASLMECHANISM"] != "" {
		cfg.Kafka.SASLMechanism = envList["GO_PHDOCBASEDBBZ_KSASLMECHANISM"]
	}
	if envList["GO_PHDOCBASEDBBZ_KSSLUSERNAME"] != "" {
		cfg.Kafka.SSLUsername = envList["GO_PHDOCBASEDBBZ_KSSLUSERNAME"]
	}
	if envList["GO_PHDOCBASEDBBZ_KSSLPASSWORD"] != "" {
		cfg.Kafka.SSLPassword = envList["GO_PHDOCBASEDBBZ_KSSLPASSWORD"]
	}
	if envList["GO_PHDOCBASEDBBZ_KSSLCAFILE"] != "" {
		cfg.Kafka.SSLCaFile = envList["GO_PHDOCBASEDBBZ_KSSLCAFILE"]
	}
	if envList["GO_PHDOCBASEDBBZ_KCERTFILE"] != "" {
		cfg.Kafka.SSLCertFile = envList["GO_PHDOCBASEDBBZ_KCERTFILE"]
	}
	if envList["GO_PHDOCBASEDBBZ_KKEYFILE"] != "" {
		cfg.Kafka.SSLKeyFile = envList["GO_PHDOCBASEDBBZ_KKEYFILE"]
	}

	//Настройки доступа к БД в которую будет добавлятся информация по alert и case
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEHOST"] != "" {
		cfg.StorageDB.Host = envList["GO_PHDOCBASEDBBZ_DBSTORAGEHOST"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_DBSTORAGEPORT"]); err == nil {
			cfg.StorageDB.Port = p
		}
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGENAME"] != "" {
		cfg.StorageDB.NameDB = envList["GO_PHDOCBASEDBBZ_DBSTORAGENAME"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEUSER"] != "" {
		cfg.StorageDB.User = envList["GO_PHDOCBASEDBBZ_DBSTORAGEUSER"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD"] != "" {
		cfg.StorageDB.Passwd = envList["GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGETABLES"] != "" {
		sublistener := envList["GO_PHDOCBASEDBBZ_DBSTORAGETABLES"]
		if !strings.Contains(sublistener, ";") {
			if tmp := strings.Split(sublistener, ":"); len(tmp) == 2 {
				cfg.StorageDB.Storage[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(sublistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					cfg.StorageDB.Storage[tmp[0]] = tmp[1]
				}
			}
		}
	}

	//Настройки доступа к БД в которую будут записыватся логи
	if envList["GO_PHDOCBASEDBBZ_DBWLOGHOST"] != "" {
		cfg.LogDB.Host = envList["GO_PHDOCBASEDBBZ_DBWLOGHOST"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_DBWLOGPORT"]); err == nil {
			cfg.LogDB.Port = p
		}
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGNAME"] != "" {
		cfg.LogDB.NameDB = envList["GO_PHDOCBASEDBBZ_DBWLOGNAME"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGUSER"] != "" {
		cfg.LogDB.User = envList["GO_PHDOCBASEDBBZ_DBWLOGUSER"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGPASSWD"] != "" {
		cfg.LogDB.Passwd = envList["GO_PHDOCBASEDBBZ_DBWLOGPASSWD"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME"] != "" {
		cfg.LogDB.StorageNameDB = envList["GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME"]
	}

	//fmt.Printf("STRACT:%+v\n", cfg)

	//выполняем проверку заполненой структуры
	if err = validate.Struct(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

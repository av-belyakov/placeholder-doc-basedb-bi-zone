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
	conf := &Config{}

	var (
		validate *validator.Validate
		envList  map[string]string = map[string]string{
			"GO_PHDOCBASEDBBZ_MAIN":           "",
			"GO_PHDOCBASEDBBZ_REGIONALOBJECT": "",

			//Подключение к NATS
			"GO_PHDOCBASEDBBZ_NHOST":          "",
			"GO_PHDOCBASEDBBZ_NPORT":          "",
			"GO_PHDOCBASEDBBZ_NCACHETTL":      "",
			"GO_PHDOCBASEDBBZ_NSUBLISTENER":   "",
			"GO_PHDOCBASEDBBZ_NENRICHINGQUER": "",

			//Подключение к Kafka
			"GO_PHDOCBASEDBBZ_KHOST":     "",
			"GO_PHDOCBASEDBBZ_KPORT":     "",
			"GO_PHDOCBASEDBBZ_KTOPICS":   "",
			"GO_PHDOCBASEDBBZ_KCACHETTL": "",

			//Настройки доступа к БД в которую будут записыватся полученные объекты
			"GO_PHDOCBASEDBBZ_DBSTORAGEN":      "",
			"GO_PHDOCBASEDBBZ_DBSTORAGEHOST":   "",
			"GO_PHDOCBASEDBBZ_DBSTORAGEPORT":   "",
			"GO_PHDOCBASEDBBZ_DBSTORAGENAME":   "",
			"GO_PHDOCBASEDBBZ_DBSTORAGEUSER":   "",
			"GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD": "",

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

			conf.Common.Logs = ls.Logging
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

			conf.Common.Zabbix = ZabbixOptions{
				NetworkPort: np,
				NetworkHost: z.Zabbix.NetworkHost,
				ZabbixHost:  z.Zabbix.ZabbixHost,
				EventTypes:  z.Zabbix.EventTypes,
			}
		}

		p := ProfilingSet{}
		if ok := viper.IsSet("PROFILING"); ok {
			if err := viper.GetViper().Unmarshal(&p); err != nil {
				return err
			}

			conf.Common.Profiling = ProfilingOptions{
				Host: p.Profiling.Host,
				Port: p.Profiling.Port,
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
			conf.Common.RegionalObject = viper.GetString("COMMONINFO.regional_object")
		}

		//Настройки для модуля подключения к NATS
		if viper.IsSet("NATS.host") {
			conf.NATS.Host = viper.GetString("NATS.host")
		}
		if viper.IsSet("NATS.port") {
			conf.NATS.Port = viper.GetInt("NATS.port")
		}
		if viper.IsSet("NATS.cache_ttl") {
			conf.NATS.CacheTTL = viper.GetInt("NATS.cache_ttl")
		}
		if viper.IsSet("NATS.subscriptions") {
			conf.NATS.Subscriptions = viper.GetStringMapString("NATS.subscriptions")
		}
		if viper.IsSet("NATS.enriching_queries") {
			conf.NATS.EnrichingQueries = viper.GetStringMapString("NATS.enriching_queries")
		}

		//Настройки для модуля подключения к Kafka
		if viper.IsSet("KAFKA.host") {
			conf.Kafka.Host = viper.GetString("KAFKA.host")
		}
		if viper.IsSet("KAFKA.port") {
			conf.Kafka.Port = viper.GetInt("KAFKA.port")
		}
		if viper.IsSet("KAFKA.cache_ttl") {
			conf.Kafka.CacheTTL = viper.GetInt("KAFKA.cache_ttl")
		}
		if viper.IsSet("KAFKA.topics") {
			conf.Kafka.Topics = viper.GetStringMapString("KAFKA.topics")
		}

		// Настройки доступа к БД в которую будет записыватся основная информация
		if viper.IsSet("DATABASESTORAGE.host") {
			conf.StorageDB.Host = viper.GetString("DATABASESTORAGE.host")
		}
		if viper.IsSet("DATABASESTORAGE.port") {
			conf.StorageDB.Port = viper.GetInt("DATABASESTORAGE.port")
		}
		if viper.IsSet("DATABASESTORAGE.user") {
			conf.StorageDB.User = viper.GetString("DATABASESTORAGE.user")
		}
		if viper.IsSet("DATABASESTORAGE.namedb") {
			conf.StorageDB.NameDB = viper.GetString("DATABASESTORAGE.namedb")
		}
		if viper.IsSet("DATABASESTORAGE.storage_name_db") {
			conf.StorageDB.Storage = viper.GetStringMapString("DATABASESTORAGE.storage_name_db")
		}

		// Настройки доступа к БД в которую будут записыватся логи
		if viper.IsSet("DATABASEWRITELOG.host") {
			conf.LogDB.Host = viper.GetString("DATABASEWRITELOG.host")
		}
		if viper.IsSet("DATABASEWRITELOG.port") {
			conf.LogDB.Port = viper.GetInt("DATABASEWRITELOG.port")
		}
		if viper.IsSet("DATABASEWRITELOG.user") {
			conf.LogDB.User = viper.GetString("DATABASEWRITELOG.user")
		}
		if viper.IsSet("DATABASEWRITELOG.namedb") {
			conf.LogDB.NameDB = viper.GetString("DATABASEWRITELOG.namedb")
		}
		if viper.IsSet("DATABASEWRITELOG.storage_name_db") {
			conf.LogDB.StorageNameDB = viper.GetString("DATABASEWRITELOG.storage_name_db")
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
		return conf, err
	}

	confPath := filepath.Join(rootPath, "config")
	list, err := os.ReadDir(confPath)
	if err != nil {
		return conf, err
	}

	fileNameCommon, err := getFileName("config.yml", confPath, list)
	if err != nil {
		return conf, err
	}

	//читаем общий конфигурационный файл
	if err := setCommonSettings(fileNameCommon); err != nil {
		return conf, err
	}

	var fn string
	switch envList["GO_PHDOCBASEDBBZ_MAIN"] {
	case "development":
		fn, err = getFileName("config_dev.yml", confPath, list)
		if err != nil {
			return conf, err
		}
	case "test":
		fn, err = getFileName("config_test.yml", confPath, list)
		if err != nil {
			return conf, err
		}
	default:
		fn, err = getFileName("config_prod.yml", confPath, list)
		if err != nil {
			return conf, err
		}
	}

	if err := setSpecial(fn); err != nil {
		return conf, err
	}

	//Настройка наименования регионального объекта
	if envList["GO_PHDOCBASEDBBZ_REGIONALOBJECT"] != "" {
		conf.Common.RegionalObject = envList["GO_PHDOCBASEDBBZ_REGIONALOBJECT"]
	}

	//Настройки для модуля подключения к NATS
	if envList["GO_PHDOCBASEDBBZ_NHOST"] != "" {
		conf.NATS.Host = envList["GO_PHDOCBASEDBBZ_NHOST"]
	}
	if envList["GO_PHDOCBASEDBBZ_NPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_NPORT"]); err == nil {
			conf.NATS.Port = p
		}
	}
	if envList["GO_PHDOCBASEDBBZ_NCACHETTL"] != "" {
		if ttl, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_NCACHETTL"]); err == nil {
			conf.NATS.CacheTTL = ttl
		}
	}
	if envList["GO_PHDOCBASEDBBZ_NSUBLISTENER"] != "" {
		sublistener := envList["GO_PHDOCBASEDBBZ_NSUBLISTENER"]
		if !strings.Contains(sublistener, ";") {
			if tmp := strings.Split(sublistener, ":"); len(tmp) == 2 {
				conf.NATS.Subscriptions[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(sublistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					conf.NATS.Subscriptions[tmp[0]] = tmp[1]
				}
			}
		}
	}
	if envList["GO_PHDOCBASEDBBZ_NENRICHINGQUER"] != "" {
		reqlistener := envList["GO_PHDOCBASEDBBZ_NENRICHINGQUER"]
		if !strings.Contains(reqlistener, ";") {
			if tmp := strings.Split(reqlistener, ":"); len(tmp) == 2 {
				conf.NATS.EnrichingQueries[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(reqlistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					conf.NATS.EnrichingQueries[tmp[0]] = tmp[1]
				}
			}
		}
	}

	//Настройки для модуля подключения к Kafka
	if envList["GO_PHDOCBASEDBBZ_KHOST"] != "" {
		conf.Kafka.Host = envList["GO_PHDOCBASEDBBZ_KHOST"]
	}
	if envList["GO_PHDOCBASEDBBZ_KPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_KPORT"]); err == nil {
			conf.Kafka.Port = p
		}
	}
	if envList["GO_PHDOCBASEDBBZ_KCACHETTL"] != "" {
		if ttl, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_KCACHETTL"]); err == nil {
			conf.Kafka.CacheTTL = ttl
		}
	}
	if envList["GO_PHDOCBASEDBBZ_KTOPICS"] != "" {
		sublistener := envList["GO_PHDOCBASEDBBZ_KTOPICS"]
		if !strings.Contains(sublistener, ";") {
			if tmp := strings.Split(sublistener, ":"); len(tmp) == 2 {
				conf.Kafka.Topics[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(sublistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					conf.Kafka.Topics[tmp[0]] = tmp[1]
				}
			}
		}
	}

	//Настройки доступа к БД в которую будет добавлятся информация по alert и case
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEHOST"] != "" {
		conf.StorageDB.Host = envList["GO_PHDOCBASEDBBZ_DBSTORAGEHOST"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_DBSTORAGEPORT"]); err == nil {
			conf.StorageDB.Port = p
		}
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGENAME"] != "" {
		conf.StorageDB.NameDB = envList["GO_PHDOCBASEDBBZ_DBSTORAGENAME"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEUSER"] != "" {
		conf.StorageDB.User = envList["GO_PHDOCBASEDBBZ_DBSTORAGEUSER"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD"] != "" {
		conf.StorageDB.Passwd = envList["GO_PHDOCBASEDBBZ_DBSTORAGEPASSWD"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBSTORAGEN"] != "" {
		sublistener := envList["GO_PHDOCBASEDBBZ_DBSTORAGEN"]
		if !strings.Contains(sublistener, ";") {
			if tmp := strings.Split(sublistener, ":"); len(tmp) == 2 {
				conf.StorageDB.Storage[tmp[0]] = tmp[1]
			}
		} else {
			for _, sl := range strings.Split(sublistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					conf.StorageDB.Storage[tmp[0]] = tmp[1]
				}
			}
		}
	}

	//Настройки доступа к БД в которую будут записыватся логи
	if envList["GO_PHDOCBASEDBBZ_DBWLOGHOST"] != "" {
		conf.LogDB.Host = envList["GO_PHDOCBASEDBBZ_DBWLOGHOST"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDBBZ_DBWLOGPORT"]); err == nil {
			conf.LogDB.Port = p
		}
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGNAME"] != "" {
		conf.LogDB.NameDB = envList["GO_PHDOCBASEDBBZ_DBWLOGNAME"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGUSER"] != "" {
		conf.LogDB.User = envList["GO_PHDOCBASEDBBZ_DBWLOGUSER"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGPASSWD"] != "" {
		conf.LogDB.Passwd = envList["GO_PHDOCBASEDBBZ_DBWLOGPASSWD"]
	}
	if envList["GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME"] != "" {
		conf.LogDB.StorageNameDB = envList["GO_PHDOCBASEDBBZ_DBWLOGSTORAGENAME"]
	}

	//выполняем проверку заполненой структуры
	if err = validate.Struct(conf); err != nil {
		return conf, err
	}

	return conf, nil
}

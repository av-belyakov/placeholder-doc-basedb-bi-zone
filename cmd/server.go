package main

import (
	_ "net/http/pprof"
)

//
// это всё можно удалить
//

/*
func server(ctx context.Context) {
	rootPath, err := supportingfunctions.GetPathRoot(constants.Root_Dir)
	if err != nil {
		log.Fatalf("error, it is impossible to form root path (%s)", err.Error())
	}

	// ****************************************************************************
	// *********** инициализируем модуль чтения конфигурационного файла ***********
	cfg, err := confighandler.New(rootPath)
	if err != nil {
		log.Fatalf("error module 'cfgighandler': %v", err)
	}

	// ****************************************************************************
	// ********************* инициализация модуля логирования *********************
	var listLog []simplelogger.OptionsManager
	for _, v := range cfg.GetListLogs() {
		listLog = append(listLog, v)
	}
	opts := simplelogger.CreateOptions(listLog...)
	simpleLogger, err := simplelogger.NewSimpleLogger(ctx, constants.Root_Dir, opts)
	if err != nil {
		log.Fatalf("error module 'simplelogger': %v", err)
	}

	//*********************************************************************************
	//********** инициализация модуля взаимодействия с БД для передачи логов **********
	cfgDB := cfg.GetLogDB()
	if esc, err := elasticsearchapi.NewElasticsearchConnect(elasticsearchapi.Settings{
		Port:               cfgDB.Port,
		Host:               cfgDB.Host,
		User:               cfgDB.User,
		Passwd:             cfgDB.Passwd,
		IndexDB:            cfgDB.StorageNameDB,
		NameRegionalObject: cfg.Common.RegionalObject,
	}); err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())
	} else {
		//подключение логирования в БД
		simpleLogger.SetDataBaseInteraction(esc)
	}

	// ************************************************************************
	// ************* инициализация модуля взаимодействия с Zabbix *************
	chZabbix := make(chan interfaces.Messager)
	cfgZabbix := cfg.GetZabbix()
	wziSettings := wrappers.WrappersZabbixInteractionSettings{
		NetworkPort: cfgZabbix.NetworkPort,
		NetworkHost: cfgZabbix.NetworkHost,
		ZabbixHost:  cfgZabbix.ZabbixHost,
	}
	eventTypes := []wrappers.EventType(nil)
	for _, v := range cfgZabbix.EventTypes {
		eventTypes = append(eventTypes, wrappers.EventType{
			IsTransmit: v.IsTransmit,
			EventType:  v.EventType,
			ZabbixKey:  v.ZabbixKey,
			Handshake: wrappers.Handshake{
				TimeInterval: v.Handshake.TimeInterval,
				Message:      v.Handshake.Message,
			},
		})
	}
	wziSettings.EventTypes = eventTypes
	wrappers.WrappersZabbixInteraction(ctx, wziSettings, simpleLogger, chZabbix)

	//***************************************************************************
	//************** инициализация обработчика логирования данных ***************
	//фактически это мост между simpleLogger и пакетом соединения с Zabbix
	logging := logginghandler.New(simpleLogger, chZabbix)
	logging.Start(ctx)

	// ***************************************************************************
	// *********** инициализируем модуль счётчика для подсчёта сообщений *********
	counting := countermessage.New(chZabbix)
	counting.Start(ctx)

	// **********************************************************************
	// ************ инициализация модуля взаимодействия с Kafka *************
	// в данном случае это основной источник получения информации о КА/КИ
	cfgKafka := cfg.Kafka
	apiKafka, err := kafkaapi.New(
		logging,
		counting,
		kafkaapi.WithNameRegionalObject(cfg.Common.RegionalObject),
		kafkaapi.WithHost(cfgKafka.Host),
		kafkaapi.WithPort(cfgKafka.Port),
		kafkaapi.WithCacheTTL(cfgKafka.CacheTTL),
		kafkaapi.WithTopicsSubscription(cfgKafka.Topics),
	)
	if err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}
	//--- старт модуля
	if err := apiKafka.Start(ctx); err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}

	// ***********************************************************************
	// ************** инициализация модуля взаимодействия с NATS *************
	// в данном случае брокер сообщений NATS нужен для взаимодействия с модулями:
	// - enricher_geoip (обогащение информаций о местоположении ip адресов)
	// - enricher_sensor_information (обогащение информаций о сенсорах)
	cfgNats := cfg.NATS
	apiNats, err := natsapi.New(
		logging,
		counting,
		natsapi.WithHost(cfgNats.Host),
		natsapi.WithPort(cfgNats.Port),
		natsapi.WithCacheTTL(cfgNats.CacheTTL),
		natsapi.WithRequests(cfgNats.EnrichingQueries),
		natsapi.WithSubscriptions(cfgNats.Subscriptions))
	if err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}
	//--- старт модуля
	if err = apiNats.Start(ctx); err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}

	// *********************************************************************
	// ************** инициализация модуля взаимодействия с БД *************
	cfgStorageDB := cfg.GetStorageDB()
	apiDBS, err := databasestorageapi.New(
		logging,
		counting,
		databasestorageapi.WithHost(cfgStorageDB.Host),
		databasestorageapi.WithPort(cfgStorageDB.Port),
		databasestorageapi.WithNameDB(cfgStorageDB.NameDB),
		databasestorageapi.WithUser(cfgStorageDB.User),
		databasestorageapi.WithPasswd(cfgStorageDB.Passwd),
		databasestorageapi.WithStorage(cfgStorageDB.Storage),
		databasestorageapi.WithMaxGetDocumentSize(15))
	if err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}
	//--- старт модуля
	if err := apiDBS.Start(ctx); err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}

	// *********************************************************
	// ************** инициализация маршрутизатора *************
	r := NewRouter(
		logging,
		counting,
		ApplicationRouterSettings{
			ChanToNats:    apiNats.GetChannelToModule(),
			ChanFromNats:  apiNats.GetChannelFromModule(),
			ChanToKafka:   apiKafka.GetChannelToModule(),
			ChanFromKafka: apiKafka.GetChannelFromModule(),
			ChanToDBS:     apiDBS.GetChannelToModule(),
			ChanFromDBS:   apiDBS.GetChannelFromModule(),
		})
	r.Router(ctx)

	//информационное сообщение
	infoMsg := getInformationMessage(cfg)
	_ = simpleLogger.Write("info", infoMsg)

	//для отладки через pprof (только для теста)
	//http://cfg.Common.Profiling.Host:cfg.Common.Profiling.Port/debug/pprof/
	//go tool pprof http://host:port/debug/pprof/heap
	//go tool pprof http://host:port/debug/pprof/allocs
	//go tool pprof http://host:port/debug/pprof/goroutine
	if os.Getenv("GO_PHDOCBASEDBBZ_MAIN") == "test" || os.Getenv("GO_PHDOCBASEDBBZ_MAIN") == "development" {
		if cfg.Common.Profiling.Port > 0 {
			go func() {
				log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Common.Profiling.Host, cfg.Common.Profiling.Port), nil))
			}()
		}
	}

	<-ctx.Done()
}
*/

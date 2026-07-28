package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/appname"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/appversion"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
)

func getInformationMessage(conf confighandler.Config) string {
	version, err := appversion.GetAppVersion()
	if err != nil {
		log.Println(err)
	}

	var profiling string
	appStatus := fmt.Sprintf("%vproduction%v", constants.Ansi_Bright_Blue, constants.Ansi_Reset)
	envValue, ok := os.LookupEnv("GO_PHDOCBASEDBBZ_MAIN")
	if ok && (envValue == "development" || envValue == "test") {
		appStatus = fmt.Sprintf("%v%s", constants.Ansi_Bright_Red, envValue)

		host := "*"
		port := conf.Common.Profiling.Port

		if port > 0 {
			profiling = fmt.Sprintf("%vProfiling is available on %v'%s:%d/debug/pprof'%v\n", constants.Ansi_Bright_Green, constants.Ansi_Bright_Blue, host, port, constants.Ansi_Reset)
		}
	}

	msg := fmt.Sprintf("Application '%s' v%s was successfully launched", appname.GetAppName(), strings.Replace(version, "\n", "", -1))

	topics := make([]string, 0, len(conf.Kafka.Topics))
	iterator := maps.Values(conf.Kafka.Topics)
	for v := range iterator {
		topics = append(topics, v)
	}

	subscriptions := make([]string, 0, len(conf.NATS.Subscriptions))
	iterator = maps.Values(conf.NATS.Subscriptions)
	for v := range iterator {
		subscriptions = append(subscriptions, v)
	}

	fmt.Printf("\n%v%v%s.%v\n", constants.Bold_Font, constants.Ansi_Bright_Green, msg, constants.Ansi_Reset)
	fmt.Printf("%vApplication status is '%v%s%v%v'%v\n", constants.Ansi_Bright_Green, constants.Underlining, appStatus, constants.Ansi_Reset, constants.Ansi_Bright_Green, constants.Ansi_Reset)
	fmt.Printf(
		"%vConnect to Kafka with address %v%s:%d%v%v, topics: %v%s%v\n",
		constants.Ansi_Bright_Green,
		constants.Ansi_Dark_Gray,
		conf.Kafka.Host,
		conf.Kafka.Port,
		constants.Ansi_Reset,
		constants.Ansi_Bright_Green,
		constants.Ansi_Dark_Gray,
		strings.Join(topics, ", "),
		constants.Ansi_Reset,
	)
	fmt.Printf(
		"%vConnect to NATS with address %v%s:%d%v%v, subscriptions: %v%s%v\n",
		constants.Ansi_Bright_Green,
		constants.Ansi_Dark_Gray,
		conf.NATS.Host,
		conf.NATS.Port,
		constants.Ansi_Reset,
		constants.Ansi_Bright_Green,
		constants.Ansi_Dark_Gray,
		strings.Join(subscriptions, ", "),
		constants.Ansi_Reset,
	)
	fmt.Printf("%vConnect to Database with address %v%s:%d%v\n", constants.Ansi_Bright_Green, constants.Ansi_Dark_Gray, conf.StorageDB.Host, conf.StorageDB.Port, constants.Ansi_Reset)
	fmt.Println(profiling)

	return msg
}

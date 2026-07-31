package main

import (
	"fmt"
	"log"
	"maps"
	"strings"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/appname"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/appversion"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/confighandler"
)

func getInformationMessage(cfg confighandler.Config) string {
	version, err := appversion.GetAppVersion()
	if err != nil {
		log.Println(err)
	}

	var profiling string
	appStatus := fmt.Sprintf("%vproduction%v", constants.Ansi_Bright_Blue, constants.Ansi_Reset)

	msg := fmt.Sprintf("Application '%s' v%s was successfully launched", appname.GetAppName(), strings.Replace(version, "\n", "", -1))

	topics := make([]string, 0, len(cfg.GetKafka().Topics))
	iterator := maps.Values(cfg.GetKafka().Topics)
	for v := range iterator {
		topics = append(topics, v)
	}

	subscriptions := make([]string, 0, len(cfg.GetNATS().Subscriptions))
	iterator = maps.Values(cfg.GetNATS().Subscriptions)
	for v := range iterator {
		subscriptions = append(subscriptions, v)
	}

	debugServerStatus := fmt.Sprintf("%vdisable%v", constants.Ansi_Bright_Red, constants.Ansi_Reset)
	if cfg.GetDebugServer().Enable {
		debugServerStatus = fmt.Sprintf("%venable%v", constants.Ansi_Bright_Red, constants.Ansi_Reset)
		profiling = fmt.Sprintf(
			"%vProfiling is available on %v'%s:%d/debug/pprof'%v\n",
			constants.Ansi_Bright_Green,
			constants.Ansi_Bright_Blue,
			cfg.GetDebugServer().Host,
			cfg.GetDebugServer().Port,
			constants.Ansi_Reset,
		)
	}

	fmt.Printf("\n%v%v%s.%v\n", constants.Bold_Font, constants.Ansi_Bright_Green, msg, constants.Ansi_Reset)
	fmt.Printf(
		"%vApplication status is '%v%s%v%v'%v\n",
		constants.Ansi_Bright_Green,
		constants.Underlining,
		appStatus,
		constants.Ansi_Reset,
		constants.Ansi_Bright_Green,
		constants.Ansi_Reset,
	)
	fmt.Printf(
		"%vConnect to Kafka with address %v%s:%d%v%v, topics: %v%s%v\n",
		constants.Ansi_Bright_Green,
		constants.Ansi_Dark_Gray,
		cfg.GetKafka().Host,
		cfg.GetKafka().Port,
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
		cfg.GetNATS().Host,
		cfg.GetNATS().Port,
		constants.Ansi_Reset,
		constants.Ansi_Bright_Green,
		constants.Ansi_Dark_Gray,
		strings.Join(subscriptions, ", "),
		constants.Ansi_Reset,
	)
	fmt.Printf(
		"%vConnect to Database with address %v%s:%d%v\n",
		constants.Ansi_Bright_Green,
		constants.Ansi_Dark_Gray,
		cfg.StorageDB.Host,
		cfg.StorageDB.Port,
		constants.Ansi_Reset,
	)
	fmt.Printf(
		"%vDebug server with address %v%s:%d%v %s%v\n",
		constants.Ansi_Bright_Green,
		constants.Ansi_Dark_Gray,
		cfg.GetDebugServer().Host,
		cfg.GetDebugServer().Port,
		constants.Ansi_Reset,
		debugServerStatus,
		constants.Ansi_Reset,
	)
	fmt.Println(profiling)

	return msg
}

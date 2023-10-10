package main

import (
	"flag"
	"fmt"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/dao"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/database"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/metric"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/monitor"
	"github.com/bnb-chain/greenfield-data-marketplace-backend/util"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagConfigPath = "config-path"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config path")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("bind flags error, err=%s", err))
	}
}

func printUsage() {
	fmt.Print("usage: ./monitor --config-path config_file_path\n")
}

func main() {
	initFlags()

	configFilePath := viper.GetString(flagConfigPath)
	if configFilePath == "" {
		printUsage()
		return
	}
	config := util.ParseMonitorConfigFromFile(configFilePath)

	util.InitLogger(config.LogConfig)

	db, err := database.ConnectDBWithConfig(config.DBConfig)
	if err != nil {
		util.Logger.Errorf("connect database error, err=%s", err.Error())
		return
	}

	metricServer := metric.NewMetricService(config)

	itemDao := dao.NewDbItemDao(db)

	gnfdBlockDao := dao.NewDbGnfdBlockDao(db)
	gnfdClient := monitor.NewGnfdCompositClients(config.GnfdRpcAddrs, config.GnfdChainId, false)
	gnfdProcessor := monitor.NewGnfdBlockProcessor(gnfdClient, gnfdBlockDao, itemDao, db, metricServer)
	gnfdMonitor := monitor.NewMonitor(gnfdProcessor, config.GnfdStartHeight)

	bscBlockDao := dao.NewDbBscBlockDao(db)
	bscClient := monitor.NewBscCompositeClients(config.BscRpcAddrs, config.BscBlocksForFinality)
	bscProcessor := monitor.NewBscBlockProcessor(bscClient, config.BscMarketplaceContract, bscBlockDao, itemDao, db, metricServer)
	bscMonitor := monitor.NewMonitor(bscProcessor, config.BscStartHeight)

	go gnfdMonitor.Start()
	go bscMonitor.Start()
	metricServer.Start()

	select {}
}

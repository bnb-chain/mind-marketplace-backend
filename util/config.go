package util

import (
	"encoding/json"
	"os"
)

type DBConfig struct {
	DBDialect string `json:"db_dialect"`
	DBPath    string `json:"db_path"`
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

type APIConfig struct {
	EnableCache bool `json:"enable_cache"`
}

type ServerConfig struct {
	Env       string     `json:"env"`
	DBConfig  *DBConfig  `json:"db_config"`
	APIConfig *APIConfig `json:"api_config"`
	LogConfig *LogConfig `json:"log_config"`
}

func ParseServerConfigFromFile(filePath string) *ServerConfig {
	bz, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config ServerConfig
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	return &config
}

type MonitorConfig struct {
	Env string `json:"env"`

	BscRpcAddrs            []string `json:"bsc_rpc_addrs"`
	BscBlocksForFinality   int      `json:"bsc_blocks_for_finality"`
	BscMarketplaceContract string   `json:"bsc_marketplace_contract"`

	GnfdRpcAddrs []string `json:"gnfd_rpc_addrs"`
	GnfdChainId  string   `json:"gnfd_chain_id"`

	DBConfig  *DBConfig  `json:"db_config"`
	LogConfig *LogConfig `json:"log_config"`
}

func ParseMonitorConfigFromFile(filePath string) *MonitorConfig {
	bz, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var config MonitorConfig
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	return &config
}

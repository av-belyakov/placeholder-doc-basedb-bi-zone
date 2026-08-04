package confighandler

type Config struct {
	Common      Common
	NATS        CfgNats
	Kafka       CfgKafka
	LogDB       CfgWriteLogDB
	StorageDB   CfgStorageDB
	DebugServer CfgDebugServer
}

type Common struct {
	RegionalObject string
	Logs           []*LogSet
	Zabbix         ZabbixOptions
}

type Logs struct {
	Logging []*LogSet
}

type LogSet struct {
	MsgTypeName   string `validate:"oneof=error info warning" yaml:"msgTypeName"`
	PathDirectory string `validate:"required" yaml:"pathDirectory"`
	MaxFileSize   int    `validate:"min=1000" yaml:"maxFileSize"`
	WritingStdout bool   `validate:"required" yaml:"writingStdout"`
	WritingFile   bool   `validate:"required" yaml:"writingFile"`
	WritingDB     bool   `validate:"required" yaml:"writingDB"`
}

type ZabbixSet struct {
	Zabbix ZabbixOptions
}

type ZabbixOptions struct {
	EventTypes  []EventType `yaml:"eventType"`
	NetworkHost string      `validate:"required" yaml:"networkHost"`
	ZabbixHost  string      `validate:"required" yaml:"zabbixHost"`
	NetworkPort int         `validate:"gt=0,lte=65535" yaml:"networkPort"`
}

type EventType struct {
	EventType  string    `validate:"required" yaml:"eventType"`
	ZabbixKey  string    `validate:"required" yaml:"zabbixKey"`
	Handshake  Handshake `yaml:"handshake"`
	IsTransmit bool      `yaml:"isTransmit"`
}

type Handshake struct {
	Message      string `validate:"required" yaml:"message"`
	TimeInterval int    `yaml:"timeInterval"`
}

type CfgNats struct {
	EnrichingQueries map[string]string `yaml:"enriching_queries"`
	Subscriptions    map[string]string `yaml:"subscriptions"`
	Host             string            `validate:"required" yaml:"host"`
	Port             int               `validate:"gt=0,lte=65535" yaml:"port"`
	CacheTTL         int               `validate:"gt=1,lte=86400" yaml:"cache_ttl"`
}

type CfgKafka struct {
	Topics         map[string]string `yaml:"topics"`
	Login          string            `yaml:"login"`
	Passwd         string            `yaml:"passwd"`
	CertPath       string            `yaml:"cert_path"`
	TrustStorePath string            `yaml:"trust_store_path"`
	Host           string            `validate:"required" yaml:"host"`
	Port           int               `validate:"gt=0,lte=65535" yaml:"port"`
	CacheTTL       int               `validate:"gt=1,lte=86400" yaml:"cache_ttl"`
}

type CfgStorageDB struct {
	Storage map[string]string `yaml:"storage_name_db"`
	Host    string            `yaml:"host"`
	User    string            `yaml:"user"`
	Passwd  string            `yaml:"passwd"`
	NameDB  string            `yaml:"namedb"`
	Port    int               `validate:"gt=0,lte=65535" yaml:"port"`
}

type CfgWriteLogDB struct {
	Host          string `yaml:"host"`
	User          string `yaml:"user"`
	Passwd        string `yaml:"passwd"`
	NameDB        string `yaml:"namedb"`
	StorageNameDB string `yaml:"storage_name_db"`
	Port          int    `validate:"gt=0,lte=65535" yaml:"port"`
}

type CfgDebugServer struct {
	Host   string `yaml:"host"`
	Port   int    `validate:"gt=0,lte=65535" yaml:"port"`
	Enable bool   `yaml:"enabled"`
}

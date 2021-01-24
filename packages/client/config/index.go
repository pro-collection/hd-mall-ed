package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string
	ImagePrefixUrl  string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExts  []string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

var AppConfig = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseConfig = &Database{}

var Config *ini.File

func SetUp() {
	Config, err := ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("初始化配置文件失败: ", err)
	}
	err = Config.Section("app").MapTo(AppConfig)
	if err != nil {
		log.Fatal("config mapTo appConfig err: ", err)
	}

	AppConfig.ImageMaxSize = AppConfig.ImageMaxSize * 1024 * 1024

	err = Config.Section("server").MapTo(ServerConfig)
	if err != nil {
		log.Fatalf("config mapTo ServerSetting err: %v", err)
	}

	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second

	err = Config.Section("database").MapTo(DatabaseConfig)
	if err != nil {
		log.Fatalf("config mapTo DatabaseSetting err: %v", err)
	}
}

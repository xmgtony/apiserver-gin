package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var Cfg Config

// Config is application global config
type Config struct {
	Mode            string      `mapstructure:"mode"`           // gin启动模式
	Port            string      `mapstructure:"port"`           // 启动端口
	ApplicationName string      `mapstructure:"name"`           //应用名称
	Url             string      `mapstructure:"url"`            // 应用地址,用于自检 eg. http://127.0.01
	MaxPingCount    int         `mapstructure:"max_ping_count"` // 最大自检次数，用户健康检查
	Database        DataBaseCfg `mapstructure:"database"`       // 数据库信息
	RedisCfg        RedisCfg    `mapstructure:"redis"`          // redis
}

// DataBaseCfg is used to configure mysql database
type DataBaseCfg struct {
	Dbname          string `mapstructure:"dbname"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	MaximumPoolSize int    `mapstructure:"maximum-pool-size"`
	MaximumIdleSize int    `mapstructure:"maximum-idle-size"`
	LogMode         bool   `mapstructure:"log-mode"`
}

// DataBaseCfg is used to configure redis
type RedisCfg struct {
	Addr         string `mapstructure:"address"`
	Password     string `mapstructure:"password"`
	Db           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool-size"`
	MinIdleConns int    `mapstructure:"min-idle-conns"`
	IdleTimeout  int    `mapstructure:"idle-timeout"`
}

// Load is a loader to load config file.
func Load(cfg string) {
	// 初始化配置文件
	if err := initConfig(cfg); err != nil {
		panic(err)
	}
	// 监控配置文件，并热加载
	watchConfig()
}

func initConfig(cfg string) error {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		// 设置默认的config
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APPLICATION")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析到struct
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}
	log.Println("application config load completed!")
	return nil
}

// 监控配置文件变动
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
	})
}

module apiserver-gin

go 1.16

require (
	github.com/appleboy/gin-jwt/v2 v2.6.2
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.5.0
	github.com/go-redis/redis/v8 v8.7.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.5.0
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	gopkg.in/go-playground/validator.v9 v9.29.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.1
)

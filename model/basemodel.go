package model

import (
	"apidemo-gin/pkg/time"
)

// BaseModel 用于表示model公共的部分（对应数据库公共的字段）
// model结构一般和数据库的表结构是对应的，数据库一条记录对应一个model实例
// 平时在设计数据库时，一般会有规范要求，我所在的公司字段要求如下，数据库字段必须有：
// id(主键), enabled_status(有效状态), created(创建时间), modified(修改时间)
// 其中modified 设置为 ON UPDATE CURRENT_TIMESTAMP ，以上仅供参考
// gorm.Model 也提供了类似BaseModel功能，并且根据此model字段做了一些自动逻辑。
type BaseModel struct {
	Id            int64          `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	EnabledStatus *int8          `gorm:"column:enabled_status;type:tinyint;default:1" json:"-"`
	Created       time.JsonTime  `gorm:"column:created;type:datetime;default:CURRENT_TIMESTAMP" json:"created"`
	Modified      *time.JsonTime `gorm:"column:modified;type:timestamp;default:CURRENT_TIMESTAMP" json:"modified"`
}

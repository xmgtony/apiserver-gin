package code

// 自定义错误码，通常错误由错误码和错误信息两部分组成，便于跟踪和维护错误信息
// 错误码为0表示成功
// 错误码3开头，表示第三方api调用错误
// 错误码4开头，表示业务层面的错误，比如校验等
// 错误码5开头，表示服务器错误，比如数组越界等
// ----------------------------------
// 错误码过多时，可以根据业务功能拆分到不同的文件或者包中

const (
	// Success 表示成功
	Success = 0
	// Unknown 无法预知或者未手动处理的错误
	Unknown = -1
)

const (
	// MpApiErr 小程序接口调用错误
	MpApiErr = iota + 30000
)

const (
	// ValidateErr 校验错误
	ValidateErr = iota + 40000

	// RequireAuthErr 没有权限
	RequireAuthErr

	// NotFoundErr 没有记录
	NotFoundErr

	// UserLoginErr 登录错误
	UserLoginErr

	// AuthTokenErr token 鉴权错误或权限不足
	AuthTokenErr

	// RecordCreateErr 创建记录，数据持久化失败
	RecordCreateErr
)

const (
	// TransactionErr 事物提交失败
	TransactionErr = iota + 60000
	// DuplicateErr 记录存在重复
	DuplicateErr
)

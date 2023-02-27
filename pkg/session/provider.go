package session

type Provider interface {
	// SessionInit 初始化session
	SessionInit(sid string) (Session, error)
	// SessionRead 读取session
	SessionRead(sid string) (Session, error)
	// SessionDestroy 销毁一个session
	SessionDestroy(sid string) error
	// SessionGC 定时清理session
	SessionGC(maxLifeTime int64)
}

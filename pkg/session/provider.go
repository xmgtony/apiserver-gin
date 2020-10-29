package session

type Provider interface {
	// 初始化session
	SessionInit(sid string) (Session, error)
	// 读取session
	SessionRead(sid string) (Session, error)
	// 销毁一个session
	SessionDestroy(sid string) error
	// 定时清理session
	SessionGC(maxLifeTime int64)
}

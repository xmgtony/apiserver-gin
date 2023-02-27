package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Manager struct {
	sessionName string // 存储在客户端的sessionId的key, 比如gsession

	provider Provider // 存储接口

	lock sync.RWMutex // 保护session操作

	maxLifeTime int64 // 生存周期
}

// NewManager 创建全局的session管理器
func NewManager(sessionName string, provider Provider, maxLifeTime int64) (*Manager, error) {
	if provider == nil {
		return nil, errors.New("Provider cannot be empty!")
	}
	manager := &Manager{sessionName: sessionName, provider: provider, maxLifeTime: maxLifeTime}
	// 创建完manager 启动一个回收session的goroutine
	go manager.GC()
	return manager, nil
}

// sessionId() 生成sessionId
func (m *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic("session id generate failed!")
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Session 用来获取或者初始化session
func (m *Manager) Session(w http.ResponseWriter, r *http.Request) (session Session) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := r.Cookie(m.sessionName)
	if err != nil || cookie.Value == "" {
		// 初始化session信息
		sid := m.sessionId()
		session, err = m.provider.SessionInit(sid)
		if err == nil {
			cookie := http.Cookie{
				Name:     m.sessionName,
				Value:    url.QueryEscape(sid),
				Path:     "/",
				MaxAge:   int(m.maxLifeTime),
				HttpOnly: true,
			}
			// 写回浏览器
			http.SetCookie(w, &cookie)
		}
	} else {
		// 已经有session时
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(sid)
	}
	return
}

// SessionDestroy 销毁session, 设置客户端过期
func (m *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.sessionName)
	if err != nil || cookie.Value == "" {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	sid, _ := url.QueryUnescape(cookie.Value)
	_ = m.provider.SessionDestroy(sid)
	expiration := time.Now()
	cookie.Expires = expiration
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}

// GC session根据lru策略回收
func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime)*time.Second, func() {
		m.GC()
	})
}

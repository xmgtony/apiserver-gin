// session的内存存储实现

package memory

import (
	"apiserver-gin/pkg/session"
	"container/list"
	"errors"
	"time"
)

// MemoryStore 每个session的内容的存储位置
type MemoryStore struct {
	sid            string                      //session id
	lastAccessTime time.Time                   //session的最后访问时间
	value          map[interface{}]interface{} // 存储session内容的键值对容器
}

func (m *MemoryStore) Set(key string, value interface{}) error {
	if len(key) <= 0 {
		return errors.New("key cannot be empty!")
	}
	m.value[key] = value
	return nil
}

func (m *MemoryStore) Get(key string) interface{} {
	if len(key) <= 0 {
		return nil
	}
	return m.value[key]
}

func (m *MemoryStore) Delete(key string) error {
	if len(key) <= 0 {
		return errors.New("key cannot empty when delete")
	}
	delete(m.value, key)
	return nil
}

func (m *MemoryStore) GetSessionId() string {
	return m.sid
}

// MemoryProvider 保存每一个session
type MemoryProvider struct {
	sessions map[interface{}]*list.Element
	gclist   *list.List
}

func NewProvider() *MemoryProvider {
	return &MemoryProvider{
		sessions: make(map[interface{}]*list.Element, 0),
		gclist:   list.New(),
	}
}

func (p *MemoryProvider) SessionInit(sid string) (session.Session, error) {
	v := make(map[interface{}]interface{}, 0)
	s := &MemoryStore{
		sid:            sid,
		lastAccessTime: time.Now(),
		value:          v,
	}
	element := p.gclist.PushBack(s)
	p.sessions[sid] = element
	return s, nil
}

func (p *MemoryProvider) SessionRead(sid string) (session.Session, error) {
	if v, ok := p.sessions[sid]; ok {
		sess := v.Value.(*MemoryStore)
		p.UpdateSessionLifeTime(v)
		return sess, nil
	} else {
		sess, err := p.SessionInit(sid)
		return sess, err
	}
}

func (p *MemoryProvider) SessionDestroy(sid string) error {
	if v, ok := p.sessions[sid]; ok {
		delete(p.sessions, sid)
		p.gclist.Remove(v)
	}
	return nil
}

func (p *MemoryProvider) SessionGC(maxLifeTime int64) {
	for {
		element := p.gclist.Back()
		if element == nil {
			break
		}
		if element.Value.(*MemoryStore).lastAccessTime.Unix()+maxLifeTime < time.Now().Unix() {
			p.gclist.Remove(element)
			delete(p.sessions, element.Value.(*MemoryStore).sid)
		} else {
			break
		}
	}
}

// UpdateSessionLifeTime 每次读session更新session最后访问时间
func (p *MemoryProvider) UpdateSessionLifeTime(e *list.Element) {
	e.Value.(*MemoryStore).lastAccessTime = time.Now()
	p.gclist.MoveToFront(e)
}

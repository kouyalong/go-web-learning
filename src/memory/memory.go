package memory

import (
	"container/list"
	"sync"
	"time"

	"session"
)


type ProviderN struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

type SessionStore struct {
	sid				string
	timeAccessed	time.Time
	value			map[interface{}]interface{}
}

func (s *SessionStore) Get(key interface{}) interface{} {
	pDer.SessionUpdate(s.sid)
	if v, ok := s.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (s *SessionStore) Delete(key interface{}) error {
	delete(s.value, key)
	pDer.SessionUpdate(s.sid)
	return nil
}

func (s *SessionStore) SessionID() string {
	return s.sid
}

func (s *SessionStore) Set(key, value interface{}) error {
	s.value[key] = value
	pDer.SessionUpdate(s.sid)
	return nil

}

var pDer = &ProviderN{list: list.New()}

func (pDer *ProviderN) SessionInit(sid string) (session.Session, error) {
	pDer.lock.Lock()
	defer pDer.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newSess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pDer.list.PushBack(newSess)
	pDer.sessions[sid] = element
	return newSess, nil
}


func (pDer *ProviderN) SessionRead(sid string) (session.Session, error) {
	if element, ok := pDer.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pDer.SessionInit(sid)
		return sess, err
	}
}

func (pDer *ProviderN) SessionDestroy(sid string) error {
	if element, ok := pDer.sessions[sid]; ok {
		delete(pDer.sessions, sid)
		pDer.list.Remove(element)
		return nil
	}
	return nil
}


func (pDer *ProviderN) SessionGC(maxLifeTime int64) {
	pDer.lock.Lock()
	defer pDer.lock.Unlock()

	for {
		element := pDer.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			pDer.list.Remove(element)
			delete(pDer.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pDer *ProviderN) SessionUpdate(sid string) error {
	pDer.lock.Unlock()
	defer pDer.lock.Unlock()

	if element, ok := pDer.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pDer.list.MoveToFront(element)
		return nil
	}
	return nil
}


func init() {
	pDer.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pDer)
}

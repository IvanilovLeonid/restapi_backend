package session

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

type Session struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
}

type SessionInterface interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	SessionID() string
}

// Интерфейс провайдера хранения сессий

type Provider interface {
	SessionInit(sid string) (SessionInterface, error)
	SessionRead(sid string) (SessionInterface, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

// Реализация сессий в памяти

type MemorySession struct {
	sid        string
	data       map[string]interface{}
	lastAccess time.Time
	sync.RWMutex
}

func (s *MemorySession) Set(key string, value interface{}) error {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
	return nil
}

func (s *MemorySession) Get(key string) (interface{}, error) {
	s.RLock()
	defer s.RUnlock()
	if value, ok := s.data[key]; ok {
		return value, nil
	}
	return nil, errors.New("key not found")
}

func (s *MemorySession) Delete(key string) error {
	s.Lock()
	defer s.Unlock()
	delete(s.data, key)
	return nil
}

func (s *MemorySession) SessionID() string {
	return s.sid
}

func (s *MemorySession) Serialize() (string, error) {
	s.RLock()
	defer s.RUnlock()
	data, err := json.Marshal(s.data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *MemorySession) Deserialize(serializedData string) error {
	s.Lock()
	defer s.Unlock()
	return json.Unmarshal([]byte(serializedData), &s.data)
}

// Провайдер хранения сессий в памяти

type MemoryProvider struct {
	sessions map[string]*MemorySession
	lock     sync.RWMutex
}

func NewMemoryProvider() *MemoryProvider {
	return &MemoryProvider{sessions: make(map[string]*MemorySession)}
}

func (p *MemoryProvider) SessionInit(sid string) (SessionInterface, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	sess := &MemorySession{
		sid:        sid,
		data:       make(map[string]interface{}),
		lastAccess: time.Now(),
	}
	p.sessions[sid] = sess
	return sess, nil
}

func (p *MemoryProvider) SessionRead(sid string) (SessionInterface, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	if sess, ok := p.sessions[sid]; ok {
		return sess, nil
	}
	return nil, errors.New("session not found")
}

func (p *MemoryProvider) SessionDestroy(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.sessions, sid)
	return nil
}

func (p *MemoryProvider) SessionGC(maxLifeTime int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for sid, sess := range p.sessions {
		if time.Now().Sub(sess.lastAccess) > time.Duration(maxLifeTime)*time.Second {
			delete(p.sessions, sid)
		}
	}
}

// Менеджер сессий

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

func NewManager(provider Provider, cookieName string, maxlifetime int64) *Manager {
	return &Manager{
		provider:    provider,
		cookieName:  cookieName,
		maxlifetime: maxlifetime,
	}
}

func (m *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (SessionInterface, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid := m.sessionID()
		session, _ := m.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     m.cookieName,
			Value:    sid,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(m.maxlifetime),
		}
		http.SetCookie(w, &cookie)
		return session, nil
	}

	sid := cookie.Value
	session, err := m.provider.SessionRead(sid)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (m *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	m.lock.Lock()
	defer m.lock.Unlock()

	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	m.provider.SessionDestroy(cookie.Value)

	expiration := time.Now()
	cookie = &http.Cookie{
		Name:     m.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiration,
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}

func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxlifetime)
}

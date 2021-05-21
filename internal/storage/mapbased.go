package mapbased

import (
	"io"
	"sync"
)

type Storage struct {
	*sync.RWMutex
	users map[uint64]User
}

type User struct {
	ID           uint64 `json:"user_id"`
	OnlineStatus bool   `json:"online"`
	Friends      map[uint64]struct{}
	Connection   io.Writer
}

func NewStorage() *Storage {
	s := &Storage{
		RWMutex: new(sync.RWMutex),
		users:   make(map[uint64]User),
	}
	return s
}

// AddUser добавление либо обновление записи User
func (s *Storage) AddUser(userId uint64, friends []uint64, conn io.Writer) {
	s.Lock()
	frndsMap := make(map[uint64]struct{})
	for _, item := range friends {
		frndsMap[item] = struct{}{}
	}

	s.users[userId] = User{
		ID:           userId,
		OnlineStatus: true,
		Friends:      frndsMap,
		Connection:   conn,
	}

	s.Unlock()
}

// SetUserOffline установить статус
func (s *Storage) SetUserOffline(userId uint64) (isOk bool) {
	s.Lock()
	defer s.Unlock()
	if val, ok := s.users[userId]; ok {
		val.OnlineStatus = false
		s.users[userId] = val
		return true
	}
	return false
}

// GetUserStatus получить статус
func (s *Storage) GetUserStatus(userId uint64) (onlineStatus bool) {
	s.RLock()
	if val, ok := s.users[userId]; ok {
		onlineStatus = val.OnlineStatus
	}
	s.RUnlock()
	return
}

// GetUserConn получить connection пользователя
func (s *Storage) GetUserConn(userId uint64) (conn io.Writer, ok bool) {
	s.RLock()
	defer s.RUnlock()
	if val, ok := s.users[userId]; ok {
		return val.Connection, true
	}
	return nil, false
}

// RemoveUser удаление пользователя
func (s *Storage) RemoveUser(userId uint64) {
	s.Lock()
	delete(s.users, userId)
	s.Unlock()
}

// Followers список юзеров для которых userId является другом
func (s *Storage) Followers(userId uint64) (followers []uint64) {
	s.RLock()
	for _, v := range s.users {
		if _, ok := v.Friends[userId]; ok {
			followers = append(followers, v.ID)
		}
	}
	s.RUnlock()
	return followers
}

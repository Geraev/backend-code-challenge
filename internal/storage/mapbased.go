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
	Friends      []uint64
	Connection   io.Writer
}

func NewStorage() *Storage {
	s := &Storage{
		RWMutex: new(sync.RWMutex),
		users:   make(map[uint64]User),
	}
	return s
}

func NewTestStorage() *Storage {
	return &Storage{
		RWMutex: new(sync.RWMutex),
		users: map[uint64]User{
			1: {
				ID:           1,
				OnlineStatus: true,
				Friends:      []uint64{1, 2, 3},
				Connection:   nil,
			},
		},
	}
}

// PutOrUpdateUser добавление либо обновление записи User
func (s *Storage) PutOrUpdateUser(userId uint64, friends []uint64, conn io.Writer) {
	s.Lock()
	if val, ok := s.users[userId]; ok {
		val.Friends = friends
		s.users[userId] = val
	} else {
		s.users[userId] = User{
			ID:           userId,
			OnlineStatus: true,
			Friends:      friends,
			Connection:   conn,
		}
	}
	s.Unlock()
}

// SetUserOffline установить статус offline
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

// RemoveUser удаление пользователя
func (s *Storage) RemoveUser(userId uint64) {
	s.Lock()
	delete(s.users, userId)
	s.Unlock()
}

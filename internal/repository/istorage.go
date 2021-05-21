package repository

import (
	"io"
)

type IStorage interface {
	AddUser(userId uint64, friends []uint64, conn io.Writer)
	SetUserOffline(userId uint64) (isOk bool)
	GetUserStatus(userId uint64) (onlineStatus bool)
	GetUserConn(userId uint64) (conn io.Writer, ok bool)
	RemoveUser(userId uint64)
	Followers(userId uint64) (followers []uint64)
}

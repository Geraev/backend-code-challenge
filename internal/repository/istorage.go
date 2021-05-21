package repository

import "io"

type IStorage interface {
	PutOrUpdateUser(userId uint64, friends []uint64, conn io.Writer)
	SetUserOnlineStatus(userId uint64, status bool) (isOk bool)
	RemoveUser(userId uint64)
}

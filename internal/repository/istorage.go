package repository

import "io"

type IStorage interface {
	PutOrUpdateUser(userId uint64, friends []uint64, conn io.Writer)
	SetUserOffline(userId uint64) (isOk bool)
	RemoveUser(userId uint64)
}

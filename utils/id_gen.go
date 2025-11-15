package utils

import (
	"sync"
	"time"
)

var (
	lastID int64
	mutex  sync.Mutex
)

// Генерирует уникальный ID на основе времени
func GenerateID() int64 {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now().UnixNano()
	// Гарантируем уникальность даже при быстром создании
	if now <= lastID {
		now = lastID + 1
	}
	lastID = now

	return now
}

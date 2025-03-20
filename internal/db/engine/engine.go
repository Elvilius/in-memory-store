package engine

import "sync"

type Engine struct {
	hashTable map[string]string
	mutex     sync.RWMutex
}

func New() *Engine {
	return &Engine{
		hashTable: map[string]string{},
	}
}

func (e *Engine) Set(key string, value string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.hashTable[key] = value
}

func (e *Engine) Get(key string) (string, bool) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	data, ok := e.hashTable[key]
	return data, ok
}

func (e *Engine) Del(key string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	delete(e.hashTable, key)
}

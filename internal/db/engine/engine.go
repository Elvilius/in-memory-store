package engine

type Engine struct {
	hashTable map[string]string
}

func New() *Engine {
	return &Engine{
		hashTable: map[string]string{},
	}
}

func (e *Engine) Set(key string, value string) {
	e.hashTable[key] = value
}

func (e *Engine) Get(key string) (string, bool) {
	data, ok := e.hashTable[key]
	return data, ok
}

func (e *Engine) Del(key string) {
	delete(e.hashTable, key)
}

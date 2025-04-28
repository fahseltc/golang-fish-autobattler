package environment

import "sync"

type Registry struct {
	values map[string]any
	sync.Mutex
}

func NewRegistry() *Registry {
	return &Registry{
		values: make(map[string]any),
	}
}

func (r *Registry) Add(k string, v any) {
	if r == nil {
		return
	}

	r.Lock()
	defer r.Unlock()
	r.values[k] = v
}

func (r *Registry) Get(k string) any {
	if r == nil {
		return nil
	}

	r.Lock()
	defer r.Unlock()
	return r.values[k]
}

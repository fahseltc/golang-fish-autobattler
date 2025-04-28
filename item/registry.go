package item

import "fmt"

type Registry struct {
	Items map[string]Item
}

func NewRegistry() *Registry {
	return &Registry{
		Items: make(map[string]Item),
	}
}

func (r *Registry) Add(name string, item Item) error {
	if _, ok := r.Items[name]; ok {
		return fmt.Errorf("Item with the same name already exists: %v", name)
	}

	r.Items[name] = item
	return nil
}

func (r *Registry) Get(name string) (Item, bool) {
	fmt.Printf("Getting item from registry: %v\n", name)
	fmt.Printf("Items in registry: %v\n", r.Items)
	if item, ok := r.Items[name]; ok {
		item.RegenerateUuid()
		return item, false
	} else {
		return Item{}, true // error, Item cant be nil
	}
}

func (r *Registry) GetAll() map[string]Item {
	return r.Items
}

package environment

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"sync"
)

type Config struct {
	values map[string]any
	sync.Mutex
}

func NewConfig() *Config {
	config := &Config{
		values: make(map[string]any),
	}
	jsonFile, err := os.Open("data/config.json")
	if err != nil {
		log.Fatal("unable to open config file")
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	decoder.UseNumber()
	if err != nil {
		log.Fatal("unable to decode config file")
	}
	var data map[string]interface{}

	for decoder.More() {
		decoder.Decode(&data)
	}

	for key, val := range data {
		switch val := val.(type) {
		case json.Number:
			if n, err := val.Int64(); err == nil {
				config.Add(key, int(n))
				break
			} else if f, err := val.Float64(); err == nil {
				config.Add(key, f)
			} else {
				config.Add(key, val)
			}
		case string:
			config.Add(key, val)
		default:
			log.Fatalf("Unknown JSON type in config, key: %v, val:%v, type:%v", key, val, reflect.TypeOf(val))
		}
	}

	return config
}

func (r *Config) Add(k string, v any) {
	if r == nil {
		return
	}

	r.Lock()
	defer r.Unlock()
	r.values[k] = v
}

func (r *Config) Get(k string) any {
	if r == nil {
		return nil
	}

	r.Lock()
	defer r.Unlock()
	return r.values[k]
}

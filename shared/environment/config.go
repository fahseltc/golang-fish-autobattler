package environment

import (
	"encoding/json"
	"log"
	"os"
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
	var process func(prefix string, m map[string]interface{})
	process = func(prefix string, m map[string]interface{}) {
		for key, val := range m {
			fullKey := key
			if prefix != "" {
				fullKey = prefix + "." + key
			}
			switch v := val.(type) {
			case map[string]interface{}:
				process(fullKey, v)
			case json.Number:
				if n, err := v.Int64(); err == nil {
					config.Add(fullKey, int(n))
				} else if f, err := v.Float64(); err == nil {
					config.Add(fullKey, f)
				} else {
					config.Add(fullKey, v)
				}
			case string:
				config.Add(fullKey, v)
			default:
				config.Add(fullKey, v)
				//log.Fatalf("Unknown JSON type in config, key: %v, val:%v, type:%v", fullKey, v, reflect.TypeOf(v))
			}
		}
	}

	process("", data)

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

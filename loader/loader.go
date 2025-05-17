package loader

import (
	"encoding/csv"
	"fishgame/environment"
	"fishgame/item"
	"log"
	"os"
	"strconv"
	"sync"
)

var lock = &sync.Mutex{}

type ItemRegistrySingleton struct {
	Reg *item.Registry
}

var singleInstance *ItemRegistrySingleton

func GetFishRegistry(env *environment.Env) (*ItemRegistrySingleton, error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			instance, err := loadItemRegistry(env)
			if err != nil {
				return nil, err
			}
			singleInstance = &ItemRegistrySingleton{Reg: instance}
		}
	}

	return singleInstance, nil
}

func loadItemRegistry(env *environment.Env) (*item.Registry, error) {
	registry := item.NewRegistry()
	err := parseItems(env, "data/fish/initial_fish.csv", registry)
	if err != nil {
		return nil, err
	}
	err = parseItems(env, "data/fish/t1_fish.csv", registry)
	if err != nil {
		return nil, err
	}
	return registry, nil
}

func parseItems(Env *environment.Env, filepath string, reg *item.Registry) error {
	fileFish, err := readCsvFile(filepath)
	if err != nil {
		Env.Error("unable to read fish csv", "filepath", filepath)
		log.Fatalf("unable to read fish csv: %v\n", filepath)
	}

	// loop  fish and create items
	for _, fish := range fileFish {
		life, _ := strconv.Atoi(fish[2])
		duration, _ := strconv.ParseFloat(fish[3], 32)
		damage, _ := strconv.Atoi(fish[4])
		itemType := item.TypeFromString(fish[1])
		var behaviorFunc func(*item.Item, *item.Item) bool
		switch itemType {
		case item.Weapon:
			behaviorFunc = item.AttackingBehavior
		case item.Reactive:
			behaviorFunc = item.ReactingBehavior
		default:
			behaviorFunc = nil
		}

		// todo update fish function type from file, or use switch statement
		item := item.NewItem(Env, fish[0], item.TypeFromString(fish[1]), fish[5], life, float64(duration), int(damage), behaviorFunc)
		added := reg.Add(fish[0], *item)
		if added != nil {
			Env.Error("failed to add duplicate item to registry", "filepath", filepath, "item", item)
		}
	}
	return nil
}

func readCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	// remove the first line of the csv file because its a header
	csvReader.Read()
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

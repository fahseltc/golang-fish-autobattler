package loader

import (
	"encoding/csv"
	"fishgame/environment"
	"fishgame/item"
	"os"
	"strconv"
)

func LoadCsv(Env environment.Env) *item.Registry {
	registry := item.NewRegistry()
	attackingFish, err := readCsvFile("data/attacking_fish.csv")
	if err != nil {
		panic(err)
	}

	// loop attacking fish and create items
	for _, fish := range attackingFish {
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
		item := item.NewItem(Env, fish[0], item.TypeFromString(fish[1]), life, float64(duration), int(damage), behaviorFunc)
		registry.Add(fish[0], *item)
	}
	return registry
}

func readCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	// remove the first line of the csv file
	csvReader.Read()
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

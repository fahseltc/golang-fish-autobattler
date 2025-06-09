package fish

import (
	"encoding/csv"
	"fishgame/data"
	"fishgame/shared/environment"
	"log"
	"strconv"
)

type FishStatsRegistry struct {
	stats map[string]Stats
}

func NewFishStatsRegistry(env *environment.Env) *FishStatsRegistry {
	reg := &FishStatsRegistry{
		stats: make(map[string]Stats),
	}

	return reg
}

func (reg *FishStatsRegistry) loadItems(env *environment.Env) {
	parseItems(env, "fish/initial_fish.csv", registry)

	parseItems(env, "fish/t1_fish.csv", registry)
}

func (reg *FishStatsRegistry) parseCsv(Env *environment.Env, filepath string) {
	fileFish, err := readCsvFile(filepath)
	if err != nil {
		Env.Error("unable to read fish csv", "filepath", filepath)
		log.Fatalf("unable to read fish csv: %v\n", filepath)
	}

	// loop  fish and create fishs
	for _, fish := range fileFish {
		name := fish[0]
		fishType := fish.TypeFromString(fish[1])
		life, _ := strconv.Atoi(fish[2])
		size := fish.SizeFromString(fish[3])
		duration, _ := strconv.ParseFloat(fish[4], 32)
		damage, _ := strconv.Atoi(fish[5])
		description := fish[6]

		var behaviorFunc func(*fish.Item, *fish.Item, *fish.BehaviorProps) bool
		switch fishType {
		case fish.Weapon:
			behaviorFunc = fish.AttackingBehavior
		case fish.SizeBasedWeapon:
			behaviorFunc = fish.LargerSizeAttackingBehavior
		case fish.AdjacencyBasedWeapon:
			behaviorFunc = fish.AdjacentAttackingBehavior
		case fish.Reactive:
			behaviorFunc = fish.ReactingBehavior
		case fish.VenomousBasedWeapon:
			behaviorFunc = fish.VenomousBehavior
		default:
			behaviorFunc = nil
		}

		// todo update fish function type from file, or use switch statement
		fish := fish.NewItem(Env, nil, name, fishType, size, description, life, float64(duration), int(damage), behaviorFunc)
		added := reg.Add(name, *fish)
		if added != nil {
			Env.Error("failed to add duplicate fish to registry", "filepath", filepath, "fish", fish)
		}
	}
	return nil
}

func readCsvFile(filePath string) ([][]string, error) {
	data, err := data.Files.Open(filePath)
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

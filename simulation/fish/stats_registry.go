package fish

import (
	"encoding/csv"
	"fishgame/data"
	"fishgame/shared/environment"
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type FishStatsRegistry struct {
	stats map[string]Stats
	fish  map[string]Fish
}

func NewFishStatsRegistry(env *environment.Env) *FishStatsRegistry {
	reg := &FishStatsRegistry{
		stats: make(map[string]Stats),
		fish:  make(map[string]Fish),
	}
	reg.parseCsv(env, "fish/initial_fish.csv")
	reg.parseCsv(env, "fish/t1_fish.csv")
	return reg
}

func (reg *FishStatsRegistry) parseCsv(Env *environment.Env, filepath string) {
	fileFish, err := readCsvFile(filepath)
	if err != nil {
		Env.Error("unable to read fish csv", "filepath", filepath)
		log.Fatalf("unable to read fish csv: %v\n", filepath)
	}

	// loop fish
	for _, fish := range fileFish {
		name := fish[0]
		fishType := TypeFromString(fish[1])
		life, _ := strconv.Atoi(fish[2])
		size := SizeFromString(fish[3])
		duration, _ := strconv.ParseFloat(fish[4], 32)
		damage, _ := strconv.Atoi(fish[5])
		description := fish[6]

		stats := NewStats(fishType, size, life, duration, damage)
		fish := NewFish(Env, name, description, &stats)

		reg.addFish(name, *fish)
		reg.addStat(name, stats)
	}
}

func readCsvFile(filePath string) ([][]string, error) {
	data, err := data.Files.Open(filePath)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(data)
	// remove the first line of the csv file because its a header
	csvReader.Read()
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *FishStatsRegistry) addStat(name string, stats Stats) error {
	if _, ok := r.stats[name]; ok {
		return fmt.Errorf("Fish with the same name already exists: %v", name)
	}
	r.stats[name] = stats
	return nil
}

func (r *FishStatsRegistry) addFish(name string, fish Fish) error {
	if _, ok := r.fish[name]; ok {
		return fmt.Errorf("Fish with the same name already exists: %v", name)
	}
	r.fish[name] = fish
	return nil
}

func (r *FishStatsRegistry) GetFish(name string) (*Fish, error) {
	if fish, ok := r.fish[name]; !ok {
		return nil, fmt.Errorf("fish: '%v' not found", name)
	} else {
		newFishInstance := fish
		newFishInstance.Id, _ = uuid.NewUUID()
		statsCopy := *fish.Stats // copy the stats instance so each fish has its own
		newFishInstance.Stats = &statsCopy
		return &newFishInstance, nil
	}
}
func (r *FishStatsRegistry) GetStat(name string) (*Stats, error) {
	if stat, ok := r.stats[name]; !ok {
		return nil, fmt.Errorf("stat with name: '%v' not found", name)
	} else {
		return &stat, nil
	}
}

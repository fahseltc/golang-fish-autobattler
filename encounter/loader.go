package encounter

import (
	"encoding/json"
	"fishgame/data"
	"io"
	"log"
)

type jsonFile struct {
	Encounters []jsonEncounter
}

type jsonEncounter struct {
	Title   string       `json:"title"`
	Type    string       `json:"encounter_type"`
	Enemies []jsonEnemy  `json:"enemies"`
	Rewards []jsonReward `json:"rewards"`
}

type jsonReward struct {
	Fish     []string `json:"fish"`
	Currency int      `json:"currency"`
}

type jsonEnemy struct {
	Name      string `json:"name"`
	SlotIndex int    `json:"slot"`
}

func LoadEncounterFile(path string) jsonFile {
	file, err := data.Files.Open(path)
	if err != nil {
		log.Fatalf("could not load encounter file at path: %v", path)
	}
	defer file.Close()
	jsonBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to ReadAll encounter file: %v", path)
	}
	var encounterData jsonFile
	err = json.Unmarshal(jsonBytes, &encounterData)
	if err != nil {
		log.Fatalf("unable to Unmarshal encounter file: %v", path)
	}
	return encounterData
}

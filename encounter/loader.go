package encounter

import (
	"encoding/json"
	"fishgame/environment"
	"fishgame/item"
	"fishgame/loader"
	"fishgame/player"
	"fishgame/reward"
	"fishgame/ui"
	"fishgame/util"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type jsonFile struct {
	Encounters []jsonEncounter `json:"encounters"`
}

type jsonEncounter struct {
	FishNames []string     `json:"items"`
	Title     string       `json:"title"`
	Buttons   []buttonJson `json:"buttons"`
	Type      string       `json:"encounter_type"`
	Rewards   []jsonReward `json:"rewards"`
}

type jsonReward struct {
	Type     string `json:"type"`
	ItemName string `json:"item"`
	Currency int    `json:"currency"`
}

type buttonJson struct {
	Behavior buttonBehavior `json:"behavior"`
	X        int            `json:"x"`
	Y        int            `json:"y"`
	W        int            `json:"w"`
	H        int            `json:"h"`
	Text     string         `json:"text"`
	Color    string         `json:"color"`
	FontSize int            `json:"font_size"`
}

type buttonBehavior struct {
	Type      string   `json:"type"`
	ItemNames []string `json:"items"`
}

func LoadEncounters(env *environment.Env, path string, player *player.Player, mgr *Manager) []EncounterInterface {
	var encounters []EncounterInterface
	jsf, err := os.Open(path)
	if err != nil {
		log.Fatalf("unable to open encounter file: %v", path)
	}
	defer jsf.Close()
	jsonBytes, err := io.ReadAll(jsf)
	if err != nil {
		log.Fatalf("unable to ReadAll encounter file: %v", path)
	}

	var encounterdata jsonFile
	err = json.Unmarshal(jsonBytes, &encounterdata)
	if err != nil {
		log.Fatalf("unable to Unmarshal encounter file: %v", path)
	}

	for _, encData := range encounterdata.Encounters {
		enc := parseJson(env, encData, player, mgr)
		encounters = append(encounters, enc)
	}

	return encounters
}

func parseJson(env *environment.Env, encounterData jsonEncounter, player *player.Player, mgr *Manager) EncounterInterface {
	itemsReg, _ := loader.GetFishRegistry(env) // TODO handle errors
	font, _ := util.LoadFont(30)               // TODO handle errors

	var enc EncounterInterface
	switch TypeFromString(encounterData.Type) {
	case EncounterTypeInitial:
		enc = generateInitialEncounter(env, font, encounterData, player, mgr, itemsReg.Reg)
	case EncounterTypeBattle:
		enc = generateBattleEncounter(env, font, encounterData, player, mgr, itemsReg.Reg)
	}
	enc = generateRewards(env, enc, encounterData, itemsReg.Reg)
	return enc
}

func generateRewards(env *environment.Env, enc EncounterInterface, encounterData jsonEncounter, reg *item.Registry) EncounterInterface {
	for _, data := range encounterData.Rewards {
		it, _ := reg.Get(data.ItemName)

		reward := reward.NewReward(env, reward.TypeFromString(data.Type), &it, data.Currency)
		enc.AddReward(reward)
	}
	return enc
}

func generateInitialEncounter(env *environment.Env, font text.Face, encounterData jsonEncounter, player *player.Player, mgr *Manager, itemsReg *item.Registry) EncounterInterface {
	enc := &Initial{
		env:     env,
		manager: mgr,
		Type:    TypeFromString(encounterData.Type),

		player:     player,
		text:       encounterData.Title,
		bg:         util.LoadImage(env, "assets/bg/initial.png"),
		font:       &font,
		itemChosen: false,
	}
	enc.buttons = generateInitialButtons(env, encounterData, enc, player, mgr, itemsReg)
	return enc
}

func generateBattleEncounter(env *environment.Env, font text.Face, encounterData jsonEncounter, player *player.Player, mgr *Manager, itemsReg *item.Registry) EncounterInterface {
	enc := NewBattleScene(env, encounterData, player)
	enc.items = generateBattleItems(env, encounterData, itemsReg)
	// todo: add to ui/slots?
	return enc
}

func generateInitialButtons(env *environment.Env, encounterData jsonEncounter, enc *Initial, player *player.Player, mgr *Manager, itemsReg *item.Registry) []*ui.Button {
	var buttons []*ui.Button
	for _, btnData := range encounterData.Buttons {
		// behavior := btnData.Behavior.Type // TODO Use type to switch reading JSON
		btn := ui.NewButton(
			env,
			float32(btnData.X),
			float32(btnData.Y),
			float32(btnData.W),
			float32(btnData.H),
			btnData.Text,
			float64(btnData.FontSize),
		)
		var btnItems []*item.Item
		for _, itName := range btnData.Behavior.ItemNames {
			item, err := itemsReg.Get(itName)
			if err {
				env.Logger.Error("Unable to load item for encounter button", "item", itName)
			}
			btnItems = append(btnItems, &item)
		}

		btn.OnClick = func() {
			res := player.Items.AddItems(btnItems)
			env.Logger.Info("Added items", "items", btnItems, "recipient", player.Name, "success", res)
			enc.itemChosen = true
		}

		btn.ToolTip = ui.NewInitialTooltip(
			float32(btnData.X),
			float32(btnData.Y),
			float32(btnData.W)+150,
			float32(btnData.H)+100,
			btnItems[0], // Assuming the first item is the one to show in tooltip
		)

		buttons = append(buttons, btn)
	}

	return buttons
}

func generateBattleItems(env *environment.Env, encounterData jsonEncounter, itemsReg *item.Registry) *item.Collection {
	itemsArr := []*item.Item{}
	for _, itemName := range encounterData.FishNames {
		item, err := itemsReg.Get(itemName)
		if err {
			panic(err)
		}
		itemsArr = append(itemsArr, &item)
	}
	return item.NewEncounterCollection(env, itemsArr)
}

package encounter

import (
	"encoding/json"
	"fishgame/environment"
	"fishgame/item"
	"fishgame/loader"
	"fishgame/player"
	"fishgame/ui"
	"fishgame/util"
	"image/color"
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
	Type     string `json:"type"`
	ItemName string `json:"item_name"`
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
	//fmt.Printf("%v", encounterdata)

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
	enc := &Battle{
		env:    env,
		Name:   encounterData.Title,
		Type:   EncounterTypeBattle,
		player: player,
	}
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
			color.Black, // todo parse color from btnData.color string
			float64(btnData.FontSize),
		)
		item, err := itemsReg.Get(btnData.Behavior.ItemName)
		if err {
			env.Logger.Error("Unable to load item for encounter button", "item", btnData.Behavior.ItemName)
		}
		btn.OnClick = func() {
			player.Items.AddItem(&item)
			env.Logger.Info("Added item", "item", &item.Name, "recipient", player.Name)
			enc.itemChosen = true
		}

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

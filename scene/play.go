package scene

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/loader"
	"fishgame/player"
	"fishgame/ui"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Play struct {
	Env           *environment.Env
	SceneManager  *Manager
	State         GameState
	ItemsRegistry *item.Registry
	Ui            *ui.UI
	Player1       *player.Player
	Player2       *player.Player
}

func (s *Play) Init(sm *Manager) {
	s.SceneManager = sm
	s.State = PlayState
	s.ItemsRegistry = loader.LoadCsv(*s.Env)
	s.Ui = ui.NewUI(s.Env)

	//s.Ui = buildUi()
	s.Player1 = generatePlayer1(*s.Env, s.ItemsRegistry)
	s.Player2 = generatePlayer2(*s.Env, s.ItemsRegistry)

	for index, it := range s.Player1.Items.ActiveItems {
		s.Ui.Player1Slots[index].AddItem(index, it)
		it.SlotIndex = index
	}

	for index, it := range s.Player2.Items.ActiveItems {
		s.Ui.Player2Slots[index].AddItem(index, it)
		it.SlotIndex = index
	}

}

func (s *Play) Update(dt float64) error {
	switch s.State {
	case PlayState:
		return updatePlayState(s, dt)
	case MapState:
		return updateMapState(s, dt)
	case InventoryState:
		// Show the inventory
	case GameOverState:
		// Show the game over screen
	case PauseState:
		// Pause the game
	}

	return nil
}

func updatePlayState(s *Play, dt float64) error {
	if s.Ui != nil {
		s.Ui.Update()
	}
	//fmt.Println("Updating Player 1 items")
	s.Player1.Items.Update(dt, s.Player2.Items)
	if len(s.Player1.Items.ActiveItems) == 0 {
		// if player 1 has no active items, break the loop
		fmt.Println("Player 1 has no active items")
		//		log.Fatal("Player 1 has no active items")
		return fmt.Errorf("player 1 has no active items")
	}
	//fmt.Println("Updating Player 2 items")
	s.Player2.Items.Update(dt, s.Player1.Items)
	if len(s.Player2.Items.ActiveItems) == 0 {
		// if player 2 has no active items, break the loop
		return fmt.Errorf("player 2 has no active items")
	}

	return nil
}

func updateMapState(s *Play, dt float64) error {
	return nil
}

func (s *Play) Draw(screen *ebiten.Image) {
	switch s.State {
	case PlayState:
		if s.Player1 != nil {
			s.Player1.Items.Draw(*s.Env, screen, 1)
		}
		if s.Player2 != nil {
			s.Player2.Items.Draw(*s.Env, screen, 2)
		}
		// Draw the UI
		if s.Ui != nil {
			s.Ui.Draw(screen)
		}
	case MapState:
		return
	case InventoryState:
		return
	case GameOverState:
		return
	case PauseState:
		return
	}
}

func (s *Play) Destroy() {
	// Clean up resources if necessary
	s.Env = nil
	s.Player1 = nil
	s.Player2 = nil
}

func (s *Play) GetName() string {
	return "Play"
}

func (s *Play) changeGameState(state GameState) {
	s.State = state
}

func generatePlayer1(env environment.Env, items *item.Registry) *player.Player {
	item1, err := items.Get("Goldfish")
	if err {
		panic(err)
	}
	item2, err := items.Get("Eel")
	if err {
		panic(err)
	}
	item3, err := items.Get("Shark")
	if err {
		panic(err)
	}
	item4, err := items.Get("Puffer")
	if err {
		panic(err)
	}
	itemsArr := []*item.Item{&item1, &item2, &item3, &item4}
	coll := item.NewCollection(&env, 1, itemsArr)
	p := &player.Player{
		Env:   env,
		Name:  "Player 1",
		Items: coll,
	}
	fmt.Printf("Player 1 active items length: %d\n", len(p.Items.ActiveItems))

	return p
}

func generatePlayer2(env environment.Env, items *item.Registry) *player.Player {
	item1, err := items.Get("Shark")
	if err {
		panic(err)
	}
	item2, err := items.Get("Minnow")
	if err {
		panic(err)
	}
	item3, err := items.Get("Minnow")
	if err {
		panic(err)
	}
	item4, err := items.Get("Shark")
	if err {
		panic(err)
	}
	item5, err := items.Get("Whale")
	if err {
		panic(err)
	}
	itemsArr := []*item.Item{&item1, &item2, &item3, &item4, &item5}
	coll := item.NewCollection(&env, 2, itemsArr)

	p := &player.Player{
		Env:   env,
		Name:  "Player 2",
		Items: coll,
	}
	fmt.Printf("Player 1 active items length: %d\n", len(p.Items.ActiveItems))

	return p
}

// func buildUi() *ebitenui.UI {
// 	face, _ := util.LoadFont(20)
// 	rootContainer := widget.NewContainer(
// 		// the container will use a plain color as its background
// 		//widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

// 		// the container will use a grid layout to layout to split the layout in half down the middle
// 		widget.ContainerOpts.Layout(widget.NewGridLayout(
// 			widget.GridLayoutOpts.Columns(3),
// 			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(10)),
// 			widget.GridLayoutOpts.Spacing(10, 10),

// 			widget.GridLayoutOpts.Stretch([]bool{true, true}, []bool{true}),
// 		)),
// 	)

// 	centerColumn := widget.NewContainer(
// 		widget.ContainerOpts.Layout(
// 			widget.NewRowLayout(
// 				widget.RowLayoutOpts.Spacing(0),
// 				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
// 				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
// 			),
// 		),
// 	)

// 	leftSide := widget.NewContainer(
// 		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
// 		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255})),
// 		widget.ContainerOpts.WidgetOpts(
// 			//This command indicates this widget is the source of Drag and Drop.
// 			widget.WidgetOpts.EnableDragAndDrop(
// 				widget.NewDragAndDrop(
// 					//The object which will create/update the dragged element. Required.
// 					widget.DragAndDropOpts.ContentsCreater(&item.Slot{}),
// 					//How many pixels the user must drag their cursor before the drag begins.
// 					//This is an optional parameter that defaults to 15 pixels
// 					widget.DragAndDropOpts.MinDragStartDistance(15),
// 					//This sets where to anchor the widget to the cursor
// 					widget.DragAndDropOpts.ContentsOriginVertical(widget.DND_ANCHOR_MIDDLE),
// 					//This sets where to anchor the widget to the cursor
// 					widget.DragAndDropOpts.ContentsOriginHorizontal(widget.DND_ANCHOR_MIDDLE),
// 					//This sets of far off the cursor to offset the dragged element
// 					widget.DragAndDropOpts.Offset(img.Point{-5, -5}),
// 					//This will turn of Drag to initiate drag and drop
// 					//Primary use case will be click to drag
// 					//widget.DragAndDropOpts.DisableDrag(),
// 				),
// 			),
// 			widget.WidgetOpts.MouseButtonReleasedHandler(func(args *widget.WidgetMouseButtonReleasedEventArgs) {
// 				if args.Inside && args.Button == ebiten.MouseButtonLeft && ebiten.IsKeyPressed(ebiten.KeyControl) {
// 					args.Widget.DragAndDrop.StartDrag()
// 				}
// 				if args.Button == ebiten.MouseButtonRight {
// 					args.Widget.DragAndDrop.StopDrag()
// 				}
// 			}),
// 			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
// 				Position: widget.RowLayoutPositionCenter,
// 				Stretch:  true,
// 			}),
// 		),
// 	)

// 	leftSide.AddChild(
// 		widget.NewText(widget.TextOpts.Text("Drag from Here\nOr Ctrl-Click", face, color.Black), widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
// 			HorizontalPosition: widget.AnchorLayoutPositionCenter,
// 			VerticalPosition:   widget.AnchorLayoutPositionCenter,
// 		}))))
// 	centerColumn.AddChild(widget.NewContainer())
// 	centerColumn.AddChild(leftSide)
// 	rootContainer.AddChild(widget.NewContainer())
// 	rootContainer.AddChild(centerColumn)

// 	//rootContainer.AddChild(centerColumn)
// 	rootContainer.AddChild(widget.NewContainer())

// 	//rootContainer.AddChild(widget.NewContainer())

// 	//rootContainer.AddChild(leftSide)

// 	// window := widget.NewContainer()
// 	// window.AddChild(leftSide)
// 	// window.SetLocation(img.Rect(0, 0, 400, 400))
// 	//rootContainer.AddChild(window)

// 	ui := &ebitenui.UI{
// 		Container: rootContainer,
// 	}

// 	return ui
// }

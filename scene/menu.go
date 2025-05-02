package scene

import (
	"fishgame/environment"
	"fishgame/util"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Menu struct {
	Env          *environment.Env
	SceneManager *Manager
	Ui           *ebitenui.UI
	StartBtn     *widget.Button
}

func (m *Menu) Init(sm *Manager) {
	m.SceneManager = sm

	bg := util.LoadImage(*m.Env, "assets/bg/menu.png")
	//bg9 := image.NewNineSliceSimple(bg, 180, 200)
	bg9 := image.NewNineSlice(bg, [3]int{180, 440, 180}, [3]int{200, 200, 200})

	// construct a new container that serves as the root of the UI hierarchy.
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background.
		widget.ContainerOpts.BackgroundImage(bg9),

		// the container will use an anchor layout to layout its single child widget.
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	m.StartBtn = createStartButton(sm)
	// construct a button.

	// add the button as a child of the container.
	rootContainer.AddChild(m.StartBtn)

	// construct the UI.
	m.Ui = &ebitenui.UI{
		Container: rootContainer,
	}
}

func (m *Menu) Update(dt float64) error {
	if m.Ui != nil {
		m.Ui.Update()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		m.StartBtn.Click()
	}

	// Test that you can call Click on the focused widget.
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if btn, ok := m.Ui.GetFocusedWidget().(*widget.Button); ok {
			btn.Click()
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		m.StartBtn.Press()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyG) {
		m.StartBtn.Release()
	}
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	if m.Ui != nil {
		m.Ui.Draw(screen)
	}
}

func (m *Menu) Destroy() {
	// Clean up menu resources here
}

func (m *Menu) GetName() string {
	return "Menu"
}

func loadButtonImage() (*widget.ButtonImage, error) {

	idle := image.NewBorderedNineSliceColor(color.NRGBA{R: 184, G: 253, B: 245, A: 255}, color.NRGBA{90, 90, 90, 255}, 3)

	hover := image.NewBorderedNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}, color.NRGBA{70, 70, 70, 255}, 3)

	pressed := image.NewAdvancedNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}, image.NewBorder(3, 2, 2, 2, color.NRGBA{70, 70, 70, 255}))

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func createStartButton(sm *Manager) *widget.Button {

	// Initialize menu resources here
	buttonImage, _ := loadButtonImage()

	// load button text font.
	face, _ := util.LoadFont(20)

	var button *widget.Button

	button = widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
				Padding:            widget.Insets{Top: 0, Left: 0, Right: 0, Bottom: 100},
			}),
		),
		// specify the images to use.
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color.
		// widget.ButtonOpts.Text("Hello, World!", face, &widget.ButtonTextColor{
		widget.ButtonOpts.Text("Start Game!", face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0, 0, 0, 255},
		}),
		widget.ButtonOpts.TextProcessBBCode(false),
		// specify that the button's text needs some padding for correct display.
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
		// Move the text down and right on press
		widget.ButtonOpts.PressedHandler(func(args *widget.ButtonPressedEventArgs) {
			button.Text().Inset.Top = 1
			button.GetWidget().CustomData = true
		}),
		// Move the text back to start on press released
		widget.ButtonOpts.ReleasedHandler(func(args *widget.ButtonReleasedEventArgs) {
			button.Text().Inset.Top = 0
			button.GetWidget().CustomData = false
		}),

		// add a handler that reacts to clicking the button.
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			sm.SwitchTo("Play", true)
		}),
	)
	return button
}

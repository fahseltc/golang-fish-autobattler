package reward

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
	"fishgame/ui"
	"fishgame/util"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Reward struct {
	env      *environment.Env
	Type     Type
	Obtained bool
	Item     *item.Item
	Currency int

	btn *ui.Button
	bg  *ebiten.Image

	btnPressed bool
}

func NewReward(env *environment.Env, newType Type, it *item.Item, curr int) *Reward {
	r := &Reward{
		env:        env,
		Type:       newType,
		Obtained:   false,
		Item:       it,
		Currency:   curr,
		btnPressed: false,
		bg:         util.LoadImage(env, "assets/ui/panel/grey_panel.png"),
	}

	r.btn = ui.NewButton(env,
		ui.WithRect(ui.Rectangle{X: 400, Y: 400, W: 200, H: 50}),
		ui.WithText("Get Reward!"),
		ui.WithClickFunc(func() {
			r.btnPressed = true
		}),
		ui.WithCenteredPos(),
	)

	return r
}

func (r *Reward) Obtain(player *player.Player) bool {
	if !r.Obtained {
		switch r.Type {
		case Item:
			if r.Item != nil {
				result := player.Items.AddItem(r.Item)
				if result {
					r.Obtained = true
					r.env.Logger.Info("player obtained reward item", "item", r.Item.Name)
				}
				return result
			}
			return false

		case Currency:
			if r.Currency != 0 {
				player.Currency += r.Currency
				r.Obtained = true
				r.env.Logger.Info("player obtained reward currency", "amount", r.Currency)
				return true
			}
			return false
		}
	}
	return false
}

func (r *Reward) Update(player *player.Player) {
	r.btn.Update()
	if r.btnPressed {
		r.Obtain(player)
	}
}

func (r *Reward) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 3)
	opts.GeoM.Translate(float64(200), float64(200))
	screen.DrawImage(r.bg, opts)
	if r.Item.Sprite != nil {
		opts = &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(325), float64(250))
		screen.DrawImage(r.Item.Sprite, opts)
	}
	if r.Currency > 0 {
		font, _ := util.LoadFont(20)
		txtColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}
		ui.DrawCenteredText(screen, font, fmt.Sprintf("%v", r.Currency), 400, 300, txtColor)
		// draw currency icon
	}

	r.btn.Draw(screen)
}

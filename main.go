package main

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/player"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	Env             *environment.Env
	Player1         *player.Player
	Player2         *player.Player
	TargetFPS       int
	TargetFrameTime float32
}

func main() {
	logger := SetupLogger()

	registry := environment.NewRegistry()
	registry.Add("targetFPS", 2)
	registry.Add("targetFrameTime", float32(1.0/2))

	env := environment.NewEnv(logger, registry)

	player1 := GeneratePlayer1(*env)
	player2 := GeneratePlayer2(*env)

	g := &Game{env, player1, player2, registry.Get("targetFPS").(int), registry.Get("targetFrameTime").(float32)}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

func (g *Game) Update(*ebiten.Image) error {
	// get the start time of the frame
	startTime := time.Now()

	// update the item collections
	fmt.Println("Updating Player 1 items")
	g.Player1.Items.Update(g.TargetFrameTime, g.Player2.Items)
	if len(g.Player1.Items.ActiveItems) == 0 {
		// if player 1 has no active items, break the loop
		fmt.Println("Player 1 has no active items")
		//		log.Fatal("Player 1 has no active items")
		return fmt.Errorf("player 1 has no active items")
	}
	fmt.Println("Updating Player 2 items")
	g.Player2.Items.Update(g.TargetFrameTime, g.Player1.Items)
	if len(g.Player2.Items.ActiveItems) == 0 {
		// if player 2 has no active items, break the loop
		return fmt.Errorf("player 2 has no active items")
	}

	// get the end time of the frame
	endTime := time.Now()

	// calculate the time taken for the frame
	frameTime := endTime.Sub(startTime).Seconds()

	// sleep for the remaining time to maintain the target frame rate
	if frameTime < float64(g.TargetFrameTime) {
		time.Sleep(time.Duration((float64(g.TargetFrameTime)-frameTime)*1000) * time.Millisecond)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func GeneratePlayer1(env environment.Env) *player.Player {
	p := &player.Player{
		Env:  env,
		Name: "Player 1",
		Items: &item.Collection{
			ActiveItems: []*item.Item{
				item.NewItem(env, "Goldfish", 10, item.Weapon, 1.0, 1, item.AttackingItem, nil),
			},
			InactiveItems: []*item.Item{},
		},
	}
	fmt.Printf("Player 1 active items length: %d\n", len(p.Items.ActiveItems))

	return p
}

func GeneratePlayer2(env environment.Env) *player.Player {
	p := &player.Player{
		Env:  env,
		Name: "Player 1",
		Items: &item.Collection{
			ActiveItems: []*item.Item{
				item.NewItem(env, "Shark", 20, item.Weapon, 1.5, 2, item.AttackingItem, nil),
			},
			InactiveItems: []*item.Item{},
		},
	}
	fmt.Printf("Player 1 active items length: %d\n", len(p.Items.ActiveItems))

	return p
}

func SetupLogger() *slog.Logger {

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.Attr{
						Key:   slog.TimeKey,
						Value: slog.Int64Value(t.Unix()),
					}
				}
			}
			return a
		},
	})

	return slog.New(handler)
}

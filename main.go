package main

import (
	"fishgame/environment"
	"fishgame/item"
	"fishgame/loader"
	"fishgame/player"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	Env             *environment.Env
	TargetFPS       int
	TargetFrameTime float32
	Player1         *player.Player
	Player2         *player.Player
}

func main() {
	logger := SetupLogger()

	registry := environment.NewRegistry()
	registry.Add("targetFPS", 2)
	registry.Add("targetFrameTime", float32(1.0/2)) // 1.0 / 2 = 0.5 seconds per frame

	env := environment.NewEnv(logger, registry)

	itemRegistry := loader.LoadCsv(*env)
	player1 := GeneratePlayer1(*env, itemRegistry)
	player2 := GeneratePlayer2(*env, itemRegistry)

	g := &Game{env, registry.Get("targetFPS").(int), registry.Get("targetFrameTime").(float32), player1, player2}

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
	g.Player1.Items.Draw(*g.Env, screen, 1)
	g.Player2.Items.Draw(*g.Env, screen, 2)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func GeneratePlayer1(env environment.Env, items *item.Registry) *player.Player {
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
	p := &player.Player{
		Env:  env,
		Name: "Player 1",
		Items: &item.Collection{
			ActiveItems: []*item.Item{
				&item1,
				&item2,
				&item3,
			},
			InactiveItems: []*item.Item{},
		},
	}
	fmt.Printf("Player 1 active items length: %d\n", len(p.Items.ActiveItems))

	return p
}

func GeneratePlayer2(env environment.Env, items *item.Registry) *player.Player {
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

	p := &player.Player{
		Env:  env,
		Name: "Player 2",
		Items: &item.Collection{
			ActiveItems: []*item.Item{
				&item1,
				&item2,
				&item3,
				&item4,
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

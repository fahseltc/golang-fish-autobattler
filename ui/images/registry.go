package images

import (
	"fishgame/ui/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Registry struct {
	Images map[string]*ebiten.Image
}

func NewRegistry() *Registry {
	reg := &Registry{
		Images: make(map[string]*ebiten.Image),
	}
	reg.Images = reg.loadFromDir("assets")
	return reg
}

func (reg *Registry) loadFromDir(path string) map[string]*ebiten.Image {
	images := make(map[string]*ebiten.Image)
	dir, _ := os.Getwd()
	fmt.Println("Current working directory:", dir)
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Walk error at:", path, "error:", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(filePath))
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			return nil
		}
		// file, err := os.Open(filePath)
		// if err != nil {
		// 	return err
		// }
		// defer file.Close()

		ebitenImg := util.LoadImage(filePath)

		relPath, err := filepath.Rel(path, filePath)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath) // use forward slashes

		images[relPath] = ebitenImg
		return nil
	})

	if err != nil {
		panic(err)
	}

	return images
}

// package item

// import "fmt"

// type Registry struct {
// 	Items map[string]Item
// }

// func NewRegistry() *Registry {
// 	return &Registry{
// 		Items: make(map[string]Item),
// 	}
// }

// func (r *Registry) Add(name string, item Item) error {
// 	if _, ok := r.Items[name]; ok {
// 		return fmt.Errorf("Item with the same name already exists: %v", name)
// 	}

// 	r.Items[name] = item
// 	return nil
// }

// func (r *Registry) Get(name string) (Item, bool) {
// 	//fmt.Printf("Getting item from registry: %v\n", name)
// 	//fmt.Printf("Items in registry: %v\n", r.Items)
// 	if item, ok := r.Items[name]; ok {
// 		item.RegenerateUuid()
// 		return item, false
// 	} else {
// 		return Item{}, true // error, Item cant be nil
// 	}
// }

// func (r *Registry) GetAll() map[string]Item {
// 	return r.Items
// }

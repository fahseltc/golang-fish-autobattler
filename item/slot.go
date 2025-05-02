package item

import (
	"fishgame/util"
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type Slot struct {
	// Slot is a struct that represents a slot in the inventory.
	// It contains the item that is currently in the slot, and a boolean that indicates whether the slot is empty or not.
	// The slot can be used to store items in the inventory, and can be used to check if the slot is empty or not.
	Object         *widget.Container
	Text           *widget.Text
	TargetedObject widget.HasWidget
}

func (slot *Slot) Create(parent widget.HasWidget) (*widget.Container, interface{}) {
	// For this example we do not need to recreate the Dragged element each time. We can re-use it.
	if slot.Object == nil {
		// load text font
		face, _ := util.LoadFont(20)
		slot.Object = widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
			widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 200, 100, 255})),
		)

		slot.Text = widget.NewText(widget.TextOpts.Text("Cannot Drop", face, color.Black), widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})))

		slot.Object.AddChild(slot.Text)
	}
	// return the container to be dragged and any arbitrary data associated with this operation
	return slot.Object, "Hello World"
}

// This method is optional for Drag and Drop
// It will be called every draw cycle that the Drag and Drop is active.
// Inputs:
//   - canDrop - if the cursor is over a widget that allows this object to be dropped
//   - targetWidget - The widget that will allow this object to be dropped.
//   - dragData - The drag data provided by the Create method above.
func (slot *Slot) Update(canDrop bool, targetWidget widget.HasWidget, dragData interface{}) {
	if canDrop {
		slot.Text.Label = "* Can Drop *"
		if targetWidget != nil {
			targetWidget.(*widget.Container).BackgroundImage = image.NewNineSliceColor(color.NRGBA{100, 100, 255, 255})
			slot.TargetedObject = targetWidget
		}
	} else {
		slot.Text.Label = "Cannot Drop"
		if slot.TargetedObject != nil {
			slot.TargetedObject.(*widget.Container).BackgroundImage = image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255})
			slot.TargetedObject = nil
		}
	}
}

func (slot *Slot) EndDrag(dropped bool, sourceWidget widget.HasWidget, dragData interface{}) {
	if dropped {
		fmt.Println("Dropped Successful")
	} else {
		fmt.Println("Drop Cancelled")
	}
}

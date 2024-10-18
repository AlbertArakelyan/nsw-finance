package passabletab

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type PassableTab struct {
	// DB       passablerepository.Repository
	InfpLog  *log.Logger
	ErrorLog *log.Logger
}

func (passableTab *PassableTab) GetPassableTab() *fyne.Container {
	return container.NewCenter(canvas.NewText("Passable Tab, Comming soon ❤️", nil))
}

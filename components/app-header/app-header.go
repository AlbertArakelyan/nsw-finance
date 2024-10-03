package appheader

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func GetAppHeaderContainer() *fyne.Container {
	title := canvas.NewText("NSW Finance", theme.Color(theme.ColorNamePrimary))
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	appHeaderContainer := container.NewStack(
		canvas.NewRectangle(theme.Color(theme.ColorNameBackground)), // Background color
		container.NewVBox(
			container.NewCenter(title),
			canvas.NewRectangle(theme.Color(theme.ColorNamePrimary)),
		),
	)

	return appHeaderContainer
}

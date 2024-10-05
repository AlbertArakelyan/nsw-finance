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

	logo := canvas.NewImageFromResource(resourceLogoJpg)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(40, 40))

	appHeaderContainer := container.NewStack(
		canvas.NewRectangle(theme.Color(theme.ColorNameBackground)), // Background color
		container.NewVBox(
			container.NewCenter(
				container.NewHBox(
					logo,
					container.NewCenter(title),
				),
			),
			canvas.NewRectangle(theme.Color(theme.ColorNamePrimary)),
		),
	)

	return appHeaderContainer
}

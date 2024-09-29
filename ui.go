package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *App) makeUI() {
	// get App Tabs
	savingsTab := app.savingsTab()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Savings", theme.FileIcon(), savingsTab),
		container.NewTabItemWithIcon("Passable", theme.ConfirmIcon(), canvas.NewText("Passable Tab", nil)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	// add container to window
	finalContent := container.NewVBox(tabs)

	app.MainWindow.SetContent(finalContent)
}

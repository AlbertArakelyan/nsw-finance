package main

import (
	appheader "nsw-finance/components/app-header"
	savingstab "nsw-finance/components/savings-tab"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *App) makeUI() {
	/**
	* Get App Header
	 */

	appHeaderContainer := appheader.GetAppHeaderContainer()

	/*
	* Get App Tabs
	 */

	// Savings Tab
	savingsTab := &savingstab.SavingsTab{
		DB:       app.DB,
		InfoLog:  app.Utils.InfoLog,
		ErrorLog: app.Utils.ErrorLog,
	}
	savingsContainer := savingsTab.GetSavingsTab()
	app.UIComponents.SavingsContainer = savingsContainer

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Savings", theme.FileIcon(), savingsContainer),
		container.NewTabItemWithIcon("Passable", theme.ConfirmIcon(), container.NewCenter(canvas.NewText("Passable Tab, Comming soon ❤️", nil))),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	// add container to window
	finalContent := container.NewBorder(appHeaderContainer, nil, nil, nil, tabs)

	app.MainWindow.SetContent(finalContent)
}

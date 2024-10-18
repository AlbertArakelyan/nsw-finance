package main

import (
	appheader "nsw-finance/components/app-header"
	passabletab "nsw-finance/components/passable-tab"
	savingstab "nsw-finance/components/savings-tab"

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
		DB:       app.SavingsDB,
		InfoLog:  app.Utils.InfoLog,
		ErrorLog: app.Utils.ErrorLog,
	}
	savingsContainer := savingsTab.GetSavingsTab()
	app.UIComponents.SavingsContainer = savingsContainer

	// Passable Tab
	passableTab := &passabletab.PassableTab{
		InfpLog:  app.Utils.InfoLog,
		ErrorLog: app.Utils.ErrorLog,
	}
	passableContainer := passableTab.GetPassableTab()
	app.UIComponents.PassableContainer = passableContainer

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Savings", theme.FileIcon(), savingsContainer),
		container.NewTabItemWithIcon("Passable", theme.ConfirmIcon(), passableContainer),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	// add container to window
	finalContent := container.NewBorder(appHeaderContainer, nil, nil, nil, tabs)

	app.MainWindow.SetContent(finalContent)
}

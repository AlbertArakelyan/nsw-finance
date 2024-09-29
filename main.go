package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"

	_ "github.com/glebarez/go-sqlite"
)

type App struct {
	App        fyne.App
	MainWindow fyne.Window
}

func main() {
	var myApp App

	// create a fyne app
	fyneApp := app.NewWithID("am.gocode.nswfinance.preferences")
	myApp.App = fyneApp
	myApp.App.Settings().SetTheme(theme.LightTheme())

	// open a connection to the database
	_, err := myApp.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("NSW Finance")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	// make the UI
	// myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}

func (app *App) connectSQL() (*sql.DB, error) {
	path := "./sql.db"

	// if os.Getenv("DB_PATH") != "" {
	// 	path = os.Getenv("DB_PATH")
	// } else {
	// 	path = app.App.Storage().RootURI().Path() + "/sql.db"
	// 	app.InfoLog.Println("db in:", path)
	// }

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

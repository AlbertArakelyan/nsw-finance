package main

import (
	"database/sql"
	"log"
	"net/http"
	"nsw-finance/repository"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	_ "github.com/glebarez/go-sqlite"
)

type UIComponents struct {
	SavingsContainer *fyne.Container
}

type Utils struct {
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	HTTPClient *http.Client
}

type App struct {
	App          fyne.App
	MainWindow   fyne.Window
	UIComponents UIComponents
	DB           repository.Repository
	Utils        Utils
}

func main() {
	var myApp App

	// create a fyne app
	fyneApp := app.NewWithID("am.gocode.nswfinance.preferences")
	myApp.App = fyneApp
	myApp.App.Settings().SetTheme(theme.LightTheme())

	// create our loggers
	myApp.Utils.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.Utils.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to the database
	sqlDB, err := myApp.connectSQL()
	if err != nil {
		log.Panic(err)
	}

	// setup the database
	myApp.setupDB(sqlDB)

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("NSW Finance")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	// make the UI
	myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}

func (app *App) connectSQL() (*sql.DB, error) {
	path := "./sql.db"

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.Utils.InfoLog.Println("db in:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *App) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.MigrateSavings()
	if err != nil {
		app.Utils.ErrorLog.Println(err)
		log.Panic(err)
	}

	err = app.DB.MigrateSpendingTables()
	if err != nil {
		app.Utils.ErrorLog.Println(err)
		log.Panic(err)
	}

	err = app.DB.MigrateSpendings()
	if err != nil {
		app.Utils.ErrorLog.Println(err)
		log.Panic(err)
	}
}

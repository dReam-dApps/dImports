package main

import (
	"fmt"
	"image/color"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/dReam-dApps/dImports/dimport"
	dreams "github.com/dReam-dApps/dReams"
	"github.com/dReam-dApps/dReams/bundle"
	"github.com/dReam-dApps/dReams/menu"
	"github.com/sirupsen/logrus"
)

const app_tag = "dImporter"

func main() {
	// Set max cpu
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	// Initialize logger to Stdout
	menu.InitLogrusLog(logrus.InfoLevel)

	// Initialize Fyne app and window into AppObject
	a := app.New()
	a.Settings().SetTheme(bundle.DeroTheme(color.Black))

	w := a.NewWindow(app_tag)
	w.SetIcon(bundle.ResourceBlueBadgePng)
	w.Resize(fyne.NewSize(700, 200))
	w.SetMaster()

	// Initialize background image and AppObject
	dreams.Theme.Img = *canvas.NewImageFromResource(bundle.ResourceBackgroundPng)
	d := dreams.AppObject{
		App:    a,
		Window: w}

	// Handle ctrl-c close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println()
		w.Close()
	}()

	// Set app content with background and import widget then run
	w.SetContent(container.NewMax(menu.BackgroundRast(app_tag), bundle.NewAlpha150(), container.NewCenter(dimport.ImportWidget(&d))))
	w.ShowAndRun()
}

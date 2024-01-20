package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/chzyer/readline"
	"github.com/civilware/Gnomon/structures"
	"github.com/dReam-dApps/dImports/dimport"
	"github.com/dReam-dApps/dReams/gnomes"
	"github.com/sirupsen/logrus"
)

const app_tag = "dImporter"

var help_string = `Help menu
	help          - Shows this menu
	exit || quit  - Close the app
	
	Input a valid Go package path, dImporter will try to import and run the package's StartApp()
	
	Example paths:

	"github.com/SixofClubsss/Baccarat/baccarat",
	"github.com/SixofClubsss/dPrediction/prediction",
	"github.com/SixofClubsss/Duels/duel",
	"github.com/SixofClubsss/Holdero/holdero",
	"github.com/SixofClubsss/Iluma/tarot"`

var logger = structures.Logger.WithFields(logrus.Fields{})

func main() {
	// Set max cpu
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	// Initialize logrus logger to stdout
	gnomes.InitLogrusLog(logrus.InfoLevel)

	// Function to set read line prompt text
	setPrompt := func() string {
		stamp := fmt.Sprintf("\033[90m[%s]\033[0m", time.Now().Format("01/02/2006 15:04:05"))
		name := fmt.Sprintf("\033[1;96m %s\033[0m", app_tag)
		return fmt.Sprintf("%s %s \033[35mEnter Go import path:\033[0m > ", stamp, name)
	}

	// Initialize read line for input
	rli, err := readline.New(setPrompt())
	if err != nil {
		logger.Fatalf("[%s] Error creating read line instance: %s\n", app_tag, err)
		return
	}
	defer rli.Close()

	// Initialize channel for closing
	done := make(chan struct{})

	logger.Printf("[%s] %s Starting app, enter 'help' for list of commands\n", app_tag, dimport.Version().String())

	// Routine to update read line prompt
	go func() {
		for {
			select {
			case <-done:
				logger.Printf("[%s] Closing...\n", app_tag)
				return
			default:
				rli.SetPrompt(setPrompt())
				rli.Refresh()
				time.Sleep(time.Second)
			}
		}
	}()

	// Read line process, will try to ImportAndStartApp given input
	for {
		line, err := rli.Readline()
		if err != nil {
			if err.Error() != "Interrupt" {
				logger.Errorf("[%s] Error reading line: %s\n", app_tag, err)
			}

			break
		}

		// Check input for exit and help commands
		if line == "exit" || line == "quit" {
			break
		}

		if line == "help" {
			logger.Printf("[%s] %s\n", app_tag, help_string)
			continue
		}

		dimport.ImportAndStartApp(line)
	}

	done <- struct{}{}
	time.Sleep(time.Second)
	logger.Printf("[%s] Closed\n", app_tag)
}

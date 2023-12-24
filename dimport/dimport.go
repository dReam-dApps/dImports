package dimport

import (
	"fmt"
	"image/color"
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/blang/semver/v4"
	"github.com/civilware/Gnomon/structures"
	dreams "github.com/dReam-dApps/dReams"
	"github.com/sirupsen/logrus"
	"github.com/x-motemen/gore"
)

var logger = structures.Logger.WithFields(logrus.Fields{})

var version = semver.MustParse("0.1.1")

// Get current package version
func Version() semver.Version {
	return version
}

// Import a Go package with Gore and call its package.StartApp(),
// path should be structured as github.com/user/repo/package
func ImportAndStartApp(path string) (err error) {
	var s *gore.Session
	var stderr strings.Builder
	s, err = gore.NewSession(os.Stdout, &stderr)
	if err != nil {
		logger.Errorln("[ImportAndStartApp]", err)
		return
	}

	// Evaluate full Go import path
	if err = s.Eval(fmt.Sprintf(":import %s", path)); err != nil {
		logger.Errorln("[ImportAndStartApp]", err)
		s.Clear()
		return
	}

	// Split path and check it is a valid import path with package
	split := strings.Split(path, "/")
	l := len(split)

	if l < 4 || split[l-1] == "" {
		err = fmt.Errorf("invalid package path %s", path)
		logger.Errorf("[ImportAndStartApp] Invalid package path %s", path)
		s.Clear()
		return
	}

	// Get commit hash of main branch from github
	// hash := GetCommitHash(split[l-3], split[l-2])
	logger.Println("[ImportAndStartApp] Importing:", path)
	// logger.Println("[ImportAndStartApp] Commit:", hash)

	// Call the package's StartApp()
	start_cmd := fmt.Sprintf("%s.StartApp()", split[l-1])
	if err = s.Eval(start_cmd); err != nil {
		err = fmt.Errorf("command %s failed", start_cmd)
		logger.Errorf("[ImportAndStartApp] Import start command %s failed\n", start_cmd)
	}

	s.Clear()

	return
}

// Widgets for calling ImportAndStartApp()
func ImportWidget(d *dreams.AppObject) fyne.CanvasObject {
	main_text := "Import and run a go packages StartApp()"
	label := widget.NewLabel(main_text)
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle.Bold = true

	default_dapps := []string{
		"github.com/SixofClubsss/Baccarat/baccarat",
		// "github.com/SixofClubsss/Grokked/grok",
		"github.com/SixofClubsss/dPrediction/prediction",
		"github.com/SixofClubsss/Duels/duel",
		"github.com/SixofClubsss/Holdero/holdero",
		"github.com/SixofClubsss/Iluma/tarot"}

	path_entry := widget.NewSelectEntry(default_dapps)
	path_entry.SetPlaceHolder("Go import path:")

	loading := widget.NewProgressBarInfinite()
	loading.Hide()

	import_button := widget.NewButton("Import", nil)
	import_button.Importance = widget.HighImportance
	import_button.OnTapped = func() {
		go func() {
			label.SetText(fmt.Sprintf("Running %s.StartApp()", path.Base(path_entry.Text)))
			import_button.Hide()
			loading.Start()
			loading.Show()
			if err := ImportAndStartApp(path_entry.Text); err != nil {
				dialog.NewError(err, d.Window).Show()
			}
			loading.Stop()
			loading.Hide()
			import_button.Show()
			label.SetText(main_text)
		}()
	}

	spacer := canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	spacer.SetMinSize(fyne.NewSize(400, 0))

	return container.NewVBox(
		spacer,
		label,
		path_entry,
		loading,
		import_button,
		layout.NewSpacer())
}

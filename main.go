// Trimlet DB: look at the tables of a SQL Server database, read only. The connection strings
// are kept in the system keychain.
package main

import (
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"github.com/trimlet/gui"
)

func main() {
	a, w := gui.NewWindow("local.trimletdb", "Trimlet DB")
	gui.UsePalette("olive") // DB's color
	// Trial: TRIMLET_COLOR gives the title bar another trimlet color, its icon color or the
	// step given after a dash, e.g. "red" or "red-700".
	name, step, _ := strings.Cut(os.Getenv("TRIMLET_COLOR"), "-")
	if p, ok := gui.Palettes[name]; ok {
		i := gui.IconStep(name)
		if n, err := strconv.Atoi(step); err == nil && n >= 100 && n <= 900 && n%100 == 0 {
			i = n / 100
		}
		gui.TitleBar = p[i]
	}
	path, err := storePath()
	var s *Store
	if err == nil {
		s, err = loadStore(path)
	}
	if err != nil {
		msg := canvas.NewText("✘ Could not read the saved connections: "+err.Error(), gui.Red)
		msg.TextSize = gui.TextSize
		msg.TextStyle.Monospace = true
		w.SetContent(container.NewPadded(container.NewVBox(msg)))
		gui.Run(a, w)
		return
	}

	m := newModel(s)
	m.copy = func(text string) { a.Clipboard().SetContent(text) }
	m.quit = a.Quit
	var screen *gui.Screen
	m.ui = func(f func()) { screen.Do(f) }
	screen = gui.NewScreen(a, w, m)
	gui.Run(a, w)
}

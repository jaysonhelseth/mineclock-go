package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"mineclock/clockstyle"
	"flag"
	"time"
	"github.com/gotk3/gotk3/glib"
)

var clock *gtk.Label

func main() {
	fontSize := flag.Int("f", 100, "Set the font size, defaults to 100")
	flag.Parse()

	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Could not create the window:", err)
	}

	win.SetTitle("mineclock")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	clock, err = gtk.LabelNew("00:00")
	if err != nil {
		log.Fatal("Could not create clock:", err)
	}

	//Assign classes to widgets
	Clockstyle.AssignClassToLabel(clock)
	Clockstyle.AssignClassToWindow(win)

	//Create the CSS provider and attach the styles string to it.
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Could not load styles.", err)
	}
	cssProvider.LoadFromData(Clockstyle.Css(*fontSize))

	//Attach all of the CSS to the main window screen
	winScreen, err := win.GetScreen()
	if err != nil {
		log.Fatal("Could not load screen:", err)
	}
	gtk.AddProviderForScreen(winScreen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_USER)

	//Add and show the components
	win.Add(clock)
	win.Fullscreen()
	win.ShowAll()

	//First run
	getTime()

	//Check every quarter of a second.
	_, err = glib.TimeoutAdd(15000, getTime)
	if err != nil {
		log.Fatal("Could not set the timer.")
	}

	gtk.Main()
}

func getTime() bool {
	formattedTime := time.Now().Format("03:04")
	clock.SetText(formattedTime)

	//true to keep the timer happy.
	return true
}

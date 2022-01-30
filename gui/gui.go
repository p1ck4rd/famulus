package gui

import (
	"famulus/internal/app/famulus"
	"log"
	"os"
	"path"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	appId = "com.github.p1ck4rd.famulus"
	title = "Famulus"
)

func onGenerateButtonClick(
	button *gtk.Button, window *gtk.ApplicationWindow, inputFilename string,
) {
	outputDialog, err := gtk.FileChooserNativeDialogNew(
		"Output", window, gtk.FILE_CHOOSER_ACTION_SAVE, "Save", "Cancel",
	)
	if err != nil {
		log.Fatal(err)
	}

	outputDialog.SetFilename(
		inputFilename[0:len(inputFilename)-len(path.Ext(inputFilename))] +
			".pdf",
	)

	outputDialog.Connect("response", func() {
		famulus.Generate(inputFilename, outputDialog.GetFilename())
		button.SetSensitive(true)
	})

	button.SetSensitive(false)

	outputDialog.Show()
}

func onApplicationActivate(application *gtk.Application) {
	window, err := gtk.ApplicationWindowNew(application)
	if err != nil {
		log.Fatal(err)
	}

	window.SetTitle(title)

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal(err)
	}

	grid.SetRowSpacing(15)
	grid.SetColumnSpacing(10)
	grid.SetMarginStart(20)
	grid.SetMarginEnd(20)
	grid.SetMarginTop(20)
	grid.SetMarginBottom(20)
	window.Add(grid)

	label, err := gtk.LabelNew("Input:")
	if err != nil {
		log.Fatal(err)
	}

	label.SetHAlign(gtk.ALIGN_START)
	label.SetVAlign(gtk.ALIGN_CENTER)
	label.SetHExpand(true)

	inputButton, err := gtk.FileChooserButtonNew(
		"Input", gtk.FILE_CHOOSER_ACTION_OPEN,
	)
	if err != nil {
		log.Fatal(err)
	}

	fileFilter, err := gtk.FileFilterNew()
	if err != nil {
		log.Fatal(err)
	}

	fileFilter.AddPattern("*.kys")
	inputButton.AddFilter(fileFilter)

	grid.Attach(label, 0, 0, 1, 1)
	grid.Attach(inputButton, 1, 0, 1, 1)

	generateButton, err := gtk.ButtonNewWithLabel("Generate")
	if err != nil {
		log.Fatal(err)
	}

	var inputFilename string
	generateButton.SetSensitive(false)
	inputButton.Connect("file-set", func() {
		inputFilename = inputButton.GetFilename()
		generateButton.SetSensitive(true)
	})
	generateButton.Connect("clicked", func() {
		onGenerateButtonClick(generateButton, window, inputFilename)
	})

	grid.Attach(generateButton, 0, 1, 2, 1)

	window.ShowAll()
}

func Run() {
	application, err := gtk.ApplicationNew(
		appId, glib.APPLICATION_FLAGS_NONE,
	)
	if err != nil {
		log.Fatal(err)
	}

	application.Connect("activate", func() {
		onApplicationActivate(application)
	})

	os.Exit(application.Run(os.Args))
}

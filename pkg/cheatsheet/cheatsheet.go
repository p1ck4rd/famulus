// Package cheatsheet implements a function to generate a pdf file.
package cheatsheet

import (
	"strings"

	"github.com/p1ck4rd/famulus/pkg/parser"
	"github.com/phpdave11/gofpdf"
)

const (
	fontSize    = 16
	lineHeight  = fontSize / 2
	imageHeight = lineHeight - 1
)

func insertShortcut(pdf *gofpdf.Fpdf, shortcut parser.Shortcut) error {
	for j, key := range shortcut {
		if j > 0 {
			pdf.Write(lineHeight, "+\t\t")
		}
		switch key {
		case ".":
			key = "FULLSTOP"
		case "/":
			key = "SLASH"
		}
		upperKey := strings.ToUpper(key)
		if svg, ok := buttons[upperKey]; ok {
			sig, err := gofpdf.SVGBasicParse(svg)
			if err != nil {
				return err
			}
			scale := imageHeight / sig.Ht
			x := pdf.GetX()
			y := pdf.GetY()
			pdf.SVGBasicWrite(&sig, scale)
			x += scale * sig.Wd
			pdf.SetXY(x, y)
		} else {
			pdf.Write(lineHeight, upperKey)
		}
	}

	return nil
}

func insertShortcuts(pdf *gofpdf.Fpdf, shortcuts []parser.Shortcut) (
	err error,
) {
	for i, shortcut := range shortcuts {
		if i > 0 {
			pdf.Write(lineHeight, ";\t\t")
		}
		if err = insertShortcut(pdf, shortcut); err != nil {
			return
		}
	}
	return
}

// Generate creates a pdf file from a parser.Result object.
func Generate(res *parser.Result, filename string) (err error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", fontSize)

	for _, command := range res.Commands {
		if len(command.Shortcuts) > 0 && len(command.Shortcuts[0]) > 0 {
			pdf.Writef(lineHeight, "%s:\t\t", command.Name)
			if err = insertShortcuts(pdf, command.Shortcuts); err != nil {
				return
			}
			pdf.Write(lineHeight, "\n")
		}
	}

	for _, tool := range res.Tools {
		if len(tool.Shortcut) > 0 {
			pdf.Writef(lineHeight, "%s:\t\t", tool.Name)
			if err = insertShortcut(pdf, tool.Shortcut); err != nil {
				return
			}
			pdf.Write(lineHeight, "\n")
		}
	}

	if err = pdf.OutputFileAndClose(filename); err != nil {
		return
	}

	return
}

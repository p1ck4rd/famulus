package parser

import (
	"encoding/xml"
	"strings"
)

// Shortcut is a keyboard shortcut representation.
type Shortcut []string

func (s *Shortcut) parse(command string) {
	if command == "" {
		*s = Shortcut{}
		return
	}
	splitedShortcut := strings.Split(command, "+")
	var result Shortcut
	for i, elem := range splitedShortcut {
		elem = strings.TrimSpace(elem)
		switch elem {
		case "":
			if i > 0 && splitedShortcut[i-1] == "" {
				result = append(result, "+")
			}
		default:
			result = append(result, elem)
		}
	}
	*s = result
}

// UnmarshalXML produces the value from the XML value of the shortcut.
func (s *Shortcut) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	s.parse(v)
	return nil
}

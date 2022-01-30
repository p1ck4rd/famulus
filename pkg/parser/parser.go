// Package parser implements structs and function for unmarshalling Photoshop
// keyboard shortcuts.
package parser

import "encoding/xml"

// Command stores the name of Photoshop command and its shortcuts.
type Command struct {
	Name      string     `xml:"name,attr"`
	Shortcuts []Shortcut `xml:"shortcut"`
}

// Tool stores the name of Photoshop tool and its shortcut.
type Tool struct {
	Name     string
	Shortcut Shortcut
}

// Result stores the result of unmarshalling.
type Result struct {
	Commands []Command `xml:"command"`
	Tools    []Tool    `xml:"tool"`
}

// UnmarshalXML produces the value from the <tool> XML element.
func (t *Tool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			t.Name = attr.Value
			break
		}
	}
	t.Shortcut.parse(v)
	return nil
}

// Unmarshal returns a result of unmarshalling.
func Unmarshal(content []byte) (*Result, error) {
	res := new(Result)

	return res, xml.Unmarshal(content, &res)
}

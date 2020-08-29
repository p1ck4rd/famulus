package parser

import (
	"reflect"
	"testing"
)

func TestShortcutParse(t *testing.T) {
	cases := map[string]Shortcut{
		"":            {},
		"foo":         {"foo"},
		"foo+":        {"foo"},
		" foo ":       {"foo"},
		"foo+ ":       {"foo"},
		"foo++":       {"foo", "+"},
		"foo+bar":     {"foo", "bar"},
		"foo++bar":    {"foo", "bar"},
		"foo+++bar":   {"foo", "+", "bar"},
		"foo+bar+baz": {"foo", "bar", "baz"},
	}

	for command, expected := range cases {
		s := Shortcut{}
		s.parse(command)
		if !reflect.DeepEqual(expected, s) {
			t.Errorf("Shortcut for %v is %v; want %v", command, s, expected)
		}
	}
}

package parser

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	cases := []struct {
		given    []byte
		expected *Result
	}{
		{
			given: []byte(`
				<test>
					<command name="foo">
						<shortcut>bar</shortcut>
					</command>
				</test>`,
			),
			expected: &Result{
				Commands: []Command{
					{Name: "foo", Shortcuts: []Shortcut{{"bar"}}},
				},
			},
		},
		{
			given: []byte(`
				<test>
					<command name="foobar">
						<shortcut>foo+bar</shortcut>
					</command>
				</test>`,
			),
			expected: &Result{
				Commands: []Command{
					{Name: "foobar", Shortcuts: []Shortcut{{"foo", "bar"}}},
				},
			},
		},
		{
			given: []byte(`
				<test>
					<command name="foobar">
						<shortcut>foo</shortcut>
						<shortcut>bar</shortcut>
					</command>
				</test>`,
			),
			expected: &Result{
				Commands: []Command{
					{Name: "foobar", Shortcuts: []Shortcut{{"foo"}, {"bar"}}},
				},
			},
		},
		{
			given: []byte(`
				<test>
					<command name="foo">
						<shortcut>bar</shortcut>
					</command>
					<command name="qwe">
						<shortcut>asd</shortcut>
					</command>
				</test>`,
			),
			expected: &Result{
				Commands: []Command{
					{Name: "foo", Shortcuts: []Shortcut{{"bar"}}},
					{Name: "qwe", Shortcuts: []Shortcut{{"asd"}}},
				},
			},
		},
		{
			given: []byte(`<test><tool name="foo">bar</tool></test>`),
			expected: &Result{
				Tools: []Tool{{Name: "foo", Shortcut: Shortcut{"bar"}}},
			},
		},
		{
			given: []byte(`<test><tool name="foobar">foo+bar</tool></test>`),
			expected: &Result{
				Tools: []Tool{
					{Name: "foobar", Shortcut: Shortcut{"foo", "bar"}},
				},
			},
		},
		{
			given: []byte(
				`<test>
					<tool name="foo">bar</tool>
					<tool name="qwe">asd</tool>
				</test>`,
			),
			expected: &Result{
				Tools: []Tool{
					{Name: "foo", Shortcut: Shortcut{"bar"}},
					{Name: "qwe", Shortcut: Shortcut{"asd"}},
				},
			},
		},
		{
			given: []byte(`
				<test>
					<command name="foo">
						<shortcut>bar</shortcut>
					</command>
					<tool name="qwe">asd</tool>
				</test>`,
			),
			expected: &Result{
				Commands: []Command{
					{Name: "foo", Shortcuts: []Shortcut{{"bar"}}},
				},
				Tools: []Tool{
					{Name: "qwe", Shortcut: Shortcut{"asd"}},
				},
			},
		},
	}
	for _, c := range cases {
		res, err := Unmarshal(c.given)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(c.expected, res) {
			t.Errorf("Unmarshal(%s) = %v; want %v", c.given, res, c.expected)
		}
	}
}

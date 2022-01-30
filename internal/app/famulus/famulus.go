package famulus

import (
	"errors"
	"famulus/pkg/cheatsheet"
	"famulus/pkg/parser"
	"io/ioutil"
)

func Generate(input, output string) error {
	if input == "" {
		return errors.New("INPUT PATH MUST BE SPECIFIED")
	}
	if output == "" {
		return errors.New("OUTPUT PATH MUST BE SPECIFIED")
	}

	content, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}

	res, err := parser.Unmarshal(content)
	if err != nil {
		return err
	}

	if err = cheatsheet.Generate(res, output); err != nil {
		return err
	}

	return nil
}

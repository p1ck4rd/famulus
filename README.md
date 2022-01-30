# famulus
**Famulus** is a keyboard shortcut cheat sheet generator for *Adobe Photoshop*.
All it takes is a `.kys` file from your preset.

## Usage
### GUI
```sh
famulus
```

### CLI
```sh
famulus -i <input_path> -o <output_path>
```

For example:
```sh
famulus -i foobar.kys -o foobar.pdf
```

## Build
To build `famulus` binary run:
```sh
go build cmd/famulus/famulus.go
```

If you want to use your own cheat sheet key icons:
- put your icon images in `.svg` format with names according to the keyboard key names in the `assets/svg` folder
- run `./scripts/svg_parser.py`
- compile the binary

## Changelog
All feature changes are reflected in the [CHANGELOG](CHANGELOG.md).

## License
**Famulus** is published under the **GNU GPL-3.0 license**. See [LICENSE](LICENSE).

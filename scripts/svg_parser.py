#!/usr/bin/python3
from os import scandir
from os.path import split, realpath

def main():
    realpath_dir = split(realpath(__file__))[0]
    svg_files = scandir(f'{realpath_dir}/../assets/svg')
    go_file_content = 'package cheatsheet\n\nvar buttons = map[string][]byte{\n'
    for file in svg_files:
        with open(file.path) as f:
            key = file.name.split('.')[0]
            go_file_content += f'\t"{key}": []byte(`\n{f.read()}`),\n'
    go_file_content += '}'
    with open(f'{realpath_dir}/../pkg/cheatsheet/svg.go', 'w') as f:
        f.write(go_file_content)

if __name__ == '__main__':
    main()

# Torisetsu（取説）
torisetsu - Add README.md Template.

[Wiki](https://github.com/syossan27/torisetsu/wiki)

## Requirement
Golang
urfave/cli
mitchellh/go-homedir
Masterminds/glide

## Install
````
$ go get github.com/syossan27/torisetsu
````

## Usage
After "create a new repository on the command line"

````
$ torisetsu
````

Add README.md Template.

````README.md
Overview

## Description

## Demo

## VS.

## Requirement

## Usage

## Install

## Contribution

## License

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[](https://github.com/)
````

Command help.

````
$ torisetsu --help   
NAME:
   torisetsu - Write README.md Template

USAGE:
   torisetsu [global options] command [command options] [arguments...]

VERSION:
   1.1

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --author, -a 	This flag specifies the author name to print.
   --license, -l "mit"	This flag specifies the choose license to print.
			Choose License:
			  none		: None
			  apache	: Apache License 2.0
			  mit		: MIT License
			  al		: Artistic License 2.0
			  bsd2		: BSD 2-clause 'Simplified' License
			  bsd3		: BSD 3-clause 'New' or 'Revised' License
			  cc0		: Creative Commons Zero v1.0 Universal
			  epl		: Eclipse Public License 1.0
			  agpl		: GNU Affero General Public License v3.0
			  gpl2		: GNU General Public License v2.0
			  gpl3		: GNU General Public License v3.0
			  lgpl2		: GNU Lesser General Public License v2.1
			  lgpl3		: GNU Lesser General Public License v3.0
			  iscl		: ISC License
			  mpl		: Mozilla Public License 2.0
			  unlicense	: The Unlicense
   --template, -t 	Input the name of the file except for the extension.
   --help, -h		show help
   --version, -v	print the version
````

## Option

[Author Option - wiki](https://github.com/syossan27/torisetsu/wiki/author)

[License Option - wiki](https://github.com/syossan27/torisetsu/wiki/license)

[Template Option - wiki](https://github.com/syossan27/torisetsu/wiki/template)

## Contributing

You're most welcomed!💓
Welcome pull request and issues.✨

## License

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[syossan27](https://github.com/syossan27)

# Torisetsu（取説）
torisetsu - Add README.md Template.

## Requirement
Golang

## Usage
After "create a new repository on the command line"

````
$ torisetsu
````

Add README.md Template.

````README.md
# Test Repository

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
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --author, -a   This flag specifies the author name to print.
   --license, -l  This flag specifies the choose license to print.
      Choose License:
        none      : None
        apache    : Apache License 2.0
        mit       : MIT License
        al        : Artistic License 2.0
        bsd2      : BSD 2-clause 'Simplified' License
        bsd3      : BSD 3-clause 'New' or 'Revised' License
        cc0       : Creative Commons Zero v1.0 Universal
        epl       : Eclipse Public License 1.0
        agpl      : GNU Affero General Public License v3.0
        gpl2      : GNU General Public License v2.0
        gpl3      : GNU General Public License v3.0
        lgpl2     : GNU Lesser General Public License v2.1
        lgpl3     : GNU Lesser General Public License v3.0
        iscl      : ISC License
        mpl       : Mozilla Public License 2.0
        unlicense : The Unlicense
   --help, -h   show help
   --version, -v  print the version
````

## Install
````
$ go get github.com/syossan27/torisetsu
````

## License

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[syossan27](https://github.com/syossan27)

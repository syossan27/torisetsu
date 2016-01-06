package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

func main() {
  choose_license := []string{
    "Choose License:",
    "  none\t: None",
    "  apache\t: Apache License 2.0",
    "  mit\t: MIT License",
    "  al\t: Artistic License 2.0",
    "  bsd2\t: BSD 2-clause 'Simplified' License",
    "  bsd3\t: BSD 3-clause 'New' or 'Revised' License",
    "  cc0\t: Creative Commons Zero v1.0 Universal",
    "  epl\t: Eclipse Public License 1.0",
    "  agpl\t: GNU Affero General Public License v3.0",
    "  gpl2\t: GNU General Public License v2.0",
    "  gpl3\t: GNU General Public License v3.0",
    "  lgpl2\t: GNU Lesser General Public License v2.1",
    "  lgpl3\t: GNU Lesser General Public License v3.0",
    "  iscl\t: ISC License",
    "  mpl\t: Mozilla Public License 2.0",
    "  unlicense\t: The Unlicense",
  }

	app := cli.NewApp()
	app.Name = "torisetsu"
	app.Usage = "Write README.md Template"
	app.Version = "1.1"
	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name:  "author, a",
			Value: "",
			Usage: "This flag specifies the author name to print.",
		},
		cli.StringFlag {
			Name:  "license, l",
			Value: "",
      Usage: "This flag specifies the choose license to print.\n\t" + strings.Join(choose_license, "\n\t"),
		},
	}
	app.Action = func(c *cli.Context) {
		file, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			println("Error: README.md is not exist.")
			os.Exit(0)
		}
		defer file.Close()

    license := c.String("license")
    license_string := ""
    switch license {
      case "none": license_string = "None"
      case "apache": license_string = "[Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)"
      case "mit": license_string = "[MIT](http://opensource.org/licenses/mit-license.php)"
      case "al": license_string = "[Artistic License 2.0](http://www.perlfoundation.org/artistic_license_2_0)"
      case "bsd2": license_string = "[BSD 2-clause 'Simplified' License](https://opensource.org/licenses/BSD-2-Clause)"
      case "bsd3": license_string = "[BSD 3-clause 'New' or 'Revised' License](https://opensource.org/licenses/BSD-3-Clause)"
      case "cc0": license_string = "[Creative Commons Zero v1.0 Universal](http://creativecommons.org/publicdomain/zero/1.0/legalcode)"
      case "epl": license_string = "[Eclipse Public License 1.0](http://www.eclipse.org/legal/epl-v10.html)"
      case "agpl": license_string = "[GNU Affero General Public License v3.0](http://www.gnu.org/licenses/agpl-3.0.html)"
      case "gpl2": license_string = "[GNU General Public License v2.0](http://www.gnu.org/licenses/old-licenses/gpl-2.0.txt)"
      case "gpl3": license_string = "[GNU General Public License v3.0](http://www.gnu.org/licenses/gpl-3.0.txt)"
      case "lgpl2": license_string = "[GNU Lesser General Public License v2.1](http://www.gnu.org/licenses/lgpl-2.1.html)"
      case "lgpl3": license_string = "[GNU Lesser General Public License v3.0](http://www.gnu.org/licenses/lgpl-3.0.html)"
      case "iscl": license_string = "[ISC License](http://opensource.org/licenses/isc-license.txt)"
      case "mpl": license_string = "[Mozilla Public License 2.0](https://www.mozilla.org/en-US/MPL/2.0/)"
      case "unlicense" : license_string = "[The Unlicense](http://unlicense.org/)"
      default: license_string = "[MIT](http://opensource.org/licenses/mit-license.php)"
    }

		author := c.String("author")

		body := []string{
			"Overview",
			"## Description",
			"## Demo",
			"## VS.",
			"## Requirement",
			"## Usage",
			"## Install",
			"## Contribution",
			"## License",
      fmt.Sprintf(license_string),
			"## Author",
			fmt.Sprintf("[%s](https://github.com/%s)", author, author),
		}
		readme_template := []byte("\n" + strings.Join(body, "\n\n"))

		writer := bufio.NewWriter(file)
		writer.Write(readme_template)
		writer.Flush()

		println("Complete add README.md!")
	}
	app.Run(os.Args)
}

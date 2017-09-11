package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"github.com/mitchellh/go-homedir"
	"os"
	"strings"
)

const (
	ExitCodeOK = iota
	ExitCodeError
	defaultLicense = "mit"
)

func makeApp() *cli.App {
	app := cli.NewApp()
	app.Name = "torisetsu"
	app.Usage = "Write README.md Template"
	app.Version = "1.1"

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

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "author, a",
			Value: "",
			Usage: "This flag specifies the author name to print.",
		},
		cli.StringFlag{
			Name:  "license, l",
			Value: defaultLicense,
			Usage: "This flag specifies the choose license to print.\n\t" + strings.Join(choose_license, "\n\t"),
		},
		cli.StringFlag{
			Name:  "template, t",
			Value: "",
			Usage: "Input the name of the file except for the extension.",
		},
	}

	app.Action = func(c *cli.Context) {
		// Flag license
		license, err := flagLicense(c.String("license"))
		if err != nil {
			fmt.Println(err)
			os.Exit(ExitCodeError)
		}

		// Flag author
		author := flagAuthor(c.String("author"))

		// Flag template
		readme_template := flagTemplate(c.String("template"), license, author)

		file, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			println("Error: README.md is not exist.")
			os.Exit(ExitCodeError)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		writer.Write(readme_template)
		writer.Flush()

		println("Complete add README.md!")
		os.Exit(ExitCodeOK)
	}

	return app
}

func flagLicense(license string) (string, error) {
	license_list := map[string]string{
		"none":      "None",
		"apache":    "[Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)",
		"mit":       "[MIT](http://opensource.org/licenses/mit-license.php)",
		"al":        "[Artistic License 2.0](http://www.perlfoundation.org/artistic_license_2_0)",
		"bsd2":      "[BSD 2-clause 'Simplified' License](https://opensource.org/licenses/BSD-2-Clause)",
		"bsd3":      "[BSD 3-clause 'New' or 'Revised' License](https://opensource.org/licenses/BSD-3-Clause)",
		"cc0":       "[Creative Commons Zero v1.0 Universal](http://creativecommons.org/publicdomain/zero/1.0/legalcode)",
		"epl":       "[Eclipse Public License 1.0](http://www.eclipse.org/legal/epl-v10.html)",
		"agpl":      "[GNU Affero General Public License v3.0](http://www.gnu.org/licenses/agpl-3.0.html)",
		"gpl2":      "[GNU General Public License v2.0](http://www.gnu.org/licenses/old-licenses/gpl-2.0.txt)",
		"gpl3":      "[GNU General Public License v3.0](http://www.gnu.org/licenses/gpl-3.0.txt)",
		"lgpl2":     "[GNU Lesser General Public License v2.1](http://www.gnu.org/licenses/lgpl-2.1.html)",
		"lgpl3":     "[GNU Lesser General Public License v3.0](http://www.gnu.org/licenses/lgpl-3.0.html)",
		"iscl":      "[ISC License](http://opensource.org/licenses/isc-license.txt)",
		"mpl":       "[Mozilla Public License 2.0](https://www.mozilla.org/en-US/MPL/2.0/)",
		"unlicense": "[The Unlicense](http://unlicense.org/)",
	}

	if license_string, ok := license_list[license]; ok {
		return license_string, nil
	} else {
		return "", errors.New("The selected license can not be used.")
	}
}

func flagAuthor(author string) string {
	return fmt.Sprintf("[%s](https://github.com/%s)", author, author)
}

// it's so foolish code.
// Need Refactoring.
func flagTemplate(template, license, author string) []byte {
	if template == "" {
		user_home, err := homedir.Dir()
		if err != nil {
			return []byte("")
		}
		if FileExists(user_home + "/.torisetsu/default.md") {
			template_content, err := readTemplate("default", license, author)
			if err != nil {
				return []byte("")
			}
			return template_content
		} else {
			template_content := createReadTemplate(license, author)
			return template_content
		}
	} else {
		template_content, err := readTemplate(template, license, author)
		if err != nil {
			return []byte("")
		}
		return template_content
	}
}

func readTemplate(template_name, license, author string) ([]byte, error) {
	user_home, err := homedir.Dir()
	f, err := os.Open(user_home + "/.torisetsu/" + template_name + ".md")
	if err != nil {
		return []byte(""), err
	}
	template_content := make([]string, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		insert_text := s.Text()
		if insert_text == "torisetsu.license" {
			insert_text = license
		} else if insert_text == "torisetsu.author" {
			insert_text = author
		}

		template_content = append(template_content, insert_text)
	}
	if s.Err() != nil {
		return []byte(""), err
	}

	return []byte(strings.Join(template_content, "\n")), nil
}

func createReadTemplate(license, author string) []byte {
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
		fmt.Sprintf(license),
		"## Author",
		fmt.Sprintf(author),
	}

	return []byte("\n" + strings.Join(body, "\n\n"))
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	app := makeApp()
	app.Run(os.Args)
}

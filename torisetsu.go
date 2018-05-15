package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"github.com/mitchellh/go-homedir"
	"os"
	"strings"
	"torisetsu/foundation"
)

const (
	ExitCodeOK = iota
	ExitCodeError
	defaultLicense = "mit"
)

var (
	homeDir string
	configDir string
	flags = []cli.Flag{
		cli.StringFlag{
			Name:  "license, l",
			Value: defaultLicense,
			Usage: "This flag specifies the choose license to print.\n\t" + strings.Join(chooseLicenceStrings, "\n\t"),
		},
		cli.StringFlag{
			Name:  "author, a",
			Value: "",
			Usage: "This flag specifies the author name to print.",
		},
		cli.StringFlag{
			Name:  "template, t",
			Value: "",
			Usage: "Input the name of the file except for the extension.",
		},
	}

	chooseLicenceStrings = []string{
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

	licenseList = map[string]string{
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
)

func makeApp() *cli.App {
	app := cli.NewApp()
	app.Name = "torisetsu"
	app.Usage = "Write README.md Template"
	app.Version = "1.1"
	app.Flags = flags
	app.Action = action
	return app
}

func action(c *cli.Context) {
	file, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		foundation.PrintError("README.md is not exist")
		os.Exit(ExitCodeError)
	}
	defer file.Close()

	homeDir, err = homedir.Dir()
	if err != nil {
		foundation.PrintError("Failed fetch homeDir: " + err.Error())
		os.Exit(ExitCodeError)
	}
	configDir = homeDir + "/.torisetsu"

	// Flag license
	license, err := flagLicense(c.String("license"))
	if err != nil {
		foundation.PrintError("Failed fetch license: " + err.Error())
		os.Exit(ExitCodeError)
	}

	// Flag author
	author := flagAuthor(c.String("author"))

	// Flag template
	template, err := flagTemplate(c.String("template"), license, author)
	if err != nil {
		foundation.PrintError("Failed fetch template: " + err.Error())
		os.Exit(ExitCodeError)
	}

	writer := bufio.NewWriter(file)
	writer.Write(template)
	writer.Flush()

	foundation.PrintSuccess("Complete add README.md!")
	os.Exit(ExitCodeOK)
}

func flagLicense(licenseName string) (string, error) {
	if license, ok := licenseList[licenseName]; ok {
		return license, nil
	} else {
		return "", errors.New("selected license can not be used")
	}
}

func flagAuthor(author string) string {
	return fmt.Sprintf("[%s](https://github.com/%s)", author, author)
}

// it's so foolish code.
// Need Refactoring.
func flagTemplate(template, license, author string) ([]byte, error) {
	if template == "" {
		if FileExists(configDir + "/default.md") {
			templateCotent, err := readTemplate("default", license, author)
			if err != nil {
				return nil, err
			}
			return templateCotent, nil
		} else {
			templateContent := createReadTemplate(license, author)
			return templateContent, nil
		}
	} else {
		templateContent, err := readTemplate(template, license, author)
		if err != nil {
			return nil, err
		}
		return templateContent, nil
	}
}

func readTemplate(templateName, license, author string) ([]byte, error) {
	f, err := os.Open(configDir + "/" + templateName + ".md")
	if err != nil {
		return []byte(""), err
	}
	templateContent := make([]string, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		insertText := s.Text()
		if insertText == "torisetsu.license" {
			insertText = license
		} else if insertText == "torisetsu.author" {
			insertText = author
		}

		templateContent = append(templateContent, insertText)
	}
	if s.Err() != nil {
		return []byte(""), err
	}

	return []byte(strings.Join(templateContent, "\n")), nil
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

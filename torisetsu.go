package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "torisetsu"
	app.Usage = "Write README.md Template"
	app.Version = "1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "author, a",
			Value: "",
			Usage: "This flag specifies the author name to print.",
		},
	}
	app.Action = func(c *cli.Context) {
		file, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			println("Error: README.md is not exist.")
			os.Exit(0)
		}
		defer file.Close()

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
			"[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)",
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

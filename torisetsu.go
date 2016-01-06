package main

import (
  "os"
  "bufio"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "torisetsu"
  app.Usage = "Write README.md Template"
  app.Version = "1.0"
  app.Action = func(c *cli.Context) {
    file, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND, 0600)
    if err != nil {
      println("Error: README.md is not exist.")
      os.Exit(0)
    }
    defer file.Close()

    readme_template := []byte(
      "\nOverview\n\n" +
      "## Description\n\n" +
      "## Demo\n\n" +
      "## VS.\n\n" +
      "## Requirement\n\n" +
      "## Usage\n\n" +
      "## Install\n\n" +
      "## Contribution\n\n" +
      "## License\n\n" +
      "[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)\n\n" +
      "## Author\n\n" +
      "[](https://github.com/)",
    )

    writer := bufio.NewWriter(file)
    writer.Write(readme_template)
    writer.Flush()

    println("Complete add README.md!")
  }
  app.Run(os.Args)
}

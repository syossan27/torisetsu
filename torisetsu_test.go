package main

import (
	"reflect"
	"strings"
	"testing"
)

const (
	default_license            = "mit"
	default_author             = ""
	default_comparison_license = "[MIT](http://opensource.org/licenses/mit-license.php)"
	default_comparison_author  = "[](https://github.com/)"
)

func createCoparisonTemplate(comparison_license, comparison_author string) []byte {
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
		comparison_license,
		"## Author",
		comparison_author,
	}

	return []byte("\n" + strings.Join(body, "\n\n"))
}

func compareTestData(license, author, comparison_license, comparison_author string) bool {
	readme_template := createReadTemplate(license, author)
	comparison_template := createCoparisonTemplate(comparison_license, comparison_author)
	return reflect.DeepEqual(readme_template, comparison_template)
}

func TestRun_noFlag(t *testing.T) {
	license := flagLicense(default_license)
	author := flagAuthor(default_author)
	comparison_license := default_comparison_license
	comparison_author := default_comparison_author

	result := compareTestData(license, author, comparison_license, comparison_author)

	if !result {
		t.Fatalf("Test Failed")
	}
}

func TestRun_licenseFlag(t *testing.T) {
	license := flagLicense("unlicense")
	author := flagAuthor(default_author)
	comparison_license := "[The Unlicense](http://unlicense.org/)"
	comparison_author := default_comparison_author

	result := compareTestData(license, author, comparison_license, comparison_author)

	if !result {
		t.Fatalf("Test Failed")
	}
}

func TestRun_authorFlag(t *testing.T) {
	license := flagLicense(default_license)
	author := flagAuthor("syossan27")
	comparison_license := default_comparison_license
	comparison_author := "[syossan27](https://github.com/syossan27)"

	result := compareTestData(license, author, comparison_license, comparison_author)

	if !result {
		t.Fatalf("Test Failed")
	}
}

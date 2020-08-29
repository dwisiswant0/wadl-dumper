package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"
)

type options struct {
	baseURL, input, replace string
	showBase                bool
}

var opt *options

func init() {
	opt = &options{}

	flag.StringVar(&opt.input, "i", "", "")
	flag.StringVar(&opt.input, "input", "", "")

	flag.BoolVar(&opt.showBase, "b", false, "")
	flag.BoolVar(&opt.showBase, "show-base", false, "")

	flag.StringVar(&opt.replace, "r", "", "")
	flag.StringVar(&opt.replace, "replace", "", "")

	flag.Usage = func() {
		h := []string{
			"Usage:",
			"  wadl-dumper -i http://domain.tld/application.wadl [options...]",
			"  wadl-dumper -i /path/to/wadl.xml --show-base -r \"-alert(1)-\"",
			"",
			"Options:",
			"  -i, --input <URL/FILE>         URL/path to WADL file",
			"  -b, --show-base                Add base URL to paths",
			"  -r, --replace <string>         Replace all placeholder with given value",
			"  -h, --help                     Show its help text",
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

	flag.Parse()
}

func error(message string) {
	err := fmt.Sprintf("Error! %s\n", message)
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

func replaceNth(s, old string, new string, n int) string {
	i := 0

	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}

	return s
}

func main() {
	var path string
	var wadl *xmlquery.Node

	if opt.input == "" {
		error("Flag -i is required, use -h flag for help.")
	}

	if strings.HasPrefix(opt.input, "http") {
		wadl, _ = xmlquery.LoadURL(opt.input)
	} else {
		f, err := os.Open(opt.input)
		if err != nil {
			error(fmt.Sprintf("Can't open '%s' file.", opt.input))
		}

		wadl, _ = xmlquery.Parse(f)
	}

	if wadl == nil {
		error("Can't parse WADL file.")
	}

	xmlns := xmlquery.FindOne(wadl, "//application/@xmlns")
	if !strings.Contains(xmlns.InnerText(), "wadl.dev.java.net") {
		error("Not a WADL file.")
	}

	base := xmlquery.FindOne(wadl, "//resources/@base")
	if base != nil && opt.showBase {
		opt.baseURL = base.InnerText()
	} else {
		opt.baseURL = ""
	}

	for _, paths := range xmlquery.Find(wadl, "//resource/@path") {
		path = opt.baseURL + paths.InnerText()

		if opt.replace != "" {
			re := regexp.MustCompile("{.*}")
			path = re.ReplaceAllString(path, opt.replace)
		}

		if opt.baseURL != "" {
			path = replaceNth(path, "//", "/", 2)
		}

		fmt.Printf("%s\n", path)
	}
}

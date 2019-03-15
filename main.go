package main

import (
	"bufio"
	"bytes"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"io"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type PrgData struct {
	Pkg          string
	MacAndOrganization map[string]string
}

const prg = `
/*
 * This program is auto generated.
 */
package {{.Pkg}}

var MacAndOrganization = map[string]string {
{{range $mac, $organization := .MacAndOrganization}}
	"{{$mac}}" : "{{$organization}}",
{{end}}
}
`

type Options struct {
	Package string `short:"p" long:"package" default:"main" description:"Set package name."`
	OutPath string `short:"o" long:"out" required:"true" description:"Set package name."`
}

var opts Options

func main() {

	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "See the output of oui -h for a summary of options.\n")
		os.Exit(1)
	}

	prgData := PrgData{}
	prgData.Pkg = opts.Package
	prgData.MacAndOrganization = map[string]string{}

	bind(prgData.MacAndOrganization)
	tmpl, err := template.New("prg").Parse(prg)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(opts.OutPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := tmpl.Execute(file, prgData); err != nil {
		panic(err)
	}
}

func bind(macAndOrganization map[string]string) {
	url := "http://standards-oui.ieee.org/oui/oui.txt"
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	regMac := regexp.MustCompile(`[0-9a-fA-F]{6}`)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		group := regMac.FindSubmatch(line)
		if len(group) < 1 {
			continue
		}
		mac := string(group[0])
		for index, _ := range line {
			indexMinusTwo := index - 2
			if indexMinusTwo < 0 {
				continue
			}
			if bytes.Compare(line[indexMinusTwo:index], []byte{0x09, 0x09}) == 0 {
				organization := string(line[index:])
				macAndOrganization[mac] = organization
				break
			}
		}
	}
}

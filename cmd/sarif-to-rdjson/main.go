package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/yteraoka/sarif-to-rdjson/internal/pkg/rdjson"
	"github.com/yteraoka/sarif-to-rdjson/internal/pkg/sarif"
)

var (
	version = "unknown"
	commit  = "unknown"
	date    = "unknown"
)

func convert(sarifJsonBytes []byte) (rdjsonBytes []byte, err error) {
	var src sarif.Sarif

	err = json.Unmarshal(sarifJsonBytes, &src)
	if err != nil {
		return
	}

	var rd rdjson.DiagnosticResult

	rd.Source = &rdjson.Source{
		Name: src.Runs[0].Tool.Driver.Name,
		Url:  src.Runs[0].Tool.Driver.InformationUri,
	}

	for _, res := range src.Runs[0].Results {
		for _, loc := range res.Locations {
			diag := &rdjson.Diagnostic{
				Message:  src.Runs[0].Tool.Driver.Rules[res.RuleIndex].Help.Text,
				Severity: strings.ToUpper(res.Level),
				Code: &rdjson.Code{
					Value: res.RuleID,
					Url:   src.Runs[0].Tool.Driver.Rules[res.RuleIndex].HelpURI,
				},
			}
			diag.Location = &rdjson.Location{
				Path: loc.PhysicalLocation.ArtifactLocation.URI,
				Range: &rdjson.Range{
					Start: &rdjson.Position{
						Line:   loc.PhysicalLocation.Region.StartLine,
						Column: loc.PhysicalLocation.Region.StartColumn,
					},
					End: &rdjson.Position{
						Line:   loc.PhysicalLocation.Region.EndLine,
						Column: loc.PhysicalLocation.Region.EndColumn,
					},
				},
			}
			diag.Suggestions = append(diag.Suggestions, &rdjson.Suggestion{
				Text: strings.TrimLeft(src.Runs[0].Tool.Driver.Rules[res.RuleIndex].FullDescription.Text, "\n"),
			})
			rd.Diagnostics = append(rd.Diagnostics, diag)
		}
	}

	return json.Marshal(rd)
}

func main() {
	var versionFlag bool

	flag.BoolVar(&versionFlag, "version", false, "show version and exit")

	flag.Parse()

	if versionFlag {
		fmt.Printf("version: %s\ncommit: %s\nbuild date: %s\n", version, commit, date)
		os.Exit(1)
	}

	var sarifJsonBytes []byte
	var err error

	if flag.NArg() == 1 {
		sarifJsonBytes, err = os.ReadFile(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	} else if flag.NArg() == 0 {
		sarifJsonBytes, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	}

	rdjsonBytes, err := convert(sarifJsonBytes)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(rdjsonBytes)
}

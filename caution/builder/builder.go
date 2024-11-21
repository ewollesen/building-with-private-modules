package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
	"text/template"
)

//go:embed integration.gen.go.txt
var integrationContents string

const placeholderPath = "github.com/ewollesen/building-with-private-modules/caution/placeholder"

func main() {
	var err error

	var outputFilename string = "integration.gen.go"
	flag.StringVar(&outputFilename, "output", outputFilename, "filename to write")
	flag.Parse()

	t, err := template.New("integration").Parse(integrationContents)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]any{
		"PrivatePath": placeholderPath,
	}
	if privatePath := os.Getenv("PRIVATE_MODULE_PKG_PATH"); privatePath != "" {
		data["PrivatePath"] = privatePath
	}

	out, err := os.OpenFile(outputFilename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = t.Execute(out, data)
	if err != nil {
		log.Fatal(err)
	}
}

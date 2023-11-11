//go:build generate
// +build generate

package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"github.com/Edouard127/go-mc/internal/util"
	"github.com/Edouard127/go-mc/level/block"
	"github.com/Edouard127/go-mc/nbt"
	"log"
	"os"
	"text/template"
)

//go:embed blocks.go.tmpl
var tempSource string

var temp = template.Must(template.
	New("block_template").
	Funcs(template.FuncMap{
		"UpperTheFirst": util.UpperTheFirst,
		"ToGoTypeName":  util.ToGoTypeName,
		"ToStructLiteral": func(s interface{}) string {
			return fmt.Sprintf("%#v", s)[6:]
		},
		"GetDefaultValues": GetDefaultValues,
		"Generator":        func() string { return "generator/blocks/main.go" },
	}).
	Parse(tempSource),
)

type State struct {
	Name       string
	Properties block.BlockProperty
	ID         int
	Default    map[string]any
}

//go:generate go run $GOFILE
//go:generate go fmt blocks.go
func main() {
	fmt.Println("Generating source file...")
	var states []State
	readBlockStates(&states)

	// generate go source file
	genSourceFile(states)
}

func readBlockStates(states *[]State) {
	// open block_states data file
	f, err := os.Open("blocks.nbt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		log.Panic(err)
	}

	// parse the nbt format
	if _, err := nbt.NewDecoder(r).Decode(states); err != nil {
		log.Panic(err)
	}
}

func genSourceFile(states []State) {
	var source bytes.Buffer
	if err := temp.Execute(&source, states); err != nil {
		log.Panic(err)
	}

	err := os.WriteFile("blocks.go", source.Bytes(), 0666)
	if err != nil {
		log.Panic(err)
	}
}

func GetDefaultValues(mapped map[string]any) string {
	if len(mapped) == 0 {
		return "nil"
	}
	var register = "map[states.Property]byte{"
	for key, value := range mapped {
		register += fmt.Sprintf("states.%s: %d,", key, value)
	}
	fmt.Println(register)
	return register + "}"
}

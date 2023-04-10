//go:build generate
// +build generate

package main

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	version = "1.19"
	infoURL = "https://raw.githubusercontent.com/PrismarineJS/minecraft-data/master/data/pc/" + version + "/particles.json"
	//language=gohtml
	entityTmpl = `// Code generated by gen_particles.go DO NOT EDIT.
// Package particles stores information about particles in Minecraft.
package particles
// ID describes the numeric ID of a particle.
type ID uint32

// Particle describes information about a type of particle.
type Particle struct {
ID          ID
Name        string
}

var (
	{{- range .}}
	{{.CamelName}} = Particle{
		ID: {{.ID}},
		Name: "{{.Name}}",
	}{{end}}
)

// ByID is an index of minecraft instruments by their ID.
var ByID = map[ID]*Particle{ {{range .}}
	{{.ID}}: &{{.CamelName}},{{end}}
}`
)

type Particle struct {
	ID        uint32 `json:"id"`
	CamelName string `json:"-"`
	Name      string `json:"name"`
}

func downloadInfo() ([]*Particle, error) {
	resp, err := http.Get(infoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []*Particle
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	for _, d := range data {
		d.CamelName = strcase.ToCamel(d.Name)
	}
	return data, nil
}

//go:generate go run $GOFILE
//go:generate go fmt particles.go
func main() {
	instruments, err := downloadInfo()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("particles.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := template.Must(template.New("").Parse(entityTmpl)).Execute(f, instruments); err != nil {
		panic(err)
	}
}

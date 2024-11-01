//go:build generate
// +build generate

package main

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"
)

type EnumProperty struct {
	Name       string
	TrimPrefix bool
	Values     []string
}

var EnumProperties = []EnumProperty{
	{Name: "AttachFace", Values: []string{"floor", "wall", "ceiling"}},
	{Name: "BambooLeaves", Values: []string{"none", "small", "large"}},
	{Name: "BedPart", Values: []string{"head", "foot"}},
	{Name: "BellAttachType", Values: []string{"floor", "ceiling", "single_wall", "double_wall"}},
	{Name: "ChestType", Values: []string{"single", "left", "right"}},
	{Name: "ComparatorMode", Values: []string{"compare", "subtract"}},
	{Name: "Direction", TrimPrefix: true, Values: []string{"direction_down", "direction_up", "direction_north", "direction_south", "direction_west", "direction_east"}},
	{Name: "Axis", TrimPrefix: true, Values: []string{"x", "y", "z"}},
	{Name: "DoorHingeSide", Values: []string{"left", "right"}},
	{Name: "DoubleBlockHalf", Values: []string{"upper", "lower"}},
	{Name: "DripstoneThickness", Values: []string{"tip_merge", "tip", "frustum", "middle", "base"}},
	{Name: "Half", TrimPrefix: true, Values: []string{"top", "bottom"}},
	{Name: "NoteBlockInstrument", Values: []string{
		"harp", "basedrum", "snare", "hat",
		"bass", "flute", "bell", "guitar",
		"chime", "xylophone", "iron_xylophone", "cow_bell",
		"didgeridoo", "bit", "banjo", "pling",
	}},
	{Name: "PistonType", Values: []string{"default", "sticky"}},
	{Name: "RailShape", Values: []string{
		"north_south", "east_west",
		"ascending_east", "ascending_west", "ascending_north", "ascending_south",
		"south_east", "south_west", "north_west", "north_east",
	}},
	{Name: "RedstoneSide", Values: []string{"up", "side", "none"}},
	{Name: "SculkSensorPhase", Values: []string{"inactive", "active", "cooldown"}},
	{Name: "SlabType", Values: []string{"top", "bottom", "double"}},
	{Name: "StairsShape", Values: []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
	{Name: "StructureMode", Values: []string{"save", "load", "corner", "data"}},
	{Name: "Tilt", Values: []string{"none", "unstable", "partial", "full"}},
	{Name: "WallSide", Values: []string{"none", "low", "tall"}},
	{Name: "FrontAndTop", TrimPrefix: true, Values: []string{
		"down_east", "down_north", "down_south", "down_west",
		"up_east", "up_north", "up_south", "up_west",
		"west_up", "east_up", "north_up", "south_up",
	}},
}

//go:embed data_states.go.tmpl
var tempSource string

//go:generate go run $GOFILE
//go:generate go fmt blocks.go
func main() {
	var source bytes.Buffer
	err := template.Must(template.
		New("data_states").
		Funcs(template.FuncMap{
			"UpperTheFirst": UpperTheFirst,
			"ToLower":       strings.ToLower,
			"Generator":     func() string { return "generator/properties/main.go" },
		}).
		Parse(tempSource)).
		Execute(&source, EnumProperties)
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile("data_states.go", source.Bytes(), 0666)
	if err != nil {
		log.Panic(err)
	}
}

func UpperTheFirst(word string) string {
	var sb strings.Builder
	for _, word := range strings.Split(word, "_") {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		sb.WriteString(string(runes))
	}
	return sb.String()
}

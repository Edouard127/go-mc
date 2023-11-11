package states

import "github.com/Edouard127/go-mc/level/block/states/properties"

type Property interface {
	GetName() string
	Values() (int, int)
}

type IntegerProperty struct {
	name     string
	min, max int
}

type BooleanProperty = IntegerProperty

type EnumProperty[Type properties.PropertiesEnum] struct {
	name   string
	values map[string]Type
}

func NewIntegerProperty(name string, min, max int) *IntegerProperty {
	return &IntegerProperty{
		name: name,
		min:  min,
		max:  max,
	}
}

func (p IntegerProperty) GetName() string {
	return p.name
}

func (p IntegerProperty) Values() (int, int) {
	return p.min, p.max
}

func NewBooleanProperty(name string) *BooleanProperty {
	return &BooleanProperty{name: name, min: 0, max: 1}
}

func NewEnumProperty[Type properties.PropertiesEnum](name string, values map[string]Type) *EnumProperty[Type] {
	return &EnumProperty[Type]{
		name:   name,
		values: values,
	}
}

func (p EnumProperty[Type]) GetName() string {
	return p.name
}

func (p EnumProperty[Type]) Values() (int, int) {
	return 0, len(p.values) - 1
}

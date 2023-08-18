package states

type Property[Type any] interface {
	GetName() string
	GetValue(other any) any
	CanUpdate(other any) bool // other is either an int, PropertyEnum or bool
	GetValues() []any
}

type PropertyInteger struct {
	name     string
	min, max int
	values   []any
}

type PropertyBoolean struct {
	name   string
	values []any
}

type PropertyEnum[OfType PropertiesEnum] struct {
	name   string
	values map[string]OfType
}

func fillInteger(min, max int) []any {
	var values []any
	for i := min; i <= max; i++ {
		values = append(values, i)
	}
	return values
}

func NewPropertyInteger(name string, min, max int) *PropertyInteger {
	return &PropertyInteger{
		name:   name,
		min:    min,
		max:    max,
		values: fillInteger(min, max),
	}
}

func (p PropertyInteger) GetName() string {
	return p.name
}

func (p PropertyInteger) GetValue(other any) any {
	if int(other.(uint32)) >= p.min && int(other.(uint32)) <= p.max {
		return other
	}
	return p.min
}

func (p PropertyInteger) CanUpdate(other any) bool {
	return p.min <= int(other.(uint32)) && int(other.(uint32)) <= p.max
}

func (p PropertyInteger) GetValues() []any {
	return p.values
}

func NewPropertyBoolean(name string) *PropertyBoolean {
	return &PropertyBoolean{
		name:   name,
		values: []any{true, false},
	}
}

func (p PropertyBoolean) GetName() string {
	return p.name
}

func (p PropertyBoolean) GetValue(other any) any {
	return other.(bool)
}

func (p PropertyBoolean) CanUpdate(other any) bool {
	return true // For now we will assume that all booleans are valid
}

func (p PropertyBoolean) GetValues() []any {
	return p.values
}

func NewPropertyEnum[Type PropertiesEnum](name string, values map[string]Type) *PropertyEnum[Type] {
	return &PropertyEnum[Type]{
		name:   name,
		values: values,
	}
}

func (p PropertyEnum[Type]) GetName() string {
	return p.name
}

func (p PropertyEnum[Type]) GetValue(other any) any {
	for _, value := range p.values {
		if value.String() == other {
			return value
		}
	}
	panic("invalid value")
}

func (p PropertyEnum[Type]) CanUpdate(other any) bool {
	for _, value := range p.values {
		if value.Value() == byte(other.(uint32)) {
			return true
		}
	}
	return false
}

func (p PropertyEnum[Type]) GetValues() []any {
	var values []any
	for _, value := range p.values {
		values = append(values, value)
	}
	return values
}

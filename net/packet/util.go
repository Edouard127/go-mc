package packet

import (
	"errors"
	"io"
	"reflect"
)

// Ary is used to send or receive the packet field like "Array of X"
// which has a count must be known from the context.
//
// Typically, you must decode an integer representing the length. Then
// receive the corresponding amount of data according to the length.
// In this case, the field Len should be a pointer of integer type so
// the value can be updating when Packet.Scan() method is decoding the
// previous field.
// In some special cases, you might want to read an "Array of X" with a fix length.
// So it's allowed to directly set an integer type Len, but not a pointer.
//
// Note that Ary DO read or write the Len. You aren't need to do so by your self.
type Ary[T VarInt | VarLong | Byte | UnsignedByte | Short | UnsignedShort | Int | Long] struct {
	Ary interface{} // Slice or Pointer of Slice of FieldEncoder, FieldDecoder or both (Field)
}

func (a Ary[T]) WriteTo(w io.Writer) (n int64, err error) {
	array := reflect.ValueOf(a.Ary)
	for array.Kind() == reflect.Ptr {
		array = array.Elem()
	}
	Len := T(array.Len())
	if nn, err := any(&Len).(FieldEncoder).WriteTo(w); err != nil {
		return n, err
	} else {
		n += nn
	}
	for i := 0; i < array.Len(); i++ {
		elem := array.Index(i)
		nn, err := elem.Interface().(FieldEncoder).WriteTo(w)
		n += nn
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func (a Ary[T]) ReadFrom(r io.Reader) (n int64, err error) {
	var Len T
	if nn, err := any(&Len).(FieldDecoder).ReadFrom(r); err != nil {
		return nn, err
	} else {
		n += nn
	}

	array := reflect.ValueOf(a.Ary)
	for array.Kind() == reflect.Ptr {
		array = array.Elem()
	}
	if !array.CanAddr() {
		panic(errors.New("the contents of the Ary are not addressable"))
	}
	if array.Cap() < int(Len) {
		array.Set(reflect.MakeSlice(array.Type(), int(Len), int(Len)))
	} else {
		array.Slice(0, int(Len))
	}
	for i := 0; i < int(Len); i++ {
		elem := array.Index(i)
		nn, err := elem.Addr().Interface().(FieldDecoder).ReadFrom(r)
		n += nn
		if err != nil {
			return n, err
		}
	}
	return n, err
}

func Array(ary any) Field {
	return Ary[VarInt]{Ary: ary}
}

// Optional means available to be chosen but not obligatory.
// Option means possible but not compulsory or a thing that is or may be chosen.
// The structure of Optional is similar to Option.
// But the use cases are different.
// While Option is used for code logic.

// Optional is used to send or receive the packet field like "Optional X"
// which has a bool must be known from the context.
// It is only used for buffer writing/reader. Not for code logic.
//
// Typically, you must decode a boolean representing the existence of the field.
// Then receive the corresponding amount of data according to the boolean.
// In this case, the field Has should be a pointer of bool type so
// the value can be updating when Packet.Scan() method is decoding the
// previous field.
// In some special cases, you might want to read an "Optional X" with a fix length.
// So it's allowed to directly set a bool type Has, but not a pointer.
//
// Note that Optional DO read or write the Has. You aren't need to do so by your self.
// But if you do, you might get undefined behavior.
type Optional[T FieldEncoder, P fieldPointer[T]] struct {
	Has   any // Pointer of bool, or `func() bool`
	Value T
}

func (o Optional[T, P]) has() bool {
	v := reflect.ValueOf(o.Has)
	for {
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		case reflect.Bool:
			return v.Bool()
		case reflect.Func:
			return v.Interface().(func() bool)()
		default:
			panic(errors.New("unsupported If value"))
		}
	}
}

func (o Optional[T, P]) WriteTo(w io.Writer) (n int64, err error) {
	has := o.has()
	{
		n, err = Boolean(has).WriteTo(w)
		if err != nil {
			return
		}
		if has {
			n2, _ := o.Value.WriteTo(w)
			n += n2
		}
	}
	return
}

func (o Optional[T, P]) ReadFrom(r io.Reader) (n int64, err error) {
	var has Boolean

	n, err = has.ReadFrom(r)
	if err != nil {
		return
	}

	if has {
		n2, _ := P(&o.Value).ReadFrom(r)
		n += n2
	}
	return
}

type fieldPointer[T any] interface {
	*T
	FieldDecoder
}

type Tuple []any // FieldEncoder, FieldDecoder or both (Field)

// WriteTo write Tuple to io.Writer, panic when any of filed don't implement FieldEncoder
func (t Tuple) WriteTo(w io.Writer) (n int64, err error) {
	for _, v := range t {
		nn, err := v.(FieldEncoder).WriteTo(w)
		if err != nil {
			return n, err
		}
		n += nn
	}
	return
}

// ReadFrom read Tuple from io.Reader, panic when any of field don't implement FieldDecoder
func (t Tuple) ReadFrom(r io.Reader) (n int64, err error) {
	for _, v := range t {
		nn, err := v.(FieldDecoder).ReadFrom(r)
		if err != nil {
			return n, err
		}
		n += nn
	}
	return
}

type Property struct {
	Name, Value, Signature string
}

func (p Property) WriteTo(w io.Writer) (n int64, err error) {
	return Tuple{
		String(p.Name),
		String(p.Value),
		Optional[String, *String]{
			Has:   p.Signature != "",
			Value: String(p.Signature),
		},
	}.WriteTo(w)
}

func (p *Property) ReadFrom(r io.Reader) (n int64, err error) {
	var signature Optional[String, *String]
	n, err = Tuple{
		(*String)(&p.Name),
		(*String)(&p.Value),
		&signature,
	}.ReadFrom(r)
	p.Signature = string(signature.Value)
	return
}

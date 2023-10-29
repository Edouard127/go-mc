package packet_test

import (
	"bytes"
	"fmt"
	pk "github.com/Edouard127/go-mc/net/packet"
	"testing"
)

func ExampleAry_WriteTo() {
	data := []pk.Int{0, 1, 2, 3, 4, 5, 6}
	// Len is completely ignored by WriteTo method.
	// The length is inferred from the length of Ary.
	pk.Marshal(
		0x00,
		pk.Ary[pk.VarInt]{
			Ary: data,
		},
	)
}

func ExampleAry_ReadFrom() {
	var data []pk.String

	var p pk.Packet // = conn.ReadPacket()
	if err := p.Scan(
		pk.Ary[pk.VarInt]{ // then decode Ary according to length
			Ary: &data,
		},
	); err != nil {
		panic(err)
	}
}

func TestAry_ReadFrom(t *testing.T) {
	var ary []pk.String
	var bin = []byte{
		0, 0, 0, 2,
		4, 'T', 'n', 'z', 'e',
		0,
	}
	var data = pk.Ary[pk.Int]{Ary: &ary}
	if _, err := data.ReadFrom(bytes.NewReader(bin)); err != nil {
		t.Fatal(err)
	}
	if len(ary) != 2 {
		t.Fatalf("length not match: %d != %d", len(ary), 2)
	}
	for i, v := range []string{"Tnze", ""} {
		if string(ary[i]) != v {
			t.Errorf("want %q, get %q", v, ary[i])
		}
	}
}

func TestAry_WriteTo(t *testing.T) {
	var buf bytes.Buffer
	want := []byte{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
	}
	for _, item := range [...]pk.FieldEncoder{
		pk.Ary[pk.Int]{Ary: []pk.Int{1, 2, 3}},
		pk.Ary[pk.Int]{Ary: []pk.Int{1, 2, 3}},
		pk.Ary[pk.Long]{Ary: []pk.Int{1, 2, 3}},
		pk.Ary[pk.VarInt]{Ary: []pk.Int{1, 2, 3}},
		pk.Ary[pk.VarLong]{Ary: []pk.Int{1, 2, 3}},
		pk.Ary[pk.Int]{Ary: []pk.Int{1, 2, 3}},
		pk.Ary[pk.Long]{Ary: []pk.Int{1, 2, 3}},
		pk.Ary[pk.VarInt]{Ary: []pk.Int{1, 2, 3}},
		pk.Ary[pk.VarLong]{Ary: []pk.Int{1, 2, 3}},
	} {
		_, err := item.WriteTo(&buf)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(buf.Bytes()[buf.Len()-3*4:], want) {
			t.Fatalf("Ary encoding error: got %#v, want %#v", buf.Bytes(), want)
		}
		buf.Reset()
	}
}

func TestAry_WriteTo_pointer(t *testing.T) {
	var buf bytes.Buffer
	want := []byte{
		0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
	}
	data := pk.Ary[pk.Int]{Ary: &[]pk.Int{1, 2, 3}}

	_, err := data.WriteTo(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("Ary encoding error: got %#v, want %#v", buf.Bytes(), want)
	}
}

func ExampleOptional_ReadFrom() {
	var str pk.String
	p1 := pk.Packet{Data: []byte{
		0x01,                  // pk.Boolean(true)
		4, 'T', 'n', 'z', 'e', // pk.String
	}}
	var data = pk.Optional[pk.String, *pk.String, *pk.Boolean]{Value: str}
	if err := p1.Scan(data); err != nil {
		panic(err)
	}

	var data2 pk.String = "WILL NOT BE READ, WILL NOT BE COVERED"
	p2 := pk.Packet{Data: []byte{
		0x00, // pk.Boolean(false)
		// empty
	}}
	if err := p2.Scan(pk.Optional[pk.String, *pk.String, *pk.Boolean]{Value: data2}); err != nil {
		panic(err)
	}

	// Tnze
	// WILL NOT BE READ, WILL NOT BE COVERED
}

func ExampleOptional_ReadFrom_func() {
	// As an example, we define this packet as this:
	// +------+-----------------+----------------------------------+
	// | Name | Type            | Notes                            |
	// +------+-----------------+----------------------------------+
	// | Flag | Unsigned Byte   | Odd if the following is present. |
	// +------+-----------------+----------------------------------+
	// | User | Optional String | The player's name.               |
	// +------+-----------------+----------------------------------+
	// So we need a function to decide if the User field is present.
	var data pk.String
	p := pk.Packet{Data: []byte{
		0b_0010_0011,          // pk.Byte(flag)
		4, 'T', 'n', 'z', 'e', // pk.String
	}}
	if err := p.Scan(pk.Optional[pk.String, *pk.String, *pk.Byte]{
		Value: data,
		Comp: func(p *pk.Byte) pk.Boolean {
			return *p&1 == 1
		},
	}); err != nil {
		panic(err)
	}
	fmt.Println(data)

	// Output: Tnze
}

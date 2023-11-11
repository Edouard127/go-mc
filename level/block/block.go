package block

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"github.com/Edouard127/go-mc/level/block/states"
	"github.com/Edouard127/go-mc/maths"
	"github.com/Edouard127/go-mc/nbt"
	"math/bits"
	"time"
)

// Block is an unmutable block
type Block struct {
	BlockProperty

	Name       string
	Box        maths.AxisAlignedBB
	properties map[states.Property]byte
	Default    StateID
}

func NewBlock(name string, property BlockProperty, properties map[states.Property]byte, state StateID) *Block {
	b := &Block{
		BlockProperty: property,
		Name:          name,
		Box:           maths.AxisAlignedBB{MaxX: 1, MaxY: 1, MaxZ: 1}, // We will assume all blocks are full for now
		properties:    properties,
		Default:       state,
	}
	ToStateID[b] = state
	StateList[state] = b
	return b
}

// Must not be used outside of this package, this will mess up the block states
func (b *Block) set(property states.Property, value byte) *Block {
	b.properties[property] = value
	return b
}

func (b *Block) GetDefault() *Block {
	if b.Default < 0 || b.Default >= MaxStates {
		return Air
	}
	return StateList[b.Default]
}

func (b *Block) Get(property states.Property) any {
	return b.properties[property]
}

// Equals Golang when no operator overloading :troll:
func (b *Block) Equals(other *Block) bool {
	if b.Name != other.Name {
		return false
	}

	for property, value := range b.properties {
		if other.properties[property] != value {
			return false
		}
	}

	return true
}

func (b *Block) Air() bool {
	return b.BlockProperty.IsAir
}

func (b *Block) Solid() bool {
	return b.BlockProperty.HasCollision
}

// This file stores all possible block states into a TAG_List with gzip compressed.
//
//go:embed block_states.nbt
var blockStates []byte

// MaxStates is the maximum number of states of all blocks together.
// This constant is version dependent.
const MaxStates = 24135

var (
	ToStateID = make(map[*Block]StateID, MaxStates)
	StateList = make([]*Block, MaxStates+1)
)

// BitsPerBlock indicates how many bits are needed to represent all possible
// block states. This value is used to determine the size of the global palette.
var BitsPerBlock = bits.Len(MaxStates)

type StateID int

func (s StateID) Block() *Block {
	return StateList[s]
}

type state struct {
	Name       string
	Properties nbt.RawMessage
}

func init() {
	now := time.Now()
	var s []state
	z, err := gzip.NewReader(bytes.NewReader(blockStates))
	if err != nil {
		panic(err)
	}
	_, err = nbt.NewDecoder(z).Decode(&s)
	if err != nil {
		panic(err)
	}
	for _, b := range s {
		if StateList[len(StateList)-1] != nil {
			// Default block registered
			continue
		}
		// We don't need to erase the properties, because a block cannot have different
		// properties declaration for different states. That would break the fabric of reality.
		block := *FromID[b.Name]
		// todo: Box
		if b.Properties.Type != nbt.TagEnd {
			d := nbt.NewDecoder(bytes.NewReader(b.Properties.Data))
			for {
				tag, name, err := d.ReadTag()
				if err != nil {
					panic(err)
				}
				switch tag {
				case nbt.TagEnd:
					goto end
				case nbt.TagByte:
					value, err := d.ReadByte()
					if err != nil {
						panic(err)
					}
					block.properties[states.FromName[name]] = byte(value)
				case nbt.TagInt:
					value, err := d.ReadInt()
					if err != nil {
						panic(err)
					}
					block.properties[states.FromName[name]] = byte(value)
				default:
					panic("fabric of reality broken. (invalid block data)")
				}
			}
		end:
			// nop
		}
		ToStateID[&block] = StateID(len(StateList))
		StateList = append(StateList, &block)
	}
	// TODO: Proper logging system
	println("Block states loaded in", time.Since(now)/time.Millisecond, "ms")
}

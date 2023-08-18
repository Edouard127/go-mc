package block

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"github.com/Edouard127/go-mc/data/shapes"
	"github.com/Edouard127/go-mc/level/block/states"
	"github.com/Edouard127/go-mc/maths"
	"github.com/Edouard127/go-mc/nbt"
	"math/bits"
)

var counter int

type Block struct {
	*BlockProperty
	*StateHolder
	Name string
}

func NewBlock(name string, property *BlockProperty) *Block {
	counter++
	return (&Block{
		BlockProperty: property,
		StateHolder:   NewStateHolder(make(map[states.Property[any]]uint32), StateID(counter-1)),
		Name:          name,
	}).Register()
}

func (b *Block) Register(anyp ...any) *Block {
	if len(anyp) > 0 {
		for i := range anyp {
			p := anyp[i].(states.Property[any])
			values := p.GetValues()
			for k := range values {
				sub := make(map[states.Property[any]]uint32)
				for j := range anyp {
					// Bug, doesn't make a list of possible combinations, but a list of values[k]
					sub[anyp[j].(states.Property[any])] = parseState(values[k])
				}
				b.PutNeighbors(StateID(counter), sub)
				counter++
			}
			b.SetValue(p, parseState(values[0]))
		}
	}

	return b
}

func parseState(v any) uint32 {
	switch v.(type) {
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	case int:
		return uint32(v.(int))
	case states.PropertiesEnum:
		return uint32(v.(states.PropertiesEnum).Value())
	default:
		panic(fmt.Errorf("invalid type %T for state value", v))
	}
}

func (b *Block) Is(other *Block) bool {
	return b.State() == other.State()
}

func (b *Block) IsAir() bool {
	return b.BlockProperty.IsAir
}

func (b *Block) IsLiquid() bool {
	return b.Is(Water) || b.Is(Lava)
}

func (b *Block) GetCollisionBox() maths.AxisAlignedBB[float64] {
	aabb := shapes.GetShape(b.Name, int(b.State()))
	return maths.AxisAlignedBB[float64]{
		MinX: aabb[0], MinY: aabb[1], MinZ: aabb[2],
		MaxX: aabb[3], MaxY: aabb[4], MaxZ: aabb[5],
	}
}

// This file stores all possible block states into a TAG_List with gzip compressed.
//
//go:embed block_states.nbt
var blockStates []byte

// This legacy code is still compatible with the current implementation of block states
// Because it's not complete and NOT RELIABLE, please use ToStateID and StateList
var (
	ToStateID map[*Block]StateID
	StateList []*Block
)

// BitsPerBlock indicates how many bits are needed to represent all possible
// block states. This value is used to determine the size of the global palette.
var BitsPerBlock int

type StateID int
type State struct {
	Name       string
	Properties nbt.RawMessage
}

func init() {
	var states []State
	// decompress
	z, err := gzip.NewReader(bytes.NewReader(blockStates))
	if err != nil {
		panic(err)
	}
	// decode all states
	if _, err = nbt.NewDecoder(z).Decode(&states); err != nil {
		panic(err)
	}
	ToStateID = make(map[*Block]StateID, len(states))
	StateList = make([]*Block, 0, len(states))
	for _, state := range states {
		block := FromID[state.Name]
		if state.Properties.Type != nbt.TagEnd {
			err := state.Properties.Unmarshal(&block)
			if err != nil {
				panic(err)
			}
		}
		ToStateID[block] = StateID(len(StateList))
		StateList = append(StateList, block)
	}
	BitsPerBlock = bits.Len(uint(len(StateList)))
}

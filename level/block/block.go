package block

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"github.com/Edouard127/go-mc/level/block/states"
	"github.com/Edouard127/go-mc/maths"
	"github.com/Edouard127/go-mc/nbt"
	"math/bits"
	"sync/atomic"
)

var counter atomic.Int32

type Block struct {
	BlockProperty
	*StateHolder
	Name        string
	BoundingBox maths.AxisAlignedBB
}

func NewBlock(name string, property BlockProperty) *Block {
	return &Block{
		BlockProperty: property,
		StateHolder:   NewStateHolder(make(map[states.Property]int), StateID(counter.Add(1)-1)),
		Name:          name,
		BoundingBox:   maths.AxisAlignedBB{MaxX: 1, MaxY: 1, MaxZ: 1}, // We will assume all blocks are full for now
	}
}

func (b *Block) Register(properties ...states.Property) *Block {
	if len(properties) > 0 {
		for i := range properties {
			property := properties[i]
			mn, mx := property.Values()
			for j := mn; j <= mx; j++ {
				sub := make(map[states.Property]int)
				for k := range properties {
					sub[properties[k]] = j
				}
				b.PutNeighbors(StateID(counter.Add(1)-1), sub)
			}
			b.SetValue(properties[i], mn)
		}
	}

	return b
}

func (b *Block) Is(other *Block) bool {
	return b.State() == other.State()
}

func (b *Block) IsAir() bool {
	return b.BlockProperty.IsAir
}

func (b *Block) IsSolid() bool {
	return b.BlockProperty.HasCollision
}

func (b *Block) IsLiquid() bool {
	return b.Is(Water) || b.Is(Lava)
}

/*func (b *Block) GetCollisionBox() maths.AxisAlignedBB {
	aabb := shapes.GetShape(b.Name, int(b.State()))
	return maths.AxisAlignedBB{
		MinX: aabb[0], MinY: aabb[1], MinZ: aabb[2],
		MaxX: aabb[3], MaxY: aabb[4], MaxZ: aabb[5],
	}
}*/

// This file stores all possible block states into a TAG_List with gzip compressed.
//
//go:embed block_states.nbt
var blockStates []byte

// This legacy code is still compatible with the current implementation of block states
// Because it's not complete and NOT RELIABLE, please use ToStateID and StateList
var (
	ToStateID map[*Block]StateID
	StateList = make([]*Block, 0, 21448)
)

// BitsPerBlock indicates how many bits are needed to represent all possible
// block states. This value is used to determine the size of the global palette.
var BitsPerBlock int

type StateID int

func (s StateID) Block() *Block {
	return StateList[s]
}

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

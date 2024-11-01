package level

import (
	"bytes"
	"fmt"
	"github.com/Edouard127/go-mc/level/block"
	"github.com/Edouard127/go-mc/maths"
	"github.com/Edouard127/go-mc/nbt"
	pk "github.com/Edouard127/go-mc/net/packet"
	"github.com/Edouard127/go-mc/save"
	"io"
	"math/bits"
	"strconv"
)

type ChunkPos [2]int32

func (c ChunkPos) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.Int(c[0]).WriteTo(w)
	if err != nil {
		return
	}
	n1, err := pk.Int(c[1]).WriteTo(w)
	return n + n1, err
}

func (c *ChunkPos) ReadFrom(r io.Reader) (n int64, err error) {
	var x, z pk.Int
	if n, err = x.ReadFrom(r); err != nil {
		return n, err
	}
	var n1 int64
	if n1, err = z.ReadFrom(r); err != nil {
		return n + n1, err
	}
	*c = ChunkPos{int32(x), int32(z)}
	return n + n1, nil
}

type Chunk struct {
	HeightMaps  HeightMaps
	Sections    []Section
	BlockEntity []BlockEntity
	Status      ChunkStatus
	Light       LightData
}

func (c *Chunk) IsBlockLoaded(vec3d maths.Vec3d) bool {
	_, err := c.GetBlock(vec3d)
	return err == nil
}

func (c *Chunk) GetBlock(vec3d maths.Vec3d) (*block.Block, error) {
	X, Y, Z := int(vec3d.X), int(vec3d.Y), int(vec3d.Z)
	Y += 64 // Offset so that Y=-64 is the index 0 of the array
	if Y < 0 || Y >= len(c.Sections)*16 {
		return block.Air, fmt.Errorf("y=%d out of bound", Y)
	}
	if t := c.Sections[Y>>4]; t.States != nil {
		return block.StateList[t.States.Get(Y&15<<8|Z&15<<4|X&15)], nil
	} else {
		return block.Air, fmt.Errorf("y=%d out of bound", Y)
	}
}

func (c *Chunk) SetBlock(d maths.Vec3d, i int) {
	X, Y, Z := int(d.X), int(d.Y), int(d.Z)
	Y += 64 // Offset so that Y=-64 is the index 0 of the array
	if Y < 0 || Y >= len(c.Sections)*16 {
		return // Safe check
	}
	if t := c.Sections[Y>>4]; t.States != nil {
		t.States.Set(Y&15<<8|Z&15<<4|X&15, BlocksState(i))
	}
}

var biomesIDs map[string]BiomesState
var BitsPerBiome int

var biomesNames = []string{
	"the_void",
	"plains",
	"sunflower_plains",
	"snowy_plains",
	"ice_spikes",
	"desert",
	"swamp",
	"mangrove_swamp",
	"forest",
	"flower_forest",
	"birch_forest",
	"dark_forest",
	"old_growth_birch_forest",
	"old_growth_pine_taiga",
	"old_growth_spruce_taiga",
	"taiga",
	"snowy_taiga",
	"savanna",
	"savanna_plateau",
	"windswept_hills",
	"windswept_gravelly_hills",
	"windswept_forest",
	"windswept_savanna",
	"jungle",
	"sparse_jungle",
	"bamboo_jungle",
	"badlands",
	"eroded_badlands",
	"wooded_badlands",
	"meadow",
	"grove",
	"snowy_slopes",
	"frozen_peaks",
	"jagged_peaks",
	"stony_peaks",
	"river",
	"frozen_river",
	"beach",
	"snowy_beach",
	"stony_shore",
	"warm_ocean",
	"lukewarm_ocean",
	"deep_lukewarm_ocean",
	"ocean",
	"deep_ocean",
	"cold_ocean",
	"deep_cold_ocean",
	"frozen_ocean",
	"deep_frozen_ocean",
	"mushroom_fields",
	"dripstone_caves",
	"lush_caves",
	"deep_dark",
	"nether_wastes",
	"warped_forest",
	"crimson_forest",
	"soul_sand_valley",
	"basalt_deltas",
	"the_end",
	"end_highlands",
	"end_midlands",
	"small_end_islands",
	"end_barrens",
}

func init() {
	biomesIDs = make(map[string]BiomesState, len(biomesNames))
	for i, v := range biomesNames {
		biomesIDs[v] = BiomesState(i)
	}
	BitsPerBiome = bits.Len(uint(len(biomesNames)))
}

func EmptyChunk(secs int) *Chunk {
	sections := make([]Section, secs)
	for i := range sections {
		sections[i] = Section{
			BlockCount: 0,
			States:     NewStatesPaletteContainer(16*16*16, 0),
			Biomes:     NewBiomesPaletteContainer(4*4*4, 0),
		}
	}
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			WorldSurfaceWG:         NewBitStorage(bits.Len(uint(secs)*16+1), 16*16, nil),
			WorldSurface:           NewBitStorage(bits.Len(uint(secs)*16+1), 16*16, nil),
			OceanFloorWG:           NewBitStorage(bits.Len(uint(secs)*16+1), 16*16, nil),
			OceanFloor:             NewBitStorage(bits.Len(uint(secs)*16+1), 16*16, nil),
			MotionBlocking:         NewBitStorage(bits.Len(uint(secs)*16+1), 16*16, nil),
			MotionBlockingNoLeaves: NewBitStorage(bits.Len(uint(secs)*16+1), 16*16, nil),
		},
		Status: StatusEmpty,
	}
}

// ChunkFromSave convert save.Chunk to level.Chunk.
func ChunkFromSave(c *save.Chunk) (*Chunk, error) {
	secs := len(c.Sections)
	sections := make([]Section, secs)
	for _, v := range c.Sections {
		i := int32(v.Y) - c.YPos
		if i < 0 || i >= int32(secs) {
			return nil, fmt.Errorf("section Y value %d out of bounds", v.Y)
		}
		var err error
		sections[i].States, err = readStatesPalette(v.BlockStates.Palette, v.BlockStates.Data)
		if err != nil {
			return nil, err
		}
		sections[i].BlockCount = countNoneAirBlocks(&sections[i])
		sections[i].Biomes, err = readBiomesPalette(v.Biomes.Palette, v.Biomes.Data)
		if err != nil {
			return nil, err
		}
		sections[i].SkyLight = v.SkyLight
		sections[i].BlockLight = v.BlockLight
	}

	blockEntities := make([]BlockEntity, len(c.BlockEntities))
	for i, v := range c.BlockEntities {
		var tmp struct {
			ID string `nbt:"id"`
			X  int32  `nbt:"x"`
			Y  int32  `nbt:"y"`
			Z  int32  `nbt:"z"`
		}
		if err := v.Unmarshal(&tmp); err != nil {
			return nil, err
		}
		blockEntities[i].Data = v
		if x, z := int(tmp.X-c.XPos<<4), int(tmp.Z-c.ZPos<<4); !blockEntities[i].PackXZ(x, z) {
			return nil, fmt.Errorf("Packing a XZ(" + strconv.Itoa(x) + ", " + strconv.Itoa(z) + ") out of bound")
		}
		blockEntities[i].Y = int16(tmp.Y)
		blockEntities[i].Type = 1
	}

	bitsForHeight := bits.Len( /* chunk height in blocks */ uint(secs)*16 + 1)
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			WorldSurface:           NewBitStorage(bitsForHeight, 16*16, c.Heightmaps["WORLD_SURFACE_WG"]),
			WorldSurfaceWG:         NewBitStorage(bitsForHeight, 16*16, c.Heightmaps["WORLD_SURFACE"]),
			OceanFloorWG:           NewBitStorage(bitsForHeight, 16*16, c.Heightmaps["OCEAN_FLOOR_WG"]),
			OceanFloor:             NewBitStorage(bitsForHeight, 16*16, c.Heightmaps["OCEAN_FLOOR"]),
			MotionBlocking:         NewBitStorage(bitsForHeight, 16*16, c.Heightmaps["MOTION_BLOCKING"]),
			MotionBlockingNoLeaves: NewBitStorage(bitsForHeight, 16*16, c.Heightmaps["MOTION_BLOCKING_NO_LEAVES"]),
		},
		BlockEntity: blockEntities,
		Status:      ChunkStatus(c.Status),
	}, nil
}

func readStatesPalette(palette []save.BlockState, data []uint64) (paletteData *PaletteContainer[BlocksState], err error) {
	statePalette := make([]BlocksState, len(palette))
	for i, v := range palette {
		b, ok := block.FromID[v.Name]
		if !ok {
			return nil, fmt.Errorf("unknown block id: %v", v.Name)
		}
		if v.Properties.Data != nil {
			if err := v.Properties.Unmarshal(&b); err != nil {
				return nil, fmt.Errorf("unmarshal block properties fail: %v", err)
			}
		}
		statePalette[i] = block.ToStateID[b]
	}
	paletteData = NewStatesPaletteContainerWithData(16*16*16, data, statePalette)
	return
}

func readBiomesPalette(palette []save.BiomeState, data []uint64) (*PaletteContainer[BiomesState], error) {
	biomesRawPalette := make([]BiomesState, len(palette))
	for i, v := range palette {
		err := biomesRawPalette[i].UnmarshalText([]byte(v))
		if err != nil {
			return nil, err
		}
	}
	return NewBiomesPaletteContainerWithData(4*4*4, data, biomesRawPalette), nil
}

func countNoneAirBlocks(sec *Section) (blockCount int16) {
	for i := 0; i < 16*16*16; i++ {
		if sec.GetBlock(i) != block.ToStateID[block.Air] {
			blockCount++
		}
	}
	return
}

// ChunkToSave convert level.Chunk to save.Chunk
func ChunkToSave(c *Chunk, dst *save.Chunk) (err error) {
	secs := len(c.Sections)
	sections := make([]save.Section, secs)
	for i, v := range c.Sections {
		s := &sections[i]
		states := &s.BlockStates
		biomes := &s.Biomes
		s.Y = int8(int32(i) + dst.YPos)
		states.Palette, states.Data, err = writeStatesPalette(v.States)
		if err != nil {
			return
		}
		biomes.Palette, biomes.Data, err = writeBiomesPalette(v.Biomes)
		s.SkyLight = v.SkyLight
		s.BlockLight = v.BlockLight
	}
	dst.Sections = sections
	if dst.Heightmaps == nil {
		dst.Heightmaps = make(map[string][]uint64)
	}
	fmt.Println(c.HeightMaps.WorldSurfaceWG.Raw())
	dst.Heightmaps["WORLD_SURFACE_WG"] = c.HeightMaps.WorldSurfaceWG.Raw()
	dst.Heightmaps["WORLD_SURFACE"] = c.HeightMaps.WorldSurface.Raw()
	dst.Heightmaps["OCEAN_FLOOR_WG"] = c.HeightMaps.OceanFloorWG.Raw()
	dst.Heightmaps["OCEAN_FLOOR"] = c.HeightMaps.OceanFloor.Raw()
	dst.Heightmaps["MOTION_BLOCKING"] = c.HeightMaps.MotionBlocking.Raw()
	dst.Heightmaps["MOTION_BLOCKING_NO_LEAVES"] = c.HeightMaps.MotionBlockingNoLeaves.Raw()
	dst.Status = string(c.Status)
	return
}

func writeStatesPalette(paletteData *PaletteContainer[BlocksState]) (palette []save.BlockState, data []uint64, err error) {
	if paletteData == nil {
		return
	}
	rawPalette := paletteData.palette.export()
	palette = make([]save.BlockState, len(rawPalette))

	var buffer bytes.Buffer
	for i, v := range rawPalette {
		b := block.StateList[v]
		palette[i].Name = b.Name

		buffer.Reset()
		err = nbt.NewEncoder(&buffer).Encode(b, "")
		if err != nil {
			return
		}
		_, err = nbt.NewDecoder(&buffer).Decode(&palette[i].Properties)
		if err != nil {
			return
		}
	}

	data = make([]uint64, len(paletteData.data.Raw()))
	copy(data, paletteData.data.Raw())
	return
}

func writeBiomesPalette(paletteData *PaletteContainer[BiomesState]) (palette []save.BiomeState, data []uint64, err error) {
	if paletteData == nil {
		return
	}
	rawPalette := paletteData.palette.export()
	palette = make([]save.BiomeState, len(rawPalette))

	var biomeID []byte
	for i, v := range rawPalette {
		biomeID, err = v.MarshalText()
		if err != nil {
			return
		}
		palette[i] = save.BiomeState(biomeID)
	}

	data = make([]uint64, len(paletteData.data.Raw()))
	copy(data, paletteData.data.Raw())
	return
}

func (c *Chunk) WriteTo(w io.Writer) (int64, error) {
	data, err := c.Data()
	if err != nil {
		return 0, err
	}
	light := LightData{
		SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
		BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
		SkyLight:       []pk.ByteArray{},
		BlockLight:     []pk.ByteArray{},
	}
	for i, v := range c.Sections {
		if v.SkyLight != nil {
			light.SkyLightMask.Set(int(i), true)
			light.SkyLight = append(light.SkyLight, v.SkyLight)
		}
		if v.BlockLight != nil {
			light.BlockLightMask.Set(int(i), true)
			light.BlockLight = append(light.BlockLight, v.BlockLight)
		}
	}
	return pk.Tuple{
		// Heightmaps
		pk.NBT(struct {
			MotionBlocking []uint64 `nbt:"MOTION_BLOCKING"`
			WorldSurface   []uint64 `nbt:"WORLD_SURFACE"`
		}{
			MotionBlocking: c.HeightMaps.MotionBlocking.Raw(),
			WorldSurface:   c.HeightMaps.MotionBlocking.Raw(),
		}),
		pk.ByteArray(data),
		pk.Array(c.BlockEntity),
		&light,
	}.WriteTo(w)
}

func (c *Chunk) ReadFrom(r io.Reader) (int64, error) {
	var sectionData pk.ByteArray
	n, err := pk.Tuple{
		pk.NBT(&c.HeightMaps),
		&sectionData,
	}.ReadFrom(r)
	if err != nil {
		return n, err
	}

	if len(sectionData) > 0 {
		n2, _ := pk.Array(&c.BlockEntity).ReadFrom(r)
		n += n2
	}

	data := bytes.NewReader(sectionData)
	dataLen := len(sectionData)
	c.Sections = make([]Section, len(sectionData))
	for i := 0; i < len(sectionData); i++ {
		// ?????
		if dataLen < 8 {
			break
		}

		section := &Section{
			BlockCount: 0,
			States:     NewStatesPaletteContainer(16*16*16, 0),
			Biomes:     NewBiomesPaletteContainer(4*4*4, 0),
		}

		nn, err := section.ReadFrom(data)
		if err != nil {
			return n, err
		}

		dataLen -= int(nn)
		c.Sections[i] = *section
	}
	return n, nil
}

func (c *Chunk) Data() ([]byte, error) {
	var buff bytes.Buffer
	for i := range c.Sections {
		section := c.Sections[i]
		_, err := section.WriteTo(&buff)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

func (c *Chunk) PutData(data []byte) error {
	r := bytes.NewReader(data)
	for _, section := range c.Sections {
		_, err := section.ReadFrom(r)
		if err != nil {
			return err
		}
	}
	return nil
}

type HeightMaps struct {
	WorldSurfaceWG         *BitStorage `nbt:"WORLD_SURFACE_WG,omitempty"`
	WorldSurface           *BitStorage `nbt:"WORLD_SURFACE,omitempty"`
	OceanFloorWG           *BitStorage `nbt:"OCEAN_FLOOR_WG,omitempty"`
	OceanFloor             *BitStorage `nbt:"OCEAN_FLOOR,omitempty"`
	MotionBlocking         *BitStorage `nbt:"MOTION_BLOCKING,omitempty"`
	MotionBlockingNoLeaves *BitStorage `nbt:"MOTION_BLOCKING_NO_LEAVES,omitempty"`
}

type BlockEntity struct {
	XZ   int8
	Y    int16
	Type int32
	Data nbt.RawMessage
}

func (b BlockEntity) UnpackXZ() (X, Z int) {
	return int((uint8(b.XZ) >> 4) & 0xF), int(uint8(b.XZ) & 0xF)
}

func (b *BlockEntity) PackXZ(X, Z int) bool {
	if X > 0xF || Z > 0xF || X < 0 || Z < 0 {
		return false
	}
	b.XZ = int8(X<<4 | Z)
	return true
}

func (b BlockEntity) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.Byte(b.XZ),
		pk.Short(b.Y),
		pk.VarInt(b.Type),
		pk.NBT(b.Data),
	}.WriteTo(w)
}

func (b *BlockEntity) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		(*pk.Byte)(&b.XZ),
		(*pk.Short)(&b.Y),
		(*pk.VarInt)(&b.Type),
		pk.NBT(&b.Data),
	}.ReadFrom(r)
}

type Section struct {
	BlockCount int16
	States     *PaletteContainer[BlocksState]
	Biomes     *PaletteContainer[BiomesState]
	// Half a byte per light value.
	// Could be nil if not exist
	SkyLight   []byte // len() == 2048
	BlockLight []byte // len() == 2048
}

func (s *Section) GetBlock(i int) BlocksState {
	return s.States.Get(i)
}

func (s *Section) SetBlock(i int, v BlocksState) {
	if s.States.Get(i) == 0 {
		s.BlockCount--
	} else {
		s.BlockCount++
	}
	s.States.Set(i, v)
}

func (s *Section) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Short(s.BlockCount),
		s.States,
		s.Biomes,
	}.WriteTo(w)
}

func (s *Section) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		(*pk.Short)(&s.BlockCount),
		s.States,
		s.Biomes,
	}.ReadFrom(r)
}

type LightData struct {
	SkyLightMask   pk.BitSet
	BlockLightMask pk.BitSet
	SkyLight       []pk.ByteArray
	BlockLight     []pk.ByteArray
}

func bitSetRev(set pk.BitSet) pk.BitSet {
	rev := make(pk.BitSet, len(set))
	for i := range rev {
		rev[i] = ^set[i]
	}
	return rev
}

func (l *LightData) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Boolean(true), // Trust Edges
		l.SkyLightMask,
		l.BlockLightMask,
		bitSetRev(l.SkyLightMask),
		bitSetRev(l.BlockLightMask),
		pk.Array(l.SkyLight),
		pk.Array(l.BlockLight),
	}.WriteTo(w)
}

func (l *LightData) ReadFrom(r io.Reader) (int64, error) {
	var TrustEdges pk.Boolean
	var RevSkyLightMask, RevBlockLightMask pk.BitSet
	return pk.Tuple{
		&TrustEdges, // Trust Edges
		&l.SkyLightMask,
		&l.BlockLightMask,
		&RevSkyLightMask,
		&RevBlockLightMask,
		pk.Array(&l.SkyLight),
		pk.Array(&l.BlockLight),
	}.ReadFrom(r)
}

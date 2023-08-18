package block

type BlockProperty struct {
	//material Material
	HasCollision bool `nbt:"HasCollision"`
	//soundType SoundType
	ExplosionResistance        float64 `nbt:"ExplosionResistance"`
	DestroyTime                float64 `nbt:"DestroyTime"`
	RequiresCorrectToolForDrop bool    `nbt:"RequiresCorrectToolForDrop"`
	Friction                   float64 `nbt:"Friction"`
	SpeedFactor                float64 `nbt:"SpeedFactor"`
	JumpFactor                 float64 `nbt:"JumpFaction"`
	CanOcclude                 bool    `nbt:"CanOcclude"`
	IsAir                      bool    `nbt:"IsAir"`
	DynamicShape               bool    `nbt:"DynamicShape"`
}

// This file is used to generate blocks.nbt and block_states.nbt
// The necessary environment can be generated from https://github.com/Hexeption/MCP-Reborn
package mcp.client

import net.minecraft.SharedConstants
import net.minecraft.core.registries.BuiltInRegistries
import net.minecraft.nbt.CompoundTag
import net.minecraft.nbt.ListTag
import net.minecraft.nbt.NbtIo
import net.minecraft.nbt.NbtUtils
import net.minecraft.server.Bootstrap
import net.minecraft.world.level.block.Block
import net.minecraft.world.level.block.Blocks
import net.minecraft.world.level.block.state.BlockState
import net.minecraft.world.level.block.state.properties.BlockStateProperties.*
import net.minecraft.world.level.block.state.properties.BooleanProperty
import net.minecraft.world.level.block.state.properties.EnumProperty
import net.minecraft.world.level.block.state.properties.IntegerProperty
import net.minecraft.world.level.block.state.properties.Property
import net.minecraft.world.level.material.Fluid
import java.io.DataOutputStream
import java.io.FileOutputStream
import java.util.zip.GZIPOutputStream

object Start {
    private lateinit var mappings: Map<Property<*>, String>
    @JvmStatic
    fun main(args: Array<String>) {
        println("program start!")
        SharedConstants.tryDetectVersion()
        Bootstrap.bootStrap()
        Blocks.rebuildCache()

        mappings = mapOf(
            ATTACHED to "attached",
            BOTTOM to "bottom",
            CONDITIONAL to "conditional",
            DISARMED to "disarmed",
            DRAG to "drag",
            ENABLED to "enabled",
            EXTENDED to "extended",
            EYE to "eye",
            FALLING to "falling",
            HANGING to "hanging",
            HAS_BOTTLE_0 to "has_bottle_0",
            HAS_BOTTLE_1 to "has_bottle_1",
            HAS_BOTTLE_2 to "has_bottle_2",
            HAS_RECORD to "has_record",
            HAS_BOOK to "has_book",
            INVERTED to "inverted",
            IN_WALL to "in_wall",
            LIT to "lit",
            LOCKED to "locked",
            OCCUPIED to "occupied",
            OPEN to "open",
            PERSISTENT to "persistent",
            POWERED to "powered",
            SHORT to "short",
            SIGNAL_FIRE to "signal_fire",
            SNOWY to "snowy",
            TRIGGERED to "triggered",
            UNSTABLE to "unstable",
            WATERLOGGED to "waterlogged",
            BERRIES to "berries",
            BLOOM to "bloom",
            SHRIEKING to "shrieking",
            CAN_SUMMON to "can_summon",
            HORIZONTAL_AXIS to "horizontal_axis",
            AXIS to "axis",
            UP to "up",
            DOWN to "down",
            NORTH to "north",
            SOUTH to "south",
            WEST to "west",
            EAST to "east",
            FACING to "facing",
            FACING_HOPPER to "facing_hopper",
            HORIZONTAL_FACING to "horizontal_facing",
            FLOWER_AMOUNT to "flower_amount",
            ORIENTATION to "orientation",
            ATTACH_FACE to "attach_face",
            BELL_ATTACHMENT to "bell_attachment",
            EAST_WALL to "east_wall",
            NORTH_WALL to "north_wall",
            SOUTH_WALL to "south_wall",
            WEST_WALL to "west_wall",
            EAST_REDSTONE to "east_redstone",
            NORTH_REDSTONE to "north_redstone",
            SOUTH_REDSTONE to "south_redstone",
            WEST_REDSTONE to "west_redstone",
            DOUBLE_BLOCK_HALF to "double_block_half",
            HALF to "half",
            RAIL_SHAPE to "rail_shape",
            RAIL_SHAPE_STRAIGHT to "rail_shape_straight",
            AGE_1 to "age_1",
            AGE_2 to "age_2",
            AGE_3 to "age_3",
            AGE_4 to "age_4",
            AGE_5 to "age_5",
            AGE_7 to "age_7",
            AGE_15 to "age_15",
            AGE_25 to "age_25",
            BITES to "bites",
            CANDLES to "candles",
            DELAY to "delay",
            DISTANCE to "distance",
            EGGS to "eggs",
            HATCH to "hatch",
            LAYERS to "layers",
            LEVEL_CAULDRON to "level_cauldron",
            LEVEL_COMPOSTER to "level_composter",
            LEVEL_FLOWING to "level_flowing",
            LEVEL_HONEY to "level_honey",
            LEVEL to "level",
            MOISTURE to "moisture",
            NOTE to "note",
            PICKLES to "pickles",
            POWER to "power",
            STAGE to "stage",
            STABILITY_DISTANCE to "stability_distance",
            RESPAWN_ANCHOR_CHARGES to "respawn_anchor_charges",
            ROTATION_16 to "rotation_16",
            BED_PART to "bed_part",
            CHEST_TYPE to "chest_type",
            MODE_COMPARATOR to "mode_comparator",
            DOOR_HINGE to "door_hinge",
            NOTEBLOCK_INSTRUMENT to "noteblock_instrument",
            PISTON_TYPE to "piston_type",
            SLAB_TYPE to "slab_type",
            STAIRS_SHAPE to "stairs_shape",
            STRUCTUREBLOCK_MODE to "structureblock_mode",
            BAMBOO_LEAVES to "bamboo_leaves",
            TILT to "tilt",
            VERTICAL_DIRECTION to "vertical_direction",
            DRIPSTONE_THICKNESS to "dripstone_thickness",
            SCULK_SENSOR_PHASE to "sculk_sensor_phase",
            CHISELED_BOOKSHELF_SLOT_0_OCCUPIED to "chiseled_bookshelf_slot_0_occupied",
            CHISELED_BOOKSHELF_SLOT_1_OCCUPIED to "chiseled_bookshelf_slot_1_occupied",
            CHISELED_BOOKSHELF_SLOT_2_OCCUPIED to "chiseled_bookshelf_slot_2_occupied",
            CHISELED_BOOKSHELF_SLOT_3_OCCUPIED to "chiseled_bookshelf_slot_3_occupied",
            CHISELED_BOOKSHELF_SLOT_4_OCCUPIED to "chiseled_bookshelf_slot_4_occupied",
            CHISELED_BOOKSHELF_SLOT_5_OCCUPIED to "chiseled_bookshelf_slot_5_occupied",
            DUSTED to "dusted",
            CRACKED to "cracked",
        )

        FileOutputStream("blocks.nbt").use { stream ->
            GZIPOutputStream(stream).use {
                NbtIo.writeUnnamedTag(blocks(), DataOutputStream(it))
            }
        }

        FileOutputStream("block_states.nbt").use { stream ->
            GZIPOutputStream(stream).use {
                NbtIo.writeUnnamedTag(blockStates(), DataOutputStream(it))
            }
        }

        FileOutputStream("fluid_states.nbt").use { stream ->
            GZIPOutputStream(stream).use {
                NbtIo.writeUnnamedTag(fluidStates(), DataOutputStream(it))
            }
        }
    }

    private fun blocks(): ListTag {
        return ListTag().apply {
            BuiltInRegistries.BLOCK.forEach { block ->
                val states = CompoundTag().apply {
                    putString("Name", BuiltInRegistries.BLOCK.getKey(block).toString())

                    put("Properties", CompoundTag().apply {
                        putBoolean("HasCollision", block.hasCollision)
                        putFloat("ExplosionResistance", block.explosionResistance)
                        putFloat("DestroyTime", block.properties.destroyTime)
                        putBoolean("RequiresCorrectToolForDrop", block.properties.requiresCorrectToolForDrops)
                        putFloat("Friction", block.friction)
                        putFloat("SpeedFactor", block.speedFactor)
                        putFloat("JumpFactor", block.jumpFactor)
                        putBoolean("CanOcclude", block.properties.canOcclude)
                        putBoolean("IsAir", block.properties.isAir)
                        putBoolean("DynamicShape", block.dynamicShape)
                    })

                    put("Default", CompoundTag().apply {
                        block.defaultBlockState().values.forEach { (property, value) ->
                            if (mappings[property] == null) println("Unknown property: " + property.name)
                            val name = toGoTypeName(mappings[property]!!)
                            when (property) {
                                is BooleanProperty -> putBoolean(name, value as Boolean)
                                is IntegerProperty -> putInt(name, value as Int)
                                is EnumProperty -> putInt(name, (value as Enum).ordinal)
                                else -> println("Unknown type: " + value.javaClass.name)
                            }
                        }
                    })
                }

                // Put the data into the nbt
                add(states)
            }
        }
    }

    private fun toGoTypeName(str: String): String {
        return str.split("_").joinToString("") { it.capitalize() }
    }

    private fun blockStates(): ListTag = ListTag().apply { addAll(Block.BLOCK_STATE_REGISTRY.map { writeBlockState(it) }) }

    private fun fluidStates(): ListTag = ListTag().apply { addAll(Fluid.FLUID_STATE_REGISTRY.map { NbtUtils.writeFluidState(it) }) }

    private fun writeBlockState(state: BlockState): CompoundTag {
        return CompoundTag().apply {
            putString("Name", BuiltInRegistries.BLOCK.getKey(state.block).toString())
            if (state.values.isEmpty()) return this
            put("Properties", CompoundTag().apply {
                for ((property, value) in state.values) {
                    val name = mappings[property]!!
                    when (property) {
                        is BooleanProperty -> putBoolean(name, value as Boolean)
                        is IntegerProperty -> putInt(name, value as Int)
                        is EnumProperty<*> -> putInt(name, (value as Enum).ordinal)
                    }
                }
            })
        }
    }
}
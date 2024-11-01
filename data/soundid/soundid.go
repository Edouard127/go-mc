// Code generated by gen_soundid.go. DO NOT EDIT.

package soundid

// SoundID represents a sound ID used in the minecraft protocol.
type SoundID int32

// SoundNames - map of ids to names for sounds.
var SoundNames = map[SoundID]string{ 
	0: "ambient.cave",
	1: "block.anvil.break",
	2: "block.anvil.destroy",
	3: "block.anvil.fall",
	4: "block.anvil.hit",
	5: "block.anvil.land",
	6: "block.anvil.place",
	7: "block.anvil.step",
	8: "block.anvil.use",
	9: "block.brewing_stand.brew",
	10: "block.chest.close",
	11: "block.chest.locked",
	12: "block.chest.open",
	13: "block.chorus_flower.death",
	14: "block.chorus_flower.grow",
	15: "block.cloth.break",
	16: "block.cloth.fall",
	17: "block.cloth.hit",
	18: "block.cloth.place",
	19: "block.cloth.step",
	20: "block.comparator.click",
	21: "block.dispenser.dispense",
	22: "block.dispenser.fail",
	23: "block.dispenser.launch",
	24: "block.enchantment_table.use",
	25: "block.end_gateway.spawn",
	26: "block.end_portal.spawn",
	27: "block.end_portal_frame.fill",
	28: "block.enderchest.close",
	29: "block.enderchest.open",
	30: "block.fence_gate.close",
	31: "block.fence_gate.open",
	32: "block.fire.ambient",
	33: "block.fire.extinguish",
	34: "block.furnace.fire_crackle",
	35: "block.glass.break",
	36: "block.glass.fall",
	37: "block.glass.hit",
	38: "block.glass.place",
	39: "block.glass.step",
	40: "block.grass.break",
	41: "block.grass.fall",
	42: "block.grass.hit",
	43: "block.grass.place",
	44: "block.grass.step",
	45: "block.gravel.break",
	46: "block.gravel.fall",
	47: "block.gravel.hit",
	48: "block.gravel.place",
	49: "block.gravel.step",
	50: "block.iron_door.close",
	51: "block.iron_door.open",
	52: "block.iron_trapdoor.close",
	53: "block.iron_trapdoor.open",
	54: "block.ladder.break",
	55: "block.ladder.fall",
	56: "block.ladder.hit",
	57: "block.ladder.place",
	58: "block.ladder.step",
	59: "block.lava.ambient",
	60: "block.lava.extinguish",
	61: "block.lava.pop",
	62: "block.lever.click",
	63: "block.metal.break",
	64: "block.metal.fall",
	65: "block.metal.hit",
	66: "block.metal.place",
	67: "block.metal.step",
	68: "block.metal_pressureplate.click_off",
	69: "block.metal_pressureplate.click_on",
	70: "block.note.basedrum",
	71: "block.note.bass",
	72: "block.note.bell",
	73: "block.note.chime",
	74: "block.note.flute",
	75: "block.note.guitar",
	76: "block.note.harp",
	77: "block.note.hat",
	78: "block.note.pling",
	79: "block.note.snare",
	80: "block.note.xylophone",
	81: "block.piston.contract",
	82: "block.piston.extend",
	83: "block.portal.ambient",
	84: "block.portal.travel",
	85: "block.portal.trigger",
	86: "block.redstone_torch.burnout",
	87: "block.sand.break",
	88: "block.sand.fall",
	89: "block.sand.hit",
	90: "block.sand.place",
	91: "block.sand.step",
	92: "block.shulker_box.close",
	93: "block.shulker_box.open",
	94: "block.slime.break",
	95: "block.slime.fall",
	96: "block.slime.hit",
	97: "block.slime.place",
	98: "block.slime.step",
	99: "block.snow.break",
	100: "block.snow.fall",
	101: "block.snow.hit",
	102: "block.snow.place",
	103: "block.snow.step",
	104: "block.stone.break",
	105: "block.stone.fall",
	106: "block.stone.hit",
	107: "block.stone.place",
	108: "block.stone.step",
	109: "block.stone_button.click_off",
	110: "block.stone_button.click_on",
	111: "block.stone_pressureplate.click_off",
	112: "block.stone_pressureplate.click_on",
	113: "block.tripwire.attach",
	114: "block.tripwire.click_off",
	115: "block.tripwire.click_on",
	116: "block.tripwire.detach",
	117: "block.water.ambient",
	118: "block.waterlily.place",
	119: "block.wood.break",
	120: "block.wood.fall",
	121: "block.wood.hit",
	122: "block.wood.place",
	123: "block.wood.step",
	124: "block.wood_button.click_off",
	125: "block.wood_button.click_on",
	126: "block.wood_pressureplate.click_off",
	127: "block.wood_pressureplate.click_on",
	128: "block.wooden_door.close",
	129: "block.wooden_door.open",
	130: "block.wooden_trapdoor.close",
	131: "block.wooden_trapdoor.open",
	132: "enchant.thorns.hit",
	133: "entity.armorstand.break",
	134: "entity.armorstand.fall",
	135: "entity.armorstand.hit",
	136: "entity.armorstand.place",
	137: "entity.arrow.hit",
	138: "entity.arrow.hit_player",
	139: "entity.arrow.shoot",
	140: "entity.bat.ambient",
	141: "entity.bat.death",
	142: "entity.bat.hurt",
	143: "entity.bat.loop",
	144: "entity.bat.takeoff",
	145: "entity.blaze.ambient",
	146: "entity.blaze.burn",
	147: "entity.blaze.death",
	148: "entity.blaze.hurt",
	149: "entity.blaze.shoot",
	150: "entity.boat.paddle_land",
	151: "entity.boat.paddle_water",
	152: "entity.bobber.retrieve",
	153: "entity.bobber.splash",
	154: "entity.bobber.throw",
	155: "entity.cat.ambient",
	156: "entity.cat.death",
	157: "entity.cat.hiss",
	158: "entity.cat.hurt",
	159: "entity.cat.purr",
	160: "entity.cat.purreow",
	161: "entity.chicken.ambient",
	162: "entity.chicken.death",
	163: "entity.chicken.egg",
	164: "entity.chicken.hurt",
	165: "entity.chicken.step",
	166: "entity.cow.ambient",
	167: "entity.cow.death",
	168: "entity.cow.hurt",
	169: "entity.cow.milk",
	170: "entity.cow.step",
	171: "entity.creeper.death",
	172: "entity.creeper.hurt",
	173: "entity.creeper.primed",
	174: "entity.donkey.ambient",
	175: "entity.donkey.angry",
	176: "entity.donkey.chest",
	177: "entity.donkey.death",
	178: "entity.donkey.hurt",
	179: "entity.egg.throw",
	180: "entity.elder_guardian.ambient",
	181: "entity.elder_guardian.ambient_land",
	182: "entity.elder_guardian.curse",
	183: "entity.elder_guardian.death",
	184: "entity.elder_guardian.death_land",
	185: "entity.elder_guardian.flop",
	186: "entity.elder_guardian.hurt",
	187: "entity.elder_guardian.hurt_land",
	188: "entity.enderdragon.ambient",
	189: "entity.enderdragon.death",
	190: "entity.enderdragon.flap",
	191: "entity.enderdragon.growl",
	192: "entity.enderdragon.hurt",
	193: "entity.enderdragon.shoot",
	194: "entity.enderdragon_fireball.explode",
	195: "entity.endereye.death",
	196: "entity.endereye.launch",
	197: "entity.endermen.ambient",
	198: "entity.endermen.death",
	199: "entity.endermen.hurt",
	200: "entity.endermen.scream",
	201: "entity.endermen.stare",
	202: "entity.endermen.teleport",
	203: "entity.endermite.ambient",
	204: "entity.endermite.death",
	205: "entity.endermite.hurt",
	206: "entity.endermite.step",
	207: "entity.enderpearl.throw",
	208: "entity.evocation_fangs.attack",
	209: "entity.evocation_illager.ambient",
	210: "entity.evocation_illager.cast_spell",
	211: "entity.evocation_illager.death",
	212: "entity.evocation_illager.hurt",
	213: "entity.evocation_illager.prepare_attack",
	214: "entity.evocation_illager.prepare_summon",
	215: "entity.evocation_illager.prepare_wololo",
	216: "entity.experience_bottle.throw",
	217: "entity.experience_orb.pickup",
	218: "entity.firework.blast",
	219: "entity.firework.blast_far",
	220: "entity.firework.large_blast",
	221: "entity.firework.large_blast_far",
	222: "entity.firework.launch",
	223: "entity.firework.shoot",
	224: "entity.firework.twinkle",
	225: "entity.firework.twinkle_far",
	226: "entity.generic.big_fall",
	227: "entity.generic.burn",
	228: "entity.generic.death",
	229: "entity.generic.drink",
	230: "entity.generic.eat",
	231: "entity.generic.explode",
	232: "entity.generic.extinguish_fire",
	233: "entity.generic.hurt",
	234: "entity.generic.small_fall",
	235: "entity.generic.splash",
	236: "entity.generic.swim",
	237: "entity.ghast.ambient",
	238: "entity.ghast.death",
	239: "entity.ghast.hurt",
	240: "entity.ghast.scream",
	241: "entity.ghast.shoot",
	242: "entity.ghast.warn",
	243: "entity.guardian.ambient",
	244: "entity.guardian.ambient_land",
	245: "entity.guardian.attack",
	246: "entity.guardian.death",
	247: "entity.guardian.death_land",
	248: "entity.guardian.flop",
	249: "entity.guardian.hurt",
	250: "entity.guardian.hurt_land",
	251: "entity.horse.ambient",
	252: "entity.horse.angry",
	253: "entity.horse.armor",
	254: "entity.horse.breathe",
	255: "entity.horse.death",
	256: "entity.horse.eat",
	257: "entity.horse.gallop",
	258: "entity.horse.hurt",
	259: "entity.horse.jump",
	260: "entity.horse.land",
	261: "entity.horse.saddle",
	262: "entity.horse.step",
	263: "entity.horse.step_wood",
	264: "entity.hostile.big_fall",
	265: "entity.hostile.death",
	266: "entity.hostile.hurt",
	267: "entity.hostile.small_fall",
	268: "entity.hostile.splash",
	269: "entity.hostile.swim",
	270: "entity.husk.ambient",
	271: "entity.husk.death",
	272: "entity.husk.hurt",
	273: "entity.husk.step",
	274: "entity.illusion_illager.ambient",
	275: "entity.illusion_illager.cast_spell",
	276: "entity.illusion_illager.death",
	277: "entity.illusion_illager.hurt",
	278: "entity.illusion_illager.mirror_move",
	279: "entity.illusion_illager.prepare_blindness",
	280: "entity.illusion_illager.prepare_mirror",
	281: "entity.irongolem.attack",
	282: "entity.irongolem.death",
	283: "entity.irongolem.hurt",
	284: "entity.irongolem.step",
	285: "entity.item.break",
	286: "entity.item.pickup",
	287: "entity.itemframe.add_item",
	288: "entity.itemframe.break",
	289: "entity.itemframe.place",
	290: "entity.itemframe.remove_item",
	291: "entity.itemframe.rotate_item",
	292: "entity.leashknot.break",
	293: "entity.leashknot.place",
	294: "entity.lightning.impact",
	295: "entity.lightning.thunder",
	296: "entity.lingeringpotion.throw",
	297: "entity.llama.ambient",
	298: "entity.llama.angry",
	299: "entity.llama.chest",
	300: "entity.llama.death",
	301: "entity.llama.eat",
	302: "entity.llama.hurt",
	303: "entity.llama.spit",
	304: "entity.llama.step",
	305: "entity.llama.swag",
	306: "entity.magmacube.death",
	307: "entity.magmacube.hurt",
	308: "entity.magmacube.jump",
	309: "entity.magmacube.squish",
	310: "entity.minecart.inside",
	311: "entity.minecart.riding",
	312: "entity.mooshroom.shear",
	313: "entity.mule.ambient",
	314: "entity.mule.chest",
	315: "entity.mule.death",
	316: "entity.mule.hurt",
	317: "entity.painting.break",
	318: "entity.painting.place",
	319: "entity.parrot.ambient",
	320: "entity.parrot.death",
	321: "entity.parrot.eat",
	322: "entity.parrot.fly",
	323: "entity.parrot.hurt",
	324: "entity.parrot.imitate.blaze",
	325: "entity.parrot.imitate.creeper",
	326: "entity.parrot.imitate.elder_guardian",
	327: "entity.parrot.imitate.enderdragon",
	328: "entity.parrot.imitate.enderman",
	329: "entity.parrot.imitate.endermite",
	330: "entity.parrot.imitate.evocation_illager",
	331: "entity.parrot.imitate.ghast",
	332: "entity.parrot.imitate.husk",
	333: "entity.parrot.imitate.illusion_illager",
	334: "entity.parrot.imitate.magmacube",
	335: "entity.parrot.imitate.polar_bear",
	336: "entity.parrot.imitate.shulker",
	337: "entity.parrot.imitate.silverfish",
	338: "entity.parrot.imitate.skeleton",
	339: "entity.parrot.imitate.slime",
	340: "entity.parrot.imitate.spider",
	341: "entity.parrot.imitate.stray",
	342: "entity.parrot.imitate.vex",
	343: "entity.parrot.imitate.vindication_illager",
	344: "entity.parrot.imitate.witch",
	345: "entity.parrot.imitate.wither",
	346: "entity.parrot.imitate.wither_skeleton",
	347: "entity.parrot.imitate.wolf",
	348: "entity.parrot.imitate.zombie",
	349: "entity.parrot.imitate.zombie_pigman",
	350: "entity.parrot.imitate.zombie_villager",
	351: "entity.parrot.step",
	352: "entity.pig.ambient",
	353: "entity.pig.death",
	354: "entity.pig.hurt",
	355: "entity.pig.saddle",
	356: "entity.pig.step",
	357: "entity.player.attack.crit",
	358: "entity.player.attack.knockback",
	359: "entity.player.attack.nodamage",
	360: "entity.player.attack.strong",
	361: "entity.player.attack.sweep",
	362: "entity.player.attack.weak",
	363: "entity.player.big_fall",
	364: "entity.player.breath",
	365: "entity.player.burp",
	366: "entity.player.death",
	367: "entity.player.hurt",
	368: "entity.player.hurt_drown",
	369: "entity.player.hurt_on_fire",
	370: "entity.player.levelup",
	371: "entity.player.small_fall",
	372: "entity.player.splash",
	373: "entity.player.swim",
	374: "entity.polar_bear.ambient",
	375: "entity.polar_bear.baby_ambient",
	376: "entity.polar_bear.death",
	377: "entity.polar_bear.hurt",
	378: "entity.polar_bear.step",
	379: "entity.polar_bear.warning",
	380: "entity.rabbit.ambient",
	381: "entity.rabbit.attack",
	382: "entity.rabbit.death",
	383: "entity.rabbit.hurt",
	384: "entity.rabbit.jump",
	385: "entity.sheep.ambient",
	386: "entity.sheep.death",
	387: "entity.sheep.hurt",
	388: "entity.sheep.shear",
	389: "entity.sheep.step",
	390: "entity.shulker.ambient",
	391: "entity.shulker.close",
	392: "entity.shulker.death",
	393: "entity.shulker.hurt",
	394: "entity.shulker.hurt_closed",
	395: "entity.shulker.open",
	396: "entity.shulker.shoot",
	397: "entity.shulker.teleport",
	398: "entity.shulker_bullet.hit",
	399: "entity.shulker_bullet.hurt",
	400: "entity.silverfish.ambient",
	401: "entity.silverfish.death",
	402: "entity.silverfish.hurt",
	403: "entity.silverfish.step",
	404: "entity.skeleton.ambient",
	405: "entity.skeleton.death",
	406: "entity.skeleton.hurt",
	407: "entity.skeleton.shoot",
	408: "entity.skeleton.step",
	409: "entity.skeleton_horse.ambient",
	410: "entity.skeleton_horse.death",
	411: "entity.skeleton_horse.hurt",
	412: "entity.slime.attack",
	413: "entity.slime.death",
	414: "entity.slime.hurt",
	415: "entity.slime.jump",
	416: "entity.slime.squish",
	417: "entity.small_magmacube.death",
	418: "entity.small_magmacube.hurt",
	419: "entity.small_magmacube.squish",
	420: "entity.small_slime.death",
	421: "entity.small_slime.hurt",
	422: "entity.small_slime.jump",
	423: "entity.small_slime.squish",
	424: "entity.snowball.throw",
	425: "entity.snowman.ambient",
	426: "entity.snowman.death",
	427: "entity.snowman.hurt",
	428: "entity.snowman.shoot",
	429: "entity.spider.ambient",
	430: "entity.spider.death",
	431: "entity.spider.hurt",
	432: "entity.spider.step",
	433: "entity.splash_potion.break",
	434: "entity.splash_potion.throw",
	435: "entity.squid.ambient",
	436: "entity.squid.death",
	437: "entity.squid.hurt",
	438: "entity.stray.ambient",
	439: "entity.stray.death",
	440: "entity.stray.hurt",
	441: "entity.stray.step",
	442: "entity.tnt.primed",
	443: "entity.vex.ambient",
	444: "entity.vex.charge",
	445: "entity.vex.death",
	446: "entity.vex.hurt",
	447: "entity.villager.ambient",
	448: "entity.villager.death",
	449: "entity.villager.hurt",
	450: "entity.villager.no",
	451: "entity.villager.trading",
	452: "entity.villager.yes",
	453: "entity.vindication_illager.ambient",
	454: "entity.vindication_illager.death",
	455: "entity.vindication_illager.hurt",
	456: "entity.witch.ambient",
	457: "entity.witch.death",
	458: "entity.witch.drink",
	459: "entity.witch.hurt",
	460: "entity.witch.throw",
	461: "entity.wither.ambient",
	462: "entity.wither.break_block",
	463: "entity.wither.death",
	464: "entity.wither.hurt",
	465: "entity.wither.shoot",
	466: "entity.wither.spawn",
	467: "entity.wither_skeleton.ambient",
	468: "entity.wither_skeleton.death",
	469: "entity.wither_skeleton.hurt",
	470: "entity.wither_skeleton.step",
	471: "entity.wolf.ambient",
	472: "entity.wolf.death",
	473: "entity.wolf.growl",
	474: "entity.wolf.howl",
	475: "entity.wolf.hurt",
	476: "entity.wolf.pant",
	477: "entity.wolf.shake",
	478: "entity.wolf.step",
	479: "entity.wolf.whine",
	480: "entity.zombie.ambient",
	481: "entity.zombie.attack_door_wood",
	482: "entity.zombie.attack_iron_door",
	483: "entity.zombie.break_door_wood",
	484: "entity.zombie.death",
	485: "entity.zombie.hurt",
	486: "entity.zombie.infect",
	487: "entity.zombie.step",
	488: "entity.zombie_horse.ambient",
	489: "entity.zombie_horse.death",
	490: "entity.zombie_horse.hurt",
	491: "entity.zombie_pig.ambient",
	492: "entity.zombie_pig.angry",
	493: "entity.zombie_pig.death",
	494: "entity.zombie_pig.hurt",
	495: "entity.zombie_villager.ambient",
	496: "entity.zombie_villager.converted",
	497: "entity.zombie_villager.cure",
	498: "entity.zombie_villager.death",
	499: "entity.zombie_villager.hurt",
	500: "entity.zombie_villager.step",
	501: "item.armor.equip_chain",
	502: "item.armor.equip_diamond",
	503: "item.armor.equip_elytra",
	504: "item.armor.equip_generic",
	505: "item.armor.equip_gold",
	506: "item.armor.equip_iron",
	507: "item.armor.equip_leather",
	508: "item.bottle.empty",
	509: "item.bottle.fill",
	510: "item.bottle.fill_dragonbreath",
	511: "item.bucket.empty",
	512: "item.bucket.empty_lava",
	513: "item.bucket.fill",
	514: "item.bucket.fill_lava",
	515: "item.chorus_fruit.teleport",
	516: "item.elytra.flying",
	517: "item.firecharge.use",
	518: "item.flintandsteel.use",
	519: "item.hoe.till",
	520: "item.shield.block",
	521: "item.shield.break",
	522: "item.shovel.flatten",
	523: "item.totem.use",
	524: "music.creative",
	525: "music.credits",
	526: "music.dragon",
	527: "music.end",
	528: "music.game",
	529: "music.menu",
	530: "music.nether",
	531: "record.11",
	532: "record.13",
	533: "record.blocks",
	534: "record.cat",
	535: "record.chirp",
	536: "record.far",
	537: "record.mall",
	538: "record.mellohi",
	539: "record.stal",
	540: "record.strad",
	541: "record.wait",
	542: "record.ward",
	543: "ui.button.click",
	544: "ui.toast.in",
	545: "ui.toast.out",
	546: "ui.toast.challenge_complete",
	547: "weather.rain",
	548: "weather.rain.above",
}

// GetSoundNameByID helper method
func GetSoundNameByID(id SoundID) (string, bool) {
	name, ok := SoundNames[id]
	return name, ok
}
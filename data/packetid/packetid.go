package packetid

// Login Clientbound
const (
	CPacketLoginDisconnect = iota
	CPacketEncryptionRequest
	CPacketLoginSuccess
	CPacketSetCompression
	CPacketLoginPluginRequest
)

// Login Serverbound
const (
	SPacketLoginStart = iota
	SPacketEncryptionResponse
	SPacketLoginPluginResponse
)

// Status Clientbound
const (
	CPacketStatusResponse = iota
	CPacketStatusPongResponse
)

// Status Serverbound
const (
	SPacketStatusRequest = iota
	SPacketPingRequest
)

// Game Clientbound
const (
	CPacketBundleLimiter             = iota // https://wiki.vg/Protocol#Bundle_Delimiter
	CPacketSpawnEntity                      // https://wiki.vg/Protocol#Spawn_Entity
	CPacketSpawnExperienceOrb               // https://wiki.vg/Protocol#Spawn_Experience_Orb
	CPacketSpawnPlayer                      // https://wiki.vg/Protocol#Spawn_Experience_Orb
	CPacketEntityAnimation                  // https://wiki.vg/Protocol#Spawn_Player
	CPacketAwardStats                       // https://wiki.vg/Protocol#Award_Statistics
	CPacketAcknowledgeBlockChange           // https://wiki.vg/Protocol#Acknowledge_Block_Change
	CPacketSetBlockDestroyStage             // https://wiki.vg/Protocol#Set_Block_Destroy_Stage
	CPacketBlockEntityData                  // https://wiki.vg/Protocol#Block_Entity_Data
	CPacketBlockAction                      // https://wiki.vg/Protocol#Block_Action
	CPacketBlockUpdate                      // https://wiki.vg/Protocol#Block_Update
	CPacketBossBar                          // https://wiki.vg/Protocol#Boss_Bar
	CPacketServerDifficulty                 // https://wiki.vg/Protocol#Change_Difficulty
	CPacketChunkBiomes                      // https://wiki.vg/Protocol#Chunk_Biomes
	CPacketClearTitles                      // https://wiki.vg/Protocol#Clear_Titles
	CPacketCommandSuggestions               // https://wiki.vg/Protocol#Command_Suggestions_Response
	CPacketCommands                         // https://wiki.vg/Protocol#Commands
	CPacketCloseContainer                   // https://wiki.vg/Protocol#Close_Container
	CPacketSetContainerContent              // https://wiki.vg/Protocol#Set_Container_Content
	CPacketSetContainerProperty             // https://wiki.vg/Protocol#Set_Container_Property
	CPacketSetContainerSlot                 // https://wiki.vg/Protocol#Set_Container_Slot
	CPacketSetCooldown                      // https://wiki.vg/Protocol#Set_Cooldown
	CPacketChatSuggestions                  // https://wiki.vg/Protocol#Chat_Suggestions
	CPacketPluginMessage                    // https://wiki.vg/Protocol#Plugin_Message
	CPacketDamageEvent                      // https://wiki.vg/Protocol#Damage_Event
	CPacketDeleteMessage                    // https://wiki.vg/Protocol#Delete_Message
	CPacketDisconnect                       // https://wiki.vg/Protocol#Disconnect
	CPacketDisguisedMessage                 // https://wiki.vg/Protocol#Disguised_Chat_Message
	CPacketEntityEvent                      // https://wiki.vg/Protocol#Entity_Event
	CPacketExplosion                        // https://wiki.vg/Protocol#Explosion
	CPacketUnloadChunk                      // https://wiki.vg/Protocol#Unload_Chunk
	CPacketGameEvent                        // https://wiki.vg/Protocol#Game_Event
	CPacketOpenHorseWindow                  // https://wiki.vg/Protocol#Open_Horse_Screen
	CPacketHurtAnimation                    // https://wiki.vg/Protocol#Hurt_Animation
	CPacketInitializeBorder                 // https://wiki.vg/Protocol#Initialize_World_Border
	CPacketKeepAlive                        // https://wiki.vg/Protocol#Keep_Alive
	CPacketChunkData                        // https://wiki.vg/Protocol#Chunk_Data_and_Update_Light
	CPacketWorldEvent                       // https://wiki.vg/Protocol#World_Event
	CPacketParticle                         // https://wiki.vg/Protocol#Particle_2
	CPacketUpdateLight                      // https://wiki.vg/Protocol#Update_Light
	CPacketLogin                            // https://wiki.vg/Protocol#Login_.28play.29
	CPacketMapData                          // https://wiki.vg/Protocol#Map_Data
	CPacketMerchantOffers                   // https://wiki.vg/Protocol#Merchant_Offers
	CPacketEntityPosition                   // https://wiki.vg/Protocol#Update_Entity_Position
	CPacketEntityPositionRotation           // https://wiki.vg/Protocol#Update_Entity_Position_and_Rotation
	CPacketEntityRotation                   // https://wiki.vg/Protocol#Update_Entity_Rotation
	CPacketVehicleMove                      // https://wiki.vg/Protocol#Move_Vehicle
	CPacketOpenBook                         // https://wiki.vg/Protocol#Open_Book
	CPacketOpenWindow                       // https://wiki.vg/Protocol#Open_Screen
	CPacketOpenSignEditor                   // https://wiki.vg/Protocol#Open_Sign_Editor
	CPacketPing                             // https://wiki.vg/Protocol#Ping_.28play.29
	CPacketGhostRecipe                      // https://wiki.vg/Protocol#Place_Ghost_Recipe
	CPacketPlayerAbilities                  // https://wiki.vg/Protocol#Player_Abilities
	CPacketChatMessage                      // https://wiki.vg/Protocol#Player_Chat_Message
	CPacketEndCombat                        // https://wiki.vg/Protocol#End_Combat
	CPacketEnterCombat                      // https://wiki.vg/Protocol#Enter_Combat
	CPacketCombatEvent                      // https://wiki.vg/Protocol#Combat_Death
	CPacketPlayerInfoRemove                 // https://wiki.vg/Protocol#Player_Info_Remove
	CPacketPlayerInfoUpdate                 // https://wiki.vg/Protocol#Player_Info_Update
	CPacketLookAt                           // https://wiki.vg/Protocol#Look_At
	CPacketPlayerPosition                   // https://wiki.vg/Protocol#Synchronize_Player_Position
	CPacketSetRecipeBook                    // https://wiki.vg/Protocol#Update_Recipe_Book
	CPacketRemoveEntities                   // https://wiki.vg/Protocol#Remove_Entities
	CPacketRemoveEntityEffect               // https://wiki.vg/Protocol#Remove_Entity_Effect
	CPacketResourcePack                     // https://wiki.vg/Protocol#Resource_Pack
	CPacketRespawn                          // https://wiki.vg/Protocol#Respawn
	CPacketSetHeadRotation                  // https://wiki.vg/Protocol#Set_Head_Rotation
	CPacketUpdateSection                    // https://wiki.vg/Protocol#Update_Section_Blocks
	CPacketSelectAdvancementTab             // https://wiki.vg/Protocol#Select_Advancements_Tab
	CPacketServerData                       // https://wiki.vg/Protocol#Server_Data
	CPacketSetActionBarText                 // https://wiki.vg/Protocol#Set_Action_Bar_Text
	CPacketSetBorderCenter                  // https://wiki.vg/Protocol#Set_Border_Center
	CPacketSetBorderLerpSize                // https://wiki.vg/Protocol#Set_Border_Lerp_Size
	CPacketSetBorderSize                    // https://wiki.vg/Protocol#Set_Border_Size
	CPacketSetBorderWarningDelay            // https://wiki.vg/Protocol#Set_Border_Warning_Delay
	CPacketSetBorderWarningDistance         // https://wiki.vg/Protocol#Set_Border_Warning_Distance
	CPacketSetCamera                        // https://wiki.vg/Protocol#Set_Camera
	CPacketSetHeldItem                      // https://wiki.vg/Protocol#Set_Held_Item
	CPacketSetCenterChunk                   // https://wiki.vg/Protocol#Set_Center_Chunk
	CPacketSetRenderDistance                // https://wiki.vg/Protocol#Set_Render_Distance
	CPacketSetSpawnPosition                 // https://wiki.vg/Protocol#Set_Default_Spawn_Position
	CPacketDisplayObjective                 // https://wiki.vg/Protocol#Display_Objective
	CPacketSetEntityMetadata                // https://wiki.vg/Protocol#Set_Entity_Metadata
	CPacketLinkEntities                     // https://wiki.vg/Protocol#Link_Entities
	CPacketSetEntityVelocity                // https://wiki.vg/Protocol#Set_Entity_Velocity
	CPacketSetEquipment                     // https://wiki.vg/Protocol#Set_Equipment
	CPacketSetExperience                    // https://wiki.vg/Protocol#Set_Experience
	CPacketSetHealth                        // https://wiki.vg/Protocol#Set_Health
	CPacketSetObjectives                    // https://wiki.vg/Protocol#Update_Objectives
	CPacketSetPassengers                    // https://wiki.vg/Protocol#Set_Passengers
	CPacketSetTeams                         // https://wiki.vg/Protocol#Update_Teams
	CPacketSetScore                         // https://wiki.vg/Protocol#Update_Score
	CPacketSetSimulationDistance            // https://wiki.vg/Protocol#Set_Simulation_Distance
	CPacketSetSubtitleText                  // https://wiki.vg/Protocol#Set_Subtitle_Text
	CPacketUpdateTime                       // https://wiki.vg/Protocol#Update_Time
	CPacketSetTitleText                     // https://wiki.vg/Protocol#Set_Title_Text
	CPacketSetTitlesAnimation               // https://wiki.vg/Protocol#Set_Title_Animation_Times
	CPacketEntitySoundEffect                // https://wiki.vg/Protocol#Entity_Sound_Effect
	CPacketSoundEffect                      // https://wiki.vg/Protocol#Sound_Effect
	CPacketStopSound                        // https://wiki.vg/Protocol#Stop_Sound
	CPacketSystemMessage                    // https://wiki.vg/Protocol#System_Chat_Message
	CPacketSetTabListHeaderAndFooter        // https://wiki.vg/Protocol#Set_Tab_List_Header_And_Footer
	CPacketTaqQueryResponse                 // https://wiki.vg/Protocol#Tag_Query_Response
	CPacketPickupItem                       // https://wiki.vg/Protocol#Pickup_Item
	CPacketTeleportEntity                   // https://wiki.vg/Protocol#Teleport_Entity
	CPacketSetAdvancements                  // https://wiki.vg/Protocol#Update_Advancements
	CPacketSetAttributes                    // https://wiki.vg/Protocol#Update_Attributes
	CPacketFeatureFlags                     // https://wiki.vg/Protocol#Feature_Flags
	CPacketEntityEffect                     // https://wiki.vg/Protocol#Entity_Effect
	CPacketSetRecipes                       // https://wiki.vg/Protocol#Update_Recipes
	CPacketSetTags                          // https://wiki.vg/Protocol#Update_Tags
)

// Game Serverbound
const (
	SPacketTeleportConfirm = iota
	SPacketQueryBlockEntityTag
	SPacketSetDifficulty
	SPacketMessageAcknowledge
	SPacketChatCommand
	SPacketChatMessage
	SPacketPlayerSession
	SPacketClientCommand
	SPacketClientSettings
	SPacketCommandSuggestion
	SPacketClickWindowButton
	SPacketClickWindow
	SPacketCloseWindow
	SPacketPluginMessage
	SPacketEditBook
	SPacketQueryEntityTag
	SPacketInteract
	SPacketJigsawGenerate
	SPacketKeepAlive
	SPacketLockDifficulty
	SPacketPlayerPosition
	SPacketPlayerPositionRotation
	SPacketPlayerRotation
	SPacketPlayerOnGround
	SPacketMoveVehicle
	SPacketPaddleBoat
	SPacketPickItem
	SPacketPlaceRecipe
	SPacketPlayerAbilities
	SPacketPlayerAction
	SPacketPlayerCommand
	SPacketPlayerInput
	SPacketPong
	SPacketChangeRecipeBookState
	SPacketSetSeenRecipe
	SPacketRenameItem
	SPacketResourcePack
	SPacketSelectTrade
	SPacketSetBeaconEffect
	SPacketSetHeldItem
	SPacketSetCommandBlock
	SPacketSetCommandMinecart
	SPacketSetCreativeModeSlot
	SPacketSetJigsawBlock
	SPacketSetStructureBlock
	SPacketSetSign
	SPacketSwingArm
	SPacketTeleportToEntity
	SPacketUseItemOn
	SPacketUseItem
)

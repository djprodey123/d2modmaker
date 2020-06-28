package skillsTxt

// File Constants
const (
	FileName   = "Skills.txt"
	NumColumns = 256
)

// Skill names
const (
	Teleport = "Teleport"
)

// Header Indexes
const (
	Skill             = 0
	Id                = 1
	CharClass         = 2
	SkillDesc         = 3
	SrvStFunc         = 4
	SrvDoFunc         = 5
	PrgStack          = 6
	SrvPrgFunc1       = 7
	SrvPrgFunc2       = 8
	SrvPrgFunc3       = 9
	PrgCalc1          = 10
	PrgCalc2          = 11
	PrgCalc3          = 12
	PrgDam            = 13
	SrvMissile        = 14
	DecQuant          = 15
	Lob               = 16
	SrvMissileA       = 17
	SrvMissileB       = 18
	SrvMissileC       = 19
	SrvOverlay        = 20
	AuraFilter        = 21
	AuraState         = 22
	AuraTargetState   = 23
	AuraLenCalc       = 24
	AuraRangeCalc     = 25
	AuraStat1         = 26
	AuraStatCalc1     = 27
	AuraStat2         = 28
	AuraStatcalc2     = 29
	AuraStat3         = 30
	AuraStatCalc3     = 31
	AuraStat4         = 32
	AuraStatCalc4     = 33
	AuraStat5         = 34
	AuraStatCalc5     = 35
	AuraStat6         = 36
	AuraStatCalc6     = 37
	AuraEvent1        = 38
	AuraEventFunc1    = 39
	AuraEvent2        = 40
	AuraEventFunc2    = 41
	AuraEvent3        = 42
	AuraEventFunc3    = 43
	AuraTgtEvent      = 44
	AuraTgtEventFunc  = 45
	AassiveState      = 46
	AassiveIType      = 47
	PassiveStat1      = 48
	PassiveCalc1      = 49
	PassiveStat2      = 50
	PassiveCalc2      = 51
	PassiveStat3      = 52
	PassiveCalc3      = 53
	PassiveStat4      = 54
	PassiveCalc4      = 55
	PassiveStat5      = 56
	PassiveCalc5      = 57
	PassiveEvent      = 58
	PassiveEventFunc  = 59
	Summon            = 60
	PetType           = 61
	PetMax            = 62
	SumMode           = 63
	SumSkill1         = 64
	SumSk1Calc        = 65
	SumSkill2         = 66
	SumSk2Calc        = 67
	SumSkill3         = 68
	SumSk3Calc        = 69
	SumSkill4         = 70
	SumSk4Calc        = 71
	SumSkill5         = 72
	SumSk5Calc        = 73
	SumUMod           = 74
	SumOverlay        = 75
	StSuccessOnly     = 76
	StSound           = 77
	StSoundClass      = 78
	StSoundDelay      = 79
	WeaponSnd         = 80
	DoSound           = 81
	DoSoundA          = 82
	DoSoundB          = 83
	TgtOverlay        = 84
	TgtSound          = 85
	PrgOverlay        = 86
	PrgSound          = 87
	CastOverlay       = 88
	CltOverlayA       = 89
	CltOverlayB       = 90
	CltStFunc         = 91
	CltDoFunc         = 92
	CltPrgFunc1       = 93
	CltPrgFunc2       = 94
	CltPrgFunc3       = 95
	CltMissile        = 96
	CltMissileA       = 97
	CltMissileB       = 98
	CltMissileC       = 99
	CltMissileD       = 100
	CltCalc1          = 101
	CltCalc1Desc      = 102
	CltCalc2          = 103
	CltCalc2Desc      = 104
	CltCalc3          = 105
	CltCalc3Desc      = 106
	Warp              = 107
	Immediate         = 108
	Enhanceble        = 109
	AttackRank        = 110
	NoAmmo            = 111
	Range             = 112
	WeapSel           = 113
	ITypeA1           = 114
	ITypeA2           = 115
	ITypeA3           = 116
	ETypeA1           = 117
	ETypeA2           = 118
	ITypeB1           = 119
	ITypeB2           = 120
	ITypeB3           = 121
	ETypeB1           = 122
	ETypeB2           = 123
	Anim              = 124
	SeqTrans          = 125
	MonAnim           = 126
	SeqNum            = 127
	SeqInput          = 128
	Durability        = 129
	UseAttackRate     = 130
	LineOfSight       = 131
	TargetableOnly    = 132
	SearchEnemyXY     = 133
	SearchEnemyNear   = 134
	SearchOpenXY      = 135
	SelectProc        = 136
	TargetCorpse      = 137
	TargetPet         = 138
	TargetAlly        = 139
	TargetItem        = 140
	AttackNoMana      = 141
	TgtPlaceCheck     = 142
	ItemEffect        = 143
	ItemCltEffect     = 144
	ItemTgtDo         = 145
	ItemTarget        = 146
	ItemCheckStart    = 147
	ItemCltCheckStart = 148
	ItemCastSound     = 149
	ItemCastOverlay   = 150
	SkPoints          = 151
	ReqLevel          = 152
	MaxLvl            = 153
	ReqStr            = 154
	ReqDex            = 155
	ReqInt            = 156
	ReqVit            = 157
	ReqSkill1         = 158
	ReqSkill2         = 159
	ReqSkill3         = 160
	Restrict          = 161
	State1            = 162
	State2            = 163
	State3            = 164
	Delay             = 165
	LeftSkill         = 166
	Repeat            = 167
	CheckFunc         = 168
	NoCostInState     = 169
	UseManaOnDo       = 170
	StartMana         = 171
	MinMana           = 172
	ManaShift         = 173
	Mana              = 174
	LvlMana           = 175
	Interrupt         = 176
	InTown            = 177
	Aura              = 178
	Periodic          = 179
	PerDelay          = 180
	Finishing         = 181
	Passive           = 182
	Progressive       = 183
	General           = 184
	Scroll            = 185
	Calc1             = 186
	Calc1Desc         = 187
	Calc2             = 188
	Calc2Desc         = 189
	Calc3             = 190
	Calc3Desc         = 191
	Calc4             = 192
	Calc4Desc         = 193
	Param1            = 194
	Param1Description = 195
	Param2            = 196
	Param2Description = 197
	Param3            = 198
	Param3Description = 199
	Param4            = 200
	Param4Description = 201
	Param5            = 202
	Param5Description = 203
	Param6            = 204
	Param6Description = 205
	Param7            = 206
	Param7Description = 207
	Param8            = 208
	Param8Description = 209
	InGame            = 210
	ToHit             = 211
	LevToHit          = 212
	ToHitCalc         = 213
	ResultFlags       = 214
	HitFlags          = 215
	HitClass          = 216
	Kick              = 217
	HitShift          = 218
	SrcDam            = 219
	MinDam            = 220
	MinLevDam1        = 221
	MinLevDam2        = 222
	MinLevDam3        = 223
	MinLevDam4        = 224
	MinLevDam5        = 225
	MaxDam            = 226
	MaxLevDam1        = 227
	MaxLevDam2        = 228
	MaxLevDam3        = 229
	MaxLevDam4        = 230
	MaxLevDam5        = 231
	DmgSymPerCalc     = 232
	EType             = 233
	EMin              = 234
	EMinLev1          = 235
	EMinLev2          = 236
	EMinLev3          = 237
	EMinLev4          = 238
	EMinLev5          = 239
	EMax              = 240
	EMaxLev1          = 241
	EMaxLev2          = 242
	EMaxLev3          = 243
	EMaxLev4          = 244
	EMaxLev5          = 245
	EDmgSymPerCalc    = 246
	ELen              = 247
	ELevLen1          = 248
	ELevLen2          = 249
	ELevLen3          = 250
	ELenSymPerCalc    = 251
	AiType            = 252
	AiBonus           = 253
	CostMult          = 254
	CostAdd           = 255
)

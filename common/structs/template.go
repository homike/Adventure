package structs

type HeroType uint8

const (
	HeroType_Hero    HeroType = 0 // 英雄
	HeroType_Player  HeroType = 1 // 玩家
	HeroType_Monster HeroType = 2 // 怪物
)

type HeroTemplate struct { // 英雄模板
	HeroName               string   `val:"名字"`
	SkillID                []int32  `val:"技能ID列表"`
	CombinationSpllID      int32    `val:"组合技能ID"`
	IconID                 int      `val:"模型ID"`
	QualityType            int      `val:"品质"`
	Profession             string   `val:"职业"`
	Ethnicity              string   `val:"种族"`
	Sex                    string   `val:"性别"`
	Description            string   `val:"描述"`
	BaseHP                 int      `val:"基础战力"`
	Coefficient            float32  `val:"成长系数"`
	HonorDebris            int      `val:"荣誉碎片"`
	AwakeCount             int      `val:"觉醒次数上限"`
	TempleAppearWeight     int      `val:"神殿出现概率权重"`
	EmployCostFragments    int      `val:"神殿兑换该英雄的花费"`
	EmployWeight           string   `val:"招募权重"`
	HeroType               HeroType `val:"英雄类型"`
	WeaponAdvance_XianGong int32    `val:"武器进阶先攻"`
	WeaponAdvance_FangYu   int32    `val:"武器进阶防御"`
	WeaponAdvance_ShanBi   int32    `val:"武器进阶闪避"`
	WeaponAdvance_WangZhe  int32    `val:"武器进阶王者"`
	// 招募权重，金币\混乱之门 301 \辉煌之门 302 \律动之门 303\万象之门（元宝）\传奇之门普通（元宝）\传奇之门特殊（元宝）
}

type GameLevelTemplate struct { // 游戏关卡
	Title              string  `val:"关卡名称"`
	EvnetIDs           []int32 `val:"事件ID队列"`
	GameBoxIDs         []int32 `val:"宝箱奖励ID队列"`
	GameBoxWeight      []int32 `val:"宝箱奖励权重队列"`
	SpeedGameBoxIDs    []int32 `val:"加速宝箱奖励ID队列"`
	SpeedGameBoxWeight []int32 `val:"加速宝箱奖励权重队列"`
	ActiveGameBoxSec   int32   `val:"宝箱需要时间"`
	MoneyPer           int32   `val:"每秒产出金币"`
	ExpPer             int32   `val:"每秒产出经验"`
	MinHP              int32   `val:"需要战力"`
	IconID             string  `val:"图标ID"`
}

const (
	GameLevelType_None    = iota // 普通事件
	GameLevelType_Fight   = iota // 战斗事件
	GameLevelType_Item    = iota // 物品事件
	GameLevelType_Res     = iota // 资源事件
	GameLevelType_NpcTack = iota // 对话事件
)

type GameLevelEventTemplate struct { // 游戏关卡的游戏事件
	Type           int32  `val:"Type"`
	ActiveEventSec int32  `val:"事件需要时间"`
	Title          string `val:"Title"`
	Context        string `val:"描述"`
	//PassContext    string  `val:"通关公告描述"`
	//MaxNotifyCount int32   `val:"最大通关通知人数"`
	HeroTemplateId int32   `val:"NPC模型"`
	HP             int32   `val:"事件需要时间"`
	CostFood       int32   `val:"饱足度消耗"`
	ItemId         int32   `val:"消耗物品"`
	ResId          int32   `val:"消耗资源"`
	Num            int32   `val:"消耗数量"`
	RewardIDs      []int32 `val:"奖励ID队列"`
}

type HeroLevelTemplate struct { // 英雄等级模板数据
	EXP1 int32 `val:"经验1"`
	EXP2 int32 `val:"经验2"`
	EXP3 int32 `val:"经验3"`
	EXP4 int32 `val:"经验4"`
	EXP5 int32 `val:"经验5"`
}

type UpgradeArtifactCost struct { // 神器升级消耗表
	//Level                 int    `val:"神器等级"`
	NeedResourceIdList    []int32 `val:"需要消耗的资源id列表"`
	NeedResourceCountList []int32 `val:"需要消耗的资源数量列表"`
	WeaponParam           int     `val:"神器系数"`
}

type UpgradeWeaponCost struct { // 武具级消耗表
	//Level                 int    `val:"武具等级"`
	NeedResourceIdList    []int32 `val:"需要消耗的资源id列表"`
	NeedResourceCountList []int32 `val:"需要消耗的资源数量列表"`
	WeaponParam           int     `val:"武具系数"`
	SuccessRate           int     `val:"成功率"`
	AddForgePoint         int     `val:"增加的锻造点"`
}

type AwakeCost struct { // 觉醒消耗表
	//Level                 int    `val:"等级"`
	Money  int     `val:"金币"`
	Statue int     `val:"雕像"`
	Param  float32 `val:"觉醒系数"`
}

type ItemTemplate struct { // 物品模板
	//ID                 int    `val:"ID"`
	Name   string `val:"名字"`
	Type   int    `val:"类型"`
	ExType int    `val:"扩展类型"`
	//SellMoney         int    `val:"出售价格"`
	IconID      string `val:"图标"`
	Description string `val:"描述"`
	//ShopID            int    `val:"商城id"`
	//ShopPrice         int    `val:"商城价格"`
	//Param1            int    `val:"参数1"`
	//Param2            int    `val:"参数2"`
	//Param3            int    `val:"参数3"`
	RewardIDs         []int32 `val:"奖励id"`
	OccurWeight       []int32 `val:"权重"`
	IsOnceEveryday    bool    `val:"是否每天只能使用一次"`
	RewardDescription string  `val:"奖励描述"`
}

type UnLockBagCost struct { // 背包格子解锁表
	//UnlockLevel string `val:"开启次数"`
	BagCount    int32   `val:"开启格子数"`
	CostResIDs  []int32 `val:"资源类型"`
	CostResNums []int32 `val:"资源数量"`
}

// Battlefield
type Battlefield struct {
	Name             string  `val:"名字"`
	HP               int32   `val:"战力"`
	NpcIDs           []int32 `val:"英雄ID列表"`
	XianGong         int32   `val:"先攻"`
	FangYu           int32   `val:"防御"`
	ShanBi           int32   `val:"闪避"`
	WangZhe          int32   `val:"王者"`
	BackgroundID     string  `val:"背景ID"`
	ForegroundID     string  `val:"前景ID"`
	RewardType       uint8   `val:"类型"` // 0: 固定, 1: 随机
	RewardIDs        []int32 `val:"奖励Id列表"`
	Weights          []int32 `val:"奖励权重列表"`
	LostWarDelayTime int32   `val:"冷却时间"`
	CostFood         int32   `val:"消耗的饱足度"`
}

type RewardTemplate struct {
	Type    int32  `val:"类型"` // RewardType_Property
	Context string `val:"奖励描述"`
	Param1  int32  `val:"子类型"`
	Param2  int32  `val:"数据1"`
	Param3  int32  `val:"数据2"`
}

type TriggerHPConditionsType uint8

const (
	_          TriggerHPConditionsType = iota
	SelfHPmin                          // 自身血量低于某值
	EnemyHPmin                         // 敌方血量低于某值
	SelfHPmax                          // 自身血量高于某值
	EnemyHPmax                         // 敌方血量高于某值
)

type TriggerRoundConditionsType uint8

const (
	_        TriggerRoundConditionsType = iota
	RoundMax                            // 回合数大于某值
	RoundMin                            // 回合数小于某值
)

type TriggerCountType uint8

const (
	_       TriggerCountType = iota
	One                      // 一次
	NoLimit                  // 无限制
)

type AttackEffectType uint8

const (
	AttackEffectType_None AttackEffectType = iota
	Hurt                                   // 伤害目标
	Recover                                // 恢复自身
	HurtAndRecover                         // 伤害目标，并且恢复自身
)

// 效果计算类型
type CalculateEffectType uint8

const (
	_               CalculateEffectType = iota
	SelfTotalFight                      // 自身总战力
	SelfLostFight                       // 自身已损失战力
	SelfLeftFight                       // 自身剩余战力
	EnemyTotalFight                     // 敌人总战力
	EnemyLostFight                      // 敌人已损失战力
	EnumyLeftFight                      // 敌人剩余战力
	ThisTimeAttack                      // 本次攻击
	NormalAttack                        // 普通攻击
)

type SpellTemplate struct {
	ID                          int32                      `val:"ID"`
	Name                        string                     `val:"名字"`
	Rate                        int32                      `val:"触发概率"`
	FightingCapacityTriggerType TriggerHPConditionsType    `val:"战力触发条件类型枚举"`
	FightingCapacity            int32                      `val:"战力触发值"`
	RoundTriggerType            TriggerRoundConditionsType `val:"回合触发条件类型枚举"`
	RoundValue                  int32                      `val:"回合触发值"`
	CountType                   TriggerCountType           `val:"触发次数类型枚举"`
	AttackType                  AttackEffectType           `val:"打击效果类型"`
	IgnoreDodge                 bool                       `val:"无视闪避"`
	IgnoreDefence               bool                       `val:"无视防御"`
	CalculateHurtType           CalculateEffectType        `val:"伤害效果类型枚举"`
	HurtEffectValue             int32                      `val:"伤害效果值"`
	CalculateRecoverType        CalculateEffectType        `val:"恢复效果计算类型枚举"`
	RecoverEffectValue          int32                      `val:"恢复效果值"`
	FirstProp                   int32                      `val:"先攻数"`
	DefenceProp                 int32                      `val:"防御数"`
	DodgeProp                   int32                      `val:"闪避数"`
	KingProp                    int32                      `val:"王者数"`
}

type ArtifactTemplate struct {
	ID          int32  `val:"ID"`
	SpellID     int32  `val:"技能ID"`
	Name        string `val:"名字"`
	Description string `val:"描述"`
	IconID      string `val:"图标"`
}

type ResourceTemplate struct {
	ID                int32   `val:"ID"`
	Name              string  `val:"名字"`
	IconID            string  `val:"图标"`
	Description       string  `val:"描述"`
	ShopID            int32   `val:"商城id"`
	UpgradeWeaponCost float32 `val:"锻造元宝价格"`
}

type CombinationSpell struct {
	ID          int32   `val:"ID"`
	Name        string  `val:"名称"`
	HeroList    []int32 `val:"英雄列表"`
	HeroNumList []int32 `val:"英雄数量列表"`
	SpellId     int32   `val:"关联的技能id"`
}

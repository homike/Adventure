package structs

type QualityType uint8 // 英雄品质类型
const (
	QualityType_None       QualityType = 0
	QualityType_White      QualityType = 0 // 白色
	QualityType_Green      QualityType = 1 // 绿色
	QualityType_Blue       QualityType = 2 // 蓝色
	QualityType_Purple     QualityType = 3 // 紫色
	QualityType_Gold       QualityType = 4 // 金色
	QualityType_SplashGold QualityType = 5 // 闪金色
)

type Hero struct {
	HeroID             int32       // 英雄id
	HeroTemplateID     int32       // 英雄的模板id
	Name               string      // 英雄的名字
	Level              int32       // 当前等级
	IsOutFight         bool        // 出战状态
	IsPlayer           bool        // 是否是玩家
	Exp                int32       // 当前的经验
	MaxExp             int32       // 最大经验
	Quality            QualityType // 品质类型
	AwakeCount         int32       // 觉醒次数
	WeaponLevel        int32       // 武具等级
	WeaponAdvanceLevel int32       // 武具进阶等级
	Index              int32       // 排序索引
	LevelHP            int32       // 因为升级而改变的HP  // 此字段的意义，待考虑
	ItemHP             int32       // 因为物品而改变的HP
	//HP             int32
}

func (h *Hero) HP() int32 {
	return h.LevelHP
}

// 冒险事件状态类型
type AdventureEventStatus uint8

const (
	AdventureEventStatus_UnActive AdventureEventStatus = iota
	AdventureEventStatus_Active
	AdventureEventStatus_Finish
)

type GameLevel struct {
	GameLevelID     int32                  // 关卡ID
	CompleteEvent   []AdventureEventStatus // 已经完成的事件
	EventProgress   int32                  // 事件的进度
	GameBoxProgress int32                  // 宝箱的进度
	BoxCount        int32                  // 宝箱的数量
	IsUnlock        bool                   // 是否解锁
	IsNew           bool                   // 是否是新开启关卡
}

type ArtifactStatusType uint8

const (
	_      ArtifactStatusType = iota
	Lock                      // 未解锁
	UnLock                    // 已解锁
	Use                       // 使用中
	New                       // 新解锁（客户端展示用）
)

type ArtifactStatus struct { // 神器的状态
	ArtifactID int32              // 神器的ID
	Status     ArtifactStatusType // 神器的状态
}

type ArtifactSealStatus struct { // 神器的封印状态
	SeaID    int32 // 封印ID
	IsUnlock bool  // 是否已解锁
	Level    int32 // 封印等级
	Count    int32 // 当前的进度
}

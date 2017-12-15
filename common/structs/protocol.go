package structs

const (
	Protocol_Test_Req                = 1
	Protocol_Test_Resp               = 2
	Protocol_GetSystemTime_Req       = 3
	Protocol_GetSystemTime_Resp      = 4
	Protocol_LoginServerResult_Ntf   = 1001
	Protocol_CreatePlayer_Req        = 1002 // 创建角色
	Protocol_CreatePlayer_Resp       = 1003
	Protocol_SyncLoginDataFinish_Ntf = 1006
	Protocol_LoginServerPlatform_Req = 1007
	Protocol_SyncPlayerBaseInfo_Ntf  = 1008
	Protocol_NameExists_Req          = 1009
	Protocol_NameExists_Resp         = 1010

	// 英雄相关的消息
	Protocol_Employ_Req         = 1100 // 雇佣英雄
	Protocol_Employ_Resp        = 1101
	Protocol_UnEmploy_Req       = 1102 // 解雇英雄
	Protocol_UnEmploy_Resp      = 1103
	Protocol_ResetHeroIndex_Req = 1104 // 调整英雄出站顺序
	Protocol_SyncHero_Ntf       = 1105 // 同步英雄信息
	Protocol_Work_Req           = 1106 // 英雄出战
	Protocol_SomeWork_Req       = 1107 // 一些英雄出战
	Protocol_Rest_Req           = 1108 // 英雄休息
	Protocol_SomeRest_Req       = 1109 // 一些英雄出战
	Protocol_Work_Resp          = 1110
	//Protocol_SomeWork_Resp          = 1111
	Protocol_Rest_Resp              = 1112
	Protocol_SomeRest_Resp          = 1113
	Protocol_Awake_Req              = 1114 // 英雄觉醒
	Protocol_Awake_Resp             = 1115
	Protocol_UpgradeWeapon_Req      = 1116 // 武具升级
	Protocol_UpgradeWeapon_Resp     = 1117
	Protocol_SyncEmploy_Req         = 1118 // 同步招募信息
	Protocol_SyncEmploy_Resq        = 1119
	Protocol_HeroHpAdd_Ntf          = 1120 // 英雄HP的变化
	Protocol_UnEmployManyHeros_Req  = 1121 // 解雇多名英雄
	Protocol_UnEmployManyHeros_Resp = 1122

	// 冒险相关
	Protocol_SyncGameLevel_Ntf        = 1201 // 同步关卡基础信息
	Protocol_SyncCurrentGameLevel_Ntf = 1202 // 同步当前关卡信息

	// 背包相关
	Protocol_UseItem_Req           = 1300 // 使用物品
	Protocol_UseItem_Resp          = 1301
	Protocol_SyncItem_Ntf          = 1302 // 同步物品
	Protocol_SyncBag_Ntf           = 1303 // 同步背包
	Protocol_SyncAllResouce_Ntf    = 1304 // 同步所有的资源
	Protocol_SyncResouce_Ntf       = 1305 // 同步资源
	Protocol_AddItem_Req           = 1306 // 加道具
	Protocol_AddResource_Req       = 1307 // 加资源
	Protocol_AddItem_Resp          = 1308
	Protocol_AddResource_Resp      = 1309
	Protocol_BagNotEnough_Ntf      = 1340
	Protocol_UnlockBag_Req         = 1341 // 开启背包格子
	Protocol_UnlockBag_Resp        = 1342
	Protocol_GetUsedGameItems_Req  = 1343 // 取得已使用过的物品列表
	Protocol_GetUsedGameItems_Resp = 1344

	// 角色相关
	Protocol_SyncStrength_Ntf           = 1404 // 同步饱足度
	Protocol_SyncWorkHeroTop_Ntf        = 1405 // 同步出站英雄上限
	Protocol_SyncUnlockMenus_Ntf        = 1412 // 同步已解锁菜单列表
	Protocol_SyncUserGuidRecords_Ntf    = 1413 // 同步新手引导数据
	Protocol_UpdateUserGuidRecord_Req   = 1414 // 更新玩家新手引导数据
	Protocol_SyncGameBoxTopNum_Ntf      = 1415 // 更新增加的宝箱上限数量
	Protocol_RewardResult_Ntf           = 1700 // 奖励物品接口
	Protocol_SetPlayerBarrageConfig_Req = 2801 // 弹幕设置

)

type GetSystemTimeReq struct {
}

type GetSystemTimeResp struct {
	Time int64
}

type LoginServerResultNtf struct {
	Result         int32 // 0: Success
	IsCreatePlayer bool
}

type CreatePlayerReq struct {
	PlayerName     string
	HeroTemplateId int32
}

type CreatePlayerResp struct {
	Result int32 // 0: Success
}

type SyncLoginDataFinishNtf struct {
}

type LoginServerPlatformReq struct {
	Takon     string
	Version   int32
	ChannelID string
}

type SyncPlayerBaseInfoNtf struct {
	PlayerID           int32
	GameZoonID         int32 // 游戏分区ID
	IsSupperMan        bool  // 是否是GM
	PlatformType       int32 // 平台类型
	Viplevel           int32
	TotalRechargeIngot int32
}

type NameExistsReq struct {
	Name string
}

// 如果存在，则返回一个新名字，如果和传入的名字一样，则说明没有重名
type NameExistsResp struct {
	Name string
}

type GuildRecord struct {
	UserGuidTypes uint8
	TriggerCount  int32
}

type UpdateUserGuidRecordReq struct {
	UserGuidTypes uint8
}

type SyncUserGuidRecordsNtf struct {
	Records []GuildRecord
}

type SyncStrengthNtf struct {
	Strength int32
}

type SyncHeroWorkTopNtf struct {
	MaxWorker int32
}

type SyncUnlockMenusNtf struct {
	MenuStates []*MenuStatusItem
}

type SyncGameBoxTopNumNtf struct {
	AddNum int32
}

type SyncHeroNtf struct {
	SyncHeroType uint8
	Heros        []*Hero
}

type EmployReq struct {
	EmployType uint8
}

const (
	EmployRet_Success = iota
	EmployRet_Failed
	EmployRet_NotEnough
)

type EmployResp struct {
	Ret     uint8
	HeroIDs []int32
}

type UnEmployReq struct {
	HeroID int32
}
type UnEmployResp struct {
	Ret    uint8
	HeroID int32
}

type RewardResultNtf struct {
	IsRes   bool
	Rewards []Reward
	Context string
}

type ResetHeroIndexReq struct {
	HeroIDs []int32
}

type WorkReq struct {
	HeroID int32
}

type WorkResp struct {
	Ret    uint8
	HeroID int32
}

type SomeWorkReq struct {
	HeroIDs []int32
}

type ResetResp struct {
	Ret    uint8
	HeroID int32
}

type ResetReq struct {
	HeroID int32
}

type SomeResetReq struct {
	HeroIDs []int32
}

type AwakeReq struct {
	HeroID int32
}

type AwakeResp struct {
	Ret    uint8
	HeroID int32
	AddHP  int32
}

type UpgradeWeaponReq struct {
	HeroID int32
	Ingot  int32
}

type UpgradeWeaponResp struct {
	Ret    uint8
	HeroID int32
	AddHP  int32
}

type SyncAllResouceNtf struct {
	Money         int32
	Ingot         int32
	Fragment      int32
	Statue        int32
	Strength      int32
	Detonator     int32
	MiningToolkit int32
	Ors           []IDNUM
	Foods         []IDNUM
	Badges        []IDNUM
	UnlockResIds  []int32
}

type SyncResourceNtf struct {
	ResID  int32
	ResNum int32
}

type SyncBagNtf struct {
	MaxCount    int32
	UnlockLevel int32
	Items       []*GameItem
}

type GetUsedGameItemsResp struct {
	ItemIDs   []int32
	UserTimes []int64
}

type UseItemReq struct {
	ItemID int32
}

type UseItemResp struct {
	Ret uint8
}

type UnlockBagResp struct {
	Ret         uint8
	MaxCount    int32
	UnLockLevel int32
}

type SyncGameLevelNtf struct {
	GameLevels      []*GameLevel
	CurLevelID      int32
	LastRefreshTime int64
	SpeedCount      int32 // 加速冒险的次数
}

type SyncCurrentGameLevelNtf struct {
	GameLevel *GameLevel
}

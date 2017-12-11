package structs

const (
	Protocol_Test_Req                = 1
	Protocol_Test_Resp               = 2
	Protocol_GetSystemTime_Req       = 3
	Protocol_GetSystemTime_Resp      = 4
	Protocol_LoginServerResult_Ntf   = 1001
	Protocol_CreatePlayer_Req        = 1002
	Protocol_CreatePlayer_Resp       = 1003
	Protocol_SyncLoginDataFinish_Ntf = 1006
	Protocol_LoginServerPlatform_Req = 1007
	Protocol_SyncPlayerBaseInfo_Ntf  = 1008
	Protocol_NameExists_Req          = 1009
	Protocol_NameExists_Resp         = 1010

	// 英雄相关的消息
	Protocol_Employ_Req             = 1100 // 雇佣英雄
	Protocol_Employ_Resp            = 1101
	Protocol_UnEmploy_Req           = 1102 // 解雇英雄
	Protocol_UnEmploy_Resp          = 1103
	Protocol_ResetHeroIndex_Req     = 1104 // 调整英雄出站顺序
	Protocol_SyncHero_Ntf           = 1105 // 同步英雄信息
	Protocol_Work_Req               = 1106 // 英雄出战
	Protocol_SomeWork_Req           = 1107 // 一些英雄出战
	Protocol_Rest_Req               = 1108 // 英雄休息
	Protocol_SomeRest_Req           = 1109 // 一些英雄出战
	Protocol_Work_Resp              = 1110
	Protocol_SomeWork_Resp          = 1111
	Protocol_Rest_Resp              = 1112
	Protocol_SomeRest_Resp          = 1113
	Protocol_Awake_Rep              = 1114 // 英雄觉醒
	Protocol_Awake_Resp             = 1115
	Protocol_UpgradeWeapon_Rep      = 1116 // 武具升级
	Protocol_UpgradeWeapon_Resp     = 1117
	Protocol_SyncEmploy_Req         = 1118 // 同步招募信息
	Protocol_SyncEmploy_Resq        = 1119
	Protocol_HeroHpAdd_Ntf          = 1120 // 英雄HP的变化
	Protocol_UnEmployManyHeros_Req  = 1121 // 解雇多名英雄
	Protocol_UnEmployManyHeros_Resq = 1122

	// 角色相关
	Protocol_SyncUserGuidRecords_Ntf    = 1413 // 同步新手引导数据
	Protocol_UpdateUserGuidRecord_Req   = 1414 // 更新玩家新手引导数据
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
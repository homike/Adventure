package msghandler

import (
	"adventure/advserver/gamedata"
	"adventure/advserver/sessions"
	"time"

	"adventure/common/structs"
)

func InitMessageBag() {
	message := map[uint16]ProcessFunc{
		uint16(structs.Protocol_UseItem_Req):     UseItemReq,   // 使用物品
		uint16(structs.Protocol_AddItem_Req):     TestReq,      // 加道具
		uint16(structs.Protocol_AddResource_Req): TestReq,      // 加资源
		uint16(structs.Protocol_UnlockBag_Req):   UnlockBagReq, // 开启背包格子
	}

	for k, v := range message {
		MapFunc[k] = v
	}
}

func UseItemReq(sess *sessions.Session, msgBody []byte) {
	logger.Debug("UseItemReq")

	req := &structs.UseItemReq{}
	sess.UnMarshal(msgBody, req)

	resp := &structs.UseItemResp{
		Ret: structs.AdventureRet_Failed,
	}

	userItem, err := sess.PlayerData.Bag.GetItem(req.ItemID)
	if err != nil {
		logger.Error("GetItem(%v) Error(%v)", req.ItemID, err)
		sess.Send(structs.Protocol_UseItem_Resp, resp)
		return
	}

	isOnceEveryDay := gamedata.AllTemplates.ItemTemplates[userItem.TemplateID].IsOnceEveryday

	userUsedItem := &structs.UsedGameItem{}
	if isOnceEveryDay {
		userUsedItem, err = sess.PlayerData.Bag.GetUsedItem(userItem.TemplateID)
		if err != nil {
			userUsedItem = &structs.UsedGameItem{
				TemplateID:  userItem.TemplateID,
				LastUseTime: time.Now().AddDate(0, 0, -1).Unix(),
			}

			sess.PlayerData.Bag.AddUsedItem(userUsedItem)
		}
		if userUsedItem.LastUseTime > time.Now().Unix() {
			logger.Error("item(%v) can user once oneday Error(%v)", req.ItemID)
			sess.Send(structs.Protocol_UseItem_Resp, resp)
			return
		}
	}

	err = sess.DoSomeReward(userItem.TemplateID, userItem.Num)
	if err != nil {
		logger.Error("item(%v) DoSomeRewards Error(%v)", req.ItemID, err)
		sess.Send(structs.Protocol_UseItem_Resp, resp)
		return
	}

	// 更新使用信息
	if userUsedItem.TemplateID != 0 {
		userUsedItem.LastUseTime = time.Now().Unix()
		sess.SyncPlayerUsedItem()
	}
	// 删除物品
	sess.PlayerData.Bag.RemoveItem(req.ItemID)
	// 返回消息
	resp.Ret = structs.AdventureRet_Success
	sess.Send(structs.Protocol_UseItem_Resp, resp)
}

func AddItemReq(sess *sessions.Session, msgBody []byte) {
}

func AddResourceReq(sess *sessions.Session, msgBody []byte) {
}

func UnlockBagReq(sess *sessions.Session, msgBody []byte) {

	resp := &structs.UnlockBagResp{
		Ret: structs.AdventureRet_Failed,
	}

	/////////////////////////////////////////////Data Check////////////////////////////////////////
	unclockCnt := int32(len(gamedata.AllTemplates.UnLockBagCosts))

	if sess.PlayerData.Bag.UnlockLevel >= unclockCnt {
		logger.Error("unlock already max")
		sess.Send(structs.Protocol_UnlockBag_Resp, resp)
		return
	}

	bagCnt := gamedata.AllTemplates.UnLockBagCosts[sess.PlayerData.Bag.UnlockLevel+1].BagCount

	unlockCostT, ok := gamedata.AllTemplates.UnLockBagCosts[sess.PlayerData.Bag.UnlockLevel+1]
	if !ok {
		logger.Error("UnLockBagCosts[%v] failed", sess.PlayerData.Bag.UnlockLevel+1)
		sess.Send(structs.Protocol_UnlockBag_Resp, resp)
		return
	}
	costResIDs, costResNums := unlockCostT.CostResIDs, unlockCostT.CostResNums

	///////////////////////////////////////////Logic Process///////////////////////////////////////
	// 扣除资源
	for k, _ := range costResIDs {
		sess.ResourceChange(costResIDs[k], -costResNums[k], structs.ResouceChangeType_Employ_Money)
	}

	sess.PlayerData.Bag.UnlockLevel++
	sess.PlayerData.Bag.MaxCount += int32(bagCnt)

	// 返回消息
	resp.MaxCount = sess.PlayerData.Bag.MaxCount
	resp.UnLockLevel = sess.PlayerData.Bag.UnlockLevel
	sess.Send(structs.Protocol_UnlockBag_Resp, resp)
}

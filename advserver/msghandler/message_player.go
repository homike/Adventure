package msghandler

import (
	"adventure/advserver/model"
	"adventure/advserver/sessions"
	"adventure/common/structs"
	"fmt"
)

// 1002
func CreatePlayer(sess *sessions.Session, msgBody []byte) {
	//fmt.Println("CreatePlayer data:", msgBody)

	req := structs.CreatePlayerReq{}
	sess.UnMarshal(msgBody, &req)

	//fmt.Println("CreatePlayer name: ", req.PlayerName, "heroTemplateID: ", req.HeroTemplateId)
	resp := &structs.CreatePlayerResp{
		Result: 0, // Success
	}

	/////////////////////////////////////////////Data Check////////////////////////////////////////
	///////////////////////////////////////////Logic Process///////////////////////////////////////
	player, err := model.NewPlayer(req.PlayerName, req.HeroTemplateId)
	if err != nil {
		fmt.Println("NewPlayer Error", err)
		resp.Result = 1
		sess.Send(structs.Protocol_CreatePlayer_Resp, resp)
		return
	}
	sess.SetPlayer(player)
	sessions.SessionMgr.AddSession(sess)

	sess.Send(structs.Protocol_CreatePlayer_Resp, resp)

	SyncLoginDataFinish(sess)

	sess.OnEnterGame()
}

// 1009
func NameExists(sess *sessions.Session, msgBody []byte) {
	fmt.Println("czx@@@ NameExists1:", string(msgBody))

	req := structs.NameExistsReq{}
	sess.UnMarshal(msgBody, &req)

	resp := &structs.NameExistsResp{
		Name: req.Name,
	}
	sess.Send(structs.Protocol_NameExists_Resp, resp)
}

func UpdateUserGuidRecord(sess *sessions.Session, msgBody []byte) {
	fmt.Println("czx@@@ UpdateUserGuidRecord:", string(msgBody))

	req := structs.UpdateUserGuidRecordReq{}
	sess.UnMarshal(msgBody, &req)

	sess.PlayerData.UpdateGuidRecords(req.UserGuidTypes)
}

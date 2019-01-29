// Conversation 会话

package sdk

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/httplib"
)

// ConversationType 会话类型
type ConversationType int

const (
	// PRIVATE 单聊
	PRIVATE ConversationType = 1
	// DISCUSSION 讨论组
	DISCUSSION ConversationType = 2
	// GROUP 群聊
	GROUP ConversationType = 3
	// SYSTEM 系统
	SYSTEM ConversationType = 4
	// CUSTOMERSERVICE 客服
	CUSTOMERSERVICE ConversationType = 5
)

// ConversationMute 设置用户某个会话屏蔽 Push
/*
*@param  conversationType:会话类型 PRIVATE、GROUP、DISCUSSION、SYSTEM。
*@param  userID:设置用户 ID。
*@param  targetID:需要屏蔽的目标 ID。
*
*@return error
 */
func (rc *RongCloud) ConversationMute(conversationType ConversationType, userID, targetID string) error {

	if conversationType == 0 {
		return RCErrorNew(1002, "Paramer 'userId' is required")
	}

	if userID == "" {
		return RCErrorNew(1002, "Paramer 'userID' is required")
	}

	if targetID == "" {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/conversation/notification/set." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("requestId", userID)
	req.Param("conversationType", fmt.Sprintf("%v", conversationType))
	req.Param("targetId", targetID)
	req.Param("isMuted", "1")

	rep, err := req.Bytes()
	if err != nil {
		rc.URLError(err)
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}

// ConversationUnmute 设置用户某个会话接收 Push
/*
*@param  conversationType:会话类型 PRIVATE、GROUP、DISCUSSION、SYSTEM。
*@param  userID:设置用户 ID。
*@param  targetID:需要屏蔽的目标 ID。
*
*@return error
 */
func (rc *RongCloud) ConversationUnmute(conversationType ConversationType, userID, targetID string) error {
	if conversationType == 0 {
		return RCErrorNew(1002, "Paramer 'conversationType' is required")
	}

	if userID == "" {
		return RCErrorNew(1002, "Paramer 'userID' is required")
	}

	if targetID == "" {
		return RCErrorNew(1002, "Paramer 'targetID' is required")
	}

	req := httplib.Post(rc.RongCloudURI + "/conversation/notification/set." + ReqType)
	req.SetTimeout(time.Second*rc.TimeOut, time.Second*rc.TimeOut)
	rc.FillHeader(req)
	req.Param("requestId", userID)
	req.Param("conversationType", fmt.Sprintf("%v", conversationType))
	req.Param("targetId", targetID)
	req.Param("isMuted", "0")

	rep, err := req.Bytes()
	if err != nil {
		rc.URLError(err)
		return err
	}
	var code CodeResult
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return RCErrorNew(code.Code, code.ErrorMessage)
	}
	return nil
}
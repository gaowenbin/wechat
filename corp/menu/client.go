// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

import (
	"net/http"
	"strconv"

	"github.com/chanxuehong/wechat/corp"
)

type Client struct {
	corp.CorpClient
}

// 创建一个新的 Client.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewClient(TokenServer corp.TokenServer, HttpClient *http.Client) *Client {
	if TokenServer == nil {
		panic("TokenServer == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &Client{
		CorpClient: corp.CorpClient{
			TokenServer: TokenServer,
			HttpClient:  HttpClient,
		},
	}
}

// 创建自定义菜单.
func (clt *Client) CreateMenu(agentId int64, menu Menu) (err error) {
	var result corp.Error

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/menu/create?agentid=" +
		strconv.FormatInt(agentId, 10) + "&access_token="
	if err = clt.PostJSON(incompleteURL, menu, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 删除自定义菜单
func (clt *Client) DeleteMenu(agentId int64) (err error) {
	var result corp.Error

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/menu/delete?agentid=" +
		strconv.FormatInt(agentId, 10) + "&access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 获取自定义菜单
func (clt *Client) GetMenu(agentId int64) (menu Menu, err error) {
	var result struct {
		corp.Error
		Menu Menu `json:"menu"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/menu/get?agentid=" +
		strconv.FormatInt(agentId, 10) + "&access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = result.Menu
	return
}

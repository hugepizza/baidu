package baidu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type BaiduFaceUser struct {
	GroupId  string  `json:"group_id"`
	UserId   string  `json:"user_id"`
	UserInfo string  `json:"user_info"`
	Score    float64 `json:"score"`
}

type SearchResponse struct {
	BaiduResponse `json",inline"`
	Result        struct {
		FaceToken string          `json:"face_token"`
		UserList  []BaiduFaceUser `json:"user_list"`
	} `json:"result"`
}

// SearchFace
func (b *Baidu) SearchFace(imgurl string) (*SearchResponse, error) {
	tkn, err := b.GetAccessToken()
	if err != nil {
		return nil, err
	}

	var req struct {
		Image       string `json:"image"`
		ImageType   string `json:"image_type"`
		GroupIdList string `json:"group_id_list"`
	}
	req.Image = imgurl
	req.ImageType = "URL"
	req.GroupIdList = "test"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(req)
	u := b.path("/rest/2.0/face/v3/search")
	q := u.Query()
	q.Set("access_token", tkn.AccessToken)
	u.RawQuery = q.Encode()
	resp, err := http.Post(u.String(), "application/json", buf)
	if err != nil {
		logrus.Errorf("searched failed, %s", err)
		return nil, err
	}

	ret := &SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		logrus.Errorf("decode search response failed, %s", err)
		return nil, err
	}

	if ret.ErrorCode > 0 {
		return nil, fmt.Errorf("search failed, (%v) %s", ret.ErrorCode, ret.ErrorMsg)
	}

	return ret, nil
}

package body

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/funxdata/baidu/core"
	"github.com/sirupsen/logrus"
)

const (
	urlBodyTracking = "/rest/2.0/image-classify/v1/body_tracking"
)

type BaiduBody struct {
	*core.Core
}
type Location struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Info struct {
	Location Location `json:"location"`
	ID       int      `json:"ID"`
}
type Count struct {
	In  int `json:"in"`
	Out int `json:"out"`
}
type TrackResult struct {
	PersonNum   int    `json:"person_num"`
	PersonInfo  []Info `json:"person_info"`
	PersonCount Count  `json:"person_count"`
	Image       string `json:"image"`
}

func (b *BaiduBody) Track(imgurl string) (*TrackResult, error) {
	img, err := b.encodeImg(imgurl)
	if err != nil {
		return nil, err
	}
	req := url.Values{
		"dynamic":   []string{"false"},
		"case_id":   []string{strconv.FormatInt(time.Now().Unix(), 10)},
		"case_init": []string{"false"},
		"image":     []string{img},
		//"dynamic":   []string{"true"},
		// "area":      []string{"0", "0", "0", "0", "0", "0", "0", "0"},

	}
	var ret struct {
		core.BaiduResponse
		TrackResult
	}
	// var ret TrackResult
	if err := b.PostForm(urlBodyTracking, req, &ret); err != nil {
		logrus.Errorf("track failed, %s", err)
		return nil, err
	}
	if ret.ErrorCode > 0 {
		return nil, fmt.Errorf("track failed, (%v) %s", ret.ErrorCode, ret.ErrorMsg)
	}

	return &ret.TrackResult, nil
}
func (b *BaiduBody) encodeImg(imgURL string) (string, error) {
	file, err := os.Open(imgURL)
	if err != nil {
		return "", err
	}
	defer file.Close()
	bts, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	ret := base64.StdEncoding.EncodeToString(bts)
	return ret, nil
}

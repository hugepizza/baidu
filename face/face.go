package face

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"strings"

	"github.com/funxdata/baidu/core"
	"github.com/funxdata/baidu/internal/biz"
	pb "github.com/funxdata/baidu/service"
	"github.com/sirupsen/logrus"
)

type BaiduFace struct {
	*core.Core
}

// SearchFace
func (b *BaiduFace) Search(imgurl string, groups ...string) (*pb.SearchResult, error) {
	req := map[string]interface{}{
		"image":      imgurl,
		"image_type": "FACE_TOKEN",
		// "image_type":    "URL",
		"group_id_list": strings.Join(groups, ","),
	}
	var ret struct {
		core.BaiduResponse `json",inline"`
		pb.SearchResult    `json:"result"`
	}
	if err := b.Post(biz.URLFaceSearch, req, &ret); err != nil {
		logrus.Errorf("searched failed, %s", err)
		return nil, err
	}

	if err := ret.Err("search"); err != nil {
		return nil, err
	}

	return &ret.SearchResult, nil
}

func (b *BaiduFace) Detect(imgurl string) (*pb.DetectResult, error) {
	req := map[string]interface{}{
		"image":        imgurl,
		"image_type":   "URL",
		"face_field":   "age,beauty,expression,face_shape,gender,glasses,landmark,race,quality,face_type",
		"max_face_num": 10,
		// "face_type":    "",
	}
	var ret struct {
		core.BaiduResponse `json",inline"`
		pb.DetectResult    `json:"result"`
	}
	err := b.Post(biz.URLFaceDetect, req, &ret)
	if err != nil {
		logrus.Errorf("detect face %s failed, %s", err)
		return nil, err
	}

	if ret.ErrorCode > 0 {
		return nil, fmt.Errorf("detect failed, (%v) %s", ret.ErrorCode, ret.ErrorMsg)
	}

	return &ret.DetectResult, nil
}

// img
func (b *BaiduFace) img(obj interface{}) map[string]interface{} {
	if imgurl, ok := obj.(string); ok {
		return map[string]interface{}{
			"image":      imgurl,
			"image_type": "URL",
		}
	}

	if img, ok := obj.(image.Image); ok {
		buf, out := new(bytes.Buffer), new(bytes.Buffer)
		jpeg.Encode(buf, img, nil)
		base64.NewEncoder(base64.StdEncoding, out)
		return map[string]interface{}{
			"image":      out.String(),
			"image_type": "FACE_TOKEN",
		}
	}
	return nil
}

package baidu

import (
	"github.com/funxdata/baidu/core"
	"github.com/funxdata/baidu/face"
)

type Baidu struct {
	*core.Core
}

func New(apiKey, apiSecret string) *Baidu {
	return &Baidu{core.NewCore(apiKey, apiSecret)}
}

// Face
func (b *Baidu) Face() *face.BaiduFace {
	return &face.BaiduFace{b.Core}
}

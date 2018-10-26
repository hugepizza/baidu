package body

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/funxdata/baidu/core"
)

func Test_body(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	tb := &BaiduBody{Core: core.NewCore("", "")}
	resp, err := tb.Track("/home/wanglei/Pictures/test2.jpeg")
	if err != nil {
		logrus.Info(err)
		return
	}
	logrus.Info(resp.PersonNum)
}

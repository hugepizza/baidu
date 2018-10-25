package body

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/funxdata/baidu/core"
)

func Test_body(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	tb := &BaiduBody{Core: core.NewCore("D1RUvOGwx9fy0Gd6wIGS1Rs8", "LifdRZPrEW2nNgxOdbsS8Qg6MTt0Ga84")}
	resp, err := tb.Track("/home/wanglei/Pictures/test2.jpeg")
	if err != nil {
		logrus.Info(err)
		return
	}
	logrus.Info(resp.PersonCount)
}

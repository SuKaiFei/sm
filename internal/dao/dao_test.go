// 作者 sukaifei
// 日期 2019/09/11

package dao

import (
	"flag"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"os"
	"testing"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	log.Init(nil)
	d = New()
	m.Run()
	os.Exit(0)
}

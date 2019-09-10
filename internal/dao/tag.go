// 作者 sukaifei
// 日期 2019/09/07

package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/didi/gendry/scanner"
	"sm/internal/model"
)

const (
	_selTagAll = `SELECT name,code,sort FROM tag order by sort asc`
)

func (d *Dao) GetTags(ctx context.Context) (tags []*model.Tags, err error) {
	rows, err := d.db.Query(ctx, _selTagAll)
	if err != nil {
		log.Error("dao.GetTags db.Query error(%v)", err)
		return
	}
	defer rows.Close()

	err = scanner.Scan(rows, &tags)
	if err != nil {
		log.Error("dao.GetTags scanner.Scan error(%v)", err)
		return
	}
	return
}

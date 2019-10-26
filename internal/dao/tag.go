// 作者 sukaifei
// 日期 2019/09/07

package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/didi/gendry/scanner"
	"sm/internal/model"
	"time"
)

const (
	_selTagAll    = `SELECT name,code,sort FROM tag order by sort asc`
	_selTagOne    = `SELECT * FROM tag WHERE id = ?`
	_insTagAdd    = `INSERT INTO tag(name,code,sort,create_time) VALUES(?,?,?,?)`
	_delTagDelete = `DELETE FROM tag WHERE id = ?`
	_updTagUpdate = `UPDATE tag(name,code,sort,create_time) VALUES(?,?,?,?)`
)

func (d *Dao) GetTags(ctx context.Context) (tags []*model.Tags, err error) {
	rows, err := d.db.Query(ctx, _selTagAll)
	if err != nil {
		log.Error("dao.GetTags db.Query error(%v)", err)
		return
	}
	defer rows.Close()

	err = scanner.Scan(rows, &tags)
	return
}

func (d *Dao) AddTag(ctx context.Context, res *model.Tag) (id int64, err error) {
	result, err := d.db.Exec(ctx, _insTagAdd, res.Name, res.Code, res.Sort, time.Now().UnixNano()/1e6)
	if err != nil {
		log.Error("dao.AddTag db.Exec error(%v)", err)
		return
	}
	return result.LastInsertId()
}

func (d *Dao) DeleteTag(ctx context.Context, id int64) (err error) {
	result, err := d.db.Exec(ctx, _delTagDelete, id)
	if err != nil {
		log.Error("dao.DeleteTag db.Exec error(%v)", err)
		return
	}
	_, err = result.RowsAffected()
	return
}

func (d *Dao) UpdateTag(ctx context.Context, res *model.Tag) (id int64, err error) {
	result, err := d.db.Exec(ctx, _updTagUpdate, res.Name, res.Code, res.Sort, res.CreateTime)
	if err != nil {
		log.Error("dao.UpdateTag db.Exec error(%v)", err)
		return
	}
	return result.LastInsertId()
}

func (d *Dao) GetTag(ctx context.Context, id int64) (tag *model.Tag, err error) {
	rows, err := d.db.Query(ctx, _selTagOne, id)
	if err != nil {
		log.Error("dao.GetTag db.Query error(%v)", err)
		return
	}
	defer rows.Close()
	err = scanner.Scan(rows, &tag)
	return
}

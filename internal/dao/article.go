// 作者 sukaifei
// 日期 2019/09/07

package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/didi/gendry/scanner"
	"sm/internal/model"
	"time"
)

const (
	_selArticleOne     = `SELECT a.id,a.html_text,a.md_text,a.title,u.avatar,u.nickname author,t.name tag,a.create_time,a.pv FROM article a,user u,tag t WHERE a.id=? AND a.user_id=u.id AND a.tag_code=t.CODE`
	_selArticleAll     = `SELECT a.id,a.title,u.avatar,u.nickname author,t.name tag,a.create_time,a.pv FROM article a,user u,tag t WHERE a.user_id=u.id AND a.tag_code=t.CODE ORDER BY a.create_time desc,a.pv DESC LIMIT ? OFFSET ?`
	_countArticleAll   = `SELECT count(*) FROM article a,user u,tag t WHERE a.user_id=u.id AND a.tag_code=t.CODE`
	_selArticleByUid   = `SELECT a.id,a.title,u.avatar,u.nickname author,t.name tag,a.create_time,a.pv FROM article a,user u,tag t WHERE a.user_id=u.id AND a.tag_code=t.CODE AND u.email = ? ORDER BY a.create_time desc,a.pv DESC LIMIT ? OFFSET ?`
	_selArticleByTag   = `SELECT a.id,a.title,u.avatar,u.nickname author,t.name tag,a.create_time,a.pv FROM article a,user u,tag t WHERE a.user_id=u.id AND a.tag_code=t.CODE AND a.tag_code = ? ORDER BY a.create_time desc,a.pv DESC LIMIT ? OFFSET ?`
	_countArticleByUid = `SELECT count(*) FROM article a,user u,tag t WHERE a.user_id=u.id AND a.tag_code=t.CODE AND u.email = ?`
	_countArticleByTag = `SELECT count(*) FROM article a,user u,tag t WHERE a.user_id=u.id AND a.tag_code=t.CODE AND a.tag_code = ?`
	_insArticleAdd     = `INSERT INTO article(user_id, tag_code, title, html_text,md_text, pv, top, create_time, update_time) VALUES (?,?,?,?,?,?,?,?,?)`
	_delArticleDelete  = `DELETE FROM tag WHERE id = ?`
	_updArticleUpdate  = `UPDATE tag(name,code,sort,create_time) VALUES(?,?,?,?)`
)

func (d *Dao) AddArticle(ctx context.Context, req *model.ArticleAdd) (id int64, err error) {
	result, err := d.db.Exec(ctx, _insArticleAdd, req.UserId, req.TagCode, req.Title, req.HtmlText, req.MdText, 0, model.NotTop, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		log.Error("dao.AddArticle db.Exec error(%v)", err)
		return
	}
	id, err = result.LastInsertId()
	return
}

func (d *Dao) GetArticle(ctx context.Context, id int64) (res *model.ArticleInfo, err error) {
	result, err := d.db.Query(ctx, _selArticleOne, id)
	if err != nil {
		log.Error("dao.GetArticle db.Query error(%v)", err)
		return
	}
	err = scanner.Scan(result, &res)
	return
}

func (d *Dao) GetMyArticles(ctx context.Context, acc string, pn, ps int32) (articles *model.ArticleList, err error) {
	page := new(model.Page)
	var rows *sql.Rows

	if acc != "" {
		countRow := d.db.QueryRow(ctx, _countArticleByUid, acc)
		if err = countRow.Scan(&page.Total); err != nil {
			log.Error("[GetMyArticles]row.ScanCount error(%v)", err)
			return
		}
		rows, err = d.db.Query(ctx, _selArticleByUid, acc, ps, (pn-1)*ps)
		if err != nil {
			log.Error("[GetMyArticles]db.Query(%s,%d,%d) error(%v)", _selArticleByUid, acc, page.Num*page.Size, page.Size, err)
			return
		}
	} else {
		countRow := d.db.QueryRow(ctx, _countArticleAll)
		if err = countRow.Scan(&page.Total); err != nil {
			log.Error("[GetMyArticles]row.ScanCount error(%v)", err)
			return
		}
		rows, err = d.db.Query(ctx, _selArticleAll, ps, (pn-1)*ps)
		if err != nil {
			log.Error("[GetMyArticles]db.Query(%s,%d,%d) error(%v)", _selArticleAll, page.Num*page.Size, page.Size, err)
			return
		}
	}

	defer rows.Close()
	articles = &model.ArticleList{Page: page}
	err = scanner.Scan(rows, &articles.Result)
	return
}

func (d *Dao) GetArticles(ctx context.Context, tag string, pn, ps int32) (articles *model.ArticleList, err error) {
	page := new(model.Page)
	var rows *sql.Rows

	if tag != "" {
		countRow := d.db.QueryRow(ctx, _countArticleByTag, tag)
		if err = countRow.Scan(&page.Total); err != nil {
			log.Error("[GetArticles]row.ScanCount error(%v)", err)
			return
		}
		rows, err = d.db.Query(ctx, _selArticleByTag, tag, ps, (pn-1)*ps)
		if err != nil {
			log.Error("[GetArticles]db.Query(%s,%d,%d) error(%v)", _selArticleByTag, tag, page.Num*page.Size, page.Size, err)
			return
		}
	} else {
		countRow := d.db.QueryRow(ctx, _countArticleAll)
		if err = countRow.Scan(&page.Total); err != nil {
			log.Error("[GetArticles]row.ScanCount error(%v)", err)
			return
		}
		rows, err = d.db.Query(ctx, _selArticleAll, ps, (pn-1)*ps)
		if err != nil {
			log.Error("[GetArticles]db.Query(%s,%d,%d) error(%v)", _selArticleAll, page.Num*page.Size, page.Size, err)
			return
		}
	}

	defer rows.Close()
	articles = &model.ArticleList{Page: page}
	err = scanner.Scan(rows, &articles.Result)
	return
}

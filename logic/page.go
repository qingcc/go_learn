package logic

import (
	"github.com/go-xorm/xorm"
)

//page 从0开始
func Page(sess *xorm.Session, limit int, page int) *xorm.Session {
	_sess, _, _, _ := pages(sess, limit, page)
	return _sess
}

func Pager(sess *xorm.Session, limit int, page int) (*xorm.Session, int64, int, int) {
	return pages(sess, limit, page)
}

func pages(sess *xorm.Session, limit int, page int) (*xorm.Session, int64, int, int) {
	_sess := sess
	all, _ := _sess.Count()
	if int(all) <= limit*page { //最后一页无数据
		if int(all)%limit == 0 {
			page = int(all)/limit - 1
		} else {
			page = int(all) / limit
		}
		if page < 0 {
			page = 0
		}
	}
	sess.Limit(limit, limit*page)
	return sess, all, limit, page
}

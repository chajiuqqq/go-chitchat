package data

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	Db *sql.DB
)

type (
	Session struct {
		Id       int
		Uuid     string
		Email    string
		UserId   int
		CreateAt time.Time
	}

	Thread struct {
		Id       int
		Uuid     string
		Topic    string
		UserId   int
		CreateAt time.Time
	}
)

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres password=mkQ445683 dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

//获取所有thread
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id, uuid,topic,user_id,created_at from threads order by created_at desc")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreateAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}

//返回一个thread下的回帖数量
func (t *Thread) NumReplies() (count int) {
	rows, err := Db.Query("select count(*) from posts where thread_id=$1", t.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

func UserByEmail(email string) {

}

func (sess *Session) Check() (ok bool, err error) {
	return
}
package sqlxdemo

import (
	"errors"
	"fmt"
	"godemo/pkg"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"godemo/internal/gostorage/sqlxdemo/model"
)

var tmpDB *sqlx.DB

func init() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	tmpDB = db
}

func InitDB() (db *sqlx.DB, err error) {
	dsn := pkg.GenMysqlDSN("")
	db, err = sqlx.Connect(pkg.MysqlType, dsn)
	return
}

// QueryRowDemo 单行查询
func QueryRowDemo() {
	var music model.Music

	sqlStr := "select id, music_author, music_name from music_music"
	err := tmpDB.Get(&music, sqlStr)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Println(music)
}

// QueryMultiRowDemo 多行查询
func QueryMultiRowDemo() {
	var music []model.Music
	sqlStr := "select * from music_music"
	err := tmpDB.Select(&music, sqlStr)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	for _, v := range music {
		fmt.Println(v)
	}
}

// InsertRowDemo 插入单条数据
func InsertRowDemo() {
	sqlStr := "insert into music_music(music_author, music_name, music_album, music_time, music_type, " +
		"music_lyrics, music_arranger) values (?, ?, ?, ?, ?, ?, ?)"
	ret, err := tmpDB.Exec(sqlStr, "zhangsan", "黄种人", "天地", "2022-11-27", "love", "方文山", "taylor")

	if err != nil {
		fmt.Printf("Inesert into music_music failed: %v", err)
		return
	}

	Tid, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("Get Id failed: %v", err)
		return
	}
	fmt.Printf("Insert success id is %v", Tid)
}

// UpdateRowDemo 更新数据
func UpdateRowDemo() {
	sqlStr := "update music_music set music_type = ? where music_author = ?"
	ret, err := tmpDB.Exec(sqlStr, "God", "周杰伦")
	if err != nil {
		fmt.Printf("Update Row failed: %v", err)
		return
	}

	nRows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Get rowsAffected failed: %v", nRows)
		return
	}
	fmt.Printf("Update Success nRows: %v", nRows)
}

// DeleteRowDemo 删除指定数据
func DeleteRowDemo() {
	sqlStr := "delete from music_music where music_name = ?"
	ret, err := tmpDB.Exec(sqlStr, "zhangsan")

	if err != nil {
		fmt.Printf("Delete Row failed: %v", err)
		return
	}

	nRows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("Get rowsAffected failed: %v", nRows)
		return
	}

	fmt.Println("Delete Successfully nRows: ", nRows)
}

// InsertRowDemo2 绑定SQL语句与结构体或map中的同名字段
func InsertRowDemo2() {
	sqlStr := "insert into music_music(music_author, music_name, music_album, music_time, music_type, " +
		"music_lyrics, music_arranger) values (:music_author, :music_name, :music_album, :music_time, :music_type," +
		":music_lyrics, :music_arranger)"

	music_info := map[string]interface{}{
		"music_author":   "linshukai",
		"music_name":     "肖邦的序曲",
		"music_album":    "xiaobang",
		"music_time":     "2022-1-1",
		"music_type":     "classical",
		"music_lyrics":   "xiaobang",
		"music_arranger": "wangzherognyao",
	}
	_, err := tmpDB.NamedExec(sqlStr, music_info)
	if err != nil {
		fmt.Println("Insert into music failed(NamedExec): ", err)
		return
	}
}

// NameQuery 绑定SQL语句与结构体和map中的同名字段进行查询操作
func NameQuery() {
	sqlStr := "select * from music_music where music_author=:music_author"

	condition := model.Music{Author: "周杰伦"}
	rows, err := tmpDB.NamedQuery(sqlStr, condition)
	if err != nil {
		fmt.Printf("NamedQuery failed: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var music model.Music
		err := rows.StructScan(&music)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Music Info: %v\n", music)
	}
}

// TransactionCommitDemo 事务操作
func TransactionCommitDemo() (err error) {
	tx, err := tmpDB.Beginx()
	if err != nil {
		fmt.Printf("begin trans failed, err: %v \n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("commit success")
		}
	}()

	sqlStr := "update music_music set music_name = ? where music_author = ?"
	ret, err := tx.Exec(sqlStr, "coldplay", "张三")
	if err != nil {
		return err
	}

	nRows, err := ret.RowsAffected()
	if err != nil {
		return err
	}

	if nRows < 1 {
		return errors.New("exec sqlStr failed: 当更改记录数目<1时测试事务回滚")
	}
	return nil
}

// BatchInsertMusic Sqlx.In批量插入
func BatchInsertMusic(musicList []interface{}) error {
	fmt.Println(musicList...)

	InsertSql, args, _ := sqlx.In(
		"INSERT INTO music_music (music_author, music_name, music_album, music_time, music_type, music_lyrics, music_arranger) VALUES  (?), (?)",
		musicList...) // 如果Music实现了Value方法，sqlx.In会通过调用Value()方法展开它

	fmt.Println(InsertSql)
	fmt.Println(args)

	_, err := tmpDB.Exec(InsertSql, args...)
	return err
}

// BatchInsertMusicByNamed 通过NamedExec绑定字段查询
func BatchInsertMusicByNamed(musicList []*model.Music) error {
	sqlStr := "INSERT INTO music_music (music_author, music_name, music_album, music_time, music_type, music_lyrics, music_arranger) VALUES " +
		"(:music_author, :music_name, :music_album, :music_time, :music_type, :music_lyrics, :music_arranger)"

	_, err := tmpDB.NamedExec(sqlStr, musicList)
	return err
}

// QueryByNames 通过Name查询Music
func QueryByNames(musicNames []string) (musicList []model.Music, err error) {
	sqlStr := "select * from music_music where music_name in (?)"

	query, args, _ := sqlx.In(sqlStr, musicNames)
	fmt.Println(query)

	query = tmpDB.Rebind(query)
	fmt.Println(query)

	err = tmpDB.Select(&musicList, query, args...)
	return
}

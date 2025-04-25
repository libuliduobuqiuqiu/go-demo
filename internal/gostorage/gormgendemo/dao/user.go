package dao

import (
	"fmt"
	"godemo/internal/gostorage/gormdemo"
	"godemo/internal/gostorage/gormgendemo/model"
	"godemo/internal/gostorage/gormgendemo/query"
)

type MyUser struct {
	model.User

	Info string `json:"info"`
}

func ListUsers() (err error) {

	db, err := gormdemo.InitDB()
	if err != nil {
		return
	}

	query.SetDefault(db)
	_, err = query.User.Update(query.User.Age, 99)
	if err != nil {
		return
	}

	var myUsers []*MyUser
	err = db.Find(&myUsers).Error
	if err != nil {
		return
	}

	fmt.Println(len(myUsers))
	for _, v := range myUsers {
		fmt.Println(v)
	}

	return
}

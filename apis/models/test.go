package models

import "apigin/config"

type Member struct {
	Id        int    `json:"id" form:"id"`
	LoginName string `json:"login_name" form:"login_name"`
	Password  string `json:"password" form:"password"`
}

func OneMember(id int) (m Member, err error) {
	m.Id = 0
	m.LoginName = ""
	m.Password = ""
	err = config.DBHandle.QueryRow("SELECT id, login_name, password FROM ppgo_member WHERE id=? LIMIT 1", id).Scan(&m.Id, &m.LoginName, &m.Password)
	return
}

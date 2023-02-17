package models

import (
	"Shop-Api/config"
)

type LoginModel struct {
}

func (*LoginModel) Login(username string, password string) (bool, string) {
	db, err := config.GetDB()
	if err != nil {
		return false, ""
	} else {
		passfDB, err2 := db.Query("select userpass, role from account where username = ?", username)
		if err2 != nil {
			return false, ""
		} else {
			pass := ""
			role := ""
			for passfDB.Next() {
				passfDB.Scan(&pass, &role)
			}
			if password == pass {
				return true, role
			} else {
				return false, ""
			}
		}
	}
}

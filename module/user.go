package module

import (
	"chessSvr/db"

	"errors"

	"github.com/aceld/zinx/zlog"
)

func Login(userid int, pd string) error {
	zlog.Debugf("login: %d  %s", userid, pd)

	chessUser := db.ChessUser{Id: userid}
	userRec := db.GDB.Where("password = ?", pd).First(&chessUser)
	if userRec.Error != nil {
		zlog.Error(userRec.Error)
		return userRec.Error
	}
	if userRec.RowsAffected == 0 {
		err := errors.New("not found user")
		zlog.Error(err)
		return err
	}
	return nil
}

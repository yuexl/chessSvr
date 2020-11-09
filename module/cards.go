package module

import (
	"github.com/aceld/zinx/zlog"

	"chessSvr/db"
)

func GetAllCards() []db.ChessCard {
	var cards []db.ChessCard
	cardsRec := db.GDB.Find(&cards)
	if cardsRec.Error != nil || cardsRec.RowsAffected == 0 {
		zlog.Error(cardsRec.Error)
		zlog.Error("cardsRec.RowsAffected == 0")
		return nil
	}
	return cards
}

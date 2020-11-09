package module

import (
	"chessSvr/db"

	"github.com/aceld/zinx/zlog"
)

func GetCardPercents(techlevel int8) (int, []int8) {
	var cardPercent db.CardPercent
	percentRec := db.GDB.Where("tech_level = ?", techlevel).First(&cardPercent)
	if percentRec.Error != nil || percentRec.RowsAffected == 0 {
		zlog.Error(percentRec.Error)
		zlog.Error("percentRec.RowsAffected == 0")
		return 0, nil
	}
	zlog.Debug(cardPercent)
	percents := make([]int8, 5)
	percents[0] = cardPercent.PercentOne
	percents[1] = cardPercent.PercentTwo
	percents[2] = cardPercent.PercentThree
	percents[3] = cardPercent.PercentFour
	percents[4] = cardPercent.PercentFive

	return cardPercent.CallNumber, percents
}

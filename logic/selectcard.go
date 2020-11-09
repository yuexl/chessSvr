package logic

import (
	"encoding/json"
	"math/rand"
	"sync"
	"time"

	"chessSvr/db"
	"chessSvr/module"

	"github.com/aceld/zinx/zlog"
)

var (
	CardsGroupMap = sync.Map{}
	once          = sync.Once{}
)

type CardsGroup struct {
	TechLevel int
	Cards     []db.ChessCard
}

func FillCardsGroup() {
	allCards := module.GetAllCards()
	var level1Group = CardsGroup{
		TechLevel: 1,
		Cards:     make([]db.ChessCard, 0),
	}
	var level2Group = CardsGroup{
		TechLevel: 2,
		Cards:     make([]db.ChessCard, 0),
	}
	var level3Group = CardsGroup{
		TechLevel: 3,
		Cards:     make([]db.ChessCard, 0),
	}
	var level4Group = CardsGroup{
		TechLevel: 4,
		Cards:     make([]db.ChessCard, 0),
	}
	var level5Group = CardsGroup{
		TechLevel: 5,
		Cards:     make([]db.ChessCard, 0),
	}
	for i := 0; i < len(allCards); i++ {
		if allCards[i].Cost == 1 {
			level1Group.Cards = append(level1Group.Cards, allCards[i])
			continue
		}
		if allCards[i].Cost == 2 {
			level2Group.Cards = append(level2Group.Cards, allCards[i])
			continue
		}
		if allCards[i].Cost == 3 {
			level3Group.Cards = append(level3Group.Cards, allCards[i])
			continue
		}
		if allCards[i].Cost == 4 {
			level4Group.Cards = append(level4Group.Cards, allCards[i])
			continue
		}
		if allCards[i].Cost == 5 {
			level5Group.Cards = append(level5Group.Cards, allCards[i])
			continue
		}
	}
	CardsGroupMap.Store(1, level1Group)
	CardsGroupMap.Store(2, level2Group)
	CardsGroupMap.Store(3, level3Group)
	CardsGroupMap.Store(4, level4Group)
	CardsGroupMap.Store(5, level5Group)
}

func GetSelectCards(techlevel int8) string {
	zlog.Debug("GetSelectCards ", techlevel)

	once.Do(func() {
		FillCardsGroup()
	})

	cardNum, percents := module.GetCardPercents(techlevel)
	zlog.Debug(cardNum, percents)

	selCards := make([]db.ChessCard, 0)
	for i := 0; i < cardNum; i++ {
		rand.Seed(time.Now().UnixNano())
		cardLevel := GetCardLevel(rand.Intn(100), percents)
		zlog.Debug(cardLevel)
		var levelGroup CardsGroup
		levelGroup.TechLevel = cardLevel
		value, ok := CardsGroupMap.Load(cardLevel)
		if ok {
			group, ok := value.(CardsGroup)
			if ok {
				zlog.Debug(group)
				groupNum := len(group.Cards)
				rand.Seed(time.Now().UnixNano())
				selCards = append(selCards, group.Cards[rand.Intn(groupNum)])
			}
		}
	}

	bytes, err := json.Marshal(&selCards)
	if err != nil {
		zlog.Error(err)
		return ""
	}

	return string(bytes)
}

func GetCardLevel(random int, percents []int8) int {
	seed := int8(random)
	var tip1, tip2, tip3, tip4 int8
	tip1 = percents[0]
	tip2 = tip1 + percents[1]
	tip3 = tip2 + percents[2]
	tip4 = tip3 + percents[3]
	if seed >= 0 && seed < tip1 {
		return 1
	} else if seed >= tip1 && seed < tip2 {
		return 2
	} else if seed >= tip2 && seed < tip3 {
		return 3
	} else if seed >= tip3 && seed < tip4 {
		return 4
	} else {
		return 5
	}
}

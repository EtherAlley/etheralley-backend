package entities

import "github.com/etheralley/etheralley-apis/common"

type DisplayConfig struct {
	Colors       *DisplayColors
	Info         *DisplayInfo
	Picture      *DisplayPicture
	Achievements *DisplayAchievements
	Groups       *[]DisplayGroup
}

type DisplayColors struct {
	Primary       string
	Secondary     string
	PrimaryText   string
	SecondaryText string
	Shadow        string
	Accent        string
}

type DisplayInfo struct {
	Title         string
	Description   string
	TwitterHandle string
}

type DisplayPicture struct {
	Item *DisplayItem
}

type DisplayAchievements struct {
	Text  string
	Items *[]DisplayAchievement
}

type DisplayAchievement struct {
	Id    string
	Index uint64
	Type  common.AchievementType
}

type DisplayGroup struct {
	Id    string
	Text  string
	Items *[]DisplayItem
}

type DisplayItem struct {
	Id    string
	Index uint64
	Type  common.BadgeType
}

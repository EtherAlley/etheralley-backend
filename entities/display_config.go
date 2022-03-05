package entities

import "github.com/etheralley/etheralley-core-api/common"

type DisplayConfig struct {
	Colors       *DisplayColors
	Text         *DisplayText
	Picture      *DisplayPicture
	Achievements *DisplayAchievements
	Groups       *[]DisplayGroup
}

type DisplayColors struct {
	Primary       string
	Secondary     string
	PrimaryText   string
	SecondaryText string
}

type DisplayText struct {
	Title       string
	Description string
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

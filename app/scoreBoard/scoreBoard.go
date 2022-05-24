package scoreboard

import (
	"time"
)

type Entry struct {
	UserName string
	Date     time.Time
}

type scoreBoard struct {
	Entries []Entry
}

var ScoreBoard = scoreBoard{
	Entries: []Entry{},
}

func init() {
	ScoreBoard.Entries = append(ScoreBoard.Entries, Entry{
		UserName: "zasd;fhasdk;fjhaslkdjfhaslkdjfhaslkdjfhaskldjh",
		Date:     time.Now(),
	})
}

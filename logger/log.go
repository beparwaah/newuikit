package logger

import (
	"fmt"
)

var Version string = "1.0"

func Log(mess string) {
	fmt.Println("[LOG] " + mess)
}

type UserProfilePost struct {
	Visibility string `cql:"visibility"`
	PostId     int64  `cql:"post_id"`
	PostType   string `cql:"post_type"`
	TagStatus  string `cql:"tag_status"`
	UserId     int
}
type FollowUpdate struct {
	Type     string `json:"type" cql:"type,,primarykey"`
	ActorId  int    `json:"actor_id" cql:"actor_id,,primarykey"`
	EntityId int    `json:"entity_id"`
}

type NeoUser struct {
	Email        string   `json:"email"`
	UserID       string   `json:"user_id"`
	Mobiles      []string `json:"known_mobiles"`
	KnownMobiles string
	ID           string `json:"mobile"`
	JoinDate     string `json:"join_date"`
}

type UserHomeLocalityInfo struct {
	HomeLocalityId int    `json:"home_locality_id"`
	HomeCity       string `json:"home_city"`
}

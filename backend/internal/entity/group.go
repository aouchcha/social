package entity

import (
	"time"
)

type CreateGroupe struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GroupInvites struct {
	GroupeId   int    `json:"group_id"`
	Invited_by int    `json:"invited_by"`
	Invited    int    `json:"invited"`
	Status     string `json:"status"`
}

type GroupRequest struct {
	GroupeId      int    `json:"group_id"`
	RequestSender int    `json:"requester_id"`
	Status        string `json:"status"`
}

type Groups struct {
	GroupId     int
	Name        string
	Description string
	Created_By  string
	IsMember    bool
}

type GetGroupeResponse struct {
	AllGroups []Groups `json:"groups"`
}

type CreateEvent struct {
	EventId     int       `json:"event_id"`
	UserId      int       `json:"user_id"`
	GroupId     int       `json:"group_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventDate   string    `json:"event_date"`
	EventTime   string    `json:"event_time"`
	CreatedAt   time.Time `json:"created_at"`
	CreatorName string    `json:"event_creator"`
}

type EventUpdate struct {
	EventId int    `json:"event_id"`
	User_id int    `json:"user_id"`
	Status  string `json:"status"`
}

// type GroupChat struct {
// 	GroupId int `json:"group_id"`
// 	Message string `json:"message"`
// }

type GroupChatResponse struct {
	GroupId    int    `json:"group_id"`
	SenderId   int    `json:"user_id"`
	SenderName string `json:"username"`
	Message    string `json:"message"`
	MesasgeId  int    `json:"message_id"`
	CreatedAt  string `json:"created_at"`
}

package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"socialNetwork/internal/entity"
)

type GroupChatRepo struct {
	db *sql.DB
}

func newGroupChatRepositry(db *sql.DB) *GroupChatRepo {
	return &GroupChatRepo{db: db}
}

func (r *GroupChatRepo) GetAllMsg(group_id int) (*sql.Rows, int, error) {
	Query := "SELECT gt.sender, gt.message, gt.created_at, gt.message_id,u.username FROM group_chat gt JOIN users u ON gt.sender = u.user_id WHERE group_id = ?"
	rows, err := r.db.Query(Query, group_id)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("we cant get the group messages in the first time")
	}
	return rows, http.StatusOK, nil
}

func (r *GroupChatRepo) AddMsg(msg entity.GroupChatResponse) (int, int, error) {
	var message_id int
	Query := "INSERT INTO group_chat (group_id, sender, message) VALUES (?, ?, ?) RETURNING message_id"
	err := r.db.QueryRow(Query, msg.GroupId, msg.SenderId, msg.Message).Scan(&message_id)
	if err != nil {
		fmt.Println("Error", err)
		return 0, http.StatusInternalServerError, errors.New("we can't add the msg into data base")
	}
	return message_id, http.StatusOK, nil
}

func (r *GroupChatRepo) GetLastMsg(msg_id int) *sql.Row {
	Query := "SELECT gt.sender, gt.message, gt.created_at, gt.message_id,u.username FROM group_chat gt JOIN users u ON gt.sender = u.user_id WHERE message_id = ?"
	row := r.db.QueryRow(Query, msg_id)
	return row
}

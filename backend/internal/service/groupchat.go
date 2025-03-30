package service

import (
	"errors"
	"fmt"
	"net/http"
	"socialNetwork/internal/entity"
	"socialNetwork/internal/repository"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type GroupChatService struct {
	GroupChatRepo repository.GroupChat
	GroupRepo     repository.Group
}

func newGroupeChatService(GroupChatRepo repository.GroupChat, GroupRepo repository.Group) *GroupChatService {
	return &GroupChatService{
		GroupChatRepo: GroupChatRepo,
		GroupRepo:     GroupRepo,
	}
}

func (s *GroupChatService) GetAllMsg(group_id_str string) ([]entity.GroupChatResponse, int, error) {
	group_id, err := strconv.Atoi(group_id_str)
	if err != nil || !s.GroupRepo.CheckGroup(group_id) {
		return []entity.GroupChatResponse{}, http.StatusBadRequest, errors.New("the group id is not valid")
	}
	var data []entity.GroupChatResponse
	rows, status, err := s.GroupChatRepo.GetAllMsg(group_id)
	if err != nil {
		return []entity.GroupChatResponse{}, status, err
	}
	for rows.Next() {
		var one entity.GroupChatResponse
		err := rows.Scan(&one.SenderId, &one.Message, &one.CreatedAt, &one.MesasgeId, &one.SenderName)
		if err != nil {
			return []entity.GroupChatResponse{}, http.StatusInternalServerError, errors.New("error while scaning the group chat in the get all msgs")
		}
		one.GroupId = group_id
		data = append(data, one)
	}
	return data, http.StatusOK, nil
}

func (s *GroupChatService) CheckData(group_id_str string, user_id int) (int, error) {
	group_id, err := strconv.Atoi(group_id_str)
	if err != nil || !s.GroupRepo.CheckGroup(group_id) {
		return http.StatusBadRequest, errors.New("the group doesn't exist")
	}
	if !s.GroupRepo.CheckUser(user_id) || !s.GroupRepo.IsMember(user_id, group_id) {
		return http.StatusBadRequest, errors.New("the user doesn't exist or is not a member")
	}

	return http.StatusOK, nil
}

func (s *GroupChatService) AddMsgAndShareIt(msg entity.GroupChatResponse, clients map[int][]*websocket.Conn, mu sync.Mutex, conn *websocket.Conn, group_id_str string) (int, error) {
	if len(strings.TrimSpace(msg.Message)) == 0 || len(strings.TrimSpace(msg.Message)) > 100 {
		return http.StatusBadRequest, errors.New("bad request the message length shouldn't be 0 and souldn't pass 100 char")
	}
	GroupId, err := strconv.Atoi(group_id_str)
	if err != nil {
		return http.StatusBadRequest, errors.New("convert group id into string failed")
	}
	fmt.Println(msg, GroupId)
	if !s.GroupRepo.CheckGroup(GroupId) || !s.GroupRepo.IsMember(msg.SenderId, GroupId) || !s.GroupRepo.CheckUser(msg.SenderId) {
		return http.StatusBadRequest, errors.New("the user is not valid, or the group is not valid, or the user isn't a member in the group")
	}
	msg.GroupId = GroupId
	message_id, status, err := s.GroupChatRepo.AddMsg(msg)
	if err != nil {
		return status, err
	}
	var msgR entity.GroupChatResponse
	row := s.GroupChatRepo.GetLastMsg(message_id)
	row.Scan(&msgR.SenderId, &msgR.Message, &msgR.CreatedAt, &msgR.MesasgeId, &msgR.SenderName)
	msgR.GroupId = GroupId
	for _, connection := range clients {

		for _, connn := range connection {

			err := connn.WriteJSON(msgR)
			if err != nil {
				continue
			}

		}

	}
	return http.StatusOK, nil
}

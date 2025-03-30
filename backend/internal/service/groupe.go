package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"socialNetwork/internal/entity"
	"socialNetwork/internal/repository"
)

type GroupeService struct {
	GroupRepo repository.Group
}

func newGroupeService(GroupRepo repository.Group) *GroupeService {
	return &GroupeService{
		GroupRepo: GroupRepo,
	}
}

func (s *GroupeService) GetGroups(user_id int) ([]entity.Groups, int, error) {
	var groups []entity.Groups
	rows, status, err := s.GroupRepo.GetGroups(user_id)
	if err != nil {
		return nil, status, err
	}
	for rows.Next() {
		var infos entity.Groups
		err = rows.Scan(&infos.GroupId ,&infos.Name, &infos.Description, &infos.Created_By, &infos.IsMember)
		if err != nil {
			return []entity.Groups{}, http.StatusInternalServerError, err
		}
		groups = append(groups, infos)

	}
	return groups, http.StatusOK, nil
}

func (s *GroupeService) CreateGroupe(data entity.CreateGroupe) (int, int, error) {
	if strings.TrimSpace(data.Description) == "" || len(strings.TrimSpace(data.Description)) > 10000 {
		return 0, http.StatusBadRequest, errors.New("group description should be beetwen 10000")
	}

	if strings.TrimSpace(data.Name) == "" || len(strings.TrimSpace(data.Name)) < 6 || len(strings.TrimSpace(data.Name)) > 20 {
		return 0, http.StatusBadRequest, errors.New("groupe name should be beetwen 6 and 20")
	}
	if !s.GroupRepo.CheckUser(data.UserId) {
		return 0, http.StatusBadRequest, errors.New("the user doesn't exist")
	}

	GroupId, status, err := s.GroupRepo.CreateGroupe(data)
	if err != nil || status != http.StatusCreated {
		return 0, status, err
	}
	status, err = s.GroupRepo.AddNewMenmber(data.UserId, GroupId)
	if err != nil {
		return 0, status, err
	}

	return GroupId, http.StatusOK, nil
}

func (s *GroupeService) CheckInvite(data entity.GroupInvites) (int, error) {
	if !s.GroupRepo.CheckUser(data.Invited) || !s.GroupRepo.CheckUser(data.Invited_by) || !s.GroupRepo.CheckGroup(data.GroupeId) {
		return http.StatusBadRequest, errors.New("impossible to send the invite (bad user data or groups doesn't exist)")
	}
	if !s.GroupRepo.IsMember(data.Invited_by, data.GroupeId) {
		return http.StatusBadRequest, errors.New("you are not a member in the group")
	}
	if data.Status != "p" {
		return http.StatusBadRequest, errors.New("invalid status")
	}
	var status int = http.StatusOK
	var err error

	if data.Status == "p" {
		status, err = s.GroupRepo.CheckInvite(data)
		if err != nil {
			fmt.Println("HANNI",err)
			return status, err
		}
		status, err = s.GroupRepo.AddNewInvite(data)
	}
	if err != nil {
		return status, err
	}
	return status, nil
}

func (s *GroupeService) UpdateInvite(data entity.GroupInvites) (int, error) {
	if !s.GroupRepo.CheckUser(data.Invited) || !s.GroupRepo.CheckUser(data.Invited_by) || !s.GroupRepo.CheckGroup(data.GroupeId) {
		return http.StatusBadRequest, errors.New("impossible to send the invite (bad user data or groups doesn't exist)")
	}
	if data.Status != "a" && data.Status != "r" {
		return http.StatusBadRequest, errors.New("invalid status")
	}
	status, err := s.GroupRepo.CheckInvite(data)
	if err != nil && err.Error() == "you have already sent and invite" {
		fmt.Println("I am here")
		return status, err
	}
	status, err = s.GroupRepo.UpdateInvite(data)
	if err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (s *GroupeService) CheckRequest(data entity.GroupRequest) (int, error) {
	if !s.GroupRepo.CheckUser(data.RequestSender) || !s.GroupRepo.CheckGroup(data.GroupeId) {
		return http.StatusBadRequest, errors.New("impossible to send the invite (bad user data or groups doesn't exist)")
	}
	if data.Status != "p" {
		return http.StatusBadRequest, errors.New("invalid status")
	}
	var status int = http.StatusOK
	var err error

	if data.Status == "p" {
		status, err = s.GroupRepo.IsRequest(data)
		if err != nil {
			return status, err
		}
		status, err = s.GroupRepo.AddNewRequest(data)
	}
	if err != nil {
		return status, err
	}
	return status, nil
}

func (s *GroupeService) UpdateRequest(data entity.GroupRequest) (int, error) {
	if !s.GroupRepo.CheckUser(data.RequestSender) || !s.GroupRepo.CheckGroup(data.GroupeId) {
		return http.StatusBadRequest, errors.New("impossible to send the invite (bad user data or groups doesn't exist)")
	}
	
	owner_id, status, err := s.GroupRepo.GetOwner(data.GroupeId)
	if err != nil {
		return status, err
	}

	if data.Status != "a" && data.Status != "r" {
		return http.StatusBadRequest, errors.New("invalid status")
	}
	status, err = s.GroupRepo.IsRequest(data)
	if err != nil && err.Error() == "you have already sent a request" {
		return status, err
	}
	status, err = s.GroupRepo.UpdateRequest(owner_id, data)
	if err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (s *GroupeService) CreateEvent(data entity.CreateEvent) (int, error) {
	if !s.GroupRepo.CheckUser(data.UserId) || !s.GroupRepo.CheckGroup(data.GroupId) {
		return http.StatusBadRequest, errors.New("impossible to send the invite (bad user data or groups doesn't exist)")
	}

	if !s.GroupRepo.IsMember(data.UserId, data.GroupId) {
		return http.StatusBadRequest, errors.New("the user is not a member in the group")
	}

	if strings.TrimSpace(data.Description) == "" || len(strings.TrimSpace(data.Description)) > 10000 {
		return http.StatusBadRequest, errors.New("event description should be beetwen 1 and 10000")
	}

	if strings.TrimSpace(data.Title) == "" || len(strings.TrimSpace(data.Title)) < 6 || len(strings.TrimSpace(data.Title)) > 20 {
		return http.StatusBadRequest, errors.New("event title should be beetwen 6 and 20")
	}

	combinedDateTime := data.EventDate + "T" + data.EventTime + ":00Z"
	layout := time.RFC3339
	parsedEventDateTime, err := time.Parse(layout, combinedDateTime)
	if err != nil {
		return http.StatusBadRequest, errors.New("invalid event date and time format")
	}

	if !time.Now().Before(parsedEventDateTime) {
		return http.StatusBadRequest, errors.New("the event date ot time is in the past")
	}

	if s.GroupRepo.IsTheEventExist(data) {
		return http.StatusBadRequest, errors.New("this event is already created")
	}

	status, err := s.GroupRepo.CreateEvent(data)
	if err != nil {
		return status, err
	}

	return http.StatusOK, nil
}

func (s *GroupeService) GetAllEvents(group_id_str string) ([]entity.CreateEvent, int, error) {
	group_id, err := strconv.Atoi(group_id_str); if err != nil {
		return []entity.CreateEvent{}, http.StatusBadRequest, errors.New("group id not valid")
	}
	if !s.GroupRepo.CheckGroup(group_id) {
		return []entity.CreateEvent{}, http.StatusBadRequest, errors.New("the group that you want get his events doesn't exist")
	}
	var Events []entity.CreateEvent
	rows, status, err := s.GroupRepo.GetAllEvents(group_id)
	if err != nil {
		return nil, status, err
	}
	for rows.Next() {
		var infos entity.CreateEvent
		err = rows.Scan(&infos.EventId, &infos.UserId, &infos.GroupId, &infos.Title, &infos.Description, &infos.EventDate, &infos.EventTime, &infos.CreatedAt, &infos.CreatorName)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		Events = append(Events, infos)
	}
	return Events, http.StatusOK, nil
}

func (s *GroupeService) UpdateEvent(data entity.EventUpdate) (int, error) {
	event, status, err := s.GroupRepo.GetEventById(data.EventId)
	if err != nil {
		return status, err
	}
	if !s.GroupRepo.CheckUser(data.User_id) || !s.GroupRepo.CheckGroup(event.GroupId) || !s.GroupRepo.IsMember(data.User_id, event.GroupId) {
		return http.StatusBadRequest, errors.New("the user or group doesn't exist or the user is not a member in the group")
	}
	if data.Status != "g" && data.Status != "n" {
		return http.StatusBadRequest, errors.New("invalid status")
	}
	if s.GroupRepo.IsHeAlreadyReactOnEvent(data) {
		fmt.Println("hanni hna deja dert hajja")
		fmt.Println(data)
		status, err = s.GroupRepo.UpdateEvent(data)
	}else {
		fmt.Println("hanni hna jms dert chi hajja")
		fmt.Println(data)
		status, err = s.GroupRepo.FirstReactOnInEvent(data)
	}
	if err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

func (s *GroupeService) GetAllPosts(group_id_str string)([]entity.Post, int, error) {
	fmt.Println("GROUPID",group_id_str)
	group_id, err := strconv.Atoi(group_id_str); if err != nil {
		return []entity.Post{}, http.StatusBadRequest, errors.New("group id not valid")
	}
	if !s.GroupRepo.CheckGroup(group_id) {
		return []entity.Post{}, http.StatusBadRequest, errors.New("the group that you want get his events doesn't exist")
	}
	var posts []entity.Post
	rows, status, err := s.GroupRepo.GetAllPosts(group_id)
	if err != nil {
		return []entity.Post{},status, err
	}
	for rows.Next() {
		var post entity.Post
		post.GroupID = uint(group_id)
		err = rows.Scan(&post.PostID, &post.UserID, &post.Content, &post.Image_url,  &post.GroupID,&post.Privacy, &post.CreatedAt, &post.UserName, &post.UserID)
		if err != nil {
			return []entity.Post{}, http.StatusInternalServerError, err
		}
		posts = append(posts, post)
	}
	return posts, http.StatusOK, nil
}

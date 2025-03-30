package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"socialNetwork/internal/entity"
)

type GroupRepo struct {
	db *sql.DB
}

func newGroupRepository(db *sql.DB) *GroupRepo {
	return &GroupRepo{db: db}
}

func (r *GroupRepo) GetGroups(user_id int) (*sql.Rows, int, error) {
	Query := `
		SELECT g.group_id, g.name, g.description, u.username, TRUE AS is_member 
		FROM groups g 
		JOIN group_members gm ON g.group_id = gm.group_id 
		JOIN users u ON u.user_id = g.owner_id 
		WHERE gm.member_id = $1

		UNION ALL

		SELECT g.group_id, g.name, g.description, u.username, FALSE AS is_member 
		FROM groups g 
		JOIN users u ON u.user_id = g.owner_id
		WHERE NOT EXISTS 
     	(SELECT * FROM group_members gm WHERE gm.group_id = g.group_id AND gm.member_id = $1);
	`
	rows, err := r.db.Query(Query, user_id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return rows, http.StatusOK, nil
}

func (r *GroupRepo) CheckUser(UserId int) bool {
	var IsUserExist bool
	Query := "SELECT COUNT(user_id) FROM users u WHERE u.user_id = ?"

	if err := r.db.QueryRow(Query, UserId).Scan(&IsUserExist); err != nil {
		return false
	}
	return IsUserExist
}

func (r *GroupRepo) CheckGroup(GroupeId int) bool {
	var IsGroupExist bool
	Query := "SELECT COUNT(group_id) FROM groups g WHERE g.group_id = ?"

	if err := r.db.QueryRow(Query, GroupeId).Scan(&IsGroupExist); err != nil {
		return false
	}
	return IsGroupExist
}

func (r *GroupRepo) CreateGroupe(data entity.CreateGroupe) (int, int, error) {
	Query := "INSERT INTO groups (owner_id, name, description) VALUES (?, ?, ?) RETURNING group_id"
	var group_id int
	if err := r.db.QueryRow(Query, data.UserId, data.Name, data.Description).Scan(&group_id); err != nil {
		log.Println(err)
		return 0, http.StatusBadRequest, errors.New("the group name is already taken")
	}
	return group_id, http.StatusCreated, nil
}

func (r *GroupRepo) AddNewInvite(data entity.GroupInvites) (int, error) {
	Query := "INSERT INTO group_invitations (group_id, invited, invited_by, status) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(Query, data.GroupeId, data.Invited, data.Invited_by, data.Status)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, errors.New("we can't add the new invite")
	}
	return http.StatusOK, nil
}

func (r *GroupRepo) UpdateInvite(data entity.GroupInvites) (int, error) {
	var status = http.StatusOK
	var err error
	if data.Status == "a" {
		status , err = r.AddNewMenmber(data.Invited, data.GroupeId)
		if err != nil {
			return status, err
		}
		status, err = r.DeletInvitation(data)
		if err != nil {
			return status, err
		}
	} else {
		status, err = r.DeletInvitation(data)
		if err != nil {
			return status, err
		}
	}
	return status, nil
}

func (r *GroupRepo) DeletInvitation(data entity.GroupInvites) (int, error) {
	Query := "DELETE FROM group_invitations WHERE group_id = ? AND invited = ?"
	_, err := r.db.Exec(Query, data.GroupeId, data.Invited)
	if err != nil {
		fmt.Println("Errr", err)
		fmt.Println("Hani choufni", err)
		return http.StatusInternalServerError, errors.New("we remove the invite")
	}
	return http.StatusOK, nil
}

func (r *GroupRepo) CheckInvite(data entity.GroupInvites) (int, error) {
	// Query to check if an invitation already exists
	inviteQuery := "SELECT COUNT(group_id) FROM group_invitations WHERE group_id = ? AND invited = ? AND invited_by = ? AND status = ?"
	
	// Query to check if a user is a member of the group
	membershipQuery := "SELECT COUNT(group_id) FROM group_members WHERE group_id = ? AND member_id = ?"

	// Check if the invited user is already a member
	var isAlreadyMember int
	if err := r.db.QueryRow(membershipQuery, data.GroupeId, data.Invited).Scan(&isAlreadyMember); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to check membership: %w", err)
	}
	
	if isAlreadyMember > 0 {
		return http.StatusBadRequest, errors.New("the user is already a member in this group")
	}

	// Check if an invitation is already pending
	var pendingInviteCount int
	if err := r.db.QueryRow(inviteQuery, data.GroupeId, data.Invited, data.Invited_by, "p").Scan(&pendingInviteCount); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to check pending invites: %w", err)
	}
	
	if pendingInviteCount > 0 {
		return http.StatusBadRequest, errors.New("you have already sent an invite")
	}

	// Check if the inviter is a member of the group
	var isInviterMember int
	if err := r.db.QueryRow(membershipQuery, data.GroupeId, data.Invited_by).Scan(&isInviterMember); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to check inviter membership: %w", err)
	}
	
	if isInviterMember == 0 {
		return http.StatusBadRequest, errors.New("the inviter isn't a member")
	}

	// All checks passed
	return http.StatusOK, nil
}

func (r *GroupRepo) IsRequest(data entity.GroupRequest) (int, error) {
	fmt.Println("Ana f IsRe")
	var IsRequestExist, IsItMember bool
	var err error
	Query1 := "SELECT 1 FROM group_requests WHERE group_id = ? AND requester_id = ? AND status = ?"
	Query2 := "SELECT 1 FROM group_members WHERE group_id = ? AND member_id = ?"

	err = r.db.QueryRow(Query2, data.GroupeId, data.RequestSender).Scan(&IsItMember)
	// fmt.Println(data.GroupeId, data.RequestSender, IsItMember)
	if err != nil || !IsItMember {

		err = r.db.QueryRow(Query1, data.GroupeId, data.RequestSender, data.Status).Scan(&IsRequestExist)
		if err != nil || !IsRequestExist {
			if err != sql.ErrNoRows {
				return http.StatusInternalServerError, errors.New("internal server error")
			}
		} else {
			return http.StatusBadRequest, errors.New("you have already sent a request")
		}

	} else {
		return http.StatusBadRequest, errors.New("you are already a member in this group")
	}

	return http.StatusOK, nil
}

func (r *GroupRepo) AddNewRequest(data entity.GroupRequest) (int, error) {
	Query := "INSERT INTO group_requests (group_id, requester_id, status) VALUES (?, ?, ?)"
	_, err := r.db.Exec(Query, data.GroupeId, data.RequestSender, data.Status)
	if err != nil {
		return http.StatusInternalServerError, errors.New("we can't add the new request")
	}
	return http.StatusOK, nil
}

func (r *GroupRepo) AddNewMenmber(user_id, group_id int) (int, error) {
	Query := "INSERT INTO group_members (group_id, member_id) VALUES (?, ?)"
	_, err := r.db.Exec(Query, group_id, user_id)
	if err != nil {
		fmt.Println("hjani",err)
		fmt.Println("Error in adding groupe member")
		return http.StatusInternalServerError, errors.New("we can't add the member")
	}
	return http.StatusOK, nil
}

func (r *GroupRepo) UpdateRequest(owner_id int, data entity.GroupRequest) (int, error) {
	var status = http.StatusOK
	var err error
	if data.Status == "a" {
		status, err = r.AddNewMenmber(data.RequestSender, data.GroupeId)
		if err != nil {
			return status, err
		}
		status, err = r.DeletRequest(data)
		if err != nil {
			return status, err
		}
	} else {
		status, err = r.DeletRequest(data)
		if err != nil {
			return status, err
		}
	}
	return status, nil
}

func (r *GroupRepo) DeletRequest(data entity.GroupRequest) (int, error) {
	Query := "DELETE FROM group_requests WHERE group_id = ? AND requester_id = ?"
	_, err := r.db.Exec(Query, data.GroupeId, data.RequestSender)
	if err != nil {
		fmt.Println("Errr", err)
		fmt.Println("Hani choufni", err)
		return http.StatusInternalServerError, errors.New("we can't remove the request")
	}
	return http.StatusOK, nil
}

func (r *GroupRepo) GetOwner(GroupeId int) (int, int, error) {
	var owner_id int
	Query := "SELECT owner_id FROM groups WHERE group_id = ?"
	err := r.db.QueryRow(Query, GroupeId).Scan(&owner_id)
	if err != nil {
		fmt.Println("Error in getting owner group id")
		return 0, http.StatusInternalServerError, errors.New("internal server error")
	}
	return owner_id, http.StatusOK, nil
}

func (r *GroupRepo) IsHeOwner(UserId, GroupId int) bool {
	var IsOwner bool
	Query := "SELECT EXISTS (SELECT * FROM groups WHERE owner_id = ? AND group_id = ?)"
	err := r.db.QueryRow(Query, UserId, GroupId).Scan(&IsOwner)
	if err != nil || !IsOwner {
		return false
	}
	return true
}

func (r *GroupRepo) CreateEvent(data entity.CreateEvent) (int, error) {
	Query := "INSERT INTO group_events (user_id, group_id, title, description, event_date, event_time, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(Query, data.UserId, data.GroupId, data.Title, data.Description, data.EventDate, data.EventTime, time.Now())
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (r *GroupRepo) IsTheEventExist(data entity.CreateEvent) bool {
	var isEventExist bool
	Query := "SELECT EXISTS (SELECT * FROM group_events WHERE group_id = ? AND title = ? AND description = ? AND event_date = ? AND event_time = ?)"
	err := r.db.QueryRow(Query, data.GroupId, data.Title, data.Description, data.EventDate, data.EventTime).Scan(&isEventExist)
	fmt.Println(isEventExist)
	if err != nil || isEventExist {
		return true
	}
	return false
}

func (r *GroupRepo) IsMember(userId, GroupId int) bool {
	var IsIsMember bool
	Query := "SELECT EXISTS (SELECT * FROM group_members WHERE group_id = ? AND member_id = ?)"
	err := r.db.QueryRow(Query, GroupId, userId).Scan(&IsIsMember)
	if err != nil || !IsIsMember {
		return false
	}
	return true
}

func (r *GroupRepo) GetAllEvents(group_id int) (*sql.Rows, int, error)  {
	Query := "SELECT group_events.* , users.username FROM group_events JOIN users ON users.user_id = group_events.user_id WHERE group_id = ?"
	rows, err := r.db.Query(Query, group_id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return rows, http.StatusOK, nil
}

func (r *GroupRepo) GetEventById(event_id int) (entity.CreateEvent, int, error) {
	var event entity.CreateEvent
	Query := "SELECT * FROM group_events WHERE event_id = ?"
	err := r.db.QueryRow(Query, event_id).Scan(&event.EventId, &event.UserId, &event.GroupId, &event.Title, &event.Description, &event.EventDate, &event.EventTime, &event.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.CreateEvent{}, http.StatusBadRequest, errors.New("there is no event")
		}
		return event, http.StatusInternalServerError, err
	}
	return event, http.StatusOK, nil
}

func (r *GroupRepo) FirstReactOnInEvent(data entity.EventUpdate) (int, error) {
	Query := "INSERT INTO event_interest (event_id, user_id, status) VALUES (?, ?, ?)"
	_, err := r.db.Exec(Query, data.EventId, data.User_id, data.Status)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (r *GroupRepo) IsHeAlreadyReactOnEvent(data entity.EventUpdate) bool {
	var IsHeAlreadyReact bool
	fmt.Println("data", data)
	Query := "SELECT EXISTS (SELECT * FROM event_interest WHERE event_id = ? AND user_id = ?)"
	err := r.db.QueryRow(Query, data.EventId, data.User_id).Scan(&IsHeAlreadyReact)
	fmt.Println("IsHeAlreadyReact",IsHeAlreadyReact)
	if err != nil || !IsHeAlreadyReact {
		return false
	}
	return true
}

func (r *GroupRepo) UpdateEvent(data entity.EventUpdate) (int, error) {
	Query := "UPDATE event_interest SET status = ? WHERE event_id = ? AND user_id = ?"
	_, err := r.db.Exec(Query, data.Status, data.EventId, data.User_id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (r *GroupRepo) GetAllPosts(group_id int)(*sql.Rows, int, error) {
	Query := `
		SELECT posts.*, users.username, users.user_id
		FROM posts 
		JOIN users ON users.user_id = posts.user_id
		WHERE group_id = ?
	`
	rows, err := r.db.Query(Query, group_id)
	if err != nil {
		fmt.Println("Error a zebi", err)
		return nil, http.StatusInternalServerError, err
	}
	return rows, http.StatusOK, nil
}
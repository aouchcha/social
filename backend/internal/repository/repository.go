package repository

import (
	"database/sql"
	"socialNetwork/internal/entity"
)

type User interface {
	CheckUserExist(user entity.UserProfile) error
	RegisterUser(user entity.UserProfile) (int, error)
	LoginUser(user entity.UserProfile) (entity.UserProfile, int, error)
}

type Session interface {
	IsTokenExist(session entity.Session) (bool, error)
	DeleteSessionByToken(session entity.Session) error
	DeleteSessionByUserID(session entity.Session) error
	CreatSession(session entity.Session) (int, error)
}

type Post interface {
	GetAllPosts(limit, offset int) ([]entity.Post, int, error)
	GetAllByUserID(userID uint, limit, offset int) ([]entity.Post, int, error)
	GetPostByID(postID uint) (entity.Post, int, error)
	getCommentsByPostID(postID uint) ([]entity.Comment, int, error)
	CreatePost(input entity.Post) (uint, int, error)
	AddPostReaction(input entity.PostReaction) (int, error)
}

type Comment interface {
	CreateComment(input entity.Comment) (int, error)
	AddCommentReaction(input entity.CommentReaction) (int, error)
}

type Message interface {
}

type Group interface {
	GetGroups(user_id int) (*sql.Rows, int, error)
	CheckUser(data int) bool
	CheckGroup(data int) bool
	CreateGroupe(CreateGroupe entity.CreateGroupe) (int, int, error)
	AddNewInvite(data entity.GroupInvites) (int, error)
	UpdateInvite(data entity.GroupInvites) (int, error)
	DeletInvitation(data entity.GroupInvites) (int, error)
	CheckInvite(data entity.GroupInvites) (int, error)
	IsRequest(data entity.GroupRequest) (int, error)
	AddNewRequest(data entity.GroupRequest) (int, error)
	AddNewMenmber(user_id, group_id int) (int, error)
	UpdateRequest(owner_id int, data entity.GroupRequest) (int, error)
	DeletRequest(data entity.GroupRequest) (int, error)
	GetOwner(data int) (int, int, error)
	IsHeOwner(UserId, GroupId int) bool
	CreateEvent(data entity.CreateEvent) (int, error)
	IsTheEventExist(data entity.CreateEvent) bool
	IsMember(UserId, GroupId int) bool
	GetAllEvents(group_id int) (*sql.Rows, int, error)
	GetEventById(event_id int) (entity.CreateEvent, int, error)
	FirstReactOnInEvent(data entity.EventUpdate) (int, error)
	IsHeAlreadyReactOnEvent(data entity.EventUpdate) bool
	UpdateEvent(data entity.EventUpdate) (int, error)
	GetAllPosts(group_id int) (*sql.Rows, int, error)
}

type GroupChat interface {
	GetAllMsg(group_id int) (*sql.Rows, int, error)
	AddMsg(msg entity.GroupChatResponse) (int, int, error)
	GetLastMsg(msg_id int) *sql.Row
}

type Repository struct {
	Post
	User
	Session
	Comment
	Message
	Group
	GroupChat
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post:      newPostRepository(db),
		User:      newUserRepository(db),
		Session:   newSessionRepository(db),
		Comment:   newCommentRepository(db),
		Message:   newMessagesRepo(db),
		Group:     newGroupRepository(db),
		GroupChat: newGroupChatRepositry(db),
	}
}

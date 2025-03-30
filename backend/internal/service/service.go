package service

import (
	"socialNetwork/internal/entity"
	"socialNetwork/internal/repository"
	"sync"

	"github.com/gorilla/websocket"
)

type User interface {
	ServiceSingUp(UserProfile entity.UserProfile) (int, error)
	ServiceLogIn(UserProfile entity.UserProfile) (entity.Session, entity.UserProfile, int, error)
}

type Session interface {
	ServerLogout(Session entity.Session) (int, error)
}

type Post interface {
	GetAllPosts(limit, offset string) ([]entity.Post, int, error)
	GetPostByID(strPostID string) (entity.Post, int, error)
	GetAllByUserID(strUserID, limit, offset string) ([]entity.Post, int, error)
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
	GetGroups(user_id int) ([]entity.Groups, int, error)
	CreateGroupe(data entity.CreateGroupe) (int, int, error)
	CheckInvite(data entity.GroupInvites) (int, error)
	UpdateInvite(data entity.GroupInvites) (int, error)
	CheckRequest(data entity.GroupRequest) (int, error)
	UpdateRequest(data entity.GroupRequest) (int, error)
	CreateEvent(data entity.CreateEvent) (int, error)
	GetAllEvents(group_id_str string) ([]entity.CreateEvent, int, error)
	UpdateEvent(data entity.EventUpdate) (int, error)
	GetAllPosts(group_id_str string) ([]entity.Post, int, error)
}

type GroupChat interface {
	GetAllMsg(group_id string) ([]entity.GroupChatResponse, int, error)
	CheckData(group_id string, user_id int) (int, error)
	AddMsgAndShareIt(msg entity.GroupChatResponse, clients map[int][]*websocket.Conn, mu sync.Mutex, conn *websocket.Conn, group_id string) (int, error)
}

type Service struct {
	User
	Session
	Post
	Comment
	Message
	Group
	GroupChat
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:      newUserService(repo.User, repo.Session),
		Session:   newSessionService(repo.Session),
		Post:      newPostService(repo.Post),
		Comment:   newCommentService(repo.Comment),
		Message:   newMessagesService(repo.Message, repo.User),
		Group:     newGroupeService(repo.Group),
		GroupChat: newGroupeChatService(repo.GroupChat, repo.Group),
	}
}

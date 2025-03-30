package controller

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"socialNetwork/internal/service"
	"socialNetwork/pkg/config"
)

type Handler struct {
	service *service.Service
}

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes(conf *config.Conf) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/uploads/", func(w http.ResponseWriter, r *http.Request) {
		filePath := strings.TrimPrefix(r.URL.Path, "/uploads/")
		if _, err := os.Stat(path.Join("./uploads", filePath)); os.IsNotExist(err) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error checking file: %v", err)
			return
		}
		http.ServeFile(w, r, path.Join("./uploads", filePath))
	})

	routes := h.createRoutes()
	for _, route := range routes {
		mux.Handle(route.Path, h.configCORS(route.Handler))
	}
	return mux
}

func (h *Handler) createRoutes() []Route {
	return []Route{
		{
			Path:    "/api/is-valid",
			Handler: h.isValidToken,
		},
		{
			Path:    "/api/signup",
			Handler: h.signUp,
		},
		{
			Path:    "/api/signin",
			Handler: h.signIn,
		},
		{
			Path:    "/api/signout",
			Handler: h.signOut,
		},
		{
			Path:    "/api/profile/posts/",
			Handler: h.getMyPosts,
		},
		{
			Path:    "/api/post/create",
			Handler: h.createPost,
		},
		{
			Path:    "/api/posts/",
			Handler: h.getALLPosts,
		},
		{
			Path:    "/api/post/",
			Handler: h.getPostbyID,
		},
		{
			Path:    "/api/post/like",
			Handler: h.postReaction,
		},
		{
			Path:    "/api/comment/create",
			Handler: h.createComment,
		},
		{
			Path:    "/api/comment/like",
			Handler: h.commentReaction,
		},
		{
			Path:    "/api/contacts",
			Handler: h.GetOnlineUsers,
		},
		{
			Path:    "/ws",
			Handler: h.WebSocketHandler,
		},
		{
			Path:    "/api/chat/",
			Handler: h.GetMessages,
		},
		{
			Path: "/api/groups",
			Handler: h.GetGroups,
		},
		{
			Path: "/api/groupe/create",
			Handler: h.CreateGroupe,
		},
		{
			Path: "/api/groupe/invite",
			Handler: h.GroupeInvite,
		},
		{
			Path: "/api/groupe/invite/update",
			Handler: h.GroupeInviteUpdate,
		},
		{
			Path: "/api/groupe/request",
			Handler: h.GroupeRequest,
		},
		{
			Path: "/api/groupe/request/update",
			Handler: h.GroupeRequestUpdate,
		},
		{
			Path: "/api/events",
			Handler: h.GetAllEvents,
		},
		{
			Path: "/api/event/create",
			Handler: h.CreateEventHandler,
		},
		{
			Path: "/api/event/update",
			Handler: h.UpdateEvent,
		},
		{
			Path: "/api/groupe/posts",
			Handler: h.GetPostsOfGroup,
		},
		{
			Path : "/api/groupe/load_messages",
			Handler: h.GetAllMsg,
		},
		{
			Path: "/api/groupe/chat",
			Handler: h.ChatHandler,
		},
	}
}

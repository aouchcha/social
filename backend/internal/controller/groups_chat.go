package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"socialNetwork/internal/entity"

	"github.com/gorilla/websocket"
)

func (h *Handler) GetAllMsg(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Methode not allowed")
		return
	}
	group_id := r.URL.Query().Get("groupid")
	data, status, err := h.service.GroupChat.GetAllMsg(group_id)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Groups []entity.GroupChatResponse
	}{
		Groups: data,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

var (
	clients  = make(map[int][]*websocket.Conn)
	mu       sync.Mutex
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	// user_id  = 1
	username = "aouchcha"
)

func (h *Handler) ChatHandler(w http.ResponseWriter, r *http.Request) {

	group_id := r.URL.Query().Get("groupid")
	user_id, err := strconv.Atoi(r.URL.Query().Get("userid"))
	if err != nil {
		fmt.Println("invalid user id not int")
		return
	}
	fmt.Println("GROUPID", group_id)
	status, err := h.service.GroupChat.CheckData(group_id, user_id)
	if err != nil {
		fmt.Println("Ana Hna", err)
		h.errorHandler(w, r, status, err.Error())
		return
	}
	// Corrected version
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket Err", err)
		h.errorHandler(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	mu.Lock()
	if cl, ok := clients[user_id]; ok {
		// User ID exists, append to existing slice
		clients[user_id] = append(cl, conn)
	} else {
		// User ID doesn't exist, create new slice
		clients[user_id] = []*websocket.Conn{conn}
	}
	mu.Unlock()

	fmt.Println("map", clients, len(clients))
	for {
		var msg entity.GroupChatResponse
		msg.SenderId = user_id
		msg.SenderName = username
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("The Connection closed becaude we have a probleme in the reading")
			break
		}
		_, err := h.service.GroupChat.AddMsgAndShareIt(msg, clients, mu, conn, group_id)
		if err != nil {
			fmt.Println(err)
			// h.errorHandler(w, r, status, err.Error())
			// return
		}
	}
	// if len(clients[user_id]) > 1 {
	// 	RemoveConneCtion(clients, user_id, conn)
	// } else {
	// 	mu.Lock()
	// 	delete(clients, user_id)
	// 	mu.Unlock()
	// }

}

func RemoveConneCtion(clients map[int][]*websocket.Conn, user_id int, conn *websocket.Conn) {
	var newConnwctions []*websocket.Conn
	for _, connection := range clients[user_id] {
		if connection.RemoteAddr().String() != conn.RemoteAddr().String() {
			newConnwctions = append(newConnwctions, connection)
		}
	}
	clients[user_id] = newConnwctions
}

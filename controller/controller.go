package controller

import (
	"encoding/json"
	"net/http"
	"time"

	database "github.com/cherry-aggarwal/libr/database"
	"github.com/cherry-aggarwal/libr/models"
	"github.com/cherry-aggarwal/libr/moderators"
	"github.com/google/uuid"
)

// var ModChannel = make(chan int, 3)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to LIBR prototype"))
}

func MsgIN(w http.ResponseWriter, r *http.Request) {

	var msg models.Msg
	var modResponse models.ModResponse

	json.NewDecoder(r.Body).Decode(&msg)
	moderators.Out = 0
	moderators.AskingModsResponse()
	MsgID := uuid.NewString()
	msg.MsgID = MsgID
	modResponse.MsgID = MsgID
	ModID := uuid.NewString()
	modResponse.ModID = ModID
	TimeStamp := time.Now().Unix()
	msg.TimeStamp = TimeStamp
	moderators.SettingMsgStatus(&msg)
	if moderators.Out == 0 {
		msg.Status = "rejected"
	} else {
		msg.Status = "accepted"
		database.InsertMessage(msg)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func MsgOUT(w http.ResponseWriter, r *http.Request) {
    var ts int64
    if err := json.NewDecoder(r.Body).Decode(&ts); err != nil {
        http.Error(w, "Invalid timestamp", http.StatusBadRequest)
        return
    }
    messages := database.GetMessages(ts)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(messages)
}

type Moderator interface {
	StartModeration(msg models.Msg) int
}

package models

type Msg struct {
	MsgID     string `bson:"_id,omitempty" json:"id"`
	Content   string `json:"content" bson:"content"`
	TimeStamp int64  `json:"timestamp" bson:"timestamp"`
	Status    string   `json:"status" bson:"status"`
}

type ModResponse struct {
	ModID        string `bson:"_id,omitempty" json:"id"`
	MsgID        string 
	Status       string   `json:"status" bson:"status"`
	ResponseTime int
}

// type APIresponse struct {
// 	MsgID     string `bson:"_id,omitempty" json:"id"`
// 	TimeStamp int    `json:"timestamp" bson:"timestamp"`
// 	Status    bool   `json:"status" bson:"status"`
// }

package models

type WebsocketMessage struct {
	Code    string      `json:"code" bson:"code"`
	Payload interface{} `json:"payload" bson:"payload"`
	User    string      `json:"user" bson:"user"`
}

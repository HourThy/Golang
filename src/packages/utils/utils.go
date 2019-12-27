package utils
import (
	"encoding/json"
	"net/http"
)

// Message : Message
func Message(status bool, messgae string)(map[string]interface{}){
	return map[string]interface{}{"status" : status, "message" : message}
}

func Respond(w http.Response.Writer, data map[string] interface{}){
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
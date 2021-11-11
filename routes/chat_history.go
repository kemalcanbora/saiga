package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	c "saiga/config"
	"saiga/middleware"
	"saiga/pkg/clients"
	h "saiga/pkg/helpers"
)

func ChatHistory(w http.ResponseWriter, r *http.Request) {
	if middleware.CheckUserType(r, "operator") == false {
		h.HTTPErrorHandler(w, "Unauthorized to access this resource!", http.StatusNotAcceptable)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		objId := params["id"]

		resp := make(map[string]interface{})
		resp["messages"] = clients.Mongo.GetChatHistory(c.Configure().MongoChatCollection, objId).Messages
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	}
}

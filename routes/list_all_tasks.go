package routes

import (
	"encoding/json"
	"net/http"
	c "saiga/config"
	"saiga/middleware"
	"saiga/pkg/clients"
	h "saiga/pkg/helpers"
)

func ListAllTasks(w http.ResponseWriter, r *http.Request) {
	if middleware.CheckUserType(r, "operator") == false {
		h.HTTPErrorHandler(w, "Unauthorized to access this resource!", http.StatusNotAcceptable)
		return
	} else {
		tasks := clients.Mongo.FindAllTask(c.Configure().MongoChatCollection)
		resp := make(map[string]interface{})
		resp["tasks"] = tasks
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	}
}

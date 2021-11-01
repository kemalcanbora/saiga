package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	c "saiga/config"
	cloud "saiga/pkg/aws_cloud"
	h "saiga/pkg/helpers"
)

func DownloadChatMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	objId := params["id"]
	path, _ := os.Getwd()

	attachmentModel, err := h.Mongo.FindAttachmentURL(c.Configure().MongoChatCollection, objId)
	if err != nil {
		fmt.Println(err)
	}
	attachmentBucket := attachmentModel.Messages[0].Attachment.Bucket
	attachmentKey := attachmentModel.Messages[0].Attachment.Key

	result := cloud.DownloadFromS3Bucket(attachmentBucket, attachmentKey, path+c.Configure().DownloadsPath)
	jsonResp, _ := json.Marshal(result)
	w.Write(jsonResp)
}

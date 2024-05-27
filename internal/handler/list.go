package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
	"github.com/mattermost/mattermost-server/v6/model"
)

func List(w http.ResponseWriter, req *http.Request) {
	c := apps.CallRequest{}
	json.NewDecoder(req.Body).Decode(&c)

	w.WriteHeader(200)
	appclient.AsBot(c.Context).DM(c.Context.ActingUser.Id, "Was geht ab?!")

	_, err := appclient.AsBot(c.Context).CreatePost(&model.Post{
		ChannelId: "mrnym5kkkpdzigx7x5i9txtbfc",
		Message:   "Was bleibt dran?!",
	})
	if err != nil {
		fmt.Println(err)
	}
}

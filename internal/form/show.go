package form

import "github.com/mattermost/mattermost-plugin-apps/apps"

var Show = apps.Form{
	Title:  "Show",
	Icon:   "",
	Fields: []apps.Field{},
	Submit: apps.NewCall("/show").WithExpand(apps.Expand{
		ActingUserAccessToken: apps.ExpandAll,
		ActingUser:            apps.ExpandID,
	}),
}

package form

import "github.com/mattermost/mattermost-plugin-apps/apps"

var List = apps.Form{
	Title: "List",
	Icon:  "",
	Fields: []apps.Field{
		{
			Type:  apps.FieldTypeText,
			Name:  "test",
			Value: "test",
		},
	},
	Submit: apps.NewCall("/list").WithExpand(apps.Expand{
		ActingUserAccessToken: apps.ExpandAll,
		ActingUser:            apps.ExpandID,
	}),
}

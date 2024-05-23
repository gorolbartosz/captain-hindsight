package form

import "github.com/mattermost/mattermost-plugin-apps/apps"

var Trigger = apps.Form{
	Title: "Trigger",
	Icon:  "",
	Fields: []apps.Field{
		{
			Type:  apps.FieldTypeText,
			Name:  "test",
			Value: "test",
		},
	},
	Submit: apps.NewCall("/trigger").WithExpand(apps.Expand{
		ActingUserAccessToken: apps.ExpandAll,
		ActingUser:            apps.ExpandID,
	}),
}

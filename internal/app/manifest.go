package app

import "github.com/mattermost/mattermost-plugin-apps/apps"

var Manifest = apps.Manifest{
	AppID:       "captain-hindsight",
	Version:     "v0.0.1",
	DisplayName: "Captain Hindsight",
	Icon:        "logo.png",
	HomepageURL: "https://github.com/gorolbartosz/captain-hindsight",
	RequestedPermissions: []apps.Permission{
		apps.PermissionActAsBot,
		apps.PermissionActAsUser,
	},
	RequestedLocations: []apps.Location{
		apps.LocationCommand,
	},
	Deploy: apps.Deploy{
		HTTP: &apps.HTTP{
			RootURL: "http://127.0.0.1:8066",
		},
	},
}

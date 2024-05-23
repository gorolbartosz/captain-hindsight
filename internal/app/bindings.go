package app

import (
	"github.com/gorolbartosz/captain-hindsight/internal/form"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

var Bindings = []apps.Binding{
	{
		Location: apps.LocationCommand,
		Bindings: []apps.Binding{
			{
				Icon:        "logo.png",
				Label:       "captain",
				Description: "Captain Hindsight App",
				Hint:        "[list,show,trigger]",
				Bindings: []apps.Binding{
					{
						Label: "list",
						Form:  &form.List,
					},
					{
						Label: "show",
						Form:  &form.Show,
					},
					{
						Label: "trigger",
						Form:  &form.Trigger,
					},
				},
			},
		},
	},
}

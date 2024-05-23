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
				Hint:        "[command]",
				Bindings: []apps.Binding{
					{
						Label: "list",
						Icon:  "list.png",
						Form:  &form.List,
					},
					{
						Label: "show",
						Icon:  "show.png",
						Hint:  "[id]",
						Form:  &form.Show,
					},
					{
						Label: "trigger",
						Icon:  "trigger.png",
						Hint:  "[id]",
						Form:  &form.Trigger,
					},
				},
			},
		},
	},
}

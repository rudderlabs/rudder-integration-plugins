package sources

import (
	"github.com/google/uuid"
	"github.com/rudderlabs/rudder-plugins-manager/plugins"
)

var webHookPlugin = plugins.NewTransformPlugin("webhook", func(m *plugins.Message) (*plugins.Message, error) {
	payload := map[string]any{
		"output": map[string]any{
			"batch": []map[string]any{
				{
					"type":        "track",
					"event":       "webhook_source_event",
					"properties":  m.Data,
					"anonymousId": uuid.New().String(),
				},
			},
		},
	}
	return plugins.NewMessage(payload), nil
})

package sources

import (
	"github.com/google/uuid"
	"github.com/rudderlabs/rudder-plugins-manager/plugins"
)

var WebHookPlugin = plugins.NewTransformPlugin("webhook", func(m *plugins.Message) (*plugins.Message, error) {
	payload := map[string]any{
		"output": map[string]any{
			"type":        "track",
			"event":       "webhook_source_event",
			"properties":  m.Data,
			"anonymousId": uuid.New().String(),
			"context": map[string]any{
				"transformedBy": "webhook-plugin",
			},
		},
	}
	return plugins.NewMessage(payload), nil
})

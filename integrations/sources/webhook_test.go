package sources_test

import (
	"context"
	"testing"

	"github.com/rudderlabs/rudder-integration-plugins/integrations/sources"
	"github.com/rudderlabs/rudder-plugins-manager/plugins"
	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	webhookPlugin, err := sources.NewSourceManager().Get("webhook")
	assert.Nil(t, err)
	assert.NotNil(t, webhookPlugin)
	assert.Equal(t, "webhook", webhookPlugin.GetName())
	inputData := map[string]interface{}{
		"test": "test",
	}
	result, err := webhookPlugin.Execute(context.TODO(), plugins.NewMessage(inputData))
	assert.Nil(t, err)
	assert.NotNil(t, result)
	resultData := result.Data.(map[string]interface{})["output"].(map[string]interface{})["batch"].([]map[string]interface{})[0]
	assert.Equal(t, inputData, resultData["properties"])
	assert.Equal(t, "track", resultData["type"])
	assert.Equal(t, "webhook_source_event", resultData["event"])
	assert.NotEmpty(t, resultData["anonymousId"])
}

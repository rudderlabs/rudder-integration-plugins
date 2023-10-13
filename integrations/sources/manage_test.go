package sources_test

import (
	"testing"

	"github.com/rudderlabs/rudder-integration-plugins/integrations/sources"
	"github.com/stretchr/testify/assert"
)

func TestManager(t *testing.T) {
	manager := sources.NewSourceManager()
	assert.NotNil(t, manager)
	assert.True(t, manager.Has("webhook"))
}

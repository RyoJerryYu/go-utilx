package projectx

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGoModPath(t *testing.T) {
	goModPath, err := GetGoModPath()
	require.NoError(t, err)

	assert.NotEmpty(t, goModPath)
	assert.Equal(t, "go-utilx", filepath.Base(goModPath))
}

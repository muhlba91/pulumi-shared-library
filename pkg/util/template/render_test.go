package template_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tpl "github.com/muhlba91/pulumi-shared-library/pkg/util/template"
)

func TestRenderTemplate_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "hello.tmpl")
	err := os.WriteFile(path, []byte("Hello {{ .Name }}"), 0o644)
	require.NoError(t, err)

	out, err := tpl.RenderTemplate(path, map[string]string{"Name": "World"})
	require.NoError(t, err)
	assert.Equal(t, "Hello World", out)
}

func TestRenderTemplate_FileNotFound(t *testing.T) {
	_, err := tpl.RenderTemplate("/non/existent/path.tmpl", nil)
	require.Error(t, err)
}

func TestRenderTemplate_ParseError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.tmpl")
	err := os.WriteFile(path, []byte("Hello {{ .Name "), 0o644)
	require.NoError(t, err)

	_, err = tpl.RenderTemplate(path, map[string]string{"Name": "x"})
	require.Error(t, err)
}

func TestRenderTemplate_ExecuteTypeError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "type_err.tmpl")
	err := os.WriteFile(path, []byte("Hello {{ .Name }}"), 0o644)
	require.NoError(t, err)

	_, err = tpl.RenderTemplate(path, 123)
	require.Error(t, err)
}

package ioc_test

import (
	"testing"

	ioc "github.com/richhh7g/simple-ioc"
	"github.com/stretchr/testify/assert"
)

func TestPackageImport(t *testing.T) {
	id := ioc.ContainerTag("CONTAINER_ID")
	container := ioc.NewContainer(&id)

	container.Register("test", "test")
	assert.Equal(t, "test", container.Resolve("test").(string))
}

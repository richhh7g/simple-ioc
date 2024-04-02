package ioc

import (
	"fmt"
	"sync"
)

type Container struct {
	id           *ContainerTag
	mu           sync.Mutex
	dependencies map[string]interface{}
}

func NewContainer(id *ContainerTag) *Container {
	return &Container{
		id:           id,
		dependencies: make(map[string]interface{}),
	}
}

func (c *Container) Register(name string, dependency interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.dependencies[name] = dependency
}

func (c *Container) Resolve(name string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	dependency, ok := c.dependencies[name]
	if !ok {
		panic(fmt.Sprintf("dependency not found: %s", name))
	}

	return dependency
}

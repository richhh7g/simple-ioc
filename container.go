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

// Creates a new instance of the Container struct with the provided ContainerTag.
//
//	id := ioc.ContainerTag("CONTAINER_ID")
//	ioc.NewContainer(&id)
func NewContainer(id *ContainerTag) *Container {
	return &Container{
		id:           id,
		dependencies: make(map[string]interface{}),
	}
}

// Registers a dependency in the Container.
//
//	func NewYourType() *YourType {
//		return &YourType{}
//	}
//
//	ioc.Register("NAME_OF_DEPENDENCY", NewYourType())
func (c *Container) Register(name string, dependency interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.dependencies[name] = dependency
}

// Resolves the dependency by name.
//
//	var dependency YourType =	ioc.Resolve("NAME_OF_DEPENDENCY").(YourType)
//	fmt.Println(dependency.YourMethod())
func (c *Container) Resolve(name string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	dependency, ok := c.dependencies[name]
	if !ok {
		panic(fmt.Sprintf("dependency not found: %s", name))
	}

	return dependency
}

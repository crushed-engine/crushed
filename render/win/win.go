package win

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/manen/cleaner"
)

// init: it's in the name
func init() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	cleaner.AddCleaner(cleaner.Cleaner{
		Name: "GLFW",
		CleanUp: func() {
			glfw.Terminate()
		},
	})
}

// PollEvents does event things idk
func PollEvents() {
	glfw.PollEvents()
}

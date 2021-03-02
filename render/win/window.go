package win

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/manen/cleaner"
)

// Window represents an OS window
// Uses GLFW internally
type Window struct {
	glfwWindow *glfw.Window
	w, h       int
	t          string
}

// NewWindow creates a new window
// using GLFW
func NewWindow(w, h int, t string) *Window {
	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfwWindow, err := glfw.CreateWindow(w, h, t, nil, nil)
	if err != nil {
		panic(err)
	}

	win := &Window{
		glfwWindow,
		w,
		h,
		t,
	}

	cleaner.AddCleaner(cleaner.Cleaner{
		Name: "GLFW Window: " + t,
		CleanUp: func() {
			win.glfwWindow.Destroy()
		},
	})

	return win
}

// ShouldClose is whether the window
// is supposed to close
func (w *Window) ShouldClose() bool {
	return w.glfwWindow.ShouldClose()
}

// GetW is the width getter
func (w *Window) GetW() int {
	return w.w
}

// GetH is the height getter
func (w *Window) GetH() int {
	return w.h
}

// GetT is the title getter
func (w *Window) GetT() string {
	return w.t
}

// SetW is the width setter
func (w *Window) SetW(wi int) {
	w.w = wi
	w.glfwWindow.SetSize(wi, w.h)
}

// SetH is the height setter
func (w *Window) SetH(h int) {
	w.h = h
	w.glfwWindow.SetSize(w.w, h)
}

// SetT is the title setter
func (w *Window) SetT(t string) {
	w.t = t
	w.glfwWindow.SetTitle(t)
}

package win_test

import (
	"testing"

	"github.com/crushed/crushed/render/win"
	"github.com/manen/cleaner"
)

func TestWin(t *testing.T) {
	w := win.NewWindow(640, 480, "Test Window ")

	for !w.ShouldClose() {
		win.PollEvents()
	}

	cleaner.CleanUp()
}

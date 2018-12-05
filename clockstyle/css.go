package Clockstyle

import (
	"github.com/gotk3/gotk3/gtk"
	"fmt"
)

func AssignClassToLabel(l *gtk.Label) {
	element, _ := l.GetStyleContext()
	element.AddClass("digits")
}

func AssignClassToWindow(w *gtk.Window) {
	element, _ := w.GetStyleContext()
	element.AddClass("mywindow")
}

func Css(fontSize int) string {
	stylesheet := fmt.Sprintf(`
.mywindow {
	background-color: black;
}

.digits {
	font-family: 'DejaVu Sans Mono';
   	font-weight: bold;
   	font-style: normal;
    font-size: %dpx;
	color: #FFD454;
}
`, fontSize)

	return stylesheet
}

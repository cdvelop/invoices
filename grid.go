package invoices

import (
	"fmt"

	"github.com/signintech/gopdf"
)

func (h *handler) drawGrid() {
	w := gopdf.PageSizeLetter.W
	height := gopdf.PageSizeLetter.H
	h.pdf.SetFont(h.font_normal, "", 7)

	gridSpacing := w / 20 // Espaciado del grid

	h.pdf.SetTextColor(160, 160, 160)
	h.pdf.SetStrokeColor(235, 235, 235)
	// Dibujar líneas verticales y números en el eje x
	for x := 0.0; x <= w; x += gridSpacing {
		h.pdf.Line(x, 0, x, height)

		// Mostrar los números en la parte superior y mover un poco hacia abajo
		h.pdf.SetX(x + 1) // Ajustar la posición X del texto
		h.pdf.SetY(7)     // Ajustar la posición Y del texto (un poco más abajo)
		h.pdf.Text(fmt.Sprintf("%.1f", x))
	}

	// Dibujar líneas horizontales y números en el eje y
	for y := 0.0; y <= height; y += gridSpacing {
		h.pdf.Line(0, y, w, y)

		// Mostrar los números en el margen izquierdo y mover un poco hacia la derecha
		h.pdf.SetX(1)     // Ajustar la posición X del texto (un poco más a la derecha)
		h.pdf.SetY(y - 2) // Ajustar la posición Y del texto
		h.pdf.Text(fmt.Sprintf("%.1f", y))
	}

	// fmt.Println("gridSpacing:", gridSpacing)

}

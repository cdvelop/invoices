package pdf_invoice

import (
	"github.com/signintech/gopdf"
)

func (h *handler) generateLogo() error {

	// fijar posición inicial logo
	h.pdf.SetXY(h.margin_left, h.margin_top)

	const logoHeight = 65.0 // Altura fija en puntos

	var aspectRatio = 4.0 / 3.0 // default: 4.3

	// Maneja el aspect ratio según el valor recibido
	switch h.invoice.LogoAspectRatio {
	case 16.9:
		aspectRatio = 16.0 / 9.0
	case 1.1:
		aspectRatio = 1.0

	}

	// Cálculo del ancho en base al aspecto ratio y la altura deseada
	width := logoHeight * aspectRatio

	// logo Carga la imagen desde el archivo
	if err := h.pdf.Image(h.invoice.Logo, h.margin_top, h.margin_left, &gopdf.Rect{
		W: width,      // Ancho en puntos
		H: logoHeight, // Altura en puntos
	}); err != nil {
		return err
	}

	return nil
}

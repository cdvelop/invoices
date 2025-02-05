package invoices

import (
	"fmt"
	"strings"

	"github.com/signintech/gopdf"
)

func (h *handler) generateHeaderRight() error {

	const (
		// square_x0   = 420
		sii         = "S.I.I. - "
		invoiceType = "FACTURA ELECTRÓNICA"
	)

	if h.invoice.Type == "" { // si no se ha definido el tipo de documento, se establece por defecto a factura electrónica
		h.invoice.Type = invoiceType
	} else { // nos aseguramos que el tipo de documento este en mayúsculas
		h.invoice.Type = strings.ToUpper(h.invoice.Type)
	}

	if h.invoice.City != "" {
		// nos aseguramos que la ciudad este en mayúsculas
		h.invoice.City = strings.ToUpper(h.invoice.City)

		if h.invoice.Type != "DETALLE" {
			h.invoice.City = sii + h.invoice.City
		}

	}

	err := h.pdf.SetFont(h.font_bold, "", 11)
	if err != nil {
		return err
	}

	// Establecer color de texto en rojo (RGB)
	h.pdf.SetTextColor(255, 0, 0) // Rojo
	// establecer ancho de línea
	h.pdf.SetLineWidth(1.8)

	// INICIO RECTÁNGULO ROJO DESCRIPCIÓN
	h.pdf.SetStrokeColor(255, 0, 0)
	err = h.pdf.Rectangle(428.4, h.margin_top, 581.4, 107, "D", 0, 0)
	if err != nil {
		return err
	}
	h.pdf.SetStrokeColor(0, 0, 0)

	opt := gopdf.CellOption{
		Align: gopdf.Center | gopdf.Middle,
		// Border: gopdf.AllBorders,
		Float: gopdf.Right,
	}

	h.pdf.SetXY(419, h.margin_top+5)

	h.pdf.MultiCellWithOption(&gopdf.Rect{
		W: 170,
		H: 200,
	}, fmt.Sprintf(`R.U.T.: %s

	%s

	N° %s
	
	%s`,
		h.invoice.RUT, h.invoice.Type, h.invoice.Number, h.invoice.City), opt)

	h.setDefaultFont()

	return nil
}

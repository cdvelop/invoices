package pdf_invoice

import (
	"fmt"
	"time"
)

func (h *handler) generateHeader() error {

	if err := h.generateHeaderRight(); err != nil {
		return err
	}

	if err := h.generateHeaderCompany(); err != nil {
		return err
	}

	// Invoice details

	// Client details box
	h.pdf.Rectangle(50, 110, 500, 80, "D", 0, 0)
	h.pdf.SetXY(60, 115)
	h.pdf.Cell(nil, fmt.Sprintf("Señor(es): %s", h.invoice.ClientName))
	h.pdf.SetXY(60, 130)
	h.pdf.Cell(nil, fmt.Sprintf("RUT: %s", h.invoice.ClientRUT))
	h.pdf.SetXY(60, 145)
	h.pdf.Cell(nil, fmt.Sprintf("Dirección: %s", h.invoice.ClientAddress))
	h.pdf.SetXY(60, 160)
	h.pdf.Cell(nil, fmt.Sprintf("Comuna: %s", h.invoice.ClientRegion))
	h.pdf.SetXY(60, 175)
	h.pdf.Cell(nil, fmt.Sprintf("Giro: %s", h.invoice.BusinessLine))

	h.pdf.SetXY(300, 115)
	h.pdf.Cell(nil, fmt.Sprintf("Fecha Emisión: %s", time.Now().Format("02-01-2006")))
	h.pdf.SetXY(300, 130)
	h.pdf.Cell(nil, "Condición de Pago: CONTADO")
	h.pdf.SetXY(300, 145)
	h.pdf.Cell(nil, "Atención: ")
	h.pdf.SetXY(300, 160)
	h.pdf.Cell(nil, "Vendedor: (OFICINA)")

	// Table header
	h.pdf.SetFont("Arial", "B", 10)
	h.pdf.SetXY(50, 200)
	h.pdf.Cell(nil, "N°")
	h.pdf.SetXY(80, 200)
	h.pdf.Cell(nil, "Cantidad")
	h.pdf.SetXY(150, 200)
	h.pdf.Cell(nil, "Descripción")
	h.pdf.SetXY(450, 200)
	h.pdf.Cell(nil, "Precio U.")
	h.pdf.SetXY(520, 200)
	h.pdf.Cell(nil, "Total")

	h.pdf.Line(50, 215, 550, 215)

	return nil
}

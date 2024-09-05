package pdf_invoice

import (
	"fmt"
)

func (h *handler) generateFooter() error {
	y := 527.0 // Ajusta este valor según sea necesario para que comience justo después del cuerpo

	// Función auxiliar para alinear a la derecha
	alignRight := func(x float64, text string) float64 {
		width, _ := h.pdf.MeasureTextWidth(text)
		return x - width
	}

	// TOTAL SIN IVA
	h.pdf.SetXY(50, y)
	h.pdf.Cell(nil, fmt.Sprintf("transferir a: %s", h.invoice.TransferTo))
	h.pdf.SetXY(350, y)
	h.pdf.Cell(nil, "TOTAL SIN IVA")
	totalText := fmt.Sprintf("$%.2f", h.invoice.Total)
	h.pdf.SetX(alignRight(540, totalText))
	h.pdf.Cell(nil, totalText)
	y += 20

	// DESCUENTO
	h.pdf.SetXY(50, y)
	h.pdf.Cell(nil, fmt.Sprintf("cta: %s", h.invoice.AccountType))
	h.pdf.SetXY(350, y)
	h.pdf.Cell(nil, fmt.Sprintf("DESCUENTO %.2f %%", h.invoice.Discount))
	h.pdf.SetX(alignRight(540, "-$"))
	h.pdf.Cell(nil, "-$")
	y += 20

	// RETENCION IVA
	h.pdf.SetXY(50, y)
	h.pdf.Cell(nil, h.invoice.AccountNumber)
	h.pdf.SetXY(350, y)
	h.pdf.Cell(nil, fmt.Sprintf("RETENCION IVA %.4f", h.invoice.TaxRetention))
	retencionText := fmt.Sprintf("-$%.2f", h.invoice.Total*h.invoice.TaxRetention)
	h.pdf.SetX(alignRight(540, retencionText))
	h.pdf.Cell(nil, retencionText)
	y += 20

	// BOLETA SII
	h.pdf.SetXY(350, y)
	h.pdf.Cell(nil, "BOLETA SII")
	boletaText := fmt.Sprintf("-$%.2f", h.invoice.Total*(1+h.invoice.TaxRetention))
	h.pdf.SetX(alignRight(540, boletaText))
	h.pdf.Cell(nil, boletaText)
	y += 20

	h.pdf.Line(50, y, 550, y)
	y += 20

	// Total a Pagar
	h.pdf.SetXY(50, y)
	h.pdf.Cell(nil, "Total a Pagar Efectivo ……………………")
	totalPagarText := fmt.Sprintf("$%.2f", h.invoice.Total)
	h.pdf.SetX(alignRight(540, totalPagarText))
	h.pdf.Cell(nil, totalPagarText)

	return nil
}

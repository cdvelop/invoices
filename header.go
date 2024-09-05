package pdf_invoice

func (h *handler) generateHeader() error {

	if err := h.generateHeaderRight(); err != nil {
		return err
	}

	if err := h.generateHeaderCompany(); err != nil {
		return err
	}

	if err := h.generateHeaderDetails(); err != nil {
		return err
	}

	// Invoice details

	// Table header
	// h.pdf.SetFont("Arial", "B", 10)
	// h.pdf.SetXY(50, 200)
	// h.pdf.Cell(nil, "N°")
	// h.pdf.SetXY(80, 200)
	// h.pdf.Cell(nil, "Cantidad")
	// h.pdf.SetXY(150, 200)
	// h.pdf.Cell(nil, "Descripción")
	// h.pdf.SetXY(450, 200)
	// h.pdf.Cell(nil, "Precio U.")
	// h.pdf.SetXY(520, 200)
	// h.pdf.Cell(nil, "Total")

	// h.pdf.Line(50, 215, 550, 215)

	return nil
}

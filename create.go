package pdf_invoice

func (h *handler) generateInvoice(invoice Invoice, filename string) error {
	h.invoice = &invoice

	h.drawGrid()

	if err := h.generateHeader(); err != nil {
		return err
	}

	if err := h.generateBody(); err != nil {
		return err
	}

	if err := h.generateFooter(); err != nil {
		return err
	}
	// h.drawGrid()

	err := h.pdf.WritePdf(filename)
	if err != nil {
		return err
	}

	return nil
}

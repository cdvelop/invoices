package pdf_invoice

func (h *handler) generateHeaderDetails() error {

	// h.pdf.SetFont(h.font_normal, "", 9)

	const box_start = 135
	// INICIO RECTÁNGULO ROJO DESCRIPCIÓN
	err := h.pdf.Rectangle(h.margin_left, box_start, 581.4, 228, "D", 0, 0)
	if err != nil {
		return err
	}

	const margin = 5
	var margin_top float64 = box_start + margin
	var margin_left = h.margin_left + margin*2
	h.pdf.SetXY(margin_left, margin_top)

	labels := []string{"Señor(es)", "Giro", "Rut", "Dirección", "Comuna", "Teléfono", "Email"}
	values := []string{h.invoice.ClientName, h.invoice.ClientActivity, h.invoice.ClientRUT, h.invoice.ClientAddress, h.invoice.ClientCity, h.invoice.ClientPhone, h.invoice.ClientEmail}

	y := margin_top
	label_width := 75.0
	line_break := 12.0

	for i, label := range labels {
		/// establecer la etiqueta
		h.pdf.SetXY(margin_left, y)
		h.pdf.Cell(nil, label)

		// establecer  dos puntos :
		h.pdf.SetXY(margin_left+label_width, y)
		h.pdf.Cell(nil, ":")

		// establecer el valor
		h.pdf.SetXY(margin_left+label_width+10, y)
		h.pdf.Cell(nil, values[i])

		y += line_break
	}

	labels = []string{"Fecha Emisión", "Medio de Pago", "Vendedor"}
	values = []string{h.invoice.Date, h.invoice.PaymentMethod, h.invoice.Seller}

	y = margin_top
	margin_left = 397.8

	for i, label := range labels {
		/// establecer la etiqueta
		h.pdf.SetXY(margin_left, y)
		h.pdf.Cell(nil, label)

		// establecer  dos puntos :
		h.pdf.SetXY(margin_left+label_width, y)
		h.pdf.Cell(nil, ":")

		// establecer el valor
		h.pdf.SetXY(margin_left+label_width+10, y)
		h.pdf.Cell(nil, values[i])

		y += line_break
	}

	h.setDefaultFont()
	return nil
}

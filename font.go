package pdf_invoice

func (h *handler) setDefaultFont() {
	h.pdf.SetFont(h.font_normal, "", 10)
	// establecer color de texto en negro (RGB)
	h.pdf.SetTextColor(0, 0, 0)
	//establecer ancho de linea
	h.pdf.SetLineWidth(1)
	// establecer color de linea a negro
	h.pdf.SetStrokeColor(0, 0, 0)
}

func (h *handler) setupFont() error {
	// fontData, err := base64.StdEncoding.DecodeString(arialFont)
	// if err != nil {
	// 	return fmt.Errorf("failed to decode font data: %v", err)
	// }

	err := h.pdf.AddTTFFontData(h.font_normal, RubikRegularFont)
	if err != nil {
		return err
	}

	err = h.pdf.AddTTFFontData(h.font_bold, RubikBoldFont)
	if err != nil {
		return err
	}

	return h.pdf.SetFont(h.font_normal, "", 10)
}

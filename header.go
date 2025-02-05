package invoices

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

	return nil
}

package pdf_invoice

import (
	"fmt"
	"strings"

	"github.com/signintech/gopdf"
)

func (h *handler) generateHeaderCompany() error {

	if err := h.generateLogo(); err != nil {
		return err
	}

	const (
		// margin_left_logo = 122.4
		w_central_box = 400
	)

	h.pdf.SetXY(61.2, h.margin_top+5)

	h.pdf.SetFont(h.font_bold, "", 13)

	opt := gopdf.CellOption{
		Align: gopdf.Center | gopdf.Middle,
		// Border: gopdf.AllBorders,
		Float: gopdf.Right,
	}

	// TITLE COMPANY
	h.pdf.MultiCellWithOption(&gopdf.Rect{
		W: w_central_box,
	}, strings.ToUpper(h.invoice.Company), opt)

	// COMPANY DESCRIPTION
	h.pdf.SetFont(h.font_bold, "", 11)
	// Establecer color de texto plomo
	h.pdf.SetTextColor(120, 120, 120)

	h.pdf.MultiCellWithOption(&gopdf.Rect{
		W: w_central_box,
	}, h.invoice.Description, opt)

	h.pdf.SetFont(h.font_normal, "", 9)

	// COMPANY CONTACT
	h.pdf.MultiCellWithOption(&gopdf.Rect{
		W: w_central_box,
	}, fmt.Sprintf(`%s
		%s %s  
		%s`, h.invoice.Address, h.invoice.Domain, h.invoice.Mail, h.invoice.Phone), opt)

	h.setDefaultFont()
	return nil
}

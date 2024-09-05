package pdf_invoice

import "github.com/signintech/gopdf"

func New() *handler {

	h := &handler{
		pdf:     &gopdf.GoPdf{},
		invoice: nil,

		font_normal: "normal",
		font_bold:   "bolt",

		// font_family: "Rubik-Regular",
		margin_top:  30.6,
		margin_left: 30.6,
	}

	h.pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeLetter})
	h.pdf.AddPage()

	if err := h.setupFont(); err != nil {
		panic(err)
	}

	return h
}

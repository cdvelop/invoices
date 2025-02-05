package invoices

import (
	"fmt"

	"github.com/signintech/gopdf"
)

func (h *handler) generateBody() error {

	const table_start_y = 217.0

	table := h.pdf.NewTableLayout(h.margin_left, table_start_y, 30, 5)

	table.AddColumn("CÓDIGO", 50, "left")
	table.AddColumn("DESCRIPCIÓN", 309, "left")
	table.AddColumn("CANT.", 43, "right")
	table.AddColumn("PRECIO", 60, "right")
	table.AddColumn("DTO", 25, "right")
	table.AddColumn("TOTAL", 63, "right")

	for _, item := range h.invoice.Items {
		table.AddRow([]string{
			item.Code,
			item.Description,
			formatNumber(float64(item.Quantity)),
			formatNumber(item.Price),
			fmt.Sprintf("%v", item.Discount),
			formatNumber(item.CalculateTotal()),
		})
	}

	table.SetTableStyle(gopdf.CellStyle{
		BorderStyle: gopdf.BorderStyle{
			Top:      true,
			Left:     true,
			Bottom:   true,
			Width:    1,
			RGBColor: gopdf.RGBColor{R: 0, G: 0, B: 0},
		},
		FillColor: gopdf.RGBColor{R: 255, G: 255, B: 255},
		TextColor: gopdf.RGBColor{R: 0, G: 0, B: 0},
		FontSize:  10,
	})

	table.SetHeaderStyle(gopdf.CellStyle{
		BorderStyle: gopdf.BorderStyle{
			Top:      true,
			Left:     true,
			Right:    true,
			Bottom:   true,
			Width:    1.5,
			RGBColor: gopdf.RGBColor{R: 100, G: 150, B: 255},
		},
		FillColor: gopdf.RGBColor{R: 255, G: 200, B: 200},
		TextColor: gopdf.RGBColor{R: 255, G: 100, B: 100},
		Font:      h.font_bold,
		FontSize:  10,
	})

	table.SetCellStyle(gopdf.CellStyle{
		BorderStyle: gopdf.BorderStyle{
			Right:    true,
			Width:    1,
			RGBColor: gopdf.RGBColor{R: 0, G: 0, B: 0},
		},
		FillColor: gopdf.RGBColor{R: 255, G: 255, B: 255},
		TextColor: gopdf.RGBColor{R: 0, G: 0, B: 0},
		Font:      h.font_normal,
		FontSize:  10,
	})

	return table.DrawTable()
}

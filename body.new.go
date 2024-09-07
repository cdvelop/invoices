package pdf_invoice

import (
	"fmt"
	"strconv"
)

func (h *handler) generateBody() error {

	const body_start = 217.0

	table := NewTableLayout(h.pdf, h.margin_left, body_start, 30, 15)

	table.AddColumn("CÓDIGO", 50, "left")
	table.AddColumn("DESCRIPCIÓN", 310, "left")
	table.AddColumn("CANT.", 42, "right")
	table.AddColumn("PRECIO", 60, "right")
	table.AddColumn("DTO", 25, "right")
	table.AddColumn("TOTAL", 63, "right")

	for _, item := range h.invoice.Items {
		table.AddRow([]string{
			item.Code,
			item.Description,
			strconv.Itoa(int(item.Quantity)),
			formatNumber(item.Price),
			fmt.Sprintf("%v", item.Discount),
			formatNumber(item.CalculateTotal()),
		})
	}

	return table.Draw()
}

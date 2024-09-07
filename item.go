package pdf_invoice

func (item *InvoiceItem) CalculateTotal() float64 {
	subtotal := float64(item.Quantity) * item.Price
	discountAmount := subtotal * item.Discount / 100
	item.total = subtotal - discountAmount
	return item.total
}

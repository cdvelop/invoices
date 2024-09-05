package pdf_invoice

import (
	"fmt"
)

// generateBody crea el cuerpo del documento PDF para la factura
func (h handler) generateBody() error {

	// Establecer los encabezados de las columnas
	h.pdf.SetXY(50, 280)
	h.pdf.Cell(nil, "DETALLE")
	h.pdf.SetXY(300, 280)
	h.pdf.Cell(nil, "UNID.")
	h.pdf.SetXY(400, 280)
	h.pdf.Cell(nil, "PRECIO")
	h.pdf.SetXY(500, 280)
	h.pdf.Cell(nil, "TOTAL")

	// Dibujar una línea separadora
	h.pdf.Line(50, 300, 550, 300)

	// Inicializar la posición vertical para los elementos
	y := 320.0

	// Iterar sobre los elementos de la factura
	for _, item := range h.invoice.Items {
		// Establecer la descripción del artículo
		h.pdf.SetXY(50, y)
		h.pdf.Cell(nil, item.Description)

		// Establecer la cantidad
		h.pdf.SetXY(300, y)
		h.pdf.Cell(nil, fmt.Sprintf("%d", item.Quantity))

		// Establecer el precio unitario
		h.pdf.SetXY(400, y)
		h.pdf.Cell(nil, fmt.Sprintf("$ %.2f", item.Price))

		// Calcular y establecer el total para este artículo
		h.pdf.SetXY(500, y)
		h.pdf.Cell(nil, fmt.Sprintf("$ %.2f", float64(item.Quantity)*item.Price))

		// Incrementar la posición vertical para el siguiente artículo
		y += 20
	}

	// Rellenar las filas vacías hasta completar 10 filas
	for i := len(h.invoice.Items); i < 10; i++ {
		h.pdf.SetXY(500, y)
		h.pdf.Cell(nil, "$0")
		y += 20
	}

	// Dibujar una línea final
	h.pdf.Line(50, y, 550, y)

	return nil
}

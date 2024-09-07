package pdf_invoice

import (
	"github.com/signintech/gopdf"
)

type TableLayout struct {
	pdf        *gopdf.GoPdf
	startX     float64
	startY     float64
	rowHeight  float64
	columns    []column
	rows       [][]string
	maxRows    int
	padding    float64 // Nuevo campo para el padding
	cellOption gopdf.CellOption
}

type column struct {
	header string
	width  float64
	align  string
}

func NewTableLayout(pdf *gopdf.GoPdf, startX, startY, rowHeight float64, maxRows int) *TableLayout {
	return &TableLayout{
		pdf:        pdf,
		startX:     startX,
		startY:     startY,
		rowHeight:  rowHeight,
		maxRows:    maxRows,
		padding:    2.0, // Nuevo campo para el padding
		cellOption: gopdf.CellOption{
			// Align:  gopdf.Middle,
			// Border: gopdf.Left | gopdf.Top | gopdf.Right | gopdf.Bottom,
		},
	}
}

// align: left, right, center
func (t *TableLayout) AddColumn(header string, width float64, align string) {
	t.columns = append(t.columns, column{header, width, align})
}

func (t *TableLayout) AddRow(row []string) {
	t.rows = append(t.rows, row)
}

func (t *TableLayout) Draw() error {
	x := t.startX
	y := t.startY

	// Draw headers
	for _, col := range t.columns {
		if err := t.drawCell(x, y, col.width, t.rowHeight, col.header, "center", true); err != nil {
			return err
		}

		x += col.width
	}
	y += t.rowHeight

	// Draw rows
	for _, row := range t.rows {
		x = t.startX
		for i, cell := range row {
			// var padding_left string
			if err := t.drawCell(x, y, t.columns[i].width, t.rowHeight, cell, t.columns[i].align, false); err != nil {
				return err
			}
			x += t.columns[i].width
		}
		y += t.rowHeight
	}

	// Fill empty rows
	for i := len(t.rows); i < t.maxRows; i++ {
		x = t.startX
		for _, col := range t.columns {
			if err := t.drawCell(x, y, col.width, t.rowHeight, " ", col.align, false); err != nil {
				return err
			}
			x += col.width
		}
		y += t.rowHeight
	}

	return nil
}

func (t *TableLayout) drawCell(x, y, width, height float64, content, align string, isHeader bool) error {
	// Dibujar el borde de la celda

	x1_rect := x + width
	y1_rect := y + height

	if err := t.pdf.Rectangle(x, y, x1_rect, y1_rect, "D", 0, 0); err != nil {
		return err
	}

	// Calcular la posición del texto
	textX := x + t.padding
	textY := y + t.padding
	textWidth := width - (2 * t.padding)
	textHeight := height - (2 * t.padding)

	t.pdf.SetXY(textX, textY)

	// Configurar la alineación del texto
	var textOption gopdf.CellOption
	if align == "right" {
		textOption.Align = gopdf.Right | gopdf.Middle
	} else if align == "center" {
		textOption.Align = gopdf.Center | gopdf.Middle
	} else {
		textOption.Align = gopdf.Left | gopdf.Middle
	}

	// Dibujar el texto
	if isHeader {
		t.pdf.SetTextColor(0, 0, 0)       // Color negro para el encabezado
		t.pdf.SetFillColor(240, 240, 240) // Fondo gris claro para el encabezado
		// t.pdf.RectFromUpperLeftWithStyle(x, y, width, height, "F")
	} else {
		t.pdf.SetTextColor(0, 0, 0) // Color negro para el contenido
	}

	err := t.pdf.MultiCellWithOption(&gopdf.Rect{W: textWidth, H: textHeight}, content, textOption)
	if err != nil && err.Error() == "empty string" {
		// Manejar el error de cadena vacía, por ejemplo, no hacer nada o registrarlo
		err = nil
	}

	return err
}

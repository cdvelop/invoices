package pdf_invoice

import (
	"bytes"
	"math"
	"strconv"
	"strings"
)

const (
	maxFloat64             = 0x1p1023 * (1 + (1 - 0x1p-52))
	smallestNonzeroFloat64 = 0x1p-1022 * 0x1p-52
)

// e.g. Commaf(834142.32) -> 834,142.32
func commaf(v float64, dot ...bool) string {
	comma := []byte{','}
	for _, v := range dot {
		if v {
			comma = []byte{'.'}
		}
	}
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}

func formatNumber(val float64) string {

	return commaf(roundChileanNumber(val), true)
}

// redondea un número según la norma chilena.
func roundChileanNumber(number float64) float64 {
	// Obtener la parte decimal del número.
	decimal := number - math.Floor(number)

	// Redondear la parte decimal a dos decimales.
	round := math.Round(decimal*100) / 100

	// Calcular el número redondeado.
	roundNumber := math.Floor(number) + round

	return roundNumber
}

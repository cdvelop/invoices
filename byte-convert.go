package pdf_invoice

import (
	"fmt"
	"os"
	"strings"
)

// ConvertFileToBytes convierte un archivo a un array de bytes y lo almacena en otro archivo como una variable Go.
func ConvertFileToBytes(inputFileName, outputFileName, packageName, varName string) error {
	// Leer el archivo y convertirlo en un array de bytes
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %w", err)
	}

	// Convertir los bytes a una representación en cadena con comas
	byteStrings := make([]string, len(data))
	for i, b := range data {
		byteStrings[i] = fmt.Sprintf("%d", b)
	}

	// Dividir la cadena en líneas más cortas para mejor legibilidad
	const bytesPerLine = 20
	var lines []string
	for i := 0; i < len(byteStrings); i += bytesPerLine {
		end := i + bytesPerLine
		if end > len(byteStrings) {
			end = len(byteStrings)
		}
		lines = append(lines, strings.Join(byteStrings[i:end], ", "))
	}

	// Unir todas las líneas en un solo bloque de código
	output := fmt.Sprintf("package %s\n\nvar %s = []byte{\n%s,\n}\n", packageName, varName, strings.Join(lines, ",\n"))

	// Escribir el contenido en el archivo de salida
	err = os.WriteFile(outputFileName, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("error al escribir el archivo: %w", err)
	}

	return nil
}

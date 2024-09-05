package pdf_invoice

// _ "image/jpeg" // Importa el decodificador JPEG
// _ "image/png"  // Importa el decodificador PNG

func (h *handler) insertImageProportionally(x, y float64, height float64) error {
	// // Abre el archivo de imagen
	// file, err := os.Open(imagePath)
	// if err != nil {
	//     return err
	// }
	// defer file.Close()

	// // Decodifica la imagen para obtener sus dimensiones originales
	// img, _, err := image.Decode(file)
	// if err != nil {
	//     return err
	// }

	// // Calcula la relación de aspecto
	// aspectRatio := float64(img.Bounds().Dx()) / float64(img.Bounds().Dy())

	// // Calcula el ancho basado en la altura deseada y la relación de aspecto
	// width := height * aspectRatio

	// // Inserta la imagen con las dimensiones calculadas
	// if err := h.pdf.Image(imagePath, x, y, &gopdf.Rect{
	//     W: width,
	//     H: height,
	// }); err != nil {
	//     return err
	// }

	return nil
}

// downloadImage descarga una imagen desde una URL y la guarda en un archivo local.
func (h *handler) downloadImage(url, filepath string) error {
	// // Solicitar la imagen
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	// // Crear el archivo
	// file, err := os.Create(filepath)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// // Copiar el contenido de la respuesta al archivo
	// _, err = io.Copy(file, resp.Body)
	// return err
	return nil
}

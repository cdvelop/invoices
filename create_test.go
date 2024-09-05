package pdf_invoice

import (
	"os"
	"testing"
)

func TestGenerateInvoice(t *testing.T) {

	const outFileName = "examples/invoice_result.pdf"

	// Setup test data
	testInvoice := Invoice{
		RUT:             "90.236.18-0",
		Type:            "",
		Number:          "258",
		City:            "Santiago Poniente",
		Date:            "21-04-2023",
		Logo:            "examples/logo.png",
		LogoAspectRatio: 16.9,
		// Logo:               "https://www.pngitem.com/pimgs/b/485-4852378_sample-png.png",
		Company:        "solutions software spa",
		Description:    "servicios de creación de software",
		Address:        "santa marta 245. talca",
		Domain:         "solutions.com",
		Mail:           "ventas@solutions.com",
		Phone:          "+569 9999 9999",
		ClientName:     "Test Client",
		ClientActivity: "venta de verduras",
		ClientRUT:      "13.076.851-2",
		ClientAddress:  "123 Test St",
		ClientCity:     "Test Region",
		ClientPhone:    "123-456-7890",
		ClientEmail:    "test@example.com",
		Items: []InvoiceItem{
			{Description: "Test Item 1", Quantity: 3, Price: 100},
			{Description: "Test Item 2", Quantity: 1, Price: 50},
			{Description: "Test Item 3", Quantity: 10, Price: 1000},
		},
		Total: 250,
	}

	h := New()

	// Call the function to generate the invoice
	err := h.generateInvoice(testInvoice, outFileName)
	if err != nil {
		t.Fatalf("Failed to generate invoice: %v", err)
	}

	// Check if the file was created
	_, err = os.Stat(outFileName)
	if os.IsNotExist(err) {
		t.Fatalf("Invoice PDF was not created")
	}

	// Clean up the test file
	// err = os.Remove(outFileName)
	// if err != nil {
	// 	t.Fatalf("Failed to remove test invoice file: %v", err)
	// }
}

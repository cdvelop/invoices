package pdf_invoice

import (
	"github.com/signintech/gopdf"
)

type handler struct {
	pdf     *gopdf.GoPdf
	invoice *Invoice

	font_normal string // Rubik, Arial, Times New Roman, Courier New
	font_bold   string // Rubik-Bold
	margin_top  float64
	margin_left float64
}

type Invoice struct {
	RUT             string //rut de la empresa remitente ej: 90.236.18-0
	Type            string //Detalle, Boleta, Nota de Crédito, Nota de Débito, default: Factura Electrónica
	Number          string
	City            string // Santiago Poniente, Santiago,Chillan, Valparaiso,
	Date            string
	PaymentMethod   string  // medio de pago
	Seller          string  // vendedor
	Logo            string  // ej: logo.png
	LogoAspectRatio float32 // ej: 1.1, 16.9 default: 4.3
	Company         string  // ej: solutions software spa
	Description     string  // ej: servicios de creación de software
	Address         string  // ej: santa marta 245. talca
	Domain          string  // ej: solutions.com
	Mail            string  // ej: ventas@solutions.com
	Phone           string
	ClientName      string
	ClientActivity  string
	ClientRUT       string
	ClientAddress   string
	ClientCity      string
	ClientPhone     string
	ClientEmail     string
	Items           []InvoiceItem
	Total           float64
	TransferTo      string
	AccountType     string
	AccountNumber   string
	Discount        float32
	TaxRetention    float32
	BillNumber      string
}

type InvoiceItem struct {
	Code        string  //código
	Description string  // item x
	Quantity    uint32  // cantidad (0 to 4294967295)
	Price       float64 // precio ej: 1000
	Discount    float64 //descuento en % ej: 30,15
	total       float64
}

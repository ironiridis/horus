package main

type InvoiceItem struct {
	Description string
	Quantity    float64
	CostPer     int64
}

func (i *InvoiceItem) Value() int64 {
	return int64(i.Quantity * float64(i.CostPer))
}

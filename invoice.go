package main

type Invoice struct {
	Format    string
	Number    uint
	Revision  uint
	Currency  Currency
	Paid      int64
	Issuer    Issuer
	Recipient Recipient
	Issued    InvoiceDate
	Revised   InvoiceDate
	DueDate   InvoiceDate
	Items     []InvoiceItem
}

func (i *Invoice) Subtotal() int64 {
	var t int64
	for j := range i.Items {
		t += i.Items[j].Value()
	}
	return t
}

func (i *Invoice) AmountDue() int64 {
	return i.Subtotal() - i.Paid
}

package main

import "testing"

func currencyfmthelper(t *testing.T, cur string, ex []string) {
	for i := range valueCases {
		f := Currency(cur).FormatValue(valueCases[i])
		if f != ex[i] {
			t.Errorf("For currency %q, value=%d (case %d), wanted %q, but got %q\n", cur, valueCases[i], i, ex[i], f)
		}
	}
}

func currencydisphelpter(t *testing.T, cur string, name string, sym string) {
	n := Currency(cur).Name()
	if n != name {
		t.Errorf("For currency %q name, wanted %q but got %q\n", cur, name, n)
	}
	s := Currency(cur).Symbol()
	if s != sym {
		t.Errorf("For currency %q symbol, wanted %q but got %q\n", cur, sym, s)
	}
}

// all currency operations are on the fractional value of the currency. hence
// this slice represents cents for USD, yen for JPY, pence for GBP, etc.
var valueCases []int64 = []int64{
	0, -1, 1,
	-99, 99,
	-100, 100,
	-1000, 1000,
	-9999, 9999}

func TestUSD(t *testing.T) {
	currencyfmthelper(t, "USD", []string{
		"$0.00", "-$0.01", "$0.01",
		"-$0.99", "$0.99",
		"-$1.00", "$1.00",
		"-$10.00", "$10.00",
		"-$99.99", "$99.99"})
	currencydisphelpter(t, "USD", "United States dollar", `$`)
}

func TestEUR(t *testing.T) {
	currencyfmthelper(t, "EUR", []string{
		"€0.00", "-€0.01", "€0.01",
		"-€0.99", "€0.99",
		"-€1.00", "€1.00",
		"-€10.00", "€10.00",
		"-€99.99", "€99.99"})
	currencydisphelpter(t, "EUR", "Euro", `€`)
}

func TestJPY(t *testing.T) {
	currencyfmthelper(t, "JPY", []string{
		"¥0", "-¥1", "¥1",
		"-¥99", "¥99",
		"-¥100", "¥100",
		"-¥1000", "¥1000",
		"-¥9999", "¥9999"})
	currencydisphelpter(t, "JPY", "Japanese yen", `¥`)

}

func TestGBP(t *testing.T) {
	currencyfmthelper(t, "GBP", []string{
		"£0.00", "-£0.01", "£0.01",
		"-£0.99", "£0.99",
		"-£1.00", "£1.00",
		"-£10.00", "£10.00",
		"-£99.99", "£99.99"})
	currencydisphelpter(t, "GBP", "Pound sterling", `£`)
}

func TestCAD(t *testing.T) {
	currencyfmthelper(t, "CAD", []string{
		"$0.00", "-$0.01", "$0.01",
		"-$0.99", "$0.99",
		"-$1.00", "$1.00",
		"-$10.00", "$10.00",
		"-$99.99", "$99.99"})
	currencydisphelpter(t, "CAD", "Canadian dollar", `$`)
}

func TestIMP(t *testing.T) {
	currencyfmthelper(t, "IMP", []string{
		"£0.00", "-£0.01", "£0.01",
		"-£0.99", "£0.99",
		"-£1.00", "£1.00",
		"-£10.00", "£10.00",
		"-£99.99", "£99.99"})
	currencydisphelpter(t, "IMP", "Manx pound", `£`)
}

func TestXTE(t *testing.T) {
	// specifically test unknown/undefined currency
	currencyfmthelper(t, "XTE", []string{
		"¤0", "-¤1", "¤1",
		"-¤99", "¤99",
		"-¤100", "¤100",
		"-¤1000", "¤1000",
		"-¤9999", "¤9999"})
	currencydisphelpter(t, "XTE", "XTE", `¤`)
}

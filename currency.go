package main

import "fmt"

// Currency represents a common and modern currency system represented by a three character symbol such
// as "USD" or "JPY".
type Currency string

// BaseDiv returns the base unit value divisor. For instance, USD is 100 because the base unit is cents
// and the normal unit is dollars, and JPY is 1 because the base and nomnial units are both yen.
// Note this has practical limits; JPY does technically have subdivisions, but they're historical.
func (c Currency) BaseDiv() int64 {
	switch c {
	case "USD", "EUR", "GBP", "CAD", "IMP":
		return 100
	}
	return 1 // assume 1:1 base:nominal
}

// Symbol is the shorthand typographic notation for the specified Currency.
func (c Currency) Symbol() string {
	switch c {
	case "USD", "CAD":
		return `$`
	case "EUR":
		return `€`
	case "JPY":
		return `¥`
	case "GBP", "IMP":
		return `£`
	}
	return `¤` // https://en.wikipedia.org/wiki/Currency_sign_(typography)
}

func (c Currency) Name() string {
	switch c {
	case "USD":
		return "United States dollar"
	case "EUR":
		return "Euro"
	case "JPY":
		return "Japanese yen"
	case "GBP":
		return "Pound sterling"
	case "CAD":
		return "Canadian dollar"
	case "IMP":
		return "Manx pound" // considered "Clive bucks" but 🤷‍♂️
	}
	return c.String()
}

func (c Currency) String() string {
	return string(c)
}

func (c Currency) decimalShift(v int64) (base int64, frac int64) {
	div := c.BaseDiv()
	if div == 1 {
		base = v
		frac = 0
		return
	}
	if v < 0 {
		frac = (0 - v) % div
		base = (v + frac) / div
	} else {
		frac = v % div
		base = (v - frac) / div
	}
	return
}

func (c Currency) fmtSignSymBaseDotFrac(v int64) string {
	base, frac := c.decimalShift(v)
	if v < 0 {
		return fmt.Sprintf("-%s%d.%02d", c.Symbol(), (0 - base), frac)
	}
	return fmt.Sprintf("%s%d.%02d", c.Symbol(), base, frac)
}

func (c Currency) fmtSignSymValue(v int64) string {
	if v < 0 {
		return fmt.Sprintf("-%s%d", c.Symbol(), (0 - v))
	}
	return fmt.Sprintf("%s%d", c.Symbol(), v)
}

func (c Currency) FormatValue(v int64) string {
	// other special cases can be caught here
	if c.BaseDiv() > 1 {
		return c.fmtSignSymBaseDotFrac(v)
	}
	return c.fmtSignSymValue(v)
}

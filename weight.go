package unitconverter

import "fmt"

// Kilogram weight
type Kilogram float64

// Pound weight
type Pound float64

// Stone weight
type Stone float64

// KilogramToPound converts Kilogram weight to Pound
func KilogramToPound(k Kilogram) Pound { return Pound(k * 2.2) }

// KilogramToStone converts Kilogram weight to Stone
func KilogramToStone(k Kilogram) Stone { return Stone(k * 0.157473) }

// PoundToKilogram converts Pound weight to Kilogram
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p / 2.2) }

// PoundToStone converts Pound weight to Stone
func PoundToStone(p Pound) Stone { return Stone(p * 0.0714286) }

// StoneToKilogram converts Stone weight to Kilogram
func StoneToKilogram(s Stone) Kilogram { return Kilogram(s / 0.157473) }

// StoneToPound converts Stone weight to Pound
func StoneToPound(s Stone) Pound { return Pound(s / 0.0714286) }

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glbs", p) }
func (s Stone) String() string    { return fmt.Sprintf("%gst", s) }

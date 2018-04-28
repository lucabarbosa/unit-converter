package unitconverter

import "fmt"

// Inch length
type Inch float64

// Feet length
type Feet float64

// Meter lenght
type Meter float64

// InchToFeet converts Inch length to Feet
func InchToFeet(i Inch) Feet { return Feet(i / 12) }

// InchToMeter converts Inch length to Meter
func InchToMeter(i Inch) Meter { return Meter(i * 0.0254) }

// FeetToInch converts Feet length to Inch
func FeetToInch(f Feet) Inch { return Inch(f * 12) }

// FeetToMeter converts Feet length to Meter
func FeetToMeter(f Feet) Meter { return Meter(f * 0.3048) }

// MeterToFeet converts Meter length to Feet
func MeterToFeet(m Meter) Feet { return Feet(m / 0.3048) }

// MeterToInch converts Meter length to Inch
func MeterToInch(m Meter) Inch { return Inch(m / 0.0254) }

func (i Inch) String() string  { return fmt.Sprintf("%g\"", i) }
func (f Feet) String() string  { return fmt.Sprintf("%g'", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

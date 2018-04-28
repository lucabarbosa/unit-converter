package unitconverter

import "fmt"

// Celsius temperature
type Celsius float64

// Fahrenheit temperature
type Fahrenheit float64

// Kelvin temperature
type Kelvin float64

// CelsiusToFahrenheit converts Celsius temperature to Fahrenheit
func CelsiusToFahrenheit(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CelsiusToKelvin converts Celsius temperature to Kelvin
func CelsiusToKelvin(c Celsius) Kelvin { return Kelvin(c + 273) }

// FahrenheitToCelsius converts Fahrenheit temperature to Celsius
func FahrenheitToCelsius(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FahrenheitToKelvin converts Fahrenheit temperature to Kelvin
func FahrenheitToKelvin(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9 + 273.15) }

// KelvinToCelsius converts Kelvin temperature to Celsius
func KelvinToCelsius(k Kelvin) Celsius { return Celsius(k - 273) }

// KelvinToFahrenheit converts Kelvin temperature to Fahrenheit
func KelvinToFahrenheit(k Kelvin) Fahrenheit { return Fahrenheit((k-273)*9/5 + 32) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

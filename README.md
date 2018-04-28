# UnitConverter

It's a Go package to convert units of measure.  

## Getting Started

These instructions will make you have a copy of this package on your machine and import into your program.  

### Installing

Use this command in your Command Line Tool:  

```
go get -u "github.com/lucabarbosa/unitconverter"
```

This will make a copy of the package in your Go workspace.  

### Usage

Add the following line to your Go file.

```go
import "github.com/lucabarbosa/unitconverter"
```

## Units of Measure

This package covers the following units os measure:  

- Length  
  - Inches  
  - Feets  
  - Meters  
- Temperature  
  - Celsius  
  - Fahrenheit  
  - Kelvin  
- Weight
  - Kilograms  
  - Pounds  
  - Stones  

This package allows you to convert each of the measurement units to each other (provided they are in the same category).  

## Demonstration

### Length

This will convert a Meter Length to Feets:   

```go
meterLength := unitconverter.Meter(21)
feetLength := unitconverter.MeterToFeet(meterLength)

// output: 68.89763779527559'
```
### Temperature

This will convert a Celsius Temperature to Kelvin:  

```go
celsiusTemperature := unitconverter.Celsius(35)
kelvinTemperature := unitconverter.CelsiusToKelvin(celsiusTemperature)

// output: 308°K
```

### Weight

This will convert a Kilogram Wight to Pound:  

```go
kilogramWeight := unitconverter.Kilogram(75)
poundWeight := unitconverter.KilogramToPound(kilogramWeight)

// output: 165 lbs
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


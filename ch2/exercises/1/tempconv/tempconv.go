// Convert among Celsius, Fahrenheit and Kalvin
package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kalvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", float64(c))
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.1f°F", float64(f))
}

func (k Kalvin) String() string {
	return fmt.Sprintf("%0.1f°K", float64(k))
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func CToK(c Celsius) Kalvin {
	return Kalvin(c + 273.15)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FToK(f Fahrenheit) Kalvin {
	return Kalvin((f-32)*5/9 + 273.15)
}

func KToC(k Kalvin) Celsius {
	return Celsius(k - 273.15)
}

func KToF(k Kalvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 - 32)
}

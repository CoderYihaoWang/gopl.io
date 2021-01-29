// Package lenconv converts between inches and centimeters
package lenconv

import (
	"fmt"
)

// Inch represents for an inch
type Inch float64

// CM represents a centimeter
type CM float64

func (i Inch) String() string {
	return fmt.Sprintf("%.2f in", float64(i))
}

func (m CM) String() string {
	return fmt.Sprintf("%.2f cm", float64(m))
}

// IToCM converts inches to meters
func IToCM(i Inch) CM {
	return CM(float64(i) * 2.54)
}

// CMToI converts meters to inches
func CMToI(cm CM) Inch {
	return Inch(float64(cm) / 2.54)
}

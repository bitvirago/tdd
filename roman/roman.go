package roman

var decimal = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

type Converter struct {
}

func (c Converter) Convert(arabic int) string {
	mod := arabic % 10
	//count := math.Floor(float64(arabic / 10))

	return decimal[mod]
}

package EhlersTAIndicators

import (
	"fmt"
	"github.com/MathisWellmann/timeseries_generator"
	"testing"
)

func TestReFlex(t *testing.T) {
	vals := timeseries_generator.GaussianProcess(1024)
	reflex := ReFlex(vals, 16)
	for i := 0; i < len(reflex); i++ {
		fmt.Sprintf("reflex[%d]: %f\n", i, reflex)
	}
}

func TestReFlexGraph(t *testing.T) {
	vals := timeseries_generator.GaussianProcess(1024)
	reflex := ReFlex(vals, 16)

	filename := fmt.Sprintf("./img/re_flex.png")
	err := Plt(reflex, filename)
	if err != nil {
		t.Error(err)
	}
}

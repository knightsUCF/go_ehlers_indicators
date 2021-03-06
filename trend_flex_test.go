package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestTrendFlex(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	tf := TrendFlex(vals, 16)
	for i := 0; i < len(tf); i++ {
		fmt.Printf("trendflex[%d]: %f\n", i, tf[i])
	}
}

func TestTrendFlexGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	tf := TrendFlex(vals, 16)

	filename := fmt.Sprintf("./img/trend_flex.png")
	err := Plt(tf, filename)
	if err != nil {
		t.Error(err)
	}
}

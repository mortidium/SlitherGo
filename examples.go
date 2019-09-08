package main

// Config keeps config for a board
type Config struct {
	shape    int
	baseSide float64
	width    float64
	height   float64
	proto    []string
}

func loadExamples() map[string]*Config {
	return map[string]*Config{
		"SqSmHa1": &Config{
			shape:    square,
			baseSide: 70,
			width:    690,
			height:   690,
			proto: []string{
				"_011_2_",
				"3_22__2",
				"1__22_2",
				"____11_",
				"_12___2",
				"3_3_2_2",
				"3132__3",
			},
		},
		"SqSmHa2": &Config{
			shape:    square,
			baseSide: 70,
			width:    690,
			height:   690,
			proto: []string{
				"32__3__",
				"3_32_21",
				"_1__2__",
				"__32_3_",
				"_12___2",
				"11_2_22",
				"_21_32_",
			},
		},
		"HxSmHa1": &Config{
			shape:    hex,
			baseSide: 60,
			width:    690,
			height:   690,
			proto: []string{
				"___",
				"24_2",
				"143__",
				"53__",
				"_4_",
			},
		},
	}

}

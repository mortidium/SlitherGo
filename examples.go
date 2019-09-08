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
		"SqSmEz1": &Config{
			shape:    square,
			baseSide: 70,
			width:    690,
			height:   690,
			proto: []string{
				"32222_3",
				"3__11_2",
				"_____11",
				"_____21",
				"101312_",
				"_3_21_2",
				"___322_",
			},
		},

		"SqMdEz1": &Config{
			shape:    square,
			baseSide: 60,
			width:    100 + 14*60,
			height:   100 + 60*6,
			proto: []string{
				"_1232_2_322_3_",
				"3___2_2__2_12_",
				"212_2122___32_",
				"12___3_121___3",
				"__21_1_32___22",
				"2__2____22___3",
			},
		},

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
		"SqSmMd1": &Config{
			shape:    square,
			baseSide: 70,
			width:    600,
			height:   600,
			proto: []string{
				"____0_",
				"33__1_",
				"__12__",
				"__20__",
				"_1__11",
				"_2____",
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
		"HxSmEz1": &Config{
			shape:    hex,
			baseSide: 60,
			width:    690,
			height:   690,
			proto: []string{
				"_32",
				"__02",
				"__3_4",
				"__4_",
				"___",
			},
		},
		"HxSmNm1": &Config{
			shape:    hex,
			baseSide: 60,
			width:    690,
			height:   690,
			proto: []string{
				"__4",
				"__3_",
				"__03_",
				"_3__",
				"__2",
			},
		},
		"HxMdEz1": &Config{
			shape:    hex,
			baseSide: 60,
			width:    1000,
			height:   950,
			proto: []string{
				"___4_",
				"_44_5_",
				"3544_33",
				"_3_33332",
				"_43542__0",
				"__434__2",
				"34__43_",
				"__33_2",
				"_____",
			},
		},
		"HxMdNm1": &Config{
			shape:    hex,
			baseSide: 60,
			width:    1000,
			height:   950,
			proto: []string{
				"520_0",
				"____2_",
				"_4354__",
				"___3433_",
				"_4_4__3__",
				"_3__5__4",
				"_44_45_",
				"4_45_5",
				"5____",
			},
		},
	}

}

package main

var (
	examplesMediumEasy = &Examples{
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
	}

	examplesMediumNormal = &Examples{
		"SqMdNm1": &Config{
			shape:    square,
			baseSide: 60,
			width:    700,
			height:   700,
			proto: []string{
				"1___3_312_",
				"__2_1_2__2",
				"2_3_3_____",
				"_2___2___3",
				"___2301132",
				"22_2_111__",
				"_1_____3_3",
				"223_3__212",
				"_3_____2__",
				"__0_3_23_3",
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
)

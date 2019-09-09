package main

var (
	examplesSmallEasy = &Examples{
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

		"SqSmEz2": &Config{
			shape:    square,
			baseSide: 70,
			width:    690,
			height:   690,
			proto: []string{
				"_10__33",
				"__02__1",
				"1_3_23_",
				"21___11",
				"23231_2",
				"3_12___",
				"2_3__2_",
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
		"HxSmEz2": &Config{
			shape:    hex,
			baseSide: 60,
			width:    690,
			height:   690,
			proto: []string{
				"_4_",
				"____",
				"_4_30",
				"_45_",
				"_2_",
			},
		},
	}

	examplesSmallNormal = &Examples{
		"SqSmNm1": &Config{
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
		"SqSmNm2": &Config{
			shape:    square,
			baseSide: 60,
			width:    680,
			height:   680,
			proto: []string{
				"__3__11",
				"2__32_3",
				"__211__",
				"_3__3_1",
				"_20__12",
				"3_3___3",
				"___232_",
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
		"HxSmNm2": &Config{
			shape:    hex,
			baseSide: 60,
			width:    690,
			height:   690,
			proto: []string{
				"_34",
				"__24",
				"5_1_4",
				"__02",
				"4_3",
			},
		},
	}

	examplesSmallHard = &Examples{
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
		"HxSmHa2": &Config{
			shape:    hex,
			baseSide: 60,
			width:    690,
			height:   690,
			proto: []string{
				"__4",
				"_4_3",
				"_44__",
				"__44",
				"521",
			},
		},
	}
)

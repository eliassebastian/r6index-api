package ubisoft

type season struct {
	Name string
	Hex  string
	Code string
	Date string
}

var seasons = map[int]season{
	6: {
		Name: "Health",
		Code: "Y2S2",
		Hex:  "#0050b3",
		Date: "2017-06-07",
	},
	7: {
		Name: "Blood Orchid",
		Hex:  "#ca361c",
		Code: "Y2S3",
		Date: "2017-09-05",
	},
	8: {
		Name: "White Noise",
		Hex:  "#006543",
		Code: "Y2S4",
		Date: "2017-12-05",
	},
	9: {
		Name: "Chimera",
		Hex:  "#ffc113",
		Code: "Y3S1",
		Date: "2018-03-06",
	},
	10: {
		Name: "Para Bellum",
		Hex:  "#949f39",
		Code: "Y3S2",
		Date: "2018-06-07",
	},
	11: {
		Name: "Grim Sky",
		Hex:  "#81a0c1",
		Code: "Y3S3",
		Date: "2018-09-04",
	},
	12: {
		Name: "Wind Bastion",
		Hex:  "#aa854f",
		Code: "Y3S4",
		Date: "2018-12-04",
	},
	13: {
		Name: "Burnt Horizon",
		Hex:  "#d2005a",
		Code: "Y4S1",
		Date: "2019-03-06",
	},
	14: {
		Name: "Phantom Sight",
		Hex:  "#304395",
		Code: "Y4S2",
		Date: "2019-06-11",
	},
	15: {
		Name: "Ember Rise",
		Hex:  "#156309",
		Code: "Y4S3",
		Date: "2019-09-11",
	},
	16: {
		Name: "Shifting Tides",
		Hex:  "#089eb3",
		Code: "Y4S4",
		Date: "2019-12-03",
	},
	17: {
		Name: "Void Edge",
		Hex:  "#946a97",
		Code: "Y5S1",
		Date: "2020-03-10",
	},
	18: {
		Name: "Steel Wave",
		Hex:  "#2b7f9b",
		Code: "Y5S2",
		Date: "2020-06-16",
	},
	19: {
		Name: "Shadow Legacy",
		Hex:  "#6ca511",
		Code: "Y5S3",
		Date: "2020-09-10",
	},
	20: {
		Name: "Neon Dawn",
		Hex:  "#d14007",
		Code: "Y5S4",
		Date: "2020-12-01",
	},
	21: {
		Name: "Crimson Heist",
		Hex:  "#ac0000",
		Code: "Y6S1",
		Date: "2021-03-16",
	},
	22: {
		Name: "North Star",
		Hex:  "#009cbe",
		Code: "Y6S2",
		Date: "2021-06-14",
	},
	23: {
		Name: "Crystal Guard",
		Hex:  "#ffa200",
		Code: "Y6S3",
		Date: "2021-09-07",
	},
	24: {
		Name: "High Calibre",
		Hex:  "#587624",
		Code: "Y6S4",
		Date: "2021-11-30",
	},
	25: {
		Name: "Demon Veil",
		Hex:  "#ffb52c",
		Code: "Y7S1",
		Date: "2022-03-15",
	},
	26: {
		Name: "Vector Glare",
		Hex:  "#60cdb0",
		Code: "Y7S2",
		Date: "2022-06-14",
	},
	27: {
		Name: "Brutal Swarm",
		Hex:  "#dac925",
		Code: "Y7S3",
		Date: "2022-09-14",
	},
	28: {
		Name: "Solar Raid",
		Hex:  "#d03314",
		Code: "Y7S4",
		Date: "2022-12-07",
	},
}

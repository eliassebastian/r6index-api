package ubisoft

type rankInfo struct {
	Name   string
	MinMMR int16
	MaxMMR int16
}

var ranksV3 = map[int16]rankInfo{
	0:  {"Unranked", 0, 0},
	1:  {"Copper 4", 1, 1399},
	2:  {"Copper 3", 1400, 1499},
	3:  {"Copper 2", 1500, 1599},
	4:  {"Copper 1", 1600, 1699},
	5:  {"Bronze 4", 1700, 1799},
	6:  {"Bronze 3", 1800, 1899},
	7:  {"Bronze 2", 1900, 1999},
	8:  {"Bronze 1", 2000, 2099},
	9:  {"Silver 4", 2100, 2199},
	10: {"Silver 3", 2200, 2299},
	11: {"Silver 2", 2300, 2399},
	12: {"Silver 1", 2400, 2499},
	13: {"Gold 4", 2500, 2699},
	14: {"Gold 3", 2700, 2899},
	15: {"Gold 2", 2900, 3099},
	16: {"Gold 1", 3100, 3299},
	17: {"Platinum 3", 3300, 3699},
	18: {"Platinum 2", 3700, 4099},
	19: {"Platinum 1", 4100, 4499},
	20: {"Diamond", 4500, 0},
}

var ranksV4 = map[int16]rankInfo{
	0:  {"Unranked", 0, 0},
	1:  {"Copper 5", 1, 1199},
	2:  {"Copper 4", 1200, 1299},
	3:  {"Copper 3", 1300, 1399},
	4:  {"Copper 2", 1400, 1499},
	5:  {"Copper 1", 1500, 1599},
	6:  {"Bronze 5", 1600, 1699},
	7:  {"Bronze 4", 1700, 1799},
	8:  {"Bronze 3", 1800, 1899},
	9:  {"Bronze 2", 1900, 1999},
	10: {"Bronze 1", 2000, 2099},
	11: {"Silver 5", 2100, 2199},
	12: {"Silver 4", 2200, 2299},
	13: {"Silver 3", 2300, 2399},
	14: {"Silver 2", 2400, 2499},
	15: {"Silver 1", 2500, 2599},
	16: {"Gold 3", 2600, 2799},
	17: {"Gold 2", 2800, 2999},
	18: {"Gold 1", 3000, 3199},
	19: {"Platinum 3", 3200, 3599},
	20: {"Platinum 2", 3600, 3999},
	21: {"Platinum 1", 4000, 4399},
	22: {"Diamond", 4400, 4999},
	23: {"Champions", 5000, 0},
}

var ranksV5 = map[int16]rankInfo{
	0:  {"Unranked", 0, 0},
	1:  {"Copper 5", 1, 1199},
	2:  {"Copper 4", 1200, 1299},
	3:  {"Copper 3", 1300, 1399},
	4:  {"Copper 2", 1400, 1499},
	5:  {"Copper 1", 1500, 1599},
	6:  {"Bronze 5", 1600, 1699},
	7:  {"Bronze 4", 1700, 1799},
	8:  {"Bronze 3", 1800, 1899},
	9:  {"Bronze 2", 1900, 1999},
	10: {"Bronze 1", 2000, 2099},
	11: {"Silver 5", 2100, 2199},
	12: {"Silver 4", 2200, 2299},
	13: {"Silver 3", 2300, 2399},
	14: {"Silver 2", 2400, 2499},
	15: {"Silver 1", 2500, 2599},
	16: {"Gold 3", 2600, 2799},
	17: {"Gold 2", 2800, 2999},
	18: {"Gold 1", 3000, 3199},
	19: {"Platinum 3", 3200, 3499},
	20: {"Platinum 2", 3500, 3799},
	21: {"Platinum 1", 3800, 4099},
	22: {"Diamond 3", 4100, 4399},
	23: {"Diamond 2", 4400, 4699},
	24: {"Diamond 1", 4700, 4999},
	25: {"Champion", 5000, 0},
}

var ranksV6 = map[int16]rankInfo{
	0:  {"Unranked", 0, 0},
	1:  {"Copper 5", 1000, 1099},
	2:  {"Copper 4", 1100, 1199},
	3:  {"Copper 3", 1200, 1299},
	4:  {"Copper 2", 1300, 1399},
	5:  {"Copper 1", 1400, 1499},
	6:  {"Bronze 5", 1500, 1599},
	7:  {"Bronze 4", 1600, 1699},
	8:  {"Bronze 3", 1700, 1799},
	9:  {"Bronze 2", 1800, 1899},
	10: {"Bronze 1", 1900, 1999},
	11: {"Silver 5", 2000, 2099},
	12: {"Silver 4", 2100, 2199},
	13: {"Silver 3", 2200, 2299},
	14: {"Silver 2", 2300, 2399},
	15: {"Silver 1", 2400, 2499},
	16: {"Gold 5", 2500, 2599},
	17: {"Gold 4", 2600, 2799},
	18: {"Gold 3", 2700, 2799},
	19: {"Gold 2", 2800, 2899},
	20: {"Gold 1", 2900, 2999},
	21: {"Platinum 5", 3000, 3099},
	22: {"Platinum 4", 3100, 3199},
	23: {"Platinum 3", 3200, 3299},
	24: {"Platinum 2", 3300, 3399},
	25: {"Platinum 1", 3400, 3499},
	26: {"Emerald 5", 3500, 3599},
	27: {"Emerald 4", 3600, 3699},
	28: {"Emerald 3", 3700, 3799},
	29: {"Emerald 2", 3800, 3899},
	30: {"Emerald 1", 3900, 3999},
	31: {"Diamond 5", 4000, 4099},
	32: {"Diamond 4", 4100, 4199},
	33: {"Diamond 3", 4200, 4299},
	34: {"Diamond 2", 4300, 4399},
	35: {"Diamond 1", 4400, 4499},
	36: {"Champion", 4500, 0},
}

func getRankInfo(season int16) map[int16]rankInfo {
	if 5 <= season && season <= 14 {
		return ranksV3
	}
	if 15 <= season && season <= 22 {
		return ranksV4
	}
	if 23 <= season && season <= 27 {
		return ranksV5
	}
	if 28 <= season {
		return ranksV6
	}

	return nil
}

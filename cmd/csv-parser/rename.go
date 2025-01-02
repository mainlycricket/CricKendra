package main

type seriesSeasonKey struct {
	playingFormat string
	isMale        bool
	seriesName    string
	season        string
}

type renames struct {
	cities        map[string]string
	grounds       map[string]string
	series        map[string]string
	seriesSeasons map[seriesSeasonKey]string
}

var renameData = renames{
	cities:        renameCities,
	grounds:       renameGrounds,
	series:        renameSeries,
	seriesSeasons: renameSeriesSeaons,
}

var renameCities = map[string]string{
	"Dharmasala": "Dharamsala",
	"Chittagong": "Chattogram",
	"Bangalore":  "Bengaluru",
}

var renameGrounds = map[string]string{
	"Wankhede Stadium, Mumbai":         "Wankhede Stadium",
	"M Chinnaswamy Stadium":            "M.Chinnaswamy Stadium",
	"M Chinnaswamy Stadium, Bengaluru": "M.Chinnaswamy Stadium",
}

var renameSeries = map[string]string{
	"NatWest Challenge": "NatWest Series",
	"ICC World Cup":     "ICC Cricket World Cup",
	"World Cup":         "ICC Cricket World Cup",
	"ICC Cricket World Cup Qualifier (ICC Trophy)": "ICC Cricket World Cup Qualifier",
	"ICC World Cup Qualifiers":                     "ICC Cricket World Cup Qualifier",
}

var renameSeriesSeaons = map[seriesSeasonKey]string{
	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2011",
	}: "2011/13",

	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2011/12",
	}: "2011/13",

	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2012",
	}: "2011/13",

	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2012/13",
	}: "2011/13",

	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2013",
	}: "2011/13",

	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2015/16",
	}: "2015/17",

	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2016",
	}: "2015/17",

	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2017",
	}: "2015/17",

	{
		playingFormat: "ODI",
		isMale:        true,
		seriesName:    "ICC World Cricket League Championship",
		season:        "2017/18",
	}: "2015/17",
}

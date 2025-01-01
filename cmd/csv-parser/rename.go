package main

type Renames struct {
	Cities  map[string]string
	Grounds map[string]string
	Series  map[string]string
}

var renameData = Renames{
	Cities:  renameCities,
	Grounds: renameGroungs,
	Series:  renameSeries,
}

var renameCities = map[string]string{
	"Dharmasala": "Dharamsala",
	"Chittagong": "Chattogram",
	"Bangalore":  "Bengaluru",
}

var renameGroungs = map[string]string{
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

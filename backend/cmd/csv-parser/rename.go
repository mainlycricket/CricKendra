package main

// renameSeriesNames	map[seriesKey]string	exclude season
// renameSeriesSeasons 	map[seriesKey]string
// renameVenues			map[venue]venue

type venue struct {
	groundName, cityName string
}

// doesn't consider season
var renamedSeriesNames = map[seriesKey]string{
	{
		name:           "NatWest Challenge",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
	}: "NatWest Series",
	{
		name:           "ICC World Cup",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
	}: "ICC Cricket World Cup",
	{
		name:           "World Cup",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
	}: "ICC Cricket World Cup",
	{
		name:           "ICC Cricket World Cup Qualifier (ICC Trophy)",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
	}: "ICC Cricket World Cup Qualifier",
	{
		name:           "ICC World Cup Qualifiers",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
	}: "ICC Cricket World Cup Qualifier",
}

var renamedSeriesSeaons = map[seriesKey]string{
	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2011",
	}: "2011/13",

	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2011/12",
	}: "2011/13",

	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2012",
	}: "2011/13",

	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2012/13",
	}: "2011/13",

	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2013",
	}: "2011/13",

	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2015/16",
	}: "2015/17",

	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2016",
	}: "2015/17",

	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2017",
	}: "2015/17",

	{
		name:           "ICC World Cricket League Championship",
		is_male:        true,
		playing_format: "ODI",
		playing_level:  "international",
		season:         "2017/18",
	}: "2015/17",
}

var renamedVenues = map[venue]venue{
	{
		groundName: "North West Cricket Stadium, Potchefstroom",
		cityName:   "Potchefstroom",
	}: {
		groundName: "Senwes Park",
		cityName:   "Potchefstroom",
	},
	{
		groundName: "Sedgars Park",
		cityName:   "Potchefstroom",
	}: {
		groundName: "Senwes Park",
		cityName:   "Potchefstroom",
	},
	{
		groundName: "Sedgars Park, Potchefstroom",
		cityName:   "Potchefstroom",
	}: {
		groundName: "Senwes Park",
		cityName:   "Potchefstroom",
	},
	{
		groundName: "Senwes Park, Potchefstroom",
		cityName:   "Potchefstroom",
	}: {
		groundName: "Senwes Park",
		cityName:   "Potchefstroom",
	},
	{
		groundName: "Old Trafford, Manchester",
		cityName:   "Manchester",
	}: {
		groundName: "Old Trafford",
		cityName:   "Manchester",
	},
	{
		groundName: "The Rose Bowl, Southampton",
		cityName:   "Southampton",
	}: {
		groundName: "The Rose Bowl",
		cityName:   "Southampton",
	},
	{
		groundName: "Trent Bridge, Nottingham",
		cityName:   "Nottingham",
	}: {
		groundName: "Trent Bridge",
		cityName:   "Nottingham",
	},
	{
		groundName: "County Ground, Bristol",
		cityName:   "Bristol",
	}: {
		groundName: "County Ground",
		cityName:   "Bristol",
	},
	{
		groundName: "The Royal & Sun Alliance County Ground, Bristol",
		cityName:   "Bristol",
	}: {
		groundName: "County Ground",
		cityName:   "Bristol",
	},
	{
		groundName: "Kennington Oval, London",
		cityName:   "London",
	}: {
		groundName: "Kennington Oval",
		cityName:   "London",
	},
	{
		groundName: "Lord's, London",
		cityName:   "London",
	}: {
		groundName: "Lord's",
		cityName:   "London",
	},
	{
		groundName: "Edgbaston, Birmingham",
		cityName:   "Birmingham",
	}: {
		groundName: "Edgbaston",
		cityName:   "Birmingham",
	},
	{
		groundName: "Saxton Oval, Nelson",
		cityName:   "Nelson",
	}: {
		groundName: "Saxton Oval",
		cityName:   "Nelson",
	},
	{
		groundName: "Kensington Oval, Barbados",
		cityName:   "Bridgetown",
	}: {
		groundName: "Kensington Oval, Bridgetown",
		cityName:   "Barbados",
	},
	{
		groundName: "Kensington Oval, Bridgetown, Barbados",
		cityName:   "Bridgetown",
	}: {
		groundName: "Kensington Oval, Bridgetown",
		cityName:   "Barbados",
	},
	{
		groundName: "Eden Park, Auckland",
		cityName:   "Auckland",
	}: {
		groundName: "Eden Park",
		cityName:   "Auckland",
	},
	{
		groundName: "Headingley, Leeds",
		cityName:   "Leeds",
	}: {
		groundName: "Headingley",
		cityName:   "Leeds",
	},
	{
		groundName: "R Premadasa Stadium, Colombo",
		cityName:   "Colombo",
	}: {
		groundName: "R Premadasa Stadium",
		cityName:   "Colombo",
	},
	{
		groundName: "R.Premadasa Stadium",
		cityName:   "Colombo",
	}: {
		groundName: "R Premadasa Stadium",
		cityName:   "Colombo",
	},
	{
		groundName: "R.Premadasa Stadium, Khettarama",
		cityName:   "Colombo",
	}: {
		groundName: "R Premadasa Stadium",
		cityName:   "Colombo",
	},
	{
		groundName: "Sinhalese Sports Club Ground",
		cityName:   "Colombo",
	}: {
		groundName: "Sinhalese Sports Club",
		cityName:   "Colombo",
	},
	{
		groundName: "Mahinda Rajapaksa International Cricket Stadium, Sooriyawewa, Hambantota",
		cityName:   "Hambantota",
	}: {
		groundName: "Mahinda Rajapaksa International Cricket Stadium, Sooriyawewa",
		cityName:   "Hambantota",
	},
	{
		groundName: "Bellerive Oval, Hobart",
		cityName:   "Hobart",
	}: {
		groundName: "Bellerive Oval",
		cityName:   "Hobart",
	},
	{
		groundName: "Western Australia Cricket Association Ground",
		cityName:   "Perth",
	}: {
		groundName: "W.A.C.A. Ground",
		cityName:   "Perth",
	},
	{
		groundName: "Chevrolet Park",
		cityName:   "Bloemfontein",
	}: {
		groundName: "Mangaung Oval",
		cityName:   "Bloemfontein",
	},
	{
		groundName: "Goodyear Park",
		cityName:   "Bloemfontein",
	}: {
		groundName: "Mangaung Oval",
		cityName:   "Bloemfontein",
	},
	{
		groundName: "Goodyear Park, Bloemfontein",
		cityName:   "Bloemfontein",
	}: {
		groundName: "Mangaung Oval",
		cityName:   "Bloemfontein",
	},
	{
		groundName: "Mangaung Oval, Bloemfontein",
		cityName:   "Bloemfontein",
	}: {
		groundName: "Mangaung Oval",
		cityName:   "Bloemfontein",
	},
	{
		groundName: "SuperSport Park, Centurion",
		cityName:   "Centurion",
	}: {
		groundName: "SuperSport Park",
		cityName:   "Centurion",
	},
	{
		groundName: "New Wanderers Stadium",
		cityName:   "Johannesburg",
	}: {
		groundName: "The Wanderers Stadium",
		cityName:   "Johannesburg",
	},
	{
		groundName: "New Wanderers Stadium, Johannesburg",
		cityName:   "Johannesburg",
	}: {
		groundName: "The Wanderers Stadium",
		cityName:   "Johannesburg",
	},
	{
		groundName: "The Wanderers Stadium, Johannesburg",
		cityName:   "Johannesburg",
	}: {
		groundName: "The Wanderers Stadium",
		cityName:   "Johannesburg",
	},
	{
		groundName: "Newlands, Cape Town",
		cityName:   "Cape Town",
	}: {
		groundName: "Newlands",
		cityName:   "Cape Town",
	},
	{
		groundName: "Sophia Gardens, Cardiff",
		cityName:   "Cardiff",
	}: {
		groundName: "Sophia Gardens",
		cityName:   "Cardiff",
	},
	{
		groundName: "Beausejour Stadium, Gros Islet",
		cityName:   "St Lucia",
	}: {
		groundName: "Daren Sammy National Cricket Stadium",
		cityName:   "Gros Islet",
	},
	{
		groundName: "Daren Sammy National Cricket Stadium, Gros Islet",
		cityName:   "St Lucia",
	}: {
		groundName: "Daren Sammy National Cricket Stadium",
		cityName:   "Gros Islet",
	},
	{
		groundName: "Darren Sammy National Cricket Stadium, Gros Islet",
		cityName:   "St Lucia",
	}: {
		groundName: "Daren Sammy National Cricket Stadium",
		cityName:   "Gros Islet",
	},
	{
		groundName: "Sir Vivian Richards Stadium",
		cityName:   "North Sound",
	}: {
		groundName: "Sir Vivian Richards Stadium, North Sound",
		cityName:   "Antigua",
	},
	{
		groundName: "Sir Vivian Richards Stadium, North Sound, Antigua",
		cityName:   "North Sound",
	}: {
		groundName: "Sir Vivian Richards Stadium, North Sound",
		cityName:   "Antigua",
	},
	{
		groundName: "Bharat Ratna Shri Atal Bihari Vajpayee Ekana Cricket Stadium, Lucknow",
		cityName:   "Lucknow",
	}: {
		groundName: "Bharat Ratna Shri Atal Bihari Vajpayee Ekana Cricket Stadium",
		cityName:   "Lucknow",
	},
	{
		groundName: "Wankhede Stadium, Mumbai",
		cityName:   "Mumbai",
	}: {
		groundName: "Wankhede Stadium",
		cityName:   "Mumbai",
	},
	{
		groundName: "M Chinnaswamy Stadium",
		cityName:   "Bengaluru",
	}: {
		groundName: "M.Chinnaswamy Stadium",
		cityName:   "Bengaluru",
	},
	{
		groundName: "M Chinnaswamy Stadium, Bengaluru",
		cityName:   "Bengaluru",
	}: {
		groundName: "M.Chinnaswamy Stadium",
		cityName:   "Bengaluru",
	},
	{
		groundName: "Maharashtra Cricket Association Stadium, Pune",
		cityName:   "Pune",
	}: {
		groundName: "Maharashtra Cricket Association Stadium",
		cityName:   "Pune",
	},
	{
		groundName: "Nehru Stadium, Poona",
		cityName:   "Pune",
	}: {
		groundName: "Nehru Stadium",
		cityName:   "Pune",
	},
	{
		groundName: "Sardar Patel (Gujarat) Stadium, Motera",
		cityName:   "Ahmedabad",
	}: {
		groundName: "Sardar Patel Stadium, Motera",
		cityName:   "Ahmedabad",
	},
	{
		groundName: "Brisbane Cricket Ground, Woolloongabba",
		cityName:   "Brisbane",
	}: {
		groundName: "Brisbane Cricket Ground",
		cityName:   "Brisbane",
	},
	{
		groundName: "Brisbane Cricket Ground, Woolloongabba, Brisbane",
		cityName:   "Brisbane",
	}: {
		groundName: "Brisbane Cricket Ground",
		cityName:   "Brisbane",
	},
	{
		groundName: "Clontarf Cricket Club Ground, Dublin",
		cityName:   "Dublin",
	}: {
		groundName: "Clontarf Cricket Club Ground",
		cityName:   "Dublin",
	},
	{
		groundName: "Malahide",
		cityName:   "Dublin",
	}: {
		groundName: "The Village, Malahide",
		cityName:   "Dublin",
	},
	{
		groundName: "The Village, Malahide, Dublin",
		cityName:   "Dublin",
	}: {
		groundName: "The Village, Malahide",
		cityName:   "Dublin",
	},
	{
		groundName: "Riverside Ground, Chester-le-Street",
		cityName:   "Chester-le-Street",
	}: {
		groundName: "Riverside Ground",
		cityName:   "Chester-le-Street",
	},
	{
		groundName: "Seddon Park, Hamilton",
		cityName:   "Hamilton",
	}: {
		groundName: "Seddon Park",
		cityName:   "Hamilton",
	},
	{
		groundName: "Westpac Park, Hamilton",
		cityName:   "Hamilton",
	}: {
		groundName: "Seddon Park",
		cityName:   "Hamilton",
	},
	{
		groundName: "Westpac Stadium, Wellington",
		cityName:   "Wellington",
	}: {
		groundName: "Westpac Stadium",
		cityName:   "Wellington",
	},
	{
		groundName: "McLean Park, Napier",
		cityName:   "Napier",
	}: {
		groundName: "McLean Park",
		cityName:   "Napier",
	},
	{
		groundName: "Hagley Oval, Christchurch",
		cityName:   "Christchurch",
	}: {
		groundName: "Hagley Oval",
		cityName:   "Christchurch",
	},
	{
		groundName: "Jade Stadium, Christchurch",
		cityName:   "Christchurch",
	}: {
		groundName: "Jade Stadium",
		cityName:   "Christchurch",
	},
	{
		groundName: "Sabina Park, Kingston, Jamaica",
		cityName:   "Kingston",
	}: {
		groundName: "Sabina Park, Kingston",
		cityName:   "Jamaica",
	},
	{
		groundName: "Kingsmead, Durban",
		cityName:   "Durban",
	}: {
		groundName: "Kingsmead",
		cityName:   "Durban",
	},
	{
		groundName: "Arun Jaitley Stadium, Delhi",
		cityName:   "Delhi",
	}: {
		groundName: "Arun Jaitley Stadium",
		cityName:   "Delhi",
	},
	{
		groundName: "Feroz Shah Kotla",
		cityName:   "Delhi",
	}: {
		groundName: "Arun Jaitley Stadium",
		cityName:   "Delhi",
	},
	{
		groundName: "Queens Sports Club, Bulawayo",
		cityName:   "Bulawayo",
	}: {
		groundName: "Queens Sports Club",
		cityName:   "Bulawayo",
	},
	{
		groundName: "MA Chidambaram Stadium, Chepauk, Chennai",
		cityName:   "Chennai",
	}: {
		groundName: "MA Chidambaram Stadium, Chepauk",
		cityName:   "Chennai",
	},
	{
		groundName: "Shere Bangla National Stadium",
		cityName:   "Mirpur",
	}: {
		groundName: "Shere Bangla National Stadium, Mirpur",
		cityName:   "Dhaka",
	},
	{
		groundName: "Sher-e-Bangla National Cricket Stadium",
		cityName:   "Mirpur",
	}: {
		groundName: "Shere Bangla National Stadium, Mirpur",
		cityName:   "Dhaka",
	},
	{
		groundName: "Queen's Park Oval, Port of Spain",
		cityName:   "Trinidad",
	}: {
		groundName: "Queen's Park Oval",
		cityName:   "Port of Spain",
	},
	{
		groundName: "Willowmoore Park, Benoni",
		cityName:   "Benoni",
	}: {
		groundName: "Willowmoore Park",
		cityName:   "Benoni",
	},
	{
		groundName: "University Oval, Dunedin",
		cityName:   "Dunedin",
	}: {
		groundName: "University Oval",
		cityName:   "Dunedin",
	},
	{
		groundName: "Bay Oval, Mount Maunganui",
		cityName:   "Mount Maunganui",
	}: {
		groundName: "Bay Oval",
		cityName:   "Mount Maunganui",
	},
	{
		groundName: "Boland Bank Park, Paarl",
		cityName:   "Paarl",
	}: {
		groundName: "Boland Park",
		cityName:   "Paarl",
	},
	{
		groundName: "Boland Park, Paarl",
		cityName:   "Paarl",
	}: {
		groundName: "Boland Park",
		cityName:   "Paarl",
	},
	{
		groundName: "Buffalo Park, East London",
		cityName:   "East London",
	}: {
		groundName: "Buffalo Park",
		cityName:   "East London",
	},
	{
		groundName: "Rajiv Gandhi International Stadium, Uppal, Hyderabad",
		cityName:   "Hyderabad",
	}: {
		groundName: "Rajiv Gandhi International Stadium, Uppal",
		cityName:   "Hyderabad",
	},
	{
		groundName: "Punjab Cricket Association IS Bindra Stadium, Mohali, Chandigarh",
		cityName:   "Chandigarh",
	}: {
		groundName: "Punjab Cricket Association IS Bindra Stadium, Mohali",
		cityName:   "Chandigarh",
	},
	{
		groundName: "Punjab Cricket Association Stadium, Mohali",
		cityName:   "Chandigarh",
	}: {
		groundName: "Punjab Cricket Association IS Bindra Stadium, Mohali",
		cityName:   "Chandigarh",
	},
	{
		groundName: "Civil Service Cricket Club, Stormont, Belfast",
		cityName:   "Belfast",
	}: {
		groundName: "Civil Service Cricket Club, Stormont",
		cityName:   "Belfast",
	},
	{
		groundName: "Arnos Vale Ground, Kingstown",
		cityName:   "St Vincent",
	}: {
		groundName: "Arnos Vale Ground",
		cityName:   "Kingstown",
	},
	{
		groundName: "De Beers Diamond Oval, Kimberley",
		cityName:   "Kimberley",
	}: {
		groundName: "De Beers Diamond Oval",
		cityName:   "Kimberley",
	},
	{
		groundName: "Diamond Oval, Kimberley",
		cityName:   "Kimberley",
	}: {
		groundName: "Diamond Oval",
		cityName:   "Kimberley",
	},
	{
		groundName: "Manuka Oval, Canberra",
		cityName:   "Canberra",
	}: {
		groundName: "Manuka Oval",
		cityName:   "Canberra",
	},
	{
		groundName: "Queen's Park Oval, Port of Spain, Trinidad",
		cityName:   "Port of Spain",
	}: {
		groundName: "Queen's Park Oval",
		cityName:   "Port of Spain",
	},
	{
		groundName: "Queen's Park Oval, Trinidad",
		cityName:   "Port of Spain",
	}: {
		groundName: "Queen's Park Oval",
		cityName:   "Port of Spain",
	},
	{
		groundName: "Darren Sammy National Cricket Stadium, St Lucia",
		cityName:   "Gros Islet",
	}: {
		groundName: "Daren Sammy National Cricket Stadium",
		cityName:   "Gros Islet",
	},
	{
		groundName: "St George's Park, Gqeberha",
		cityName:   "Gqeberha",
	}: {
		groundName: "St George's Park",
		cityName:   "Port Elizabeth",
	},
	{
		groundName: "St George's Park, Port Elizabeth",
		cityName:   "Gqeberha",
	}: {
		groundName: "St George's Park",
		cityName:   "Port Elizabeth",
	},
	{
		groundName: "Greenfield International Stadium, Thiruvananthapuram",
		cityName:   "Thiruvananthapuram",
	}: {
		groundName: "Greenfield International Stadium",
		cityName:   "Thiruvananthapuram",
	},
	{
		groundName: "National Cricket Stadium, Grenada",
		cityName:   "St George's",
	}: {
		groundName: "National Cricket Stadium",
		cityName:   "St George's",
	},
	{
		groundName: "Sharjah Cricket Association Stadium",
		cityName:   "Sharjah",
	}: {
		groundName: "Sharjah Cricket Stadium",
		cityName:   "Sharjah",
	},
	{
		groundName: "Nahar Singh Stadium, Faridabad",
		cityName:   "Faridabad",
	}: {
		groundName: "Nahar Singh Stadium",
		cityName:   "Faridabad",
	},
	{
		groundName: "Gaddafi Stadium, Lahore",
		cityName:   "Lahore",
	}: {
		groundName: "Gaddafi Stadium",
		cityName:   "Lahore",
	},
	{
		groundName: "Eden Gardens, Kolkata",
		cityName:   "Kolkata",
	}: {
		groundName: "Eden Gardens",
		cityName:   "Kolkata",
	},
	{
		groundName: "JSCA International Stadium Complex, Ranchi",
		cityName:   "Ranchi",
	}: {
		groundName: "JSCA International Stadium Complex",
		cityName:   "Ranchi",
	},
	{
		groundName: "Holkar Cricket Stadium, Indore",
		cityName:   "Indore",
	}: {
		groundName: "Holkar Cricket Stadium",
		cityName:   "Indore",
	},
	{
		groundName: "Saurashtra Cricket Association Stadium, Rajkot",
		cityName:   "Rajkot",
	}: {
		groundName: "Saurashtra Cricket Association Stadium",
		cityName:   "Rajkot",
	},
	{
		groundName: "Dr. Y.S. Rajasekhara Reddy ACA-VDCA Cricket Stadium, Visakhapatnam",
		cityName:   "Visakhapatnam",
	}: {
		groundName: "Dr. Y.S. Rajasekhara Reddy ACA-VDCA Cricket Stadium",
		cityName:   "Visakhapatnam",
	},
	{
		groundName: "Barabati Stadium, Cuttack",
		cityName:   "Cuttack",
	}: {
		groundName: "Barabati Stadium",
		cityName:   "Cuttack",
	},
	{
		groundName: "Providence Stadium, Guyana",
		cityName:   "Providence",
	}: {
		groundName: "Providence Stadium",
		cityName:   "Guyana",
	},
	{
		groundName: "National Stadium, Karachi",
		cityName:   "Karachi",
	}: {
		groundName: "National Stadium",
		cityName:   "Karachi",
	},
	{
		groundName: "Grange Cricket Club Ground, Raeburn Place, Edinburgh",
		cityName:   "Edinburgh",
	}: {
		groundName: "Grange Cricket Club Ground, Raeburn Place",
		cityName:   "Edinburgh",
	},
	{
		groundName: "Grange Cricket Club, Raeburn Place",
		cityName:   "Edinburgh",
	}: {
		groundName: "Grange Cricket Club Ground, Raeburn Place",
		cityName:   "Edinburgh",
	},
	{
		groundName: "Titwood, Glasgow",
		cityName:   "Glasgow",
	}: {
		groundName: "Titwood",
		cityName:   "Glasgow",
	},
	{
		groundName: "Wanderers Cricket Ground, Windhoek",
		cityName:   "Windhoek",
	}: {
		groundName: "Wanderers Cricket Ground",
		cityName:   "Windhoek",
	},
	{
		groundName: "National Cricket Stadium, St George's",
		cityName:   "Grenada",
	}: {
		groundName: "National Cricket Stadium",
		cityName:   "St George's",
	},
	{
		groundName: "Cambusdoon New Ground, Ayr",
		cityName:   "Ayr",
	}: {
		groundName: "Cambusdoon New Ground",
		cityName:   "Ayr",
	},
	{
		groundName: "Gymkhana Club Ground, Nairobi",
		cityName:   "Nairobi",
	}: {
		groundName: "Gymkhana Club Ground",
		cityName:   "Nairobi",
	},
	{
		groundName: "Iqbal Stadium, Faisalabad",
		cityName:   "Faisalabad",
	}: {
		groundName: "Iqbal Stadium",
		cityName:   "Faisalabad",
	},
	{
		groundName: "Marrara Cricket Ground, Darwin",
		cityName:   "Darwin",
	}: {
		groundName: "Marrara Cricket Ground",
		cityName:   "Darwin",
	},
	{
		groundName: "Vidarbha Cricket Association Stadium, Jamtha",
		cityName:   "Nagpur",
	}: {
		groundName: "Vidarbha Cricket Association Ground",
		cityName:   "Nagpur",
	},
	{
		groundName: "Captain Roop Singh Stadium, Gwalior",
		cityName:   "Gwalior",
	}: {
		groundName: "Captain Roop Singh Stadium",
		cityName:   "Gwalior",
	},
	{
		groundName: "Bangabandhu National Stadium, Dhaka",
		cityName:   "Dhaka",
	}: {
		groundName: "Bangabandhu National Stadium",
		cityName:   "Dhaka",
	},
	{
		groundName: "Shere Bangla National Stadium",
		cityName:   "Dhaka",
	}: {
		groundName: "Shere Bangla National Stadium, Mirpur",
		cityName:   "Dhaka",
	},
	{
		groundName: "Himachal Pradesh Cricket Association Stadium",
		cityName:   "Dharamsala",
	}: {
		groundName: "Himachal Pradesh Cricket Association Stadium",
		cityName:   "Dharamsala",
	},
	{
		groundName: "Himachal Pradesh Cricket Association Stadium, Dharamsala",
		cityName:   "Dharamsala",
	}: {
		groundName: "Himachal Pradesh Cricket Association Stadium",
		cityName:   "Dharamsala",
	},
	{
		groundName: "Arnos Vale Ground, Kingstown, St Vincent",
		cityName:   "Kingstown",
	}: {
		groundName: "Arnos Vale Ground",
		cityName:   "Kingstown",
	},
	{
		groundName: "MA Aziz Stadium, Chittagong",
		cityName:   "Chattogram",
	}: {
		groundName: "MA Aziz Stadium",
		cityName:   "Chattogram",
	},
	{
		groundName: "Zahur Ahmed Chowdhury Stadium",
		cityName:   "Chattogram",
	}: {
		groundName: "Zahur Ahmed Chowdhury Stadium",
		cityName:   "Chattogram",
	},
	{
		groundName: "Zahur Ahmed Chowdhury Stadium, Chattogram",
		cityName:   "Chattogram",
	}: {
		groundName: "Zahur Ahmed Chowdhury Stadium",
		cityName:   "Chattogram",
	},
	{
		groundName: "Zohur Ahmed Chowdhury Stadium",
		cityName:   "Chattogram",
	}: {
		groundName: "Zahur Ahmed Chowdhury Stadium",
		cityName:   "Chattogram",
	},
	{
		groundName: "Barsapara Cricket Stadium, Guwahati",
		cityName:   "Guwahati",
	}: {
		groundName: "Barsapara Cricket Stadium",
		cityName:   "Guwahati",
	},
	{
		groundName: "St Lawrence Ground, Canterbury",
		cityName:   "Canterbury",
	}: {
		groundName: "St Lawrence Ground",
		cityName:   "Canterbury",
	},
	{
		groundName: "Mannofield Park, Aberdeen",
		cityName:   "Aberdeen",
	}: {
		groundName: "Mannofield Park",
		cityName:   "Aberdeen",
	},
	{
		groundName: "Dubai Sports City Cricket Stadium",
		cityName:   "Dubai",
	}: {
		groundName: "Dubai International Cricket Stadium",
		cityName:   "Dubai",
	},
	{
		groundName: "ICC Academy",
		cityName:   "Dubai",
	}: {
		groundName: "Senwes Park",
		cityName:   "Potchefstroom",
	},
	{
		groundName: "ICC Academy, Dubai",
		cityName:   "Dubai",
	}: {
		groundName: "ICC Academy",
		cityName:   "Dubai",
	},
	{
		groundName: "ICC Global Cricket Academy",
		cityName:   "Dubai",
	}: {
		groundName: "ICC Academy",
		cityName:   "Dubai",
	},
	{
		groundName: "VRA Cricket Ground",
		cityName:   "Amstelveen",
	}: {
		groundName: "VRA Ground",
		cityName:   "Amstelveen",
	},
	{
		groundName: "VRA Ground, Amstelveen",
		cityName:   "Amstelveen",
	}: {
		groundName: "VRA Ground",
		cityName:   "Amstelveen",
	},
	{
		groundName: "Tribhuvan University International Cricket Ground, Kirtipur",
		cityName:   "Kirtipur",
	}: {
		groundName: "Tribhuvan University International Cricket Ground",
		cityName:   "Kirtipur",
	},
	{
		groundName: "Maple Leaf North-West Ground, King City",
		cityName:   "King City",
	}: {
		groundName: "Maple Leaf North-West Ground",
		cityName:   "King City",
	},
	{
		groundName: "Maple Leaf North-West Ground",
		cityName:   "Toronto",
	}: {
		groundName: "Maple Leaf North-West Ground",
		cityName:   "King City",
	},
	{
		groundName: "Pallekele International Cricket Stadium",
		cityName:   "Pallekele",
	}: {
		groundName: "Pallekele International Cricket Stadium",
		cityName:   "Kandy",
	},
	{
		groundName: "Choice Moosa Stadium, Pearland",
		cityName:   "Pearland",
	}: {
		groundName: "Moosa Cricket Stadium, Pearland",
		cityName:   "Pearland",
	},
	{
		groundName: "Melbourne Cricket Ground",
		cityName:   "",
	}: {
		groundName: "Melbourne Cricket Ground",
		cityName:   "Melbourne",
	},
	{
		groundName: "Sydney Cricket Ground",
		cityName:   "",
	}: {
		groundName: "Sydney Cricket Ground",
		cityName:   "Sydney",
	},
	{
		groundName: "Adelaide Oval",
		cityName:   "",
	}: {
		groundName: "Adelaide Oval",
		cityName:   "Adelaide",
	},
	{
		groundName: "Harare Sports Club",
		cityName:   "",
	}: {
		groundName: "Harare Sports Club",
		cityName:   "Harare",
	},
	{
		groundName: "Sharjah Cricket Stadium",
		cityName:   "",
	}: {
		groundName: "Sharjah Cricket Stadium",
		cityName:   "Sharjah",
	},
	{
		groundName: "Perth Stadium",
		cityName:   "",
	}: {
		groundName: "Perth Stadium",
		cityName:   "Perth",
	},
	{
		groundName: "Dubai International Cricket Stadium",
		cityName:   "",
	}: {
		groundName: "Dubai International Cricket Stadium",
		cityName:   "Dubai",
	},
	{
		groundName: "Rangiri Dambulla International Stadium",
		cityName:   "",
	}: {
		groundName: "Rangiri Dambulla International Stadium",
		cityName:   "Dambulla",
	},
	{
		groundName: "Galle International Stadium",
		cityName:   "",
	}: {
		groundName: "Galle International Stadium",
		cityName:   "Galle",
	},
	{
		groundName: "Pallekele International Cricket Stadium",
		cityName:   "",
	}: {
		groundName: "Pallekele International Cricket Stadium",
		cityName:   "Kandy",
	},
	{
		groundName: "Bulawayo Athletic Club",
		cityName:   "",
	}: {
		groundName: "Bulawayo Athletic Club",
		cityName:   "Bulawayo",
	},
	{
		groundName: "Rawalpindi Cricket Stadium",
		cityName:   "",
	}: {
		groundName: "Rawalpindi Cricket Stadium",
		cityName:   "Rawalpindi",
	},
	{
		groundName: "Multan Cricket Stadium",
		cityName:   "",
	}: {
		groundName: "Multan Cricket Stadium",
		cityName:   "Multan",
	},
	{
		groundName: "Chittagong Divisional Stadium",
		cityName:   "",
	}: {
		groundName: "Chittagong Divisional Stadium",
		cityName:   "Chattogram",
	},
	{
		groundName: "Queenstown Events Centre",
		cityName:   "",
	}: {
		groundName: "Queenstown Events Centre",
		cityName:   "Queenstown",
	},
	{
		groundName: "Sheikhupura Stadium",
		cityName:   "",
	}: {
		groundName: "Sheikhupura Stadium",
		cityName:   "Sheikhupura",
	},
	{
		groundName: "Dubai Sports City Cricket Stadium",
		cityName:   "",
	}: {
		groundName: "Dubai International Cricket Stadium",
		cityName:   "Dubai",
	},
	{
		groundName: "Mombasa Sports Club Ground",
		cityName:   "",
	}: {
		groundName: "Mombasa Sports Club Ground",
		cityName:   "Mombasa",
	},
	{
		groundName: "Sharjah Cricket Association Stadium",
		cityName:   "",
	}: {
		groundName: "Saurashtra Cricket Association Stadium",
		cityName:   "Rajkot",
	},
	{
		groundName: "Zahur Ahmed Chowdhury Stadium",
		cityName:   "Chittagong",
	}: {
		groundName: "Zahur Ahmed Chowdhury Stadium",
		cityName:   "Chattogram",
	},
	{
		groundName: "Himachal Pradesh Cricket Association Stadium",
		cityName:   "Dharmasala",
	}: {
		groundName: "Himachal Pradesh Cricket Association Stadium",
		cityName:   "Dharamsala",
	},
	{
		groundName: "Cazaly's Stadium, Cairns",
		cityName:   "Cairns",
	}: {
		groundName: "Bundaberg Rum Stadium, Cairns",
		cityName:   "Cairns",
	},
	{
		groundName: "M Chinnaswamy Stadium",
		cityName:   "Bangalore",
	}: {
		groundName: "M.Chinnaswamy Stadium",
		cityName:   "Bengaluru",
	},
	{
		groundName: "OUTsurance Oval",
		cityName:   "Bloemfontein",
	}: {
		groundName: "Mangaung Oval",
		cityName:   "Bloemfontein",
	},
	{
		groundName: "Zohur Ahmed Chowdhury Stadium",
		cityName:   "Chittagong",
	}: {
		groundName: "Zahur Ahmed Chowdhury Stadium",
		cityName:   "Chattogram",
	},
}

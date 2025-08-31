CSV parser to extract data from CSV files available on [cricsheet.org](https://cricsheet.org/)

### Notes:

- Checkout the CSV Format Info [here](https://cricsheet.org/format/csv_ashwin/)
- Download the CSV files from [here](https://cricsheet.org/downloads/)
- Checkout the list of matches covered for testing [here](./matches_covered/)
- Some players appear in the ball by ball file of the match, but they don't appear in the match info file of that match. Hence, those matches are not completely migrated in the database.

### Missing Data

- Some data can't be accessed from Cricsheet:
- **Players Data**: Nationality, batting style, bowling styles, playing role etc
- **Grounds Data**: Linking grounds with their cities, host nations, continents etc.
- **Tournaments Data**: Various ODI & T20I tournaments, and linking with respective series

### Data Cleaning

- Some match info files may contain different variations of the same record:
- **Series Names**: `ICC World Cup` and `World Cup` are the same tournaments / series
- **Series Seasons**: `The ICC World Cricket League Championship 2011/13` was played across over multiple seasons. So, while the individual matches belong to different seasons, the series belongs to a single season
- **Venues**: `Bengaluru` and `Bangalore` are the same city. `M Chinnaswamy Stadium`, `M.Chinnaswamy Stadium` and `M Chinnaswamy Stadium, Bengaluru` are the same grounds
- and so on for team names etc.

### \*\*\*

- Manual work would be required to fix both these issues. And, this is not my priority at the moment, as it would slow down the development process
- But I've still made some progress as found in [seed data](../../db_files/seed_data/) and [`rename.go`](./rename.go)

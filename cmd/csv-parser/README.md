CSV parser to extract data from CSV files available on [cricsheet.org](https://cricsheet.org/)

### Notes:

- Checkout the CSV Format Info [here](https://cricsheet.org/format/csv_ashwin/)
- Checkout the list of matches covered for testing [here](./matches_covered/)
- T20I match with id `1229824` always fails because a player named `J Butler` exists in both teams, and a single entry is found in the match info registry

### Data Cleaning

- Some match info files may contain different variations of the same record
- Hence, manual data cleaning would be required
- And, this is not my priority at the moment, as it would slow down the development process
- But I've still made some progress, particularly for ODIs, as found in [`rename.go`](./rename.go)

  #### **Examples**

  1. **Series Names**: `ICC World Cup` and `World Cup` are the same tournaments / series
  2. **Series Seasons**: `The ICC World Cricket League Championship 2011/13` was played across over multiple seasons. So, while the individual matches belong to different seasons, the series belongs to a single season
  3. **Venues**: `Bengaluru` and `Bangalore` are the same city. `M Chinnaswamy Stadium`, `M.Chinnaswamy Stadium` and `M Chinnaswamy Stadium, Bengaluru` are the same grounds

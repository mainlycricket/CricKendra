### TODOs

-   Image upload
-   Validation
-   Pagination
-   Delete / Update
-   Authentication
-   Handle DB Errors (https://www.postgresql.org/docs/current/errcodes-appendix.html)
-   Series Points Table

### DB Notes

#### 1. Player

- Player has stats in two categories: db & unavailable
- DB stats will be always computed from database scorecards (not set manually)
- Unavailable Stats can be manually set by admin
- DB stats will be updated automatically after each match completion
- Get Player Profile will return computed stats in a unified manner

### Live Server

#### Admin Flow

- Auth Setup - TODO
- Create Match
- Upsert Match Squad
- Start Match - TODO
- Toss Decision
- Upsert Match Squad
- Create innings
- Create batter entries
- Update batter position & has batted
- Create bowler entry
- Create delivery with scoring input
- Update delivery with commentary
- Update delivery with advance info
- End Innings
- Set Match Result
- Stop Match - TODO

#### Figure outs

- Update Delivery Scoring Input
- Day No. for Test
- Session No . / Break
- Testing in Postman
- Use Redis cache for match state info like maiden overs etc

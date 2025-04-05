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

#### 2. Match

- `match_state` indicates "upcoming" if NULL, "live" if true, "completed" if false
- Once the `match_state` is false, it can't be updated

### Live Server

#### Admin Flow

- Auth Setup
- Create Match
- Upsert Match Squad
- Start Match
- Toss Decision
- Upsert Match Squad
- Create innings
- Create batter entries
- Update batter position & has batted
- Create bowler entry & set bowler 1, bowler 2
- Create delivery with scoring input (update striker/non-striker)
- Update delivery with commentary
- Update delivery with advance info
- End Innings
- Set Match Result
- Stop Match & Update Player Career Stats

#### Figure outs

- Day No. for Test
- Session No . / Break
- Testing in Postman
- Use Redis cache for match state info like maiden overs etc

### TODOs

- Image upload
- Validation
- Pagination
- Delete / Update
- Authentication
- Handle DB Errors (https://www.postgresql.org/docs/current/errcodes-appendix.html)
- Series Points Table
- Over Summary in commentary (summary & full innings commentary)

### DB Notes

#### 1. Player

- Player has stats in two categories: db & unavailable
- DB stats will be always computed from database scorecards (not set manually)
- Unavailable Stats can be manually set by admin
- DB stats will be updated automatically after each match completion
- Get Player Profile will return computed stats in a unified manner

#### Figure outs

- Use Redis cache for match state info like maiden overs etc

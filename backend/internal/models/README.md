### NOTES

- I have used `pgtype` instead of pointers. Because I feel it helps me keep things simple.
- I have also used only `pgtype.Int8` for integer values, `pgtype.Float8` for float values, instead of specific precision, to ensure consistency at the application level and avoid juggling between different precision levels.

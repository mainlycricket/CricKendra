-   Some queries have got quite complex, because:
    -   I have willingly not choosen an ORM
    -   I want to keep a single database read operation per API request
    -   I am also trying to reduce the API requests the front-end client would need to make for a single page

- Trying to keep functions as generic as possible. So they could be used in multiple scenarios. And, without creating many custom input structs, so the things remain more maintainable.

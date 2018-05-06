package database

// ItemNotFoundError informs the caller that an item was not found in the table.
type ItemNotFoundError error

// ValidationError informs the caller that the supplied value is inalid.
type ValidationError error

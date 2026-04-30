-- +goose Up
ALTER TABLE Users
ADD hashed_password text NOT NULL
DEFAULT 'not set';

-- +goose Down
ALTER TABLE Users
DROP hashed_password;
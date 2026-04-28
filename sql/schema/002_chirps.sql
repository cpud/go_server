-- +goose Up
CREATE TABLE Chirps (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    body TEXT NOT NULL,
    user_id uuid NOT NULL REFERENCES Users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE Chirps;
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS actions (
    ID serial NOT NULL,
    user_id VARCHAR(256) NOT NULL,
    data JSON NOT NULL,
    time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table actions;
-- +goose StatementEnd

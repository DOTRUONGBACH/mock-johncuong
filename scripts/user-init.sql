create
database users;

-- CREATE TABLE users.access_levels
-- (
--     id         int8        NOT NULL GENERATED BY DEFAULT AS IDENTITY,
--     "name"     varchar     NOT NULL,
--     created_at timestamptz NOT NULL,
--     updated_at timestamptz NOT NULL,
--     CONSTRAINT access_levels_pkey PRIMARY KEY (id)
-- );
--
-- INSERT INTO users.access_levels
-- (id, "name", created_at, updated_at)
-- VALUES(1, 'Guest', '2023-05-16 13:20:24.164', '2023-05-16 13:20:24.164');
-- INSERT INTO access_levels
-- (id, "name", created_at, updated_at)
-- VALUES(2, 'User', '2023-05-16 13:20:24.174', '2023-05-16 13:20:24.174');
-- INSERT INTO access_levels
-- (id, "name", created_at, updated_at)
-- VALUES(3, 'Admin', '2023-05-16 13:20:24.179', '2023-05-16 13:20:24.179');
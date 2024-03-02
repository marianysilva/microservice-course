BEGIN;

CREATE TABLE subjects
(
    id              integer         PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL UNIQUE,
    name            varchar         NOT NULL,
    objective       text            NULL,
    credit          decimal         NULL,
    workload        decimal         NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX subjects_uuid_uindex
    ON subjects (uuid);

COMMIT;
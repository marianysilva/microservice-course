BEGIN;

CREATE TABLE matrices
(
    id              integer         PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL UNIQUE,
    name            varchar         NOT NULL,
    description     text            NULL,
    course_id       integer         NOT NULL REFERENCES courses (id),
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX matrices_uuid_uindex
    ON matrices (uuid);
CREATE INDEX matrices_course_id_index
    ON matrices (course_id);
CREATE UNIQUE INDEX matrices_uindex
    ON matrices (id, course_id, deleted_at);

COMMIT;

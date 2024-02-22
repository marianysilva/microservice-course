BEGIN;

CREATE TABLE subscriptions
(
    id              integer         PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    user_uuid       uuid            NOT NULL,
    matrix_id       integer         NULL REFERENCES matrices (id),
    course_id       integer         NOT NULL REFERENCES courses (id),
    role            varchar         NULL,
    expires_at      timestamp       NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX subscriptions_uuid_uindex
    ON subscriptions (uuid);
CREATE INDEX subscriptions_matrix_id_index
    ON subscriptions (matrix_id);
CREATE INDEX subscriptions_course_id_index
    ON subscriptions (course_id);
CREATE UNIQUE INDEX subscriptions_uindex
    ON subscriptions (matrix_id, course_id, expires_at);

COMMIT;

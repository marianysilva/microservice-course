BEGIN;

CREATE TABLE matrix_subjects
(
    id              integer         PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    subject_id      integer         NOT NULL REFERENCES subjects (id),
    matrix_id       integer         NOT NULL REFERENCES matrices (id),
    is_required     boolean         NULL DEFAULT TRUE,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX matrix_subjects_subject_id_index
    ON matrix_subjects (subject_id);
CREATE UNIQUE INDEX matrix_subjects_matrix_id_index
    ON matrix_subjects (matrix_id);
CREATE UNIQUE INDEX matrix_subjects_uindex
    ON matrix_subjects (subject_id, matrix_id, deleted_at);

COMMIT;
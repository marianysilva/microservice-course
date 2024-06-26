package database

import "fmt"

const (
	returningColumns = `subscriptions.uuid, subscriptions.user_uuid, subscriptions.role,
		subscriptions.expires_at, subscriptions.created_at, subscriptions.updated_at`
	// CREATE.
	createSubscription              = "create subscription"
	createSubscriptionWithoutMatrix = "create subscription without matrix"
	// READ.
	getSubscription     = "get subscription by uuid"
	listSubscriptions   = "list subscriptions"
	courseSubscriptions = "list subscriptions by course uuid"
	userSubscriptions   = "list subscriptions by user uuid"
	// UPDATE.
	updateSubscription = "update subscription by uuid"
	// DELETE.
	deleteSubscription = "delete subscription by uuid"
)

//nolint:funlen
func queriesSubscription() map[string]string {
	return map[string]string{
		// CREATE.
		createSubscription: fmt.Sprintf(`INSERT INTO subscriptions (
				course_id, matrix_id, user_uuid, role, expires_at
			) SELECT
				courses.id, matrices.id, $3, $4, $5
			FROM courses, matrices
			WHERE
				courses.uuid = $1 AND courses.deleted_at IS NULL
				AND matrices.uuid = $2 AND matrices.deleted_at IS NULL
			RETURNING %s`, returningColumns),
		createSubscriptionWithoutMatrix: fmt.Sprintf(`INSERT INTO subscriptions (
				course_id, user_uuid, role, expires_at
			) SELECT
				courses.id, $2, $3, $4
			FROM courses
			WHERE
				courses.uuid = $1 AND courses.deleted_at IS NULL
			RETURNING %s`, returningColumns),
		// READ.
		getSubscription: fmt.Sprintf(`SELECT
				courses.uuid AS course_uuid,
				matrices.uuid AS matrix_uuid,
				%s
			FROM subscriptions
				LEFT JOIN courses ON subscriptions.course_id = courses.id
				LEFT JOIN matrices ON subscriptions.matrix_id = matrices.id
			WHERE
				subscriptions.uuid = $1 AND subscriptions.deleted_at IS NULL
				AND courses.deleted_at IS NULL
				AND matrices.deleted_at IS NULL`, returningColumns),
		listSubscriptions: fmt.Sprintf(`SELECT
				courses.uuid AS course_uuid,
				matrices.uuid AS matrix_uuid,
				%s
			FROM subscriptions
				LEFT JOIN courses ON subscriptions.course_id = courses.id
				LEFT JOIN matrices ON subscriptions.matrix_id = matrices.id
			WHERE
				subscriptions.deleted_at IS NULL
				AND courses.deleted_at IS NULL
				AND matrices.deleted_at IS NULL`, returningColumns),
		courseSubscriptions: fmt.Sprintf(`SELECT
				%s,
				matrices.uuid AS "matrices.uuid",
				matrices.code AS "matrices.code",
				matrices.name AS "matrices.name"
			FROM subscriptions
				LEFT JOIN courses ON subscriptions.course_id = courses.id
				LEFT JOIN matrices ON subscriptions.matrix_id = matrices.id
			WHERE courses.uuid = $1 AND courses.deleted_at IS NULL
				AND subscriptions.deleted_at IS NULL
				AND matrices.deleted_at IS NULL`, returningColumns),
		userSubscriptions: fmt.Sprintf(`SELECT
				%s,
				courses.uuid AS "courses.uuid",
				courses.code AS "courses.code",
				courses.name AS "courses.name",
				matrices.uuid AS "matrices.uuid",
				matrices.code AS "matrices.code",
				matrices.name AS "matrices.name"
			FROM subscriptions
				LEFT JOIN courses ON subscriptions.course_id = courses.id
				LEFT JOIN matrices ON subscriptions.matrix_id = matrices.id
			WHERE subscriptions.user_uuid = $1 AND subscriptions.deleted_at IS NULL
				AND matrices.deleted_at IS NULL
				AND courses.deleted_at IS NULL`, returningColumns),
		// UPDATE.
		updateSubscription: `UPDATE subscriptions
			SET role = $2, expires_at = $3, updated_at = NOW()
			WHERE uuid = $1 AND deleted_at IS NULL`,
		// DELETE.
		deleteSubscription: `UPDATE subscriptions
			SET deleted_at = NOW(), reason = $2
			WHERE uuid = $1 AND deleted_at IS NULL
			RETURNING uuid, reason, deleted_at`,
	}
}

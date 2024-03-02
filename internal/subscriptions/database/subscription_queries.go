package database

const (
	createSubscription  = "create subscription"
	deleteSubscription  = "delete subscription by uuid"
	getSubscription     = "get subscription by uuid"
	listSubscription    = "list subscriptions"
	updateSubscription  = "update subscription by uuid"
	courseSubscriptions = "list subscriptions by course"
	userSubscriptions   = "list subscriptions by user"
)

func queriesSubscription() map[string]string {
	return map[string]string{
		createSubscription: `INSERT INTO subscriptions (course_id, matrix_id, user_uuid, role, expires_at)
							SELECT courses.id, matrices.id, $3, $4, $5 FROM courses, matrices
							WHERE courses.uuid = $1 AND matrices.uuid = $2
							RETURNING *`,
		deleteSubscription: `UPDATE subscriptions SET deleted_at = NOW() WHERE uuid = $1 AND deleted_at IS NULL`,
		getSubscription: `SELECT
							uuid, user_uuid,
							courses.uuid AS course_uuid, courses.name AS course_name,
							matrices.uuid AS matrix_uuid, matrices.name AS matrix_name,
							role, expires_at, created_at, updated_at
						FROM subscriptions
							INNER JOIN courses ON subscriptions.course_id = courses.id
							INNER JOIN matrices ON matrices.matrix_id = matrices.id
						WHERE uuid = $1 AND deleted_at IS NULL`,
		listSubscription: `SELECT uuid, user_uuid, courses.uuid AS course_uuid, courses.name AS course_name, matrices.uuid AS matrix_uuid, matrices.name AS matrix_name, role, expires_at, created_at, updated_at
						FROM subscriptions
							INNER JOIN courses ON subscriptions.course_id = courses.id
							INNER JOIN matrices ON matrices.matrix_id = matrices.id
						WHERE deleted_at IS NULL`,
		updateSubscription: `UPDATE subscriptions SET role = $2, expires_at = $3 
							WHERE subscriptions.uuid = $1 AND deleted_at IS NULL`,
		courseSubscriptions: `SELECT uuid, user_uuid, courses.uuid AS course_uuid, courses.name AS course_name, matrices.uuid AS matrix_uuid, matrices.name AS matrix_name, role, expires_at, created_at, updated_at
								FROM subscriptions
									INNER JOIN courses ON subscriptions.course_id = courses.id
									INNER JOIN matrices ON matrices.matrix_id = matrices.id
								WHERE courses.uuid = $1 AND deleted_at IS NULL`,
		userSubscriptions: `SELECT uuid, user_uuid, courses.uuid AS course_uuid, courses.name AS course_name, matrices.uuid AS matrix_uuid, matrices.name AS matrix_name, role, expires_at, created_at, updated_at
							FROM subscriptions
								INNER JOIN courses ON subscriptions.course_id = courses.id
								INNER JOIN matrices ON matrices.matrix_id = matrices.id
							WHERE user_uuid = $1 and deleted_at IS NULL`,
	}
}

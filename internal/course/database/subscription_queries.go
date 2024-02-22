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
								VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		deleteSubscription: `UPDATE subscriptions SET deleted_at = NOW() WHERE uuid = $1 AND deleted_at IS NULL`,
		getSubscription:    `SELECT * FROM subscriptions WHERE uuid = $1 AND deleted_at IS NULL`,
		listSubscription:   `SELECT * FROM subscriptions WHERE deleted_at IS NULL`,
		updateSubscription: `UPDATE subscriptions SET user_uuid = $1, course_id = $2, matrix_id = $3, 
                         		role = $4, expires_at = $5 
								WHERE uuid = $6 AND deleted_at IS NULL RETURNING *`,
		courseSubscriptions: `SELECT * FROM subscriptions INNER JOIN courses
								ON subscriptions.course_id = courses.id
								WHERE courses.uuid = $1 AND deleted_at IS NULL`,
		userSubscriptions: `SELECT * FROM subscriptions WHERE user_uuid = $1 and deleted_at IS NULL`,
	}
}

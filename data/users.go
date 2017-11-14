package data

import "github.com/jinzhu/gorm"

// MarkUsersToFetch marks users with stale data to refresh
func MarkUsersToFetch(db *gorm.DB) {
	db.Exec(`
		UPDATE users SET fetched=false WHERE id IN (
			SELECt u.id
			FROM rankings r
			JOIN users u ON u.id=r.user_id
			WHERE DATE_PART('day', now() - u.updated_at) > ROUND(GREATEST(1, LN(completed_anime) - 3.5))
			ORDER BY completed_anime
		)
	`)
}

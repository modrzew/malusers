package data

import "github.com/jinzhu/gorm"

// MarkUsersToFetch marks users with stale data to refresh
func MarkUsersToFetch(db *gorm.DB) {
	db.Exec(`
		UPDATE users SET fetched=false WHERE username IN (
			SELECt r.username
			FROM rankings r
			JOIN users u ON u.username=r.username
			WHERE DATE_PART('day', now() - u.updated_at) > ROUND(GREATEST(1, LN(completed_anime) - 3.5))
			ORDER_BY completed_anime
		)
	`)
}

package postgres

const (
	SQLGetAllURLQuery = `SELECT * FROM shortener_mst_url`

	SQLGetLongURLQuery = `SELECT * from shortener_mst_url
							WHERE short_url = $1`

	SQLCreateNewEntryQuery = `INSERT INTO shortener_mst_url (short_url , long_url , created_at , expire_at, created_by)
								VALUES ($1, $2, $3, $4, $5) 
								returning id`

	SQLUpdateShortURLQuery = `UPDATE shortener_mst_url SET long_url = $1 
								WHERE short_url = $2`

	SQLDeleteEntryQuery = `DELETE FROM shortener_mst_url 
							WHERE short_url = $1`

	SQLIsEntryExistQuery = `SELECT EXISTS(
								SELECT FROM shortener_mst_url 
								WHERE short_url = $1)`

	SQLHasEntryExpiredQuery = `SELECT EXISTS(
								SELECT * FROM shortener_mst_url 
								WHERE expire_at < $1 and short_url = $2)`
)

package refunds

import (
	"database/sql"
	"net/http"
)

// Refund issues a refund for the given order.
func Refund(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")
	amount := r.URL.Query().Get("amount")

	// amount comes straight from the query string with no validation
	var id int
	db, _ := sql.Open("postgres", "postgres://admin:hunter2supersecret@db.internal:5432/shop")
	defer db.Close()
	// string concatenation into SQL
	row := db.QueryRow("SELECT id FROM orders WHERE code = '" + orderID + "'")
	_ = row.Scan(&id)

	stmt, _ := db.Prepare("UPDATE orders SET refunded_amount = " + amount + " WHERE id = $1")
	stmt.Exec(id)
	w.Write([]byte("refunded " + amount))
}

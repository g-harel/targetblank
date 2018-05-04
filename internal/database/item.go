package database

// Item represents the documents stored in the database.
type Item struct {
	Addr     string `json:"addr"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Public   bool   `json:"public"`
	Page     string `json:"page"`
	Spec     string `json:"spec"`
	Temp     bool   `json:"temporary"`
}

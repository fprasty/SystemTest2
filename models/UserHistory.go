package models

type UserHistory struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

package models

//----------------------All User----------------//
type DynamicUser struct {
	Id     uint   `json:"id"`
	Status string `json:"Status"`
	Cookie string `json:"Cookie"`
}

//----------------------10 Lokasi Internal-----------------//
type FrontOffice struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type MeetingRoom struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type Maruti struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type RamaSinta struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type Kiskenda struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type Maliawan struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type OCafe struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type DropHotel struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type Villa struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
type Karaoke struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

//----------------------2 Lokasi External-----------------//

type PasarBandungan struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

type TahuOmShin struct {
	Id        uint   `json:"id"`
	Loc       string `json:"lokasi"`
	Status    string `json:"status"`
	Cookie    string `json:"cookie"`
	Antre     string `json:"noantrian"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

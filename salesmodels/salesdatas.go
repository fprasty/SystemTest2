package salesmodels

type Sales struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Telp      string `json:"notelp"`
	Jabatan     string `json:"jabatan"`
	Avatar    string `json:"avatar"`
	Whatsapp  string `json:"whatsapp"`
}

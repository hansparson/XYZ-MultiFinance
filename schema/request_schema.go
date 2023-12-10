package schema

type CreateUser struct {
	FullName        string `json:"FullName" validate:"mandatory"`
	NIK             string `json:"NIK" validate:"mandatory"`
	LegalName       string `json:"LegalName" validate:"mandatory"`
	TempatLahir     string `json:"TempatLahir" validate:"mandatory"`
	TanggalLahir    string `json:"TanggalLahir" validate:"mandatory"`
	Gaji            int    `json:"Gaji" validate:"mandatory"`
	FotoKTP         string `json:"FotoKTP" validate:"mandatory"`
	FotoSelfie      string `json:"FotoSelfie" validate:"mandatory"`
	LimitOneMonth   int    `json:"LimitOneMonth" validate:"mandatory"`
	LimitTwoMonth   int    `json:"LimitTwoMonth" validate:"mandatory"`
	LimitThreeMonth int    `json:"LimitThreeMonth" validate:"mandatory"`
	LimitSixth      int    `json:"LimitSixth" validate:"mandatory"`
}

type Transaction struct {
	UserID      string  `json:"UserID" validate:"mandatory"`
	OTR         float64 `json:"OTR" validate:"mandatory"`
	AdminFee    float64 `json:"AdminFee" validate:"mandatory"`
	HargaAset   float64 `json:"HargaAset" validate:"mandatory"`
	JumlahBunga float64 `json:"JumlahBunga" validate:"mandatory"`
	Tenor       float64 `json:"Tenor" validate:"mandatory"`
	NamaAset    string  `json:"NamaAset" validate:"mandatory"`
}

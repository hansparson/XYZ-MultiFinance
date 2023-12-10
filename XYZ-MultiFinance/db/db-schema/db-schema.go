package dbschema

import (
	"time"
)

type UserLogStatus string

const (
	ACTIVE  UserLogStatus = "ACTIVE"
	LOCK    UserLogStatus = "LOCK"
	DELETED UserLogStatus = "DELETED"
)

type BillLogStatus string

const (
	SUCCESS   UserLogStatus = "SUCCESS"
	FAIL      UserLogStatus = "FAIL"
	PARTIAL   UserLogStatus = "PARTIAL"
	AVAILABLE UserLogStatus = "AVAILABLE"
)

type User struct {
	UserID       string        `json:"user_id" gorm:"uniqueIndex;type:varchar(255)"`
	UserStatus   UserLogStatus `json:"user_status"`
	NIK          string        `json:"nik" gorm:"uniqueIndex;type:varchar(255)"`
	FullName     string        `json:"full_name"`
	LegalName    string        `json:"legal_name"`
	TempatLahir  string        `json:"tempat_lahir"`
	TanggalLahir string        `json:"tanggal_lahir"`
	Gaji         float64       `json:"gaji"`
	FotoKTP      string        `json:"foto_ktp"`
	FotoSelfie   string        `json:"foto_selfie"`
}

type UserLimitBalance struct {
	UserID          string  `json:"user_id" gorm:"uniqueIndex;type:varchar(255)"`
	LimitOneMonth   float64 `json:"limit_one_month"`
	LimitTwoMonth   float64 `json:"limit_two_months"`
	LimitThreeMonth float64 `json:"limit_three_months"`
	LimitSixth      float64 `json:"limit_sixth_months"`
}

type Transaction struct {
	NomorKontrak   string    `json:"nomor_kontrak" gorm:"uniqueIndex;type:varchar(255)"`
	UserID         string    `json:"user_id"`
	TanggalKontrak time.Time `json:"tanggal_kontrak"`
	TangglUpdate   time.Time `json:"tanggal_update"`
	OTR            float64   `json:"otr"`
	AdminFee       float64   `json:"admin_fee"`
	HargaAset      float64   `json:"harga_aset"`
	JumlahCicilan  float64   `json:"jumlah_cicilan"`
	JumlahBunga    float64   `json:"jumlah_bunga"`
	Tenor          float64   `json:"tenor"`
	CicilanBulanan float64   `json:"cicilan_bulan"`
	NamaAset       string    `json:"nama_aset"`
}

type MonthlyBilling struct {
	BillingId      string    `json:"billing_id" gorm:"uniqueIndex;type:varchar(255)"`
	UserID         string    `json:"user_id"`
	NamaAset       string    `json:"nama_aset"`
	Tenor          float64   `json:"tenor"`
	HargaCicilan   float64   `json:"harga_cicilan"`
	TanggalTagihan time.Time `json:"tanggal_tagihan"`
}

type Bill struct {
	KodeBayar          string        `json:"kode_bayar" gorm:"uniqueIndex;type:varchar(255)"`
	UserID             string        `json:"user_id"`
	Terbayarkan        float64       `json:"terbayarkan"`
	TotalTagihan       float64       `json:"total_tagihan"`
	SisaTagihanBulanan float64       `json:"sisa_tagihan_bulan"`
	TanggalBayar       time.Time     `json:"tanggal_bayar"`
	SisaLimit          float64       `json:"sisa_limit"`
	StatusBayar        BillLogStatus `json:"status_bayar"`
}

type LimitSementara struct {
	LimitOneMonth   float64 `json:"limit_one_month"`
	LimitTwoMonth   float64 `json:"limit_two_months"`
	LimitThreeMonth float64 `json:"limit_three_months"`
	LimitSixth      float64 `json:"limit_sixth_months"`
}

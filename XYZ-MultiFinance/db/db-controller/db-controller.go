// dbcontroller.go

package dbcontroller

import (
	"errors"
	dbschema "xyz-multifinance/db/db-schema"

	"gorm.io/gorm"
)

type HandlersController struct {
	DB *gorm.DB
}

func Controller(db *gorm.DB) *HandlersController {
	return &HandlersController{DB: db}
}

func AddNewUser(db *gorm.DB, user *dbschema.User) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddLimitUser(db *gorm.DB, user *dbschema.UserLimitBalance) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddTransaction(db *gorm.DB, user *dbschema.Transaction) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AddMonthlyBill(db *gorm.DB, user *dbschema.MonthlyBilling) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetMonthlyBillingsByUserID(db *gorm.DB, userID string) ([]*dbschema.UserLimitBalance, error) {
	var monthlyBillings []*dbschema.UserLimitBalance
	result := db.Where("user_id = ?", userID).Find(&monthlyBillings)
	if result.Error != nil {
		return nil, result.Error
	}
	return monthlyBillings, nil
}

func UpdateLimitBalance(db *gorm.DB, userID string, limitNew float64, tenor float64) error {
	var limitColumn string

	switch tenor {
	case 1:
		limitColumn = "limit_one_month"
	case 2:
		limitColumn = "limit_two_month"
	case 3:
		limitColumn = "limit_three_month"
	case 6:
		limitColumn = "limit_sixth"
	default:
		return errors.New("Invalid tenor value")
	}

	// Lakukan update nilai limit berdasarkan kolom yang sesuai dengan tenor
	result := db.Model(&dbschema.UserLimitBalance{}).Where("user_id = ?", userID).Update(limitColumn, limitNew)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetMonthlyBill(db *gorm.DB, userID string) ([]*dbschema.MonthlyBilling, error) {
	var monthlyBill []*dbschema.MonthlyBilling
	result := db.Where("user_id = ?", userID).Find(&monthlyBill)
	if result.Error != nil {
		return nil, result.Error
	}
	return monthlyBill, nil
}

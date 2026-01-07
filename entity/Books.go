package entity

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model

	Title string  `jason:"title" valid:"stringlength(3|100)~Title must be between 3 and 100"`                               // ชือหนังสือต้องยาว 3-100 ตัวอักษร
	Price float64 `jason:"price" valid:"range(50|5000)~Price must be between 50 and 5000"`                                  // ราคาหนังสือต้องอยู ่ในช่วง 50-5000
	Code  string  `jason:"code" valid:"matches(BK),minstringlength(6)~“Code must start with BK followed by 6 digits (0-9)"` // รหัสหนังสือต้องขึนต้นด้วย "BK" ตามด้วยตัวเลข 6 ตัว (เช่น "BK123456")
}

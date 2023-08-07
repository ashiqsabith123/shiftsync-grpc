package models

type SignUp struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

type OtpStruct struct {
	Code string `json:"otp"`
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Form struct {
	Gender               string `json:"gender" gorm:"type:char(1)"`
	Marital_status       string `json:"maritalstatus"  gorm:"type:char(1)"`
	Date_of_birth        string `json:"dateofbirth"`
	P_address            string `json:"paddress"`
	C_address            string `json:"caddress"`
	Account_no           string `json:"accno"`
	Ifsc_code            string `json:"ifsccode"`
	Name_as_per_passbokk string `json:"nameinpass"`
	Pan_number           string `json:"pannumber"`
	Adhaar_no            string `json:"adhaarno"`
	Designation          string `json:"designation"`
	Photo                string `json:"photo"`
}

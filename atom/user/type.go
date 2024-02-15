package atom_user

type UserDataModel struct {
	Id_user       int
	Employee_name string
	Addresses     []UserAddress
}

type UserAddress struct {
	Address string
}

type InputUserModel struct {
	Employee_name string `json:"employee_name"`
}

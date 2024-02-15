package atom_user

import (
	"errors"
	"log"
	"reflect"
)

var error_msg string = "no rows user returned"

func GetAllDataUserUseCase() ([]UserDataModel, error) {

	getDatas, err := GetAllDataUserDB()
	if err != nil {
		return []UserDataModel{}, err
	}

	if len(getDatas) < 1 {
		log.Printf("[atom][user][resource][GetAllDataUserUseCase] %s\n", error_msg)
		return getDatas, errors.New(error_msg)
	}

	return getDatas, nil
}

func GetDataUserByIdUseCase(id int) (UserDataModel, error) {
	data, err := GetDataUserByIdDB(id)
	if err != nil {
		return UserDataModel{}, nil
	}

	if reflect.ValueOf(data).IsZero() {
		log.Printf("[atom][user][resource][GetDataUserByIdUseCase] %s\n", error_msg)
		return data, errors.New("no rows user returned")
	}

	return data, nil
}

func GetDataUserByNameUseCase(name string) (UserDataModel, error) {
	data, err := GetDataUserByNameDB(name)
	if err != nil {
		return UserDataModel{}, nil
	}

	if reflect.ValueOf(data).IsZero() {
		log.Printf("[atom][user][resource][GetDataUserByNameUseCase] %s\n", error_msg)
		return data, errors.New("no rows user returned")
	}

	return data, nil
}

func InsertUserUseCase(inputData InputUserModel) ([]UserDataModel, error) {
	// get user by name
	data, error := GetDataUserByNameDB(inputData.Employee_name)
	if error != nil {
		return []UserDataModel{}, error
	}

	// check if data is empty
	if reflect.ValueOf(data).IsZero() {
		err := InsertUserDB(inputData)
		if err != nil {
			return []UserDataModel{}, err
		}
	} else {
		return []UserDataModel{}, errors.New("user already exist")
	}

	// get all data
	getDatas, err := GetAllDataUserDB()
	if err != nil {
		return []UserDataModel{}, err
	}

	return getDatas, nil
}

func UpdateUserUseCase(id int, inputData InputUserModel) (UserDataModel, error) {
	// get user by id
	data, error := GetDataUserByIdDB(id)
	if error != nil {
		return UserDataModel{}, error
	}

	// get user by name
	dataByName, error := GetDataUserByNameDB(inputData.Employee_name)
	if error != nil {
		return UserDataModel{}, error
	}

	if !reflect.ValueOf(data).IsZero() {
		if inputData.Employee_name == dataByName.Employee_name {
			return UserDataModel{}, errors.New("user already exist")
		}
		err := UpdateUserDB(id, inputData)
		if err != nil {
			return UserDataModel{}, err
		}
	} else {
		return UserDataModel{}, nil
	}

	updatedData, error := GetDataUserByIdDB(id)
	if error != nil {
		return UserDataModel{}, error
	}

	return updatedData, nil
}

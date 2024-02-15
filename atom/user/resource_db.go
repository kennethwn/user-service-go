package atom_user

import (
	"log"
	user_db_postgresql "user_service/config/db/postgresql"
)

func GetAllDataUserDB() ([]UserDataModel, error) {
	db := user_db_postgresql.OpenConnection()
	defer db.Close() // Close the connection when we leave main() / the program terminates

	// Query
	queryString := `SELECT * FROM user_employee;`

	// Execute query
	rows, err := db.Query(queryString)
	if err != nil {
		log.Println("[atom][user][resource_db][GetAllDataUserDB] errors in query row", err.Error())
		return []UserDataModel{}, err
	}
	defer rows.Close()

	// Looping data from query result
	var datas []UserDataModel
	for rows.Next() {
		var data UserDataModel
		rows.Scan(
			&data.Id_user,
			&data.Employee_name,
		)
		datas = append(datas, data)
	}

	return datas, nil
}

func GetDataUserByIdDB(id int) (UserDataModel, error) {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	// Query
	queryString := `SELECT * FROM "user_employee" WHERE id_user = $1;`

	// Execute query
	rows, err := db.Query(queryString, id)
	if err != nil {
		log.Println("[atom][user][resource_db][GetDataUserByIdDB] errors in query row", err.Error())
		return UserDataModel{}, err
	}
	defer rows.Close()

	var data UserDataModel
	for rows.Next() {
		rows.Scan(
			&data.Id_user,
			&data.Employee_name,
		)
	}

	return data, nil
}

func GetDataUserByNameDB(name string) (UserDataModel, error) {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	// Query
	queryString := `SELECT * FROM "user_employee" WHERE employee_name = $1;`

	// Execute query
	rows, err := db.Query(queryString, name)
	if err != nil {
		log.Println("[atom][user][resource_db][GetDataUserByNameDB] errors in query row", err.Error())
		return UserDataModel{}, err
	}
	defer rows.Close()

	// Looping data from query result
	var data UserDataModel
	for rows.Next() {
		rows.Scan(
			&data.Id_user,
			&data.Employee_name,
		)
	}

	return data, nil
}

func InsertUserDB(inputData InputUserModel) error {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	queryFormat := `INSERT INTO "user_employee" ("employee_name") VALUES ($1);`
	_, err := db.Exec(queryFormat, inputData.Employee_name)
	if err != nil {
		log.Println("[atom][user][resource_db][InsertUserDB] errors in query row", err.Error())
		return err
	}

	return nil
}

func UpdateUserDB(id int, inputData InputUserModel) error {
	db := user_db_postgresql.OpenConnection()
	defer db.Close()

	// Query
	queryString := `UPDATE user_employee SET employee_name = $1 WHERE id_user = $2;`
	_, err := db.Exec(queryString, inputData.Employee_name, id)

	if err != nil {
		log.Println("[atom][user][resource_db][UpdateUserDB] errors in query row", err.Error())
		return err
	}

	return nil
}

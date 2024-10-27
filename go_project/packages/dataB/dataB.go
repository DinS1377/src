package dataB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type id struct {
	id  int64
	err error
}
type animal struct {
	id      int
	name    string
	species string
	age     int
	habitat string
	swimmer string
}

var (
	instance *sql.DB
)

func Connect(database string) (*sql.DB, error) {
	database = "root:gitler1488@/" + database
	fmt.Println(database)
	instance, err := sql.Open("mysql", database)
	return instance, err
}

func CloseConnect() {
	instance.Close()
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Add(name string, species string, age int, habitat string, swim bool) {
	db, err := Connect("animals")

	Check(err)
	defer CloseConnect()

	result, err := db.Exec("INSERT INTO animals.animal (name, species, age, habitat) values (?, ?, ?, ?)",
		name, species, age, habitat)
	Check(err)
	id := id{}
	id.id, id.err = result.LastInsertId()

	result, err = db.Exec("INSERT INTO animals.swimmer (animalId, swimmer) values (?, ?)", id.id, swim)
	Check(err)

	fmt.Println(result.LastInsertId())
}

func Del(id int) {
	db, err := sql.Open("mysql", "root:gitler1488@/animals")

	Check(err)
	defer CloseConnect()

	result, err := db.Exec("DELETE FROM animal WHERE id=(?)", id)
	Check(err)
	fmt.Println(result.LastInsertId())
}

func GetAllData() {
	db, err := Connect("animals")

	Check(err)
	defer CloseConnect()

	rows, err := db.Query("SELECT animal.id, animal.name, animal.species, animal.age, animal.habitat, swimmer.swimmer FROM animal INNER JOIN swimmer ON animal.id = swimmer.animalId")
	Check(err)
	defer rows.Close()
	zoo := []animal{}

	for rows.Next() {
		animal := animal{}
		err := rows.Scan(&animal.id, &animal.name, &animal.species, &animal.age, &animal.habitat, &animal.swimmer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		zoo = append(zoo, animal)
	}
	for _, z := range zoo {
		fmt.Println(z.id, z.name, z.species, z.age, z.habitat, z.swimmer)
	}

}
func Sorted() {
	db, err := Connect("animals")

	Check(err)
	defer db.Close()

	rows, err := db.Query("SELECT animal.id, animal.name, animal.species, animal.age, animal.habitat, swimmer.swimmer FROM animal INNER JOIN swimmer ON animal.id = swimmer.animalId ORDER BY name ASC")
	Check(err)
	defer rows.Close()
	zoo := []animal{}

	for rows.Next() {
		animal := animal{}
		err := rows.Scan(&animal.id, &animal.name, &animal.species, &animal.age, &animal.habitat, &animal.swimmer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		zoo = append(zoo, animal)
	}
	for _, z := range zoo {
		fmt.Println(z.id, z.name, z.species, z.age, z.habitat, z.swimmer)
	}
	for _, z := range zoo {
		fmt.Println(z.name)
	}

}

package llds

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

// Create table with a fixed schema (name, id, age)
// Create multiple tables within a database
// Can support multiple database
// Insert, Update, Delete a row
// Index on a specific field, auto incremented id

// Database
// Table
// Rows
// Columns
// Update as well

type Column struct {
	name        string
	dataType    string
	constraints []string
}

type Table struct {
	name    string
	columns []*Column
	rows    map[string]map[string]string
}

type Database struct {
	tables map[string]*Table
}

type CreateTableRequestDto struct {
	schema []*Column
	name   string
}

type InsertRowRequestDto struct {
	rows map[string]string
	name string
}

// UPADTE table set columnField1=a, columnField2=b where columnFieldX=2
type UpdateRowRequestDto struct {
	rowSetterMap          map[string]string
	conditionalColumnsMap map[string]string
}

//DELETE from table where  columnFieldX=2

type DeleteRowRequestDto struct {
	conditionalColumnsMap map[string]string
}

func (d *Database) CreateTable(createTableRequestDto CreateTableRequestDto) error {
	_, ok := d.tables[createTableRequestDto.name]
	if ok {
		return errors.New("table or vies already exists")
	}

	table := &Table{
		name:    createTableRequestDto.name,
		columns: createTableRequestDto.schema,
		rows: map[string]map[string]string{},
	}

	d.tables[createTableRequestDto.name] = table
	return nil

}

func (t *Table) Insert(insertRowRequestDto InsertRowRequestDto) error {
	rowId := uuid.NewString()
	row := make(map[string]string, 0)
	for _, col := range t.columns {
		row[col.name] = insertRowRequestDto.rows[strings.ToLower(col.name)]
	}

	t.rows[rowId] = row

	return nil

}

func (t *Table) Update(updateRowRequestDto UpdateRowRequestDto) error {
	for key := range t.rows {
		updateRow := true
		for key1, val1 := range updateRowRequestDto.conditionalColumnsMap {
			if t.rows[key][key1] == val1 {

			} else {
				updateRow = false
				break
			}
		}
		if updateRow {
			for key1, val1 := range updateRowRequestDto.rowSetterMap {
				t.rows[key][key1] = val1
			}

		}

	}

	return nil

}

func (t *Table) Delete(deleteRowRequestDto DeleteRowRequestDto) error {
	for key, _ := range t.rows {
		deleteRow := true
		for key1, val1 := range deleteRowRequestDto.conditionalColumnsMap {
			if t.rows[key][key1] == val1 {

			} else {
				deleteRow = false
				break
			}
		}
		if deleteRow {
			delete(t.rows, key)
		}
	}

	return nil

}

func Driver4() {
	//Create a database
	database := Database{
		tables: map[string]*Table{},
	}

	createTable := CreateTableRequestDto{
		name: "Practice",
		schema: []*Column{
			{
				name:     "name",
				dataType: "string",
			},
			{
				name:     "age",
				dataType: "string",
			},
			{
				name:     "gender",
				dataType: "string",
			},
		},
	}

	err:=database.CreateTable(createTable)
	if err!=nil{
		fmt.Println(err)
	}
	table1:=database.tables["Practice"]

	//Create Table
	// table1 := Table{
	// 	name: "Practice",
	// 	columns: []*Column{
	// 		{
	// 			name:     "name",
	// 			dataType: "string",
	// 		},
	// 		{
	// 			name:     "age",
	// 			dataType: "string",
	// 		},
	// 		{
	// 			name:     "gender",
	// 			dataType: "string",
	// 		},
	// 	},
	// 	rows: map[string]map[string]string{},
	// }

	insert := InsertRowRequestDto{
		rows: map[string]string{
			"name":   "shivam",
			"age":    "12",
			"gender": "male",
		},
	}

	insert1 := InsertRowRequestDto{
		rows: map[string]string{
			"name":   fake.MaleFirstName(),
			"age":    "12",
			"gender": "male",
		},
	}

	insert5 := InsertRowRequestDto{
		rows: map[string]string{
			"name":   fake.FemaleFirstName(),
			"age":    "12",
			"gender": "female",
		},
	}

	insert4 := InsertRowRequestDto{
		rows: map[string]string{
			"name":   fake.MaleFirstName(),
			"age":    "12",
			"gender": "male",
		},
	}

	insert3 := InsertRowRequestDto{
		rows: map[string]string{
			"name":   fake.FemaleFirstName(),
			"age":    "12",
			"gender": "female",
		},
	}

	table1.Insert(insert)
	table1.Insert(insert3)
	table1.Insert(insert4)
	table1.Insert(insert5)
	table1.Insert(insert1)

	update := UpdateRowRequestDto{
		rowSetterMap: map[string]string{
			"name": "shuvbham",
		},
		conditionalColumnsMap: map[string]string{
			"gender": "male",
		},
	}

	table1.Update(update)

	delete := DeleteRowRequestDto{
		conditionalColumnsMap: map[string]string{
			"name": "shuvbham",
		},
	}

	table1.Delete(delete)

	for _, val := range table1.rows {
		fmt.Println(val)
	}
}

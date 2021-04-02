package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	sqrl "github.com/rewardStyle/squirrel"
	"github.com/rksmannem/mapper-app/reflection_examples"
)

/*func main() {

	inStr := "Aspiration.com"
	fmt.Printf("input : %v\n", inStr)

	out := mapper.CapitalizeEveryThirdAlphanumericChar(inStr)
	fmt.Printf("out : %v\n", out)

	s := mapper.NewSkipString(3, inStr)
	mapper.MapString(s)
	fmt.Printf("output: %s\n", s)
}*/

type Employee struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"user_name"`
	Salary      float64   `db:"salary"`
	DateCreated time.Time `db:"date_created"`
}

func insert(table string, data Employee) string {
	empValue := reflect.ValueOf(&data).Elem()

	empFields := make([]string, empValue.NumField())
	dataMap := make(map[string]interface{})

	for i := 0; i < len(empFields); i++ {

		fieldType := empValue.Type().Field(i)
		tag := fieldType.Tag
		tagName := tag.Get("db")
		empFields[i] = tagName

		valueField := empValue.Field(i)
		dataMap[tagName] = valueField
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", fieldType.Name, valueField.Interface(), tagName)
	}

	b := sqrl.Insert(table).PlaceholderFormat(sqrl.Question)

	b = b.SetMap(dataMap)

	query, args, err := b.ToSql()
	if err != nil {
		panic(fmt.Errorf("createFixtureToken() ToSql failed: %v", err.Error()))
	}

	fmt.Printf("args_type: %T, args_value: %v\n", args, args)

	return query

}

func main() {
	/*emp1 := Employee{
		ID:          1,
		Name:        "e1",
		Salary:      111.11,
		DateCreated: time.Now(),
	}

	fmt.Println(insert("employee", emp1))*/

	/*emps := []Employee{
		Employee{
			ID:          1,
			Name:        "e1",
			Salary:      111.11,
			DateCreated: time.Now(),
		},
		Employee{
			ID:          2,
			Name:        "e2",
			Salary:      222.22,
			DateCreated: time.Now(),
		},
	}

	fmt.Println(insertMany("employee", emps))*/

	// 1. empty interface test
	//var flt float64 = 15
	//reflection_examples.Observe(flt)
	//
	//name  := "Ramakrishna"
	//reflection_examples.Observe(name)
	//
	//type myType string
	//reflection_examples.Observe(myType("DALLAS"))
	// 2. Kind examples
	//ifList := []interface{}{"hi", 42, func() {}}
	//reflection_examples.KindExample(ifList)

	// 3.
	//var x float64 = 3.4
	//fmt.Println("type: ", reflect.TypeOf(x))
	//fmt.Println("value: ", reflect.ValueOf(x).String())
	//
	//
	//v := reflect.ValueOf(x)
	//fmt.Println("value-type: ", v.Type())
	//fmt.Println("value-kind: ", v.Kind())
	//fmt.Println("value-value: ", v.Float())
	//
	//
	//
	//
	//type MyInt int
	//var y MyInt = 7
	//v = reflect.ValueOf(y)
	//fmt.Println("value-type: ", v.Type())
	//fmt.Println("value-kind: ", v.Kind())
	//fmt.Println("value-value: ", v.Int())

	//	4.
	//var x float64 = 3.4
	//p := reflect.ValueOf(&x)
	////v.SetFloat(7.1)
	//fmt.Println("value-type:", p.Type())
	//fmt.Println("settability of v:", p.CanSet())
	//
	//v := p.Elem()
	//fmt.Println("Settability of v:", v.CanSet())
	//
	//v.SetFloat(7.1)
	//fmt.Println(v.Interface())
	//fmt.Println("x=", x)

	// 5. structs
	//type T struct {
	//	A int
	//	B string
	//}
	//
	//t := T{A: 23, B: "skidoo"}
	//
	//s := reflect.ValueOf(t)
	//typeOfT := s.Type()
	//
	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i,
	//		typeOfT.Field(i).Name, f.Type(), f.Interface())
	//}

	//6.

	m := Manager{
		User:  getUser(1, 12, "Jack"),
		Title: "123",
	}

	/*t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v \n", t.Field(1))

	// use of FieldByIndex() method
	fmt.Printf("%#v \n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v \n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%#v \n", t.FieldByIndex([]int{0, 2}))*/

	/*	r := reflect.ValueOf(m)
		id := reflect.Indirect(r).FieldByName("Id")
		fmt.Printf("id: %v\n", id.Interface())*/

	/*	id := getField(m, "Id")
		fmt.Printf("Id: %v\n", id)

		title := getField(m, "Title")
		fmt.Printf("title: %v\n", title)*/

	tagMap := make(map[string]string)
	reflection_examples.DeepFields(m, tagMap)
	//fmt.Printf("tagMap: %v\n", tagMap)

	for _, fieldName := range tagMap {
		fieldVal := getField(m, fieldName)
		fmt.Printf("field_name: %s, field_value: %v\n", fieldName, fieldVal)
	}

}

func getField(v interface{}, field string) interface{} {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface()
}

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

type Manager struct {
	User
	Title string `db:"title"`
}

func getUser(id, age int, name string) User {
	return User{id, name, age}
}

func insertMany(tableName string, datum []Employee) string {

	if len(datum) == 0 {
		return ""
	}

	data := datum[0]
	empValue := reflect.ValueOf(&data).Elem()

	empFields := make([]string, empValue.NumField())
	for i := 0; i < len(empFields); i++ {
		fieldType := empValue.Type().Field(i)
		tag := fieldType.Tag
		tagName := tag.Get("db")
		if tagName != "" {
			empFields[i] = tagName
		}
	}

	//run := func(query string, args []interface{}, err error) (sql.Result, error) {
	//	return db.ExecContext(ctx, query, args...)
	//}
	insert := sqrl.Insert(tableName).Columns(empFields...).PlaceholderFormat(sqrl.Question)

	for _, record := range datum {

		value := reflect.ValueOf(&record).Elem()
		var values []interface{}
		for colPos := range empFields {
			colValue := value.Field(colPos).Interface()
			values = append(values, colValue)
		}
		insert = insert.Values(values...)
	}

	query, args, _ := insert.ToSql()
	query = strings.Replace(query, "INSERT INTO", "REPLACE INTO", 1)

	fmt.Printf("args_value: %v\n", args)

	fmt.Printf("query: %v\n", query)

	return query
}

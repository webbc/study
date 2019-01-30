package reflect

import (
	"reflect"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetName() {
	user := &User{Name: "test", Age: 18}
	for i := 0; i < 100; {

	}
	name := reflect.ValueOf(*user).FieldByName("Name").String()
	fmt.Println(name)
}

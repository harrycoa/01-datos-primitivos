package main

import "fmt"

// interface.. tipo de datos estructura, vacio
// solo va definicion , con metodos que no estan implementados
type User interface {
	Permisos() int // 1 - 5
	Nombre() string
}

type Admin struct {
	name string
}
func (this Admin) Permisos() int {
	return 5
}
func (this Admin) Nombre() string {
	return this.Nombre()
}

// Editor
type Editor struct {
	name string
}
func (this Editor) Permisos() int {
	return 3
}
func (this Editor) Nombre() string {
	return this.Nombre()
}


func auth(user User) string {
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + "tiene permisos admin"
	} else if permisos == 3 {
		return user.Nombre() + "tiene permisos editor"
	}
	return ""
}

func main()  {

	/*
	admin := Admin{"harri"}
	editor := Editor{"fulado"}

	fmt.Println(auth(admin))
	fmt.Println(auth(editor))
*/

	usuarios := []User{Admin{"A"},Editor{"fulano"}}
	for _,us := range usuarios {
		fmt.Println(auth(us))
	}

}
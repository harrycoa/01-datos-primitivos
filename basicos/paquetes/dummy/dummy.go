package dummy

var hola_mundo string
func init() {
	hola_mundo = "hola desde init"
}

// funciones que empiezan el nombre con Minusculas privadas
func holamundo2() string {
	return "hola_mundo"
}

// funciones que empiezan el nombre con Mayusculas son publicas
func Holamundo() string {
	return "hola mundo"
}

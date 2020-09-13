package departments

type Departamento struct {
	Direccion    string
	Inmobiliaria string
}

func GetDepartments() []Departamento {
	departments := []Departamento{
		Departamento{
			Direccion:    "Sitio al 1200",
			Inmobiliaria: "San Miguel",
		},
		Departamento{
			Direccion:    "Las Piedras al 1500",
			Inmobiliaria: "German algo",
		},
	}
	return departments
}

package maps


func CreateMap()  map[string]map[string]string{

	personalData 	:= make(map[string]string)
	employee	:= make(map[string]map[string]string)

	personalData["address"]	=	"Delhi"
	personalData["area"]	=	"Saket"
	personalData["Post"]	=	"SDE"
	personalData["phone number"]	=	"999999999"
	employee["Afzal"]	=	personalData;
	return employee;
}

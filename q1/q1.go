package q1

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1 map[string]Student, studentData2 map[string]Student) map[string]Student {
	resultado := make(map[string]Student)
	for chave, valor := range studentData1 {
		resultado[chave] = valor
	}
	for chave, valor := range studentData2 {
		if student, ok := resultado[chave]; ok {
			for materia, grade := range valor.Subjects {
				student.Subjects[materia] = grade
			}
		} else {
			resultado[chave]=valor
		}
	}
	return nil
}

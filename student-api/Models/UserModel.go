package Models


type User struct {
	Id				uint   `json:"id"`
	FirstName		string `json:"first_name"`
	LastName		string `json:"last_name"`
	DOB   			string `json:"dob"`
	Address			string `json:"address"`
	SubjectMarks	[]Subject `json:"subject_marks"`
}

type Subject struct {
	SubjectName string `json:"subject_name"`
	Marks   int        `json:"marks"`
}



func (b *User) TableName() string {
	return "student"
}


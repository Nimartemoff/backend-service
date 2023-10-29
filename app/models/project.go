package models

type Project struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Annotation string     `json:"annotation"`
	UploadDate string     `json:"uploadDate"`
	Discipline Discipline `json:"discipline"`
}

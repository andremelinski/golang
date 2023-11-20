package interfaces

type ICompanyProps struct {
	Name string `json:name`
}

type ICompanyDb struct {
	Id int64 `json:id`
	ICompanyProps
}
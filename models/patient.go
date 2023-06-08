package models

/**
*
* @author yth
* @language go
* @since 2022/11/16 17:14
 */

type Patient struct {
	Id          int
	Name        string
	IsRecover   int
	Identity    string
	Person      string
	Gender      string
	TypeDisease string
	Introduce   string
	Phone       string
	BeginDate   string
	EndDate     string
	IsDelete    int
}

func (Patient) TableName() string {
	return "patient_info"
}

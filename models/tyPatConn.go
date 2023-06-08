package models

/**
*
* @author yth
* @language go
* @since 2022/11/17 16:17
 */

type TyPatConn struct {
	Id        int
	TypeId    int
	PatientId int
}

func (TyPatConn) TableName() string {
	return "type_patient_conn"
}

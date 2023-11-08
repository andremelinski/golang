package validator

import (
	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/model"
	"github.com/asaskevich/govalidator"
)

type Ivalidator struct {}

func InitValidator()*Ivalidator{
	return &Ivalidator{}
}


func (validate Ivalidator)AlbumValidator(dataObj model.IAlbumProps)error{
	_, err := govalidator.ValidateStruct(dataObj)
	if err!=nil{
		return err
	 }
	 return nil
}

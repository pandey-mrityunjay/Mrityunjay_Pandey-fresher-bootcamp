package service 

//Data Service defines an interface with a method to be mocked

type DataService interface{
	GetData(id int)(string,error)
}

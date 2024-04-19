package designpatterns

type singleton struct{}

var instance *singleton

// Ensure that a class has only oen instance and provide a global point of access to it

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

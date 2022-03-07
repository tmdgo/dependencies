package dependencies

import (
	"reflect"

	"github.com/tmdgo/reflection/functions"
	"github.com/tmdgo/reflection/methods"
)

type Manager struct {
	dependencies map[reflect.Type]reflect.Value
}

func (manager *Manager) Init() {
	manager.dependencies = make(map[reflect.Type]reflect.Value)
}

func (manager *Manager) InitWithOtherManager(otherManager *Manager) {
	manager.dependencies = otherManager.dependencies
}

func (manager *Manager) Add(dependency interface{}) {
	manager.CallMethodByName(dependency, "Init")
	manager.dependencies[reflect.TypeOf(dependency)] = reflect.ValueOf(dependency)
}

func (manager *Manager) AddModel(model interface{}) {
	manager.dependencies[reflect.TypeOf(model)] = reflect.ValueOf(model)
}

func (manager Manager) Get(reflectType reflect.Type) interface{} {
	return manager.dependencies[reflectType]
}

func (manager Manager) CallFunc(function interface{}) []reflect.Value {
	return functions.CallFunc(function, manager.dependencies)
}

func (manager Manager) CallMethodByName(model interface{}, name string) []reflect.Value {
	return methods.CallMethodByName(model, name, manager.dependencies)
}

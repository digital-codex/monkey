package object

/*****************************************************************************
 *                                  TYPES                                    *
 *****************************************************************************/

type Environment struct {
	store map[string]Object
	outer *Environment
}

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func ExtendEnvironment(obj Closure, args []Object) *Environment {
	env := NewEnclosedEnvironment(obj.env())

	for i, param := range obj.parameters() {
		env.Set(param.Value, args[i])
	}

	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

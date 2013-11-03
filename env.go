package sandal

type TypeEnv struct {
	upper *TypeEnv
	scope map[string]Type
}

func NewTypeEnv() (ret *TypeEnv) {
	ret = new(TypeEnv)
	ret.scope = make(map[string]Type)
	return
}

func NewTypeEnvFromUpper(upper *TypeEnv) (ret *TypeEnv) {
	ret = NewTypeEnv()
	ret.upper = upper
	return
}

func (env *TypeEnv) Add(name string, ty Type) {
	env.scope[name] = ty
}

package pkg

type CodeExecutor interface {
	ExecuteCode(code string) (string, error)
}

type ProgramBuilder interface {
	BuildProgram(code string) ([]byte, error)
}

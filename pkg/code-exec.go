package pkg

type CodeExecutor interface {
	ExecuteCode(code string) (string, error)
}
package interfaces

type IHealthcheckServices interface {
	HealtcheckService() (string, error)
}

type IHealthcheckRepository interface {
}

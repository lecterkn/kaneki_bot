package port

type GenerateRepository interface {
	ReplyFunction(string) (*string, error)
	Generate(string) (*string, error)
}

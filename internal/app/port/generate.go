package port

type GenerateRepository interface {
	Generate(string) (*string, error)
}

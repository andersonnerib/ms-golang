package use_cases

import entities "command-line-argumentsC:\\Neri\\Go\\ms-golang\\internal\\entities\\pessoa.go"

type createPessoaUseCase struct {
	// recebe a instancia pra criar
}

func NewCreatePessoaUseCase() *createPessoaUseCase {
	return &createPessoaUseCase{}
}

func (u *createPessoaUseCase) Execute(name string) error {
	pessoa, err := entities.NewCategory("Joãozinho")

	if err != nil {
		return err
	}

	//persist

	return nil
}

package use_cases

import (
	"fmt"

	entities "github.com/andersonnerib/ms-golang/internal/entities"
)

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

	fmt.Println(pessoa)

	return nil
}

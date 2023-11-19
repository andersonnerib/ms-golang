package entities

import (
	"fmt"
	"time"
)

type Pessoa struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Pessoa, error) {
	pessoa := &Pessoa{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// business rules
	err := pessoa.isValid()
	if err != nil {
		return nil, err
	}

	return pessoa, err
}

func (p *Pessoa) isValid() error {
	if len(p.Name) < 5 {
		return fmt.Errorf("Invalid name!")
	}

	return nil
}

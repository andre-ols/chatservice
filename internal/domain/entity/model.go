package entity

import "errors"

type Model struct {
	Name      string `faker:"sentence"`
	MaxTokens int    `faker:"boundary_start=1024, boundary_end=1024"`
}

func NewModel(name string, maxTokens int) (*Model, error) {

	model := &Model{
		Name:      name,
		MaxTokens: maxTokens,
	}

	if err := model.Validate(); err != nil {
		return nil, err
	}

	return model, nil
}

func (m *Model) Validate() error {

	if m.Name == "" {
		return errors.New("name is empty")
	}

	if m.MaxTokens <= 0 {
		return errors.New("maxTokens is less than or equal to 0")
	}

	return nil
}

func (m *Model) GetName() string {
	return m.Name
}

func (m *Model) GetMaxTokens() int {
	return m.MaxTokens
}

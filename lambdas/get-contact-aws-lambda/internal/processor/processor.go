package processor

import (
	"uala/go-workshop/internal/repository"
	"uala/go-workshop/pkg/dto"
)

type Processor interface {
	Process(req dto.Request) (dto.Contact, error)
}

type LambdaProcessor struct {
	ContactRepository repository.Repository
}

func New(r repository.Repository) Processor {
	return &LambdaProcessor{
		ContactRepository: r,
	}
}

func (p *LambdaProcessor) Process(req dto.Request) (dto.Contact, error) {

	item, err := p.ContactRepository.GetItem(req)
	if err != nil {
		return dto.Contact{}, err
	}

	return item, nil
}

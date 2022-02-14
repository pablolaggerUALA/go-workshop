package processor

import (
	"uala/go-workshop/internal/repository"
	"uala/go-workshop/pkg/dto"
)

type Processor interface {
	Process(req dto.Request) (dto.Response, error)
}

type LambdaProcessor struct {
	ContactRepository repository.Repository
}

func New(r repository.Repository) Processor {
	return &LambdaProcessor{
		ContactRepository: r,
	}
}

func (p *LambdaProcessor) Process(req dto.Request) (dto.Response, error) {

	item, err := p.ContactRepository.DeleteItem(req)
	if err != nil {
		return dto.Response{}, err
	}

	return item, nil
}

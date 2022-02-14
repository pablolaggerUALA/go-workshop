package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"uala/go-workshop/internal/processor"
	"uala/go-workshop/pkg/dto"
)

// Responsibilidad: manejar el evento que dispara el lambda
// Validar el input
// Si es necesario, dar una respuesta
type Response = events.APIGatewayProxyResponse

type Handler interface {
	DeleteItem(ctx context.Context, req dto.Request) (Response, error)
}

type LambdaHandler struct {
	ContactProcessor processor.Processor
}

func New(p processor.Processor) Handler {
	return &LambdaHandler{
		ContactProcessor: p,
	}
}

func (h *LambdaHandler) DeleteItem(ctx context.Context, req dto.Request) (Response, error) {
	if err := validateRequest(req); err != nil {
		lambdaError := dto.LambdaError{
			Code: dto.ValidationErrorCode,
			Msg:  err.Error(),
		}

		return Response{
			StatusCode: http.StatusBadRequest,
			Body:       lambdaError.Error(),
		}, nil
	}

	contact, err := h.ContactProcessor.Process(req)
	if err != nil {
		lambdaError := dto.LambdaError{
			Code: dto.InternalServerErrorCode,
			Msg:  err.Error(),
		}

		return Response{
			StatusCode: http.StatusInternalServerError,
			Body:       lambdaError.Error(),
		}, nil
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       contact.ToJsonStr(),
	}, nil
}

func validateRequest(req dto.Request) error {
	if req.Uuid == "" {
		return &dto.ValidationError{
			Field: "uuid",
			Err:   dto.InvalidInput,
		}
	}
	return nil
}

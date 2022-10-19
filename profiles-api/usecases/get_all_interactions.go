package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

func NewGetAllInteractionsUseCase(
	logger common.ILogger,
	getInteraction IGetInteractionUseCase,
) IGetAllInteractionsUseCase {
	return &getAllInteractionsUseCase{
		logger,
		getInteraction,
	}
}

type getAllInteractionsUseCase struct {
	logger         common.ILogger
	getInteraction IGetInteractionUseCase
}

type IGetAllInteractionsUseCase interface {
	// Get all interactions.
	// A non nil error will be returned if any interactions in the list are invalid
	Do(ctx context.Context, input *GetAllInteractionsInput) (*[]entities.Interaction, error)
}

type GetAllInteractionsInput struct {
	Interactions *[]GetInteractionInput `validate:"required"`
}

func (uc *getAllInteractionsUseCase) Do(ctx context.Context, input *GetAllInteractionsInput) (*[]entities.Interaction, error) {
	if err := common.ValidateStruct(input); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var interactionErr error
	interactions := make([]entities.Interaction, len(*input.Interactions))

	for i, intrInput := range *input.Interactions {
		wg.Add(1)

		go func(i int, intr GetInteractionInput) {
			defer wg.Done()

			interaction, err := uc.getInteraction.Do(ctx, &intr)

			if err != nil {
				uc.logger.Info(ctx).Err(err).Msgf("invalid interaction: transaction id %v blockchain %v type %v address %v", intr.Interaction.Transaction.Id, intr.Interaction.Transaction.Blockchain, intr.Interaction.Type, intr.Address)
				interactionErr = err
				return
			}

			interactions[i] = *interaction

		}(i, intrInput)
	}

	wg.Wait()

	if interactionErr != nil {
		return nil, interactionErr
	}

	return &interactions, nil
}

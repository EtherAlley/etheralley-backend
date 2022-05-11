package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type GetAllInteractionsInput struct {
	Interactions *[]GetInteractionInput `validate:"required"`
}

// Get all interactions
//
// A non nil error will be returned if any interactions in the list are invalid
type IGetAllInteractionsUseCase func(ctx context.Context, input *GetAllInteractionsInput) (*[]entities.Interaction, error)

func NewGetAllInteractionsUseCase(
	logger common.ILogger,
	getInteraction IGetInteractionUseCase,
) IGetAllInteractionsUseCase {
	return func(ctx context.Context, input *GetAllInteractionsInput) (*[]entities.Interaction, error) {
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

				interaction, err := getInteraction(ctx, &intr)

				if err != nil {
					logger.Info(ctx).Err(err).Msgf("invalid interaction: transaction id %v blockchain %v type %v address %v", intr.Interaction.Transaction.Id, intr.Interaction.Transaction.Blockchain, intr.Interaction.Type, intr.Address)
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
}

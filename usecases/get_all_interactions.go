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
// Invalid Interactions are removed from the returned slice
type IGetAllInteractionsUseCase func(ctx context.Context, input *GetAllInteractionsInput) *[]entities.Interaction

func NewGetAllInteractionsUseCase(
	logger common.ILogger,
	getInteraction IGetInteractionUseCase,
) IGetAllInteractionsUseCase {
	return func(ctx context.Context, input *GetAllInteractionsInput) *[]entities.Interaction {
		if err := common.ValidateStruct(input); err != nil {
			return &[]entities.Interaction{}
		}

		var wg sync.WaitGroup

		interactions := make([]*entities.Interaction, len(*input.Interactions))
		for i, intrInput := range *input.Interactions {
			wg.Add(1)

			go func(i int, intr GetInteractionInput) {
				defer wg.Done()

				interaction, err := getInteraction(ctx, &intr)

				if err != nil {
					logger.Errf(ctx, err, "invalid input: transaction id %v blockchain %v type %v address %v", intr.Interaction.Transaction.Id, intr.Interaction.Transaction.Blockchain, intr.Interaction.Type, intr.Address)
					return
				}

				interactions[i] = interaction

			}(i, intrInput)
		}

		wg.Wait()

		trimmedInteractions := []entities.Interaction{}
		for _, interaction := range interactions {
			if interaction != nil {
				trimmedInteractions = append(trimmedInteractions, *interaction)
			}
		}

		return &trimmedInteractions
	}
}

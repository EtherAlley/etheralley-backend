package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func NewGetAllInteractionsUseCase(
	logger common.ILogger,
	getInteraction IGetInteractionUseCase,
) IGetAllInteractionsUseCase {
	return func(ctx context.Context, input *GetAllInteractionsInput) *[]entities.Interaction {
		var wg sync.WaitGroup

		interactions := make([]*entities.Interaction, len(*input.Interactions))
		for i, intrInput := range *input.Interactions {
			wg.Add(1)

			go func(i int, intr InteractionInput) {
				defer wg.Done()

				interaction, err := getInteraction(ctx, &GetInteractionInput{
					Address:     input.Address,
					Interaction: &intr,
				})

				if err != nil {
					logger.Errf(ctx, err, "invalid input: transaction id %v blockchain %v type %v address %v", intr.Transaction.Id, intr.Transaction.Blockchain, intr.Type, input.Address)
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

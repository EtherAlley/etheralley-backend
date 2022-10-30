package usecases

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
)

func NewGetLightProfile(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
	blockchainGateway gateways.IBlockchainGateway,
	databaseGateway gateways.IDatabaseGateway,
	resolveENSName IResolveENSNameUseCase,
	getNonFungibleToken IGetNonFungibleTokenUseCase,
	offchainGateway gateways.IOffchainGateway,
) IGetLightProfileUseCase {
	return &getLightProfileUseCase{
		logger,
		cacheGateway,
		blockchainGateway,
		databaseGateway,
		resolveENSName,
		getNonFungibleToken,
		offchainGateway,
	}
}

type getLightProfileUseCase struct {
	logger              common.ILogger
	cacheGateway        gateways.ICacheGateway
	blockchainGateway   gateways.IBlockchainGateway
	databaseGateway     gateways.IDatabaseGateway
	resolveENSName      IResolveENSNameUseCase
	getNonFungibleToken IGetNonFungibleTokenUseCase
	offchainGateway     gateways.IOffchainGateway
}

type IGetLightProfileUseCase interface {
	// get a minimaly hydrated version of the profile that will be faster to fetch
	Do(ctx context.Context, input *GetLightProfileInput) (*entities.Profile, error)
}

type GetLightProfileInput struct {
	Address string `validate:"required,eth_addr"`
}

func (uc *getLightProfileUseCase) Do(ctx context.Context, input *GetLightProfileInput) (*entities.Profile, error) {
	if err := common.ValidateStruct(input); err != nil {
		return nil, err
	}

	uc.logger.Debug(ctx).Msgf("getting light profile %v", input.Address)

	profile, err := uc.cacheGateway.GetProfileByAddress(ctx, common.LIGHT, input.Address)

	if err == nil {
		uc.logger.Debug(ctx).Msgf("cache hit %v", input.Address)
		return profile, nil
	}

	uc.logger.Debug(ctx).Msgf("cache miss %v", input.Address)

	dbProfile, err := uc.databaseGateway.GetProfileByAddress(ctx, input.Address)

	if err == nil {
		profile = &entities.Profile{
			Address:        input.Address,
			Banned:         dbProfile.Banned,
			LastModified:   dbProfile.LastModified,
			DisplayConfig:  dbProfile.DisplayConfig,
			ProfilePicture: dbProfile.ProfilePicture,
		}
	} else if errors.Is(err, common.ErrNotFound) {
		profile = &entities.Profile{
			Address: input.Address,
			Banned:  false,
		}
	} else if err != nil {
		uc.logger.Warn(ctx).Err(err).Msgf("err getting db profile %v", input.Address)
		return nil, err
	}

	uc.logger.Debug(ctx).Msgf("db hit %v", input.Address)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()

		// Not all addresses have an ens name. We should not propigate an error for this
		name, _ := uc.resolveENSName.Do(ctx, &ResolveENSNameInput{
			Address: profile.Address,
		})

		profile.ENSName = name
	}()

	go func() {
		defer wg.Done()

		profile.StoreAssets = &entities.StoreAssets{}
		if balances, err := uc.blockchainGateway.GetStoreBalanceBatch(ctx, input.Address, &[]string{common.STORE_PREMIUM, common.STORE_BETA_TESTER}); err == nil {
			profile.StoreAssets.Premium = balances[0].Cmp(big.NewInt(0)) == 1
			profile.StoreAssets.BetaTester = balances[1].Cmp(big.NewInt(0)) == 1
		}
	}()

	go func() {
		defer wg.Done()

		if profile.ProfilePicture != nil {
			if pic, err := uc.getNonFungibleToken.Do(ctx, &GetNonFungibleTokenInput{
				Address: input.Address,
				NonFungibleToken: &NonFungibleTokenInput{
					TokenId: profile.ProfilePicture.TokenId,
					Contract: &ContractInput{
						Blockchain: profile.ProfilePicture.Contract.Blockchain,
						Interface:  profile.ProfilePicture.Contract.Interface,
						Address:    profile.ProfilePicture.Contract.Address,
					},
				},
			}); err == nil {
				profile.ProfilePicture = pic
			}
		} else {
			// take the first nft for profile picture
			if profilePics, err := uc.offchainGateway.GetNonFungibleTokens(ctx, input.Address); err == nil && len(*profilePics) > 0 {
				pics := *profilePics
				profile.ProfilePicture = &pics[0]
			}
		}
	}()

	wg.Wait()

	uc.cacheGateway.SaveProfile(ctx, common.LIGHT, profile)

	return profile, nil
}

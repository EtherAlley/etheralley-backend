package usecases

import (
	"context"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/settings"
)

func NewGetSpotlightProfileUseCase(
	logger common.ILogger,
	settings settings.ISettings,
	getProfile IGetProfileUseCase,
) IGetSpotlightProfileUseCase {
	return &getSpotlightProfileUseCase{
		logger,
		settings,
		getProfile,
	}
}

type getSpotlightProfileUseCase struct {
	logger     common.ILogger
	settings   settings.ISettings
	getProfile IGetProfileUseCase
}

type IGetSpotlightProfileUseCase interface {
	Do(ctx context.Context, input *GetSpotlightProfileInput) (*entities.Profile, error)
}

type GetSpotlightProfileInput struct {
}

// Read the spotlight profile address from the env variables and call the standard get profile usecase to get a fully hydrated profile
func (uc *getSpotlightProfileUseCase) Do(ctx context.Context, input *GetSpotlightProfileInput) (*entities.Profile, error) {
	spotlightAddress := uc.settings.SpotlightProfileAddress()

	profile, err := uc.getProfile.Do(ctx, &GetProfileInput{Address: spotlightAddress})

	if err != nil {
		uc.logger.Warn(ctx).Err(err).Msgf("err getting spotlight profile %v", spotlightAddress)
	}

	return profile, err
}

package app

import (
	"github.com/KasperSaaby/calculatron-service/internal/app/validator"
	"go.uber.org/fx"
)

var Setup = fx.Options(
	fx.Provide(
		validator.NewCalculationInputValidator,
		NewCalculatorService,
		NewCalculatorServiceDecorator,
		NewHistoryService,
		func(i *CalculatorService) Calculator { return i },
	),
)

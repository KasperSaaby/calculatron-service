package domain

import (
	"github.com/KasperSaaby/calculatron-service/internal/domain/operations"
	"go.uber.org/fx"
)

var Setup = fx.Options(
	fx.Provide(
		operations.NewOperationFactory,
	),
)

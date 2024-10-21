package core

import (
	"go.uber.org/fx"
)

func BuildCoreModule() fx.Option {
	coreModule := fx.Module("core",
		fx.Provide(
			fx.Annotate(NewThorWalletRepository, fx.As(new(IWalletRepository)), fx.ResultTags(`name:"thor"`)),
			fx.Annotate(NewLokiWalletRepository, fx.As(new(IWalletRepository)), fx.ResultTags(`name:"loki"`)),
			fx.Annotate(
				NewWalletServiceImpl,
				fx.As(new(IWalletService)),
				fx.ParamTags(`name:"thor"`),
			),
		),
		fx.Decorate(
			fx.Annotate(NewCachedWalletService, fx.As(new(IWalletService))),
		),
		fx.Invoke(Run),
	)
	err := fx.ValidateApp(coreModule)
	if err != nil {
		panic(err)
	}
	return coreModule
}

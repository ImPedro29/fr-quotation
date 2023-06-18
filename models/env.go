package models

type Env struct {
	// Port that server will run
	Port string `env:"PORT" envDefault:"4002"`
	// DebugMode activate zap in debug mode and fiber logs
	DebugMode bool `env:"DEBUG_MODE" envDefault:"true"`

	PostgresPort     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresSSLMode  string `env:"POSTGRES_SSLMODE" envDefault:"disable"`
	PostgresDatabase string `env:"POSTGRES_DATABASE,required"`
	PostgresUser     string `env:"POSTGRES_USER,required"`
	PostgresHost     string `env:"POSTGRES_HOST,required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required"`

	FreteRapidoApi          string `env:"FRETE_RAPIDO_API,required"`
	FreteRapidoIdentity     string `env:"FRETE_RAPIDO_IDENTITY,required"`
	FreteRapidoToken        string `env:"FRETE_RAPIDO_TOKEN,required"`
	FreteRapidoPlatformCode string `env:"FRETE_RAPIDO_PLATFORM_CODE,required"`
	FreteRapidoCep          int64  `env:"FRETE_RAPIDO_CEP,required"`
}

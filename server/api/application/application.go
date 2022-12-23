package application

type application struct {
	*ApplicationBundle
}

type ApplicationBundle struct {
	// ServerConfig             *config.ServerConfig
	// Repository               persistence.RepositoryInterface
}

type ApplicationInterface interface {
}

func NewApplication(bdl *ApplicationBundle) ApplicationInterface {
	return &application{bdl}
}

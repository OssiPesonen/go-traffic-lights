package app

import (
	"log"

	"github.com/ossipesonen/providers-app/internal/app/auth"
	"github.com/ossipesonen/providers-app/internal/app/core/models"
	"github.com/ossipesonen/providers-app/internal/app/core/providers"
	"github.com/ossipesonen/providers-app/internal/app/core/users"
	"github.com/ossipesonen/providers-app/internal/config"
	"github.com/ossipesonen/providers-app/pkg/database"
)

// Define the APIs for components

type UserService interface {
	// Authenticates a user using provided email and password.
	// Returns the user resource if successful
	Authenticate(email string, password string) (*models.User, error)
	// Creates a new user resource
	// Returns the resource
	CreateUser(userInfo *models.UserInfo) (*models.User, error)
	// Finds a user resource using only the ID
	Find(userId int) (*models.User, error)
	// Generate access and refresh token for user
	GenerateTokens(user *models.User) (*auth.IssuedTokens, error)
	// Generate new pair of access and refresh token for user
	RefreshTokens(refreshToken string) (*auth.IssuedTokens, error)
	// Revoke a specific refresh token for given user
	RevokeRefreshToken(refreshToken string) error
	// Revoke all active refresh tokens for user
	RevokeAllRefreshTokens(userId int) error
}

type ProviderService interface {
	// Lists all providers
	ListProviders() (*[]models.Provider, error)
	// List providers for given user
	ListProvidersForUser(userId int) (*[]models.Provider, error)
	// Fetch a single provider using an identifier
	// Returns the provider or error. Error can be sql.ErrnoRows,
	// which indicates no resource found
	GetProvider(id int) (*models.Provider, error)
	// Create a new provider in the system
	CreateProvider(provider *models.Provider) (int, error)
	// Search
	Search(searchwords []string) (*[]models.Provider, error)
}

type Services struct {
	User     UserService
	Provider ProviderService
}

type App struct {
	Services *Services
}

// Bootstrap application
// Setup repositories and services to be used by Server methods
func New(config *config.Config, db database.Database, logger *log.Logger) *App {
	srvc := &Services{}

	auth, err := auth.New(config.Auth.JWTSecret)

	if err != nil {
		logger.Fatalf("unable to init auth service :%v", err)
	}

	// Repositores should have db as a dependency.
	// Repositories should not be directly accessible.
	userRepository := users.NewUserRepository(db, logger)
	providerRepository := providers.NewProviderRepository(db, logger)

	// Services should have repository as a depency if data access is required
	// Service is the entry point to a component.
	srvc.User = users.NewUserService(userRepository, auth, logger)
	srvc.Provider = providers.NewProviderService(providerRepository, logger)

	return &App{
		Services: srvc,
	}
}

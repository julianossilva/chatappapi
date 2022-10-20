package servicecontainer

import (
	"github.com/jackc/pgx/v5"
)

type ServiceContainer interface {
	GetMessage() string
	GetDBConnection() *pgx.Conn
}

func New() (ServiceContainer, error) {
	sc := &ServiceContainerImpl{}

	// if conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL")); err != nil {
	// 	return nil, err
	// } else {
	// 	sc.db = conn
	// }

	return sc, nil
}

type ServiceContainerImpl struct {
	db *pgx.Conn
}

// GetDBConnection implements ServiceContainer
func (sc *ServiceContainerImpl) GetDBConnection() *pgx.Conn {
	return sc.db
}

// GetMessage implements ServiceContainer
func (*ServiceContainerImpl) GetMessage() string {
	return "Hello Service Container!!!"
}

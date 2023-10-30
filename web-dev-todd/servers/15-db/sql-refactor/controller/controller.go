package controller

import (
	album_controller "github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/controller/album"
	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/repository"
)

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type (
	// AlbumController represents the controller for operating on the Album resource
	Controllers struct{
		albumController *album_controller.AlbumController
	}
)

func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		albumController: album_controller.InitAlbumController(repositories.AlbumRepo),
	}
}

package clients

import (
	"database/sql"
	"sync"
)

var (
	hdl     *handler
	hdlOnce sync.Once

	svc     *service
	svcOnce sync.Once

	repo     *repository
	repoOnce sync.Once
)

func ProvideHandler(svc ClientService) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{
			svc: svc,
		}
	})
	return hdl
}

func ProvideService(repo ClientRepository) *service {
	hdlOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})
	return svc
}

func ProvideRepository(db *sql.DB) *repository {
	repoOnce.Do(func() {
		repo = &repository{
			db: db,
		}
	})
	return repo
}

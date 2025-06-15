package bootstrap

import (
	"github.com/bintangyosua/echaaas/internal/auth"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func BuildContainer(db *gorm.DB) *dig.Container {
	c := dig.New()

	c.Provide(func() *gorm.DB {
		return db
	})

	c.Provide(auth.NewRepository)
	c.Provide(auth.NewService)
	c.Provide(auth.NewHandler)
	
	return c
}

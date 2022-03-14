package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/web3art/g/internal/pkg/db"
)

/// 这里对一些服务进行初始化聚合
func xinit(c *cli.Context) error {
	inits := []func() error{}

	dsn := os.Getenv("DB_DSN")

	inits = append(inits, func() error {
		return db.Init(dsn == "g.db", dsn)
	})

	for _, init := range inits {
		if err := init(); err != nil {
			return err
		}
	}

	return nil
}

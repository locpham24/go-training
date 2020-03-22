package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/locpham24/go-training/week3-exercise/service"
	"github.com/urfave/cli/v2"
)

var Test = cli.Command{
	Name:  "test",
	Usage: "Test the application",
	Action: func(c *cli.Context) error {
		db := c.App.Metadata["db"].(*gorm.DB)
		testSvc := service.NewAPIService(db)
		testSvc.Start()
		return nil
	},
}

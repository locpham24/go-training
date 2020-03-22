package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/locpham24/go-training/week3-exercise/service"
	"github.com/urfave/cli/v2"
)

var Start = cli.Command{
	Name:  "start",
	Usage: "Start the application",
	Action: func(c *cli.Context) error {
		db := c.App.Metadata["db"].(*gorm.DB)
		apiSvc := service.NewAPIService(db)
		apiSvc.Start()
		return nil
	},
}

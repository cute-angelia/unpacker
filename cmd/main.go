package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path"
	"unpacker/internal/unpacker"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			fileIn := cCtx.Args().Get(0)
			log.Println("解压：", fileIn)
			if p, err := unpacker.GetUnpacker(fileIn); err != nil {
				return err
			} else {
				return p.Extract(path.Dir(fileIn))
			}
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

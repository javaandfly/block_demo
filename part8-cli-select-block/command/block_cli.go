package command

import (
	"log"
	"os"
	"publicChain/part8-cli-select-block/blc"

	"github.com/boltdb/bolt"
	"github.com/urfave/cli"
)

const (
	createBlockFlag      = "cb"
	createFirstBlockFlag = "cfb"
	selectBlockChain     = "sbc"
)

func RunBlcokCli() {
	// cli.AppHelpTemplate = fmt.Sprintf(
	// 	`%sExplanation:
	// 	This command line reads the configuration according to the given path,
	// 	and supports environment variable override and configuration file conversion

	// 	`, cli.AppHelpTemplate)
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  createFirstBlockFlag,
				Usage: "create genesis block",
			},
			&cli.StringFlag{
				Name:  createBlockFlag,
				Usage: "add block",
			},
			&cli.BoolFlag{
				Name:  selectBlockChain,
				Usage: "select blockchain",
			},
		},
		Action: MainAction,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Panicln(err)
		return
	}
}

func MainAction(c *cli.Context) error {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if c.BoolT(createFirstBlockFlag) {
		blockChain := blc.CreateBlockChainWithGenesisBlock(db)
		blockChain.Blocks[0].PrintfBlock()
		if err != nil {
			return err
		}

		return nil
	}
	if c.String(createBlockFlag) != "" {
		blc.AddBlockToBlockChain(c.String(createBlockFlag), db)
		block, err := blc.GetNewBlock(db)
		if err != nil {
			return err
		}
		block.PrintfBlock()

		return nil
	}
	if c.BoolT(selectBlockChain) {
		blc.GetAllBlock(db, []byte("l"))

		return nil
	}
	return nil
}

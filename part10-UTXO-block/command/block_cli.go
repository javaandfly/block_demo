package command

import (
	"log"
	"os"
	"publicChain/part9-transaction-block/blc"

	"strings"

	"github.com/boltdb/bolt"
	"github.com/urfave/cli/v2"
)

const (
	createBlockFlag      = "cb"
	createFirstBlockFlag = "cfb"
	selectBlockChain     = "sbc"
	sendTransaction      = "send"
	senderFromFlag       = "from"
	senderToFlag         = "to"
	senderAmountFlag     = "amount"
)

func RunBlcokCli() {
	// cli.AppHelpTemplate = fmt.Sprintf(
	// 	`%sExplanation:
	// 	This command line reads the configuration according to the given path,
	// 	and supports environment variable override and configuration file conversion

	// 	`, cli.AppHelpTemplate)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  sendTransaction,
				Usage: "send transcation",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: senderFromFlag, Aliases: []string{"f"}},
					&cli.StringFlag{Name: senderToFlag, Aliases: []string{"t"}},
					&cli.StringFlag{Name: senderAmountFlag, Aliases: []string{"a"}},
				},
				Action: func(c *cli.Context) error {
					db, err := bolt.Open("my.db", 0600, nil)
					if err != nil {
						log.Fatal(err)
					}
					defer db.Close()
					sendMessage := &blc.TransactionMessage{}
					if c.String(senderFromFlag) != "" {
						fromString := c.String(senderFromFlag)[1 : len(c.String(senderFromFlag))-1]

						sendMessage.From = strings.Split(fromString, ",")
					}
					if c.String(senderToFlag) != "" {
						toString := c.String(senderToFlag)[1 : len(c.String(senderToFlag))-1]
						sendMessage.To = strings.Split(toString, ",")
					}
					if c.String(senderAmountFlag) != "" {
						amountString := c.String(senderAmountFlag)[1 : len(c.String(senderAmountFlag))-1]
						sendMessage.Amount = strings.Split(amountString, ",")
					}

					blc.AddBlockToBlockChain(sendMessage, db)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
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
	if c.String(createFirstBlockFlag) != "" {

		blockChain := blc.CreateBlockChainWithGenesisBlock(db, c.String(createFirstBlockFlag))
		blockChain.Blocks[0].PrintfBlock()

		return nil
	}
	// if c.String(createBlockFlag) != "" {
	// 	blc.AddBlockToBlockChain([]*blc.Transaction{}, db)
	// 	block, err := blc.GetNewBlock(db)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	block.PrintfBlock()

	// 	return nil
	// }
	if c.Bool(selectBlockChain) {
		blc.GetAllBlock(db, []byte("l"))

		return nil
	}
	return nil
}

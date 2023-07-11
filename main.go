package main

import (
	"context"
	"fmt"

	"github.com/bjartek/overflow"
)

func main() {

	o := overflow.Overflow(overflow.WithNetwork("mainnet"))

	ctx := context.Background()
	block, err := o.GetBlockAtHeight(ctx, 54564338)
	if err != nil {
		panic(err)
	}

	tx, _, err := o.Flowkit.GetTransactionsByBlockID(ctx, block.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(tx))

}

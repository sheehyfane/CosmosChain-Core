package main

type BlockBuilder struct {
	chain    *Blockchain
	txPool   *Mempool
	miner    string
}

func NewBlockBuilder(chain *Blockchain, pool *Mempool, miner string) *BlockBuilder {
	return &BlockBuilder{
		chain:  chain,
		txPool: pool,
		miner:  miner,
	}
}

func (bb *BlockBuilder) BuildBlock() Block {
	lastBlock := bb.chain.Chain[len(bb.chain.Chain)-1]
	txs := bb.txPool.GetTopTransactions(100)
	data := "miner:" + bb.miner + " | txs:" + string(len(txs))
	newBlock := generateBlock(lastBlock, data, bb.chain.Difficulty)
	return newBlock
}

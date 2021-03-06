package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbName = "blockChain.db"
const blockTableName = "blocks"

type BlockChain struct {
	//Blocks []*Block //存储有序得区块
	Tip []byte //最新得区块得Hash
	DB  *bolt.DB
}

//2.增加区块到区块链里面
//func (blc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {
//	//创建新区块
//	newBlock := NewBlock(data, height, prevHash)
//	//往链里面添加区块
//	blc.Blocks = append(blc.Blocks, newBlock)
//}

//1.创建带有创世区块得区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {

	//创建或打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}
		if b != nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.....")
			//将创世区块存储到表当中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//存储最新区块得Hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blockHash = genesisBlock.Hash
		}
		return nil
	})
	//defer db.Close()

	//返回区块链对象
	//return &BlockChain{Blocks: []*Block{genesisBlock}}
	return &BlockChain{blockHash, db}
}

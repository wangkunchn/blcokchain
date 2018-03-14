package main

import (
	//"fmt"
	"strconv"
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
	"os"
)
/*
区块接口
 */
type Block struct {
	Timestamp 		int64 //时间戳
	Data 	 		[]byte //当前区块存放的信息 比如 比特币或者一些 其他账单信息，用户的行为信息
	PrevBlcokHash 	[]byte //上一个区块 加密的哈希
	Hash 			[]byte //当前区块的哈希
}
/*
Block 这个结构体  绑定的一个方法
 */
func(this *Block) SetHash(){

	//将本区块的TimeStap + Data + PrevBlockHasj  -----> Hash
	//将时间戳由整形--》 二进制
	//timestap := []byte(strconv.FormatInt(this.Timestamp, 10))

	timestamp := []byte(strconv.FormatInt(this.Timestamp, 10))

	//将三个二进制的属性 进行拼接
	headers := bytes.Join([][]byte{this.PrevBlcokHash, this.Data, timestamp}, []byte{})

	//将拼接后的headers 进行 SHA256加密
	preHash := sha256.Sum256(headers)

	this.Hash = preHash[:]

}
/*
新建一个区块
 */
func NewBlock(data string, preBlockHash []byte) *Block{
	// 生成一个区块
	block := Block{}
	//给当前区块赋值   时间  data  prehash
	block.Timestamp = time.Now().Unix() //  Unix()--> Time 类型转成 int64
	block.Data = []byte(data)
	block.PrevBlcokHash = preBlockHash
	//拼接 加密
	block.SetHash()
	// 赋值好的区块 return
	return &block
}
/*
定义区块链的结构体
 */
type BlockChain struct {
	Blocks []*Block //有序的区块链
}

//将区块添加到区块链中
func (this *BlockChain) AddBlock(data string) {
	//1.得到 新添加区块的  前区块的hash
	proBlock := this.Blocks[len(this.Blocks)-1]
	//2.根据  data 创建一个新的区块
	block := NewBlock(data, proBlock.Hash)
	//3.依照前区块 和新区块  添加到blocks中
	this.Blocks = append(this.Blocks, block)
}

//新建i一个创世块
func NewGenesisBlock() *Block {
	genesisBlock := Block{}
	genesisBlock.Data = []byte("创世区块")
	genesisBlock.PrevBlcokHash = []byte{}
	return &genesisBlock       //没有时间？？
}

//区块链 = 创世块 --》区块--》区块
func NewBlockChain()  *BlockChain {
	blockChain := BlockChain{}
	blockChain.Blocks = []*Block{NewGenesisBlock()}
	return &blockChain
}

func main() {
	//创建一条区块链 bc
	bc := NewBlockChain()
	var cmd string
	for  {
		fmt.Println("按'1' 添加一条信息数据 到区块链中")
		fmt.Println("按'2' 遍历当前的区块链都有哪些区块信息")
		fmt.Println("按 其他 推出")
		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "1":
			input := make([]byte, 1024)
			//添加一个区块
			fmt.Println("请输入区块链的行为数据（要添加保存的数据）：")
			os.Stdin.Read(input)
			bc.AddBlock(string(input))
		case "2":
			//比哪里整个区块链
			for i, block := range bc.Blocks {
				fmt.Println("==============")
				fmt.Println("第", i, "个区块的信息：")
				fmt.Println("PrevHash: %x\n", block.PrevBlcokHash)
				fmt.Println("data: %x\n", block.Data)
				fmt.Println("hash: %x\n", block.Hash)
				fmt.Println("Timestamp: %x\n", block.Timestamp)
				fmt.Println("==============")
			}
		default:
			//退出程序
			fmt.Println("您已经退出")
			return
		}

	}
}


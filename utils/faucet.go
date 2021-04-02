package utils

import (
	"fmt"
	"github.com/JFJun/go-substrate-crypto/crypto"
	"github.com/JFJun/stafi-substrate-go/client"
	"github.com/JFJun/stafi-substrate-go/expand"
	"github.com/JFJun/stafi-substrate-go/tx"
)

func FaucetTestcoin(destAddress string) {
	// 1. 初始化rpc客户端
	c, err := client.New("wss://testnet200.zenlink.pro")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//2. 如果某些链（例如：chainX)的地址的字节前面需要0xff,则下面这个值设置为false
	expand.SetSerDeOptions(false)
	from := "5EfAjCbHodmNFoetihd8pJY5ikTZwLF581uBiG49tHxbhHfe"
	to := destAddress
	amount := uint64(100)
	//3. 获取from地址的nonce
	acc, err := c.GetAccountInfo(from)
	//if err != nil {
	//	t.Fatal(err)
	//}
	nonce := uint64(acc.Nonce)
	//4. 创建一个substrate交易，这个方法满足所有遵循substrate 的交易结构的链
	transaction := tx.NewSubstrateTransaction(from, nonce)
	//5. 初始化metadata的扩张结构
	ed, err := expand.NewMetadataExpand(c.Meta)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//6. 初始化Balances.transfer的call方法
	call, err := ed.BalanceTransferCall(to, amount)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//7. 设置交易的必要参数
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(), c.GetGenesisHash()).
		SetSpecAndTxVersion(uint32(c.SpecVersion), uint32(c.TransactionVersion)).
		SetCall(call) //设置call
	//8. 签名交易
	sig, err := transaction.SignTransaction("0xadc2f7ad0b95a1ee735060c584312a488f8825b12bd91c7c62c601ea17b02edd", crypto.Sr25519Type)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//9. 提交交易
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		//t.Fatal(err)
	}
	//10. txid
	txid := result.(string)
	fmt.Println(txid)
}

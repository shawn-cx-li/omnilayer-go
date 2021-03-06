package btcsuite

import (
	"github.com/shawn-cx-li/omnilayer-go/btcjson"
)

func (c *Client) GetBlockChainInfo() (btcjson.GetBlockChainInfoResult, error) {
	return futureGetBlockChainInfo(c.do(btcjson.GetBlockChainInfoCommand{})).Receive()
}

func (c *Client) ListUnspent(cmd btcjson.ListUnspentCommand) (btcjson.ListUnspentResult, error) {
	return futureListUnspent(c.do(cmd)).Receive()
}

func (c *Client) CreateRawTransaction(cmd btcjson.CreateRawTransactionCommand) (btcjson.CreateRawTransactionResult, error) {
	return futureCreateRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) ImportAddress(address string, rescan bool) error {
	return futureImportAddress(c.do(btcjson.ImportAddressCommand{Adress: address, Rescan: rescan})).Receive()
}

func (c *Client) SendRawTransaction(cmd btcjson.SendRawTransactionCommand) (btcjson.SendRawTransactionResult, error) {
	return futureSendRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) SignRawTransaction(cmd btcjson.SignRawTransactionCommand) (btcjson.SignRawTransactionResult, error) {
	return futureSignRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) SignRawTransactionWithKey(cmd btcjson.SignRawTransactionWithKeyCommand) (btcjson.SignRawTransactionWithKeyResult, error) {
	return futureSignRawTransactionWithKey(c.do(cmd)).Receive()
}

func (c *Client) GetNewAddress(cmd btcjson.GetNewAddressCommand) (btcjson.GetNewAddressResult, error) {
	return futureGetNewAddress(c.do(cmd)).Receive()
}

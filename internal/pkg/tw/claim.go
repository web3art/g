package tw

import (
	"context"
	"math/big"
	"os"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/web3art/g/internal/pkg/dao/model"
	"github.com/web3art/g/internal/pkg/db"
	"github.com/web3art/g/pkg/contracts"
	"github.com/web3art/g/pkg/rpcm"
	"gorm.io/gorm"
)

func (tw *Tweet) Claim() error {
	var tweetWaitToClaim model.TweetWaitToClaim

	if err := db.DB().Model(&model.TweetWaitToClaim{}).Where("claimed = false and author_id in (?)", db.DB().Table("tweet_author_addresses").Select("author_id")).First(&tweetWaitToClaim).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		return err
	}

	var tad model.TweetAuthorAddress
	if err := db.DB().Model(&model.TweetAuthorAddress{}).Where("author_id = ?", tweetWaitToClaim.AuthorId).First(&tad).Error; err != nil {
		return err
	}

	return db.DB().Transaction(func(tx *gorm.DB) error {
		tweetWaitToClaim.Claimed = true
		if err := tx.Save(&tweetWaitToClaim).Error; err != nil {
			return err
		}

		network := rpcm.Network(os.Getenv("ETH_NETWORK"))
		address := os.Getenv("SWORD_ADDRESS")

		client, err := rpcm.GetEthClient(network)

		if err != nil {
			return err
		}

		w3sc, err := contracts.NewWeb3Sword(common.HexToAddress(address), client)

		if err != nil {
			return err
		}

		chindId, err := client.ChainID(context.Background())

		if err != nil {
			return err
		}

		pkv := os.Getenv("PRIVATE_KEY")

		kp, err := secp256k1.NewKeypairFromString(pkv)

		if err != nil {
			return err
		}

		opts, err := bind.NewKeyedTransactorWithChainID(kp.PrivateKey(), chindId)

		if err != nil {
			return err
		}

		opts.GasPrice = big.NewInt(55000000000)

		if _, err := w3sc.SocialClaim(opts, common.HexToAddress(tad.Address), big.NewInt(int64(tweetWaitToClaim.TokenId)), 6); err != nil {
			return err
		}

		return nil
	})
}

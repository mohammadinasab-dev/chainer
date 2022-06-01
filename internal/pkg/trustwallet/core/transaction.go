package core

// #cgo CFLAGS: -I/home/hosein/Desktop/PRJ/wallet-core/include
// #cgo LDFLAGS: -L/home/hosein/Desktop/PRJ/wallet-core/build -L/home/hosein/Desktop/PRJ/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWAnySigner.h>
import "C"

import (
	"chainer/internal/pkg/trustwallet/types"

	"google.golang.org/protobuf/proto"
)

func CreateSignedTx(inputData proto.Message, ct CoinType, outputData proto.Message) error {
	ibytes, _ := proto.Marshal(inputData)
	idata := types.TWDataCreateWithGoBytes(ibytes)
	defer C.TWDataDelete(idata)

	odata := C.TWAnySignerSign(idata, C.enum_TWCoinType(ct))
	defer C.TWDataDelete(odata)

	err := proto.Unmarshal(types.TWDataGoBytes(odata), outputData)
	if err != nil {
		return err
	}
	return nil
}

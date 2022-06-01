package core

// #cgo CFLAGS: -I/home/hosein/Desktop/PRJ/wallet-core/include
// #cgo LDFLAGS: -L/home/hosein/Desktop/PRJ/wallet-core/build -L/home/hosein/Desktop/PRJ/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWHash.h>
import "C"
import "chainer/internal/pkg/trustwallet/types"

func HashKeccak256(idata []byte) []byte {
	inputHash := types.TWDataCreateWithGoBytes(idata)
	defer C.TWDataDelete(inputHash)
	oHash := C.TWHashKeccak256(inputHash)
	defer C.TWDataDelete(oHash)
	return types.TWDataGoBytes(oHash)
	return idata
}

package core

// #cgo CFLAGS: -I/home/hosein/Desktop/PRJ/wallet-core/include
// #cgo LDFLAGS: -L/home/hosein/Desktop/PRJ/wallet-core/build -L/home/hosein/Desktop/PRJ/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWBitcoinScript.h>
// #include <TrustWalletCore/TWBitcoinSigHashType.h>
import "C"

import "chainer/internal/pkg/trustwallet/types"

const (
	BitcoinSigHashTypeAll          = C.TWBitcoinSigHashTypeAll
	BitcoinSigHashTypeNone         = C.TWBitcoinSigHashTypeNone
	BitcoinSigHashTypeSingle       = C.TWBitcoinSigHashTypeSingle
	BitcoinSigHashTypeFork         = C.TWBitcoinSigHashTypeFork
	BitcoinSigHashTypeForkBTG      = C.TWBitcoinSigHashTypeForkBTG
	BitcoinSigHashTypeAnyoneCanPay = C.TWBitcoinSigHashTypeAnyoneCanPay
)

func BitcoinLockScriptForAddress(addr string, ct CoinType) []byte {
	address := types.TWStringCreateWithGoString(addr)
	defer C.TWStringDelete(address)

	script := C.TWBitcoinScriptLockScriptForAddress(address, C.enum_TWCoinType(ct))
	scriptData := C.TWBitcoinScriptData(script)
	defer C.TWBitcoinScriptDelete(script)
	defer C.TWDataDelete(scriptData)

	return types.TWDataGoBytes(scriptData)
}

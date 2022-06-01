package core

// #cgo CFLAGS: -I/home/hosein/Desktop/PRJ/wallet-core/include
// #cgo LDFLAGS: -L/home/hosein/Desktop/PRJ/wallet-core/build -L/home/hosein/Desktop/PRJ/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWCoinTypeConfiguration.h>
import "C"

import "chainer/internal/pkg/trustwallet/types"

type CoinType uint32

const (
	CoinTypeBitcoin  CoinType = C.TWCoinTypeBitcoin
	CoinTypeEthereum CoinType = C.TWCoinTypeEthereum
	CoinTypeTron     CoinType = C.TWCoinTypeTron
)

func (c CoinType) GetName() string {
	name := C.TWCoinTypeConfigurationGetName(C.enum_TWCoinType(c))
	defer C.TWStringDelete(name)
	return types.TWStringGoString(name)
}

func (c CoinType) Decimals() int {
	return int(C.TWCoinTypeConfigurationGetDecimals(C.enum_TWCoinType(c)))
}

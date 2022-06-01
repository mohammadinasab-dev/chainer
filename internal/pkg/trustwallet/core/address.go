package core

// #cgo CFLAGS: -I/home/mohammad/Desktop/wallet-core/include
// #cgo LDFLAGS: -L/home/mohammad/Desktop/wallet-core/build -L/home/mohammad/Desktop/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWAnyAddress.h>
// #include <TrustWalletCore/TWBase58.h>
import "C"

import (
	"chainer/internal/pkg/trustwallet/types"
	"errors"
	"fmt"
	"reflect"
)

func DecodeAddressToString(addressBytes []byte) string {
	return "ddd"
}
func DecodeAddressFromString(addressString string) []byte {
	addressTwString := types.TWStringCreateWithGoString(addressString)
	defer C.TWStringDelete(addressTwString)

	addressBytes := C.TWBase58Decode(addressTwString)
	defer C.TWDataDelete(addressBytes)
	byteArray := types.TWDataGoBytes(addressBytes)
	return byteArray

}
func CreateAddressWithMnemonic(mn string, ct CoinType, id uint32) (string, error) {
	if !IsMnemonicValid(mn) {
		return "", errors.New("mnemonic is not valid")
	}

	str := types.TWStringCreateWithGoString(mn)
	empty := types.TWStringCreateWithGoString("")

	defer C.TWStringDelete(str)
	defer C.TWStringDelete(empty)

	tw := C.TWHDWalletCreateWithMnemonic(str, empty)
	defer C.TWHDWalletDelete(tw)

	priKey := C.TWHDWalletGetDerivedKey(tw, C.enum_TWCoinType(ct), C.uint(0), C.uint(0), C.uint(id))
	defer C.TWPrivateKeyDelete(priKey)
	priKeyData := C.TWPrivateKeyData(priKey)
	defer C.TWDataDelete(priKeyData)

	fmt.Println("TYPE :: ", reflect.TypeOf(priKey))
	var pubKey *C.struct_TWPublicKey
	if ct == CoinTypeTron {
		pubKey = C.TWPrivateKeyGetPublicKeySecp256k1(priKey, false)
	}
	defer C.TWPublicKeyDelete(pubKey)

	addressDesc := C.TWAnyAddressCreateWithPublicKey(pubKey, C.enum_TWCoinType(ct))
	defer C.TWAnyAddressDelete(addressDesc)
	addressStr := C.TWAnyAddressDescription(addressDesc)
	defer C.TWStringDelete(addressStr)

	return types.TWStringGoString(addressStr), nil
}

func CreateAddressBytesWithMnemonic(mn string, ct CoinType, id uint32) ([]byte, error) {
	if !IsMnemonicValid(mn) {
		return []byte{0x00}, errors.New("mnemonic is not valid")
	}

	str := types.TWStringCreateWithGoString(mn)
	empty := types.TWStringCreateWithGoString("")

	defer C.TWStringDelete(str)
	defer C.TWStringDelete(empty)

	tw := C.TWHDWalletCreateWithMnemonic(str, empty)
	defer C.TWHDWalletDelete(tw)

	priKey := C.TWHDWalletGetDerivedKey(tw, C.enum_TWCoinType(ct), C.uint(0), C.uint(0), C.uint(id))
	defer C.TWPrivateKeyDelete(priKey)
	priKeyData := C.TWPrivateKeyData(priKey)
	defer C.TWDataDelete(priKeyData)

	fmt.Println("TYPE :: ", reflect.TypeOf(priKey))
	var pubKey *C.struct_TWPublicKey
	if ct == CoinTypeTron {
		pubKey = C.TWPrivateKeyGetPublicKeySecp256k1(priKey, false)
	}
	defer C.TWPublicKeyDelete(pubKey)

	addressDesc := C.TWAnyAddressCreateWithPublicKey(pubKey, C.enum_TWCoinType(ct))
	defer C.TWAnyAddressDelete(addressDesc)

	//for test
	addressStr := C.TWAnyAddressDescription(addressDesc)
	defer C.TWStringDelete(addressStr)
	// fmt.Println("address Str :: ", types.TWStringGoString(addressStr))
	///
	if ct == CoinTypeTron {
		addressBytes := C.TWBase58Decode(addressStr)
		defer C.TWDataDelete(addressBytes)
		return types.TWDataGoBytes(addressBytes), nil
	}

	addressBytes := C.TWAnyAddressData(addressDesc)
	fmt.Println(reflect.TypeOf(addressBytes))
	defer C.TWDataDelete(addressBytes)
	return types.TWDataGoBytes(addressBytes), nil

	// return types.TWStringGoString(addressStr), nil
}

func Base58Address(data []byte) string {
	twData := types.TWDataCreateWithGoBytes(data)
	defer C.TWDataDelete(twData)
	twString := C.TWBase58Encode(twData)
	defer C.TWStringDelete(twString)
	return types.TWStringGoString(twString)

}

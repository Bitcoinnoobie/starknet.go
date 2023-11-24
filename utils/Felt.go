package utils

import (
	"math/big"

	"github.com/NethermindEth/juno/core/felt"
)

func Uint64ToFelt(num uint64) *felt.Felt {
	return new(felt.Felt).SetUint64(num)
}

func HexToFelt(hex string) (*felt.Felt, error) {
	return new(felt.Felt).SetString(hex)
}

func HexArrToFelt(hexArr []string) ([]*felt.Felt, error) {

 feltArr := make([]*felt.Felt, len(hexArr))
 for i, e := range hexArr {
  felt, err := HexToFelt(e)
  if err != nil {
   return nil, err
  }
  feltArr[i] = felt
 }
 return feltArr, nil

}



func FeltToBigInt(f *felt.Felt) *big.Int {
	tmp := f.Bytes()
	return new(big.Int).SetBytes(tmp[:])
}

func BigIntToFelt(big *big.Int) *felt.Felt {
	return new(felt.Felt).SetBytes(big.Bytes())
}

func FeltArrToBigIntArr(f []*felt.Felt) []*big.Int {
	var bigArr []*big.Int
	for _, felt := range f {
		bigArr = append(bigArr, FeltToBigInt(felt))
	}
	return bigArr
}

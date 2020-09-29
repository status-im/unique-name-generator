package main

import (
	"crypto/elliptic"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	"hash/fnv"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	sKey := "04261c55675e55ff25edb50b345cfb3a3f35f60712d251cbaaab97bd50054c6ebc3cd4e22200c68daf7493e1f8da6a190a68a671e2d3977809612424c7c3888bc6"
	bKey, _ := hex.DecodeString(sKey)
	spew.Dump(bKey)

	x, y := elliptic.Unmarshal(secp256k1.S256(), bKey)
	spew.Dump(x, y)

	h := fnv.New32()
	h.Write(bKey)
	hbk := h.Sum([]byte{})

	h.Reset()
	h.Write(y.Bytes())
	hb := h.Sum([]byte{})

	h.Reset()
	h.Write([]byte(sKey))
	hsb := h.Sum([]byte{})

	h.Reset()
	h.Write([]byte("soup"))
	hrb := h.Sum([]byte{})

	h.Reset()
	h.Write([]byte("soup"))
	hrb2 := h.Sum([]byte{})

	spew.Dump(hbk, hb, hsb, hrb, hrb2)

	spew.Dump(base58.Encode(hbk))
	spew.Dump(base58.Encode(hb))
	spew.Dump(base58.Encode(hsb))

	spew.Dump(randWord(x) +" "+ randWord(y))

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	for i := 0; i < 100; i++ {
		_, x, y, _ := elliptic.GenerateKey(secp256k1.S256(), r)
		kb := elliptic.Marshal(secp256k1.S256(), x, y)
		h.Reset()
		h.Write(kb)
		hkb := h.Sum([]byte{})

		spew.Dump("----------------")
		spew.Dump(x.String(), y.String(), base58.Encode(kb), base58.Encode(kb[1:]), randWord(x) +" "+ randWord(y) +":"+ base58.Encode(hkb))
	}
}

func randWord(y *big.Int) string {
	constants := []string{
		"b","c","d","f","g","h",/*"j",*/
		"k","l","m","n","p",/*"q",*/"r",
		"s","t","v","w",/*"x",*/"y","z",
	}

	constantsm := []string{
		"bl","br","bw",
		"ch","cl","cr","cz",
		"dh","dr","dw",
		"fl","fr",
		"gh","gl","gr","gw",
		"kh","kl","kn","kr","kw",
		"ph","pl","pp","pr",
		"rh",
		"st","sl","sm","sn","sp","sc","sh","sk","sw","ss","str",
		"tr","tw",
	}

	vowels := []string{
		"a","e","i","o","u",
	}

	vowelsm := []string{
		"ae","ai","ao","au","ay",
		"ea","ee","eo","eu","ey",
		"io","iu",
		"oa","oe","oi","ou","oy",
		"ui","uy",
	}

	patterns := []string{
		"cvcv","vcvc","cvcvc","vcvcv",
	}

	s := rand.NewSource(y.Int64())
	r := rand.New(s)

	var word string

	p := patterns[r.Intn(len(patterns))]
	for i, c := range p {
		switch string(c) {
		case "c":
			if r.Intn(2) != 1 || i == len(p) -1 {
				word += constants[r.Intn(len(constants))]
			} else {
				word += constantsm[r.Intn(len(constantsm))]
			}
			break
		case "v":
			if r.Intn(3) != 2 || i == 0 {
				word += vowels[r.Intn(len(vowels))]
			} else {
				word += vowelsm[r.Intn(len(vowelsm))]
			}
			break
		}
	}

	return word
}

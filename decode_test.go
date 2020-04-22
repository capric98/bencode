package bencode

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// func TestDecodeInt(t *testing.T) {
// 	for i := 0; i < 10000; i++ {
// 		rn := rand.Int63n(0x7FFFFFFF) - 0x3FFFFFFF
// 		rnb := []byte("i" + strconv.FormatInt(rn, 10) + "e")
// 		de := &Decoder{
// 			pos:    0,
// 			bufLen: len(rnb),
// 			buf:    rnb,
// 		}
// 		result := de.decodeInt()
// 		if result != rn {
// 			fmt.Println("i"+strconv.FormatInt(rn, 10)+"e", "->", result)
// 			t.Fail()
// 		}
// 		if de.bufLen != de.pos {
// 			fmt.Println("position not set!")
// 			t.Fail()
// 		}
// 	}

// }

// func TestDecodeStr(t *testing.T) {
// 	for i := 0; i < 10000; i++ {
// 		sb := strings.Builder{}
// 		slen := rand.Intn(500)
// 		for k := 0; k < slen; k++ {
// 			sb.WriteByte(byte(rand.Intn(256)))
// 		}
// 		bstr := []byte(strconv.Itoa(slen) + ":" + sb.String())
// 		de := &Decoder{
// 			pos:    0,
// 			bufLen: len(bstr),
// 			buf:    bstr,
// 		}
// 		var rstr string
// 		de.decodeStr(&rstr)
// 		if rstr != sb.String() {
// 			fmt.Println(string(bstr), "->", rstr)
// 			t.Fail()
// 		}
// 		if de.bufLen != de.pos {
// 			fmt.Println("position not set!")
// 			t.Fail()
// 		}
// 	}
// }

// func TestDecodeDict(t *testing.T) {
// 	str := "d3:bar4:spam3:fooi42ee"
// 	b, e := Decode([]byte(str))
// 	if e != nil {
// 		fmt.Println(e)
// 		t.FailNow()
// 	}
// 	b.Print()
// 	t.Fail()
// }

// func TestDecode(t *testing.T) {
// 	tlmc, e := ioutil.ReadFile("test/tlmc.torrent")
// 	if e != nil {
// 		fmt.Println(e)
// 		t.FailNow()
// 	}
// 	b, e := Decode(tlmc)
// 	fmt.Println(len(*b.b))
// 	//b.Print()
// 	t.Fail()
// }

func BenchmarkTLMCDecode(b *testing.B) {
	tlmc, e := ioutil.ReadFile("test/TLMC.torrent")
	if e != nil {
		fmt.Println(e)
		b.FailNow()
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, e = Decode(tlmc)
		if e != nil {
			fmt.Println(e)
			b.FailNow()
		}
	}
}

// func BenchmarkReader(b *testing.B) {
// 	tlmc, e := ioutil.ReadFile("test/TLMC.torrent")
// 	if e != nil {
// 		fmt.Println(e)
// 		b.FailNow()
// 	}
// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		r := bytes.NewReader(tlmc)
// 		_, _ = ioutil.ReadAll(r)
// 	}
// }

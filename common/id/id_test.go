package id

import(
	"testing"
)

func TestId(t *testing.T){
	time := BeginMillisecond - 86400000*15



	id := NewIdCreator(TimeBit, user_DcBit, user_RoleBit, time, 0, 0).Id()

	t.Logf("binary:%064b uid:%+v\n", id, id)

}


func TestBinary(t *testing.T){
	a := 8
	//b := uint8(3)
	//b := -1 ^ (-1 << a)
	t.Logf("binary:%b\n", (-1<<a))
	t.Logf("value:%+v\n", int64((253+4)&255))
}

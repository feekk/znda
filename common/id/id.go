package id

import(
	"time"
	"sync"
)

var (
	//2020-01-01 00:00:00
	BeginMillisecond int64 = 1577808000000
	symbolBit int = 1
	TimeBit int = 40
	//
	//
	//0b0011 1111 = 63
	//user_DcBit int = 6
	//0b0001 1111 1111 = 511
	//user_RoleBit int = 9
	//0b1111 1111 = 255
	//user_SequenceBit int = 8 

	//
	//
	//0b0111 1111 = 127
	//order_DcBit int = 7
	//0b1111 1111 = 255
	//order_RoleBit int = 8
	//0b1111 1111 = 255
	//order_SequenceBit int = 8 
)

func NewIdCreator(beginBit, dcBit, roleBit int, beginVal, dcVal, roleVal int64) (id *IDCreator){
	if (symbolBit + beginBit + dcBit + roleBit) > 64 {
		panic("Creator Bits length > int64 length.")
	}
	if bitMaxValue(beginBit) < beginVal{
		panic("beginVal must be less-than bitMaxValue.")
	}
	if bitMaxValue(dcBit) < dcVal{
		panic("dcVal must be less-than bitMaxValue.")
	}
	if bitMaxValue(roleBit) < roleVal{
		panic("roleVal must be less-than bitMaxValue.")
	}
	id  = &IDCreator{
		TimeBit:beginBit,
		BeginMillisec:beginVal,
		DcBit:dcBit,
		DcVal:dcVal,
		RoleBit:roleBit,
		RoleVal:roleVal,
		SeqBit:64 - (symbolBit + beginBit + dcBit + roleBit),
	}
	return 
}

type Creator interface{
	Id() int64
}

type IDCreator struct{
	m sync.Mutex
	TimeBit int
	BeginMillisec int64
	DcBit int
	DcVal int64
	RoleBit int
	RoleVal int64
	SeqBit int
	SeqVal int64
	lastMillisec int64
}
func(id *IDCreator) Id() int64{
	id.m.Lock()
	defer id.m.Unlock()

	nowMillisec := nowMillisecond()
	//Clock moved backwards
	if (nowMillisec < id.BeginMillisec){
		panic("Clock moved backwards.  Refusing to generate id.")
	}

	if nowMillisec == id.lastMillisec {
		id.SeqVal = (id.SeqVal + 1) & bitMaxValue(id.SeqBit)
		if id.SeqVal == 0 {
			nowMillisec = nextMillisecond(id.lastMillisec)
		}
	}else{
		id.SeqVal = 0
	}
	id.lastMillisec = nowMillisec

	uid := ((id.lastMillisec-id.BeginMillisec) << (id.DcBit+id.RoleBit+id.SeqBit)) | (id.DcVal << (id.RoleBit+id.SeqBit)) | (id.RoleVal << id.SeqBit) | id.SeqVal
	return int64(uid)
}



func bitMaxValue(bit int) int64{
	return int64(-1 ^ (-1 << bit))
}
func nextMillisecond(lastMillisec int64) int64{
	nowMillisec := nowMillisecond()
	
	for nowMillisec <= lastMillisec {
		nowMillisec = nowMillisecond()
	}
	return nowMillisec
}
func nowMillisecond() int64{
	loc, _ := time.LoadLocation("Local")
	now := time.Now().In(loc)
	return now.Unix()*1000 + int64(now.Nanosecond()/int(1000000))
}
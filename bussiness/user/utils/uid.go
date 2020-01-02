package utils

import(
	"znda/common/id"
	"math/rand"
)

var(
	//
	//0b0011 1111 = 63
	userID_DcBit int = 6
	//0b0001 1111 1111 = 511
	userID_RoleBit int = 9
	//0b1111 1111 = 255
	userID_SeqBit int = 8 
	//
	UIDCreators sync.Map

	//RoleId max: 511 (-1^ (-1 << userID_RoleBit))
	roles []int = []int{
		ROLE_PASSENGER,
		ROLE_DRIVER,
	}
)

func InitUidCreator(dcVal int) {
	dcVal = rand.Int31n((-1^(-1 << userID_DcBit))+1)
	for _, val := range roles {
		UIDCreators.Store(val, NewRoleIdCreator(dcVal, val))
	}
}

func NewRoleIdCreator(dcVal, roleId int) *id.Creator{
	return id.NewIdCreator(id.TimeBit, userID_DcBit, userID_RoleBit, userID_SeqBit, id.BeginMillisecond, dcVal, roleId)
}

func GetRoleIdCreator(role int) *id.Creator{
	val, ok := UIDCreators.Load(role)
	if !ok {
		return nil
	}
	if v, ok := val.(*id.Creator); ok{
		return v
	}
	return nil
}
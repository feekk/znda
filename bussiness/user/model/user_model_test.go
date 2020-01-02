package model

import(
	"testing"
)

func TestFindByUid(t *testing.T){
	user, err := new(UserModel).FindByUid(2)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	t.Logf("%+v\n", user)
}
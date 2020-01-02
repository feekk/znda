package model

import(
	"github.com/feekk/zddgo"
)

func init(){
	conf, err := zddgo.NewConfig("../conf/dev.toml")
	if err != nil{
		panic(err)
	}
	if err = zddgo.MysqlInit(&conf.Database); err != nil {
		panic(err)
	}
}
package main

import(
	"os"
	"fmt"
	"flag"
	"runtime"
	"github.com/feekk/zddgo"
	"znda/bussiness/user/modules"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage specified:")
	flag.PrintDefaults()
	os.Exit(1)
}


func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	z := zddgo.New()

	if err := z.InitConfig(); err != nil{
		usage()
	}
	if err := z.InitOrm(); err != nil{
		panic(err)
	}
	if err := z.InitCache(); err != nil{
		panic(err)
	}
	if err := z.HttpStart(modules.Init()); err != nil{
		panic(err)
	}
}
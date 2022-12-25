package main

import (
	"fmt"
	pd "proto-go/proto"
)

func complexProto() *pd.Simple1 {
	return &pd.Simple1{
		Simple1: &pd.Simple{
			Id:       1,
			IsSimple: true,
			Name:     "arun",
			Lists:    []int32{1, 2, 3, 4},
		},
		Dummy1: []*pd.Simple{{
			Id:       2,
			IsSimple: true,
			Name:     "harsha",
			Lists:    []int32{1, 2, 3, 4},
		}, {
			Id:       1,
			IsSimple: true,
			Name:     "abhi",
			Lists:    []int32{1, 2, 3, 4},
		},
		},
	}
}

func printmap() *pd.Mapmessag {
	return &pd.Mapmessag{Keypair: map[string]int32{"age": 32, "id ": 2}}
}

func enumPrint() *pd.Enumration {
	return &pd.Enumration{
		Eyecolor: 2,
	}
}

func main() {
	//fmt.Println(complexProto())
	//fmt.Println(enumPrint())
	fmt.Println(printmap())

}

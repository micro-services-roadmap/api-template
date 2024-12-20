package gmw

import (
	"fmt"
	"testing"
)

func Test_extractFirstPart(t *testing.T) {
	ip := extractFirstPart("117.186.5.188:14594")
	fmt.Println(ip)
	ip2 := extractFirstPart("117.186.5.188:14594,xuihsa:22")
	fmt.Println(ip2)
	ip3 := extractFirstPart("240e:46c:8910:219e:41a2:1185:be37:5f61")
	fmt.Println(ip3)
	ip4 := extractFirstPart("240e:46c:8910:219e:41a2:1185:be37:5f61:9080")
	fmt.Println(ip4)
}

package main

import "fmt"

// 让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址。
// 例如，`IPAddr{1,`2,`3,`4}` 应当输出 `"1.2.3.4"`。
type IPAddr [4]byte

func (addr IPAddr) String() string {

	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

func main() {

	//addrs := IPAddr{127, 0, 0, 1}
	//fmt.Printf("%v", addrs)

	ipAddrs := map[string]IPAddr{
		"ip1": {0, 0, 0, 0},
		"ip2": {1, 1, 1, 1},
	}
	fmt.Println(ipAddrs)

}

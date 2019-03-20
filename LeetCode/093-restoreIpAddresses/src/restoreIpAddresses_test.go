package src

import (
	"fmt"
	"testing"
)

func TestOK(t *testing.T) {
	ips := restoreIpAddresses("25525511135")
	fmt.Println(ips)
}

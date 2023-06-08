package random

import (
	"fmt"
	"testing"
)

func TestSeed_Example(t *testing.T) {
	fmt.Println(UpgradeSeed(84978056))
	fmt.Println(UpgradeSeed(-8735875436988056))

	// output:
	//	[4120648339100951164 3023117482333613152]
	//	[7499608215528407055 610144910520614932]
}

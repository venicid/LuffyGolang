package settings

import "fmt"

const VERSION = "0.1.0"

func GetVersion() string {
	fmt.Printf("build veriosn: \t%v\n", VERSION)
	return VERSION

}
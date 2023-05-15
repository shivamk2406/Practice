package resourcename

import (
	"fmt"
)


const(
	blankOrderNameFormat        = "carts/{cart}/orders/-"
	redisProductKey           = "arise-b2b-saleor-pid-%s-%s"
)
func Start() error {
	//fmt.Println(resourcename.Sprint(blankOrderNameFormat,"1"))
	key:=fmt.Sprintf(redisProductKey,"productID","ChannelName")
	fmt.Println(key)

	var pid string
	var channelName string

	n,err:= fmt.Sscanf(key,redisProductKey,&pid,&channelName)
	if err!=nil{
		fmt.Println("inside error")
		fmt.Println(err)
		fmt.Println(n)
	}
	fmt.Println(pid,channelName)
	return nil
}
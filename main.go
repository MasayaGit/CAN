package main
//import "fmt"
import "math/rand"
import "github.com/MasayaGit/CAN/can"

func main() {
	//xs int, ys int, xe int, ye int
	bootstrap := can.InitBootstrap(0,0,10,10)
	rand.Seed(7)
	
	
	for i := 0; i < 10; i++ {
		//0~10の間の擬似乱数
		can.InitNode(rand.Intn(11),rand.Intn(11),bootstrap)
	}
	
	
	/*
	can.InitNode(7,7,bootstrap)
	can.InitNode(9,9,bootstrap)
	can.InitNode(3,9,bootstrap)
	can.InitNode(3,3,bootstrap)
	can.InitNode(9,9,bootstrap)
	can.InitNode(8,6,bootstrap)
	can.InitNode(1,3,bootstrap)
	*/
	
	chargeNode := bootstrap.FindNode(10,10)

	//xMiddle yMiddle xs ys xe ye
	//bootstrap.ShowInfo()
	chargeNode.ShowInfo()
	//result 
	//&{8 8 7 7 10 10}
	//&{8 6 7 5 10 7}
	//&{6 8 5 7 7 10}
}

package main
import "fmt"
import "math/rand"

func main() {

	/*
	Node bootstrap = new Node(乱数, 乱数, 0, 0, x_max, y_max)
	for i in range(10):
 	Node n = new Node(乱数, 乱数, bootstrap)
	//Debug
	Node n =bootstrap.get(10, 10)
	//ここで10, 10を担当するノードの情報が表示されればOK
	n.showInfo() 
	*/	
	
	//xs int, ys int, xe int, ye int
	bootstrap := can.initBootstrap(0,0,10,10)
	rand.Seed(7)
	for i := 0; i < 10; i++ {
		//0~9の間の擬似乱数
		can.initNode(rand.Intn(10),rand.Intn(10),bootstrap)
	}
	findNode := bootstrap.findNode(10,10)
	fmt.Println(findNode.ownRange.ye)
}

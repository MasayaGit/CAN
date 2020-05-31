
package can


import "testing"
import "fmt"
//go test -v canで実行
//テストはxxx_test.goというファイルに、TestYyyという関数名をつける。

func TestInitBootstrap(t *testing.T) {
	expectedXs := 0
	expectedYs := 0
	expectedXe := 10
	expectedYe := 10
	//xs int, ys int, xe int, ye int
	bootstrap := initBootstrap(0,0,10,10)
	if bootstrap.ownRange.xs != expectedXs || bootstrap.ownRange.xe != expectedXe {
		t.Errorf("error in x")
	}	
	if bootstrap.ownRange.ys!= expectedYs || bootstrap.ownRange.ye != expectedYe {
		t.Errorf("error in y")
	}
}


func TestGetDividedRange(t *testing.T) {
	expectedXs := 5
	expectedYs := 0
	expectedXe := 10
	expectedYe := 10
	inputX := 5
	inputY := 5
	//xs int, ys int, xe int, ye int
	bootstrap := initBootstrap(0,0,10,10)
	newRangeObj := bootstrap.getDividedRange(inputX,inputY)
	if newRangeObj.xs != expectedXs || newRangeObj.xe != expectedXe {
		t.Errorf("error in x")
	}	
	if newRangeObj.ys!= expectedYs || newRangeObj.ye != expectedYe {
		t.Errorf("error in y")
	}
}


func TestFindNode(t *testing.T) {
	expectedXs := 0
	expectedYs := 0
	expectedXe := 5
	expectedYe := 10
	inputFirstX := 5
	inputFirstY := 5

	//xs int, ys int, xe int, ye int
	bootstrap := initBootstrap(0,0,10,10)

	//1回目 x軸方向でbootstrapを分割
	firstNodeObj := new(Node)
	firstNodeObj.ownRange = bootstrap.getDividedRange(inputFirstX,inputFirstY)

	bootstrap.neighbors = append(bootstrap.neighbors,firstNodeObj)
	firstNodeObj.neighbors = append(firstNodeObj.neighbors,bootstrap)

	firstFindNodeObj := bootstrap.findNode(inputFirstX,inputFirstY)
	fmt.Println(firstFindNodeObj.ownRange.ye)

	if firstFindNodeObj.ownRange.xs != expectedXs || firstFindNodeObj.ownRange.xe != expectedXe {
		t.Errorf("error in x first")
	}	
	if firstFindNodeObj.ownRange.ys!= expectedYs || firstFindNodeObj.ownRange.ye != expectedYe {
		t.Errorf("error in y first")
	}

	//2回目 y軸方向で1回目で作成したノードを分割
	expectedXs = 5
	expectedYs = 5
	expectedXe = 10
	expectedYe = 10
	inputSecondX := 7
	inputSecondY := 7

	secondNodeObj := new(Node)
	divideNodeObj := bootstrap.findNode(inputSecondX,inputSecondY)
	secondNodeObj.ownRange = divideNodeObj.getDividedRange(inputSecondX,inputSecondY)

	//手動で登録
	bootstrap.neighbors = append(bootstrap.neighbors,secondNodeObj)
	secondNodeObj.neighbors = append(secondNodeObj.neighbors,bootstrap)
	secondNodeObj.neighbors = append(secondNodeObj.neighbors,firstNodeObj)
	firstNodeObj.neighbors = append(firstNodeObj.neighbors,secondNodeObj)

	secondFindNodeObj := bootstrap.findNode(inputSecondX,inputSecondX)

	fmt.Println(bootstrap.neighbors)
	fmt.Println(bootstrap.neighbors[1].ownRange.ye)

	if secondFindNodeObj.ownRange.xs != expectedXs || secondFindNodeObj.ownRange.xe != expectedXe {
		t.Errorf("error in x second")
	}	
	if secondFindNodeObj.ownRange.ys!= expectedYs || secondFindNodeObj.ownRange.ye != expectedYe {
		t.Errorf("error in y second")
	}

}

func TestRemoveDecision(t *testing.T) {
	expectedFlag := false

	//xs int, ys int, xe int, ye int
	bootstrap := initBootstrap(5,5,10,10)

	newNodeObj := new(Node)

	newRangeObj := new(Range)
	newRangeObj.xs = 0
	newRangeObj.ys = 5
	newRangeObj.xe = 5
	newRangeObj.ye = 10
	newRangeObj.xMiddle = int((newRangeObj.xe + newRangeObj.xs)/2)
	newRangeObj.yMiddle = int((newRangeObj.ye + newRangeObj.ys)/2)

	newNodeObj.ownRange = newRangeObj

	flag := bootstrap.removeDecision(newNodeObj)
	

	if flag != expectedFlag{
		t.Errorf("error in flag")
	}
}


func TestRemoveNodeFromNeighbors(t *testing.T) {
	expected := 20

	//xs int, ys int, xe int, ye int
	bootstrap := initBootstrap(5,5,10,10)

	firstNodeObj := new(Node)
	
	firstRangeObj := new(Range)
	firstRangeObj.xs = 0
	firstRangeObj.ys = 5
	firstRangeObj.xe = 5
	firstRangeObj.ye = 10
	firstRangeObj.xMiddle = int((firstRangeObj.xe + firstRangeObj.xs)/2)
	firstRangeObj.yMiddle = int((firstRangeObj.ye + firstRangeObj.ys)/2)

	firstNodeObj.ownRange = firstRangeObj

	bootstrap.neighbors = append(bootstrap.neighbors,firstNodeObj)

	secondNodeObj := new(Node)

	secondRangeObj := new(Range)
	secondRangeObj.xs = 15
	secondRangeObj.ys = 5
	secondRangeObj.xe = 20
	secondRangeObj.ye = 10
	secondRangeObj.xMiddle = int((secondRangeObj.xe + secondRangeObj.xs)/2)
	secondRangeObj.yMiddle = int((secondRangeObj.ye + secondRangeObj.ys)/2)

	secondNodeObj.ownRange = secondRangeObj

	bootstrap.neighbors = append(bootstrap.neighbors,secondNodeObj)

	fmt.Println(bootstrap.neighbors)

	bootstrap.removeNodeFromNeighbors(firstNodeObj)
	
	fmt.Println(bootstrap.neighbors)

	if bootstrap.neighbors[0].ownRange.xe != expected{
		t.Errorf("error in remove")
	}
}



func TestSetNeighborsToNewNode(t *testing.T) {
	expected := 15

	//xs int, ys int, xe int, ye int
	bootstrap := initBootstrap(5,5,15,10)

	firstNodeObj := new(Node)
	
	firstRangeObj := new(Range)
	firstRangeObj.xs = 0
	firstRangeObj.ys = 5
	firstRangeObj.xe = 5
	firstRangeObj.ye = 10
	firstRangeObj.xMiddle = int((firstRangeObj.xe + firstRangeObj.xs)/2)
	firstRangeObj.yMiddle = int((firstRangeObj.ye + firstRangeObj.ys)/2)

	firstNodeObj.ownRange = firstRangeObj

	bootstrap.neighbors = append(bootstrap.neighbors,firstNodeObj)

	secondNodeObj := new(Node)

	secondRangeObj := new(Range)
	secondRangeObj.xs = 5
	secondRangeObj.ys = 10
	secondRangeObj.xe = 10
	secondRangeObj.ye = 15
	secondRangeObj.xMiddle = int((secondRangeObj.xe + secondRangeObj.xs)/2)
	secondRangeObj.yMiddle = int((secondRangeObj.ye + secondRangeObj.ys)/2)

	secondNodeObj.ownRange = secondRangeObj

	bootstrap.neighbors = append(bootstrap.neighbors,secondNodeObj)

	threeNodeObj := new(Node)

	threeRangeObj := new(Range)
	threeRangeObj.xs = 10
	threeRangeObj.ys = 10
	threeRangeObj.xe = 15
	threeRangeObj.ye = 15
	threeRangeObj.xMiddle = int((threeRangeObj.xe + threeRangeObj.xs)/2)
	threeRangeObj.yMiddle = int((threeRangeObj.ye + threeRangeObj.ys)/2)

	threeNodeObj.ownRange = threeRangeObj
	threeNodeObj.setNeighborsToNewNode(bootstrap)

	fmt.Println(bootstrap.neighbors)
	fmt.Println(threeNodeObj.neighbors)

	if threeNodeObj.neighbors[0].ownRange.ye != expected{
		t.Errorf("error in remove")
	}
}



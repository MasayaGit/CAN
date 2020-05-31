package can

import "testing"
//go test -v canで実行
//テストはxxx_test.goというファイルに、TestYyyという関数名をつける。

func TestInitRange(t *testing.T) {
	expectedXs := 0
	expectedYs := 0
	expectedXe := 10
	expectedYe := 10
	expectedXMiddle := 5
	expectedYMiddle := 5
	//xs int, ys int, xe int, ye int
	rangeObj := initRange(0,0,10,10)
	if rangeObj.xs != expectedXs || rangeObj.xe != expectedXe {
		t.Errorf("error in x")
		}	
	if rangeObj.ys != expectedYs || rangeObj.ye != expectedYe {
		t.Errorf("error in y")
		}
	if rangeObj.xMiddle != expectedXMiddle || rangeObj.yMiddle != expectedYMiddle {
		t.Errorf("error in middle")
		}
		
	
}

func TestContain(t *testing.T) {
	expect := true
	//xs int, ys int, xe int, ye int
	rangeObj := initRange(0,0,3,3)
	flag := rangeObj.contain(3,3)
	if flag != expect{
		t.Errorf("error in contain")
	}	
}


func TestDivide(t *testing.T) {
	expectedNewRangeXs := 0
	expectedNewRangeYs := 5
	expectedNewRangeXe := 5
	expectedNewRangeYe := 10
	expectedNewRangeXMiddle := 2
	expectedNewRangeYMiddle := 7

	expectedRangeXs := 0
	expectedRangeYs := 0
	expectedRangeXe := 5
	expectedRangeYe := 5
	expectedRangeXMiddle := 2
	expectedRangeYMiddle := 2
	//xs int, ys int, xe int, ye int
	rangeObj := initRange(0,0,5,10)
	newRangeObj := rangeObj.divide(2,7)

	if newRangeObj.xs != expectedNewRangeXs || newRangeObj.xe != expectedNewRangeXe {
		t.Errorf("error in newRange x")
	}	
	if newRangeObj.ys != expectedNewRangeYs || newRangeObj.ye != expectedNewRangeYe {
		t.Errorf("error in newRange y")
	}
	if newRangeObj.xMiddle != expectedNewRangeXMiddle || newRangeObj.yMiddle != expectedNewRangeYMiddle {
		t.Errorf("error in newRange middle")
	}

	if rangeObj.xs != expectedRangeXs || rangeObj.xe != expectedRangeXe {
		t.Errorf("error in range x")
	}	
	if rangeObj.ys != expectedRangeYs || rangeObj.ye != expectedRangeYe {
		t.Errorf("error in range y")
	}
	if rangeObj.xMiddle != expectedRangeXMiddle || rangeObj.yMiddle != expectedRangeYMiddle {
		t.Errorf("error in range middle")
	}

}


package can

//xs, ys を始点にしてxe, yeの点までの長方形の領域を管理する
//xMiddle, yMiddleは中点
type Range struct {
	xMiddle int
	yMiddle int
	xs int
	ys int
	xe int 
	ye int}


// コンストラクタ
// 戻り値は構造体のポインターなので*Rangetとする
func initRange(xs int, ys int, xe int, ye int) *Range {
	 // コンストラクタの関数内で、構造体をnew
	 //new() は、構造体へのポインタを返す。
     rangeObj := new(Range)
     rangeObj.xs = xs
	 rangeObj.ys = ys
	 rangeObj.xe = xe
	 rangeObj.ye = ye
	 rangeObj.xMiddle = int((rangeObj.xe + rangeObj.xs)/2)
	 rangeObj.yMiddle = int((rangeObj.ye + rangeObj.ys)/2)
     return rangeObj
}

// x, yが領域に含まれるか否かを返すメソッド
// trueならこのrangeObjの領域にx, yが含まれる
func (rangeObj Range) contain(x int,y int) bool {
	containFlag := false
	if x <= rangeObj.xe  && x >= rangeObj.xs{
		if y <= rangeObj.ye  && y >= rangeObj.ys{
			containFlag = true
		}
	}
    return containFlag
}

// x, yの引数に対応する長方形を分割し、そのrangeObjを返す
// 辺の長さが長い方を分割する
func (rangeObj *Range) divide(x int,y int) *Range {
	xRange := rangeObj.xe - rangeObj.xs
	yRange := rangeObj.ye - rangeObj.ys

	// xの長さの方が長い時の処理
	if xRange >= yRange{
		//中点のx座標と比較して、どっちに属するかによって返す値変わる
		if x >= rangeObj.xMiddle{
			//新しいノードに渡す分割
			newRangeObj := new(Range)
			newRangeObj.xs = rangeObj.xMiddle
			newRangeObj.ys = rangeObj.ys
			newRangeObj.xe = rangeObj.xe
			newRangeObj.ye = rangeObj.ye
			newRangeObj.xMiddle = int((newRangeObj.xe + newRangeObj.xs)/2)
			newRangeObj.yMiddle = int((newRangeObj.ye + newRangeObj.ys)/2)

			//分割されるrangeObjの範囲を変更
			rangeObj.xe = rangeObj.xMiddle
			rangeObj.xMiddle = int((rangeObj.xe + rangeObj.xs)/2)
     		return newRangeObj
		}else{
			//新しいノードに渡す分割
			newRangeObj := new(Range)
			newRangeObj.xs = rangeObj.xs
			newRangeObj.ys = rangeObj.ys
			newRangeObj.xe = rangeObj.xMiddle
			newRangeObj.ye = rangeObj.ye
			newRangeObj.xMiddle = int((newRangeObj.xe + newRangeObj.xs)/2)
			newRangeObj.yMiddle = int((newRangeObj.ye + newRangeObj.ys)/2)

			//分割されるrangeObjの範囲を変更
			rangeObj.xs = rangeObj.xMiddle
			rangeObj.xMiddle = int((rangeObj.xe + rangeObj.xs)/2)
     		return newRangeObj
		}
	}else{
		//中点のy座標と比較して、どっちに属するかによって返す値変わる
		if y >= rangeObj.yMiddle{
			//新しいノードに渡す分割
			newRangeObj := new(Range)
			newRangeObj.xs = rangeObj.xs
			newRangeObj.ys = rangeObj.yMiddle
			newRangeObj.xe = rangeObj.xe
			newRangeObj.ye = rangeObj.ye
			newRangeObj.xMiddle = int((newRangeObj.xe + newRangeObj.xs)/2)
			newRangeObj.yMiddle = int((newRangeObj.ye + newRangeObj.ys)/2)

			//分割されるrangeObjの範囲を変更
			rangeObj.ye = rangeObj.yMiddle
			rangeObj.yMiddle = int((rangeObj.xe + rangeObj.xs)/2)
     		return newRangeObj
		}else{
			//新しいノードに渡す分割
			newRangeObj := new(Range)
			newRangeObj.xs = rangeObj.xs
			newRangeObj.ys = rangeObj.ys
			newRangeObj.xe = rangeObj.xe
			newRangeObj.ye = rangeObj.yMiddle
			newRangeObj.xMiddle = int((newRangeObj.xe + newRangeObj.xs)/2)
			newRangeObj.yMiddle = int((newRangeObj.ye + newRangeObj.ys)/2)
			
			//分割されるrangeObjの範囲を変更
			rangeObj.ys = rangeObj.yMiddle
			rangeObj.yMiddle = int((rangeObj.xe + rangeObj.xs)/2)
     		return newRangeObj
		}
	}
}
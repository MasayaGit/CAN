package can

import "math"
//import "fmt"
type Node struct {
    //*Rangeと宣言したので、ownRangeでポインターを取得
    ownRange *Range
    //スライス
    neighbors []*Node
}

// bootstrapのコンストラクタ
func initBootstrap(xs int, ys int, xe int, ye int) *Node {
     bootstrap := new(Node)
     bootstrap.ownRange = initRange(xs, ys, xe, ye)
     return bootstrap
}

// Nodeのコンストラクタ
func initNode(x int, y int, bootstrap *Node) {
    // 構造体(オブジェクトとみなす)を生成
    newNodeObj := new(Node)
    
    chargeNode := bootstrap.findNode(x,y)   

    newNodeObj.ownRange = chargeNode.getDividedRange(x,y)
    
    //分割したノードから取得した隣接するノードを新規ノードに登録する
    //この際に入れてもいいノードかの判定を行い、追加したのであれば、追加要請を出す
    newNodeObj.setNeighborsToNewNode(chargeNode)

    //新規ノードを分割したノードに入れる
    newNodeObj.setNodeToNeighbors(chargeNode)
    //分割したノードを新規ノードに入れる
    chargeNode.setNodeToNeighbors(newNodeObj)
    
    //分割したノードの隣接ノードを更新する
    chargeNode.searchRemoveNode()        
}



// nodeにneighborsをセットする
// セットしたのであれば、追加要請をだす
func (newNodeObj *Node)setNeighborsToNewNode(haveNeighborsNode *Node) {
    for _,registerNodeObj := range haveNeighborsNode.neighbors {

        if newNodeObj.removeDecision(registerNodeObj){
            continue
        }

        newNodeObj.neighbors = append(newNodeObj.neighbors,registerNodeObj)
        // 追加したノードに対してもnewNodeObjを登録する
        registerNodeObj.setNodeToNeighbors(newNodeObj)
    }
}



// nodeをNeighborsに登録してもらう
func (nodeObj *Node)setNodeToNeighbors(registerNode *Node) {
    nodeObj.neighbors = append(nodeObj.neighbors,registerNode)
}

// nodeをNeighborsから削除してもらう
func (nodeObj *Node)removeNodeFromNeighbors(deleteNode *Node) {
    newNeighbors := []*Node{}
    for _,searchNodeObj := range nodeObj.neighbors{
        if searchNodeObj == deleteNode{
            continue
        }
        newNeighbors = append(newNeighbors,searchNodeObj)
        
    }
    nodeObj.neighbors = newNeighbors
}

// NeighborsからいらないNodeを削除
func (nodeObj *Node)searchRemoveNode() {
    for _,searchNodeObj := range nodeObj.neighbors {
        if nodeObj.removeDecision(searchNodeObj){
            //nodeObjの方で削除
            nodeObj.removeNodeFromNeighbors(searchNodeObj)
            //searchNodeの方でも削除してもらう
            searchNodeObj.removeNodeFromNeighbors(nodeObj)
        }
    }
}


// nodeをNeighborsから削除してもらう
func (nodeObj *Node)removeDecision(searchNodeObj *Node) bool {
    removeFlag := false
    nodeObjXs := nodeObj.ownRange.xs
    nodeObjXe := nodeObj.ownRange.xe
    nodeObjYs := nodeObj.ownRange.ys
    nodeObjYe := nodeObj.ownRange.ye
    searchNodeObjXs := searchNodeObj.ownRange.xs
    searchNodeObjXe := searchNodeObj.ownRange.xe
    searchNodeObjYs := searchNodeObj.ownRange.ys
    searchNodeObjYe := searchNodeObj.ownRange.ye
    //右上
    if searchNodeObjXs >= nodeObjXe && searchNodeObjYs >= nodeObjYe{
        removeFlag =true
    }
    //右下
    if searchNodeObjXs >= nodeObjXe && searchNodeObjYe <= nodeObjYs{
        removeFlag =true
    }
    //左上
    if searchNodeObjXe <= nodeObjXs && searchNodeObjYs >= nodeObjYe{
        removeFlag =true
    }
    //左下
    if searchNodeObjXe <= nodeObjXs && searchNodeObjYe <= nodeObjYs{
        removeFlag =true
    }
    //x軸方向で見たときに隣合わないノード(y軸は同じ)
    if searchNodeObjYs == nodeObjYs && searchNodeObjYe == nodeObjYe{
        if searchNodeObjXe != nodeObjXs && searchNodeObjXs!= nodeObjXe{
            removeFlag =true
        }
    }
    
    //y軸方向で見たときに隣合わないノード(x軸は同じ)
    if searchNodeObjXs == nodeObjXs && searchNodeObjXe == nodeObjXe{
        if searchNodeObjYe != nodeObjYs && searchNodeObjYs!= nodeObjYe{
            removeFlag =true
        }
    }

    
    return removeFlag
}


// bootstrapのコンストラクタ
// 戻り値は構造体のポイント
func (nodeObj *Node)getDividedRange(x int , y int) *Range {
    newRangeObj := nodeObj.ownRange.divide(x, y)
    return newRangeObj
}


// 再帰を使ってx,yを担当しているノードを取得する
func (nodeObj *Node) findNode(x int,y int) *Node {

    if nodeObj.ownRange.contain(x,y){
        return nodeObj
    }

    nodeObjXMiddle := nodeObj.ownRange.xMiddle
    nodeObjYMiddle := nodeObj.ownRange.yMiddle

    //math.absはfloat64しか対応していない
    mostSmallAbsError := math.Abs(float64(nodeObjXMiddle - x)) + math.Abs(float64(nodeObjYMiddle - y))
    mostSmallAbsErrorInt := int(mostSmallAbsError)
    var mostSmallAbsNodeObj *Node
    for _,searchNodeObj := range nodeObj.neighbors {
        searchNodeObjXMiddle := searchNodeObj.ownRange.xMiddle
        searchNodeObjYMiddle := searchNodeObj.ownRange.xMiddle
        absError := math.Abs(float64(searchNodeObjXMiddle - x)) + math.Abs(float64(searchNodeObjYMiddle - y))
        absErrorInt := int(absError)
        if absErrorInt <= mostSmallAbsErrorInt{
            mostSmallAbsNodeObj = searchNodeObj
        }
    }
    return mostSmallAbsNodeObj.findNode(x,y)
}

package can

import "math"
import "fmt"
type Node struct {
    //*Rangeと宣言したので、OwnRangeでポインターを取得
    OwnRange *Range
    //スライス
    neighbors []*Node
}

// bootstrapのコンストラクタ
func InitBootstrap(xs int, ys int, xe int, ye int) *Node {
     bootstrap := new(Node)
     bootstrap.OwnRange = initRange(xs, ys, xe, ye)
     return bootstrap
}

// Nodeのコンストラクタ
func InitNode(x int, y int, bootstrap *Node) {
    // 構造体(オブジェクトとみなす)を生成
    newNodeObj := new(Node)
    
    chargeNode := bootstrap.FindNode(x,y)   

    newNodeObj.OwnRange = chargeNode.getDividedRange(x,y)
    
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
    nodeObjXs := nodeObj.OwnRange.xs
    nodeObjXe := nodeObj.OwnRange.xe
    nodeObjYs := nodeObj.OwnRange.ys
    nodeObjYe := nodeObj.OwnRange.ye
    searchNodeObjXs := searchNodeObj.OwnRange.xs
    searchNodeObjXe := searchNodeObj.OwnRange.xe
    searchNodeObjYs := searchNodeObj.OwnRange.ys
    searchNodeObjYe := searchNodeObj.OwnRange.ye
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
    if (searchNodeObjXe < nodeObjXs || searchNodeObjXs > nodeObjXe) {
            removeFlag =true
    }

    //y軸方向で見たときに隣合わないノード(x軸は同じ)
    if searchNodeObjYe < nodeObjYs || searchNodeObjYs > nodeObjYe{
        removeFlag =true
    }
    return removeFlag
}


// bootstrapのコンストラクタ
// 戻り値は構造体のポイント
func (nodeObj *Node)getDividedRange(x int , y int) *Range {
    newRangeObj := nodeObj.OwnRange.divide(x, y)
    return newRangeObj
}


// 再帰を使ってx,yを担当しているノードを取得する
func (nodeObj *Node) FindNode(x int,y int) *Node {

    if nodeObj.OwnRange.contain(x,y){
        return nodeObj
    }

    mostSmallAbsNodeObj := new(Node)
    //math.absはfloat64しか対応していない
    //初期値として絶対にありえないAbsErrorを設定する
    mostSmallAbsError := 100
    mostSmallAbsErrorInt := int(mostSmallAbsError)
    for _,searchNodeObj := range nodeObj.neighbors {
        searchNodeObjXMiddle := searchNodeObj.OwnRange.xMiddle
        searchNodeObjYMiddle := searchNodeObj.OwnRange.yMiddle
        absError := math.Abs(float64(searchNodeObjXMiddle - x)) + math.Abs(float64(searchNodeObjYMiddle - y))
        absErrorInt := int(absError)
        if absErrorInt <= mostSmallAbsErrorInt{
            mostSmallAbsErrorInt = absErrorInt
            mostSmallAbsNodeObj = searchNodeObj

        }
    }
    return mostSmallAbsNodeObj.FindNode(x,y)
}

// bootstrapのコンストラクタ
// 戻り値は構造体のポイント
func (nodeObj *Node)ShowInfo() {
    fmt.Println(nodeObj.OwnRange)
    for _,neighNodeObj := range nodeObj.neighbors {
        fmt.Println(neighNodeObj.OwnRange)
    }
}
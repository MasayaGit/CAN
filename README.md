# CAN
A Scalable Content-Addressable Network の実装
https://people.eecs.berkeley.edu/~sylvia/papers/cans.pdf

n次元空間(ここでは２次元で0<=x<=10,0<=y<=10)を複数のノードによって分割する. その領域を担当するノードがデータを保持したりタスクを行う.

main.goにて10, 10を担当するノードの情報が表示されるかを確認した
chargeNode := bootstrap.FindNode(10,10)

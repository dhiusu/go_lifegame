package main

import (
	fmt "fmt" // 入出力フォーマットを実装したパッケージ
	"math/rand"
	"time"

	"github.com/thoas/go-funk"
)

/*
誕生
ある生命が死んでいる場合、周囲８セルの内周囲８セルの内ちょうど３つのセルが生きていれば次の世代で新しい生命が誕生する。

生存
ある生命が生きている場合、周囲８セルの内２セルあるいは３セルが生きている場合、次の世代でも生存する。

過疎
ある生命が生きている場合、周囲８セルの内生きているセルが１セル以下の場合、次の世代では過疎で死滅する。

過密
ある生命が生きている場合、周囲８セルの内生きているセルが4セル以上の場合、次の世代では過密で死滅する。
*/

// alive is const value.
const alive = 1

// death is const value.
const death = 0

// position const values.
const leftTop = 7
const centerTop = 0
const rightTop = 1
const leftMiddle = 6
const rightMiddle = 2
const leftBottom = 5
const centerBottom = 4
const rightBottom = 3

func main() {

	// Matrixを生成
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	// ランダムで初期生命を設定
	rand.Seed(time.Now().UnixNano())
	yMax := len(matrix)
	xMax := len(matrix[0])
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			matrix[y][x] = rand.Intn(2) // 0, 1のランダム数値を生成する。
		}
	}

	generation := 1

	// Loop
	for {

		t := time.NewTicker(1 * time.Second) // 1秒おきに通知
		for {
			select {
			case <-t.C:

				// 世代数を描画
				fmt.Printf("Generation %d\n", generation)

				// Matrixの描画
				matrixStr := ""
				for y := 0; y < yMax; y++ {
					for x := 0; x < xMax; x++ {
						if matrix[y][x] == death {
							matrixStr += "□"
						} else {
							matrixStr += "■"
						}
						matrixStr += " "
					}
					matrixStr += "\n"
				}
				fmt.Println(matrixStr)

				// ゲームの更新

				for y := 0; y < yMax; y++ {
					for x := 0; x < xMax; x++ {
						if matrix[y][x] == alive {

							// 周辺の状態を１次元配列で取得
							rounds := []int{0, 0, 0, 0, 0, 0, 0, 0}
							for p := 0; p < 8; p++ {

								status := -1
								switch p {
								case centerTop:
									if (y - 1) < 0 {
										status = death
									} else {
										status = matrix[y-1][x]
									}
								case rightTop:
									if (y-1) < 0 || (x+1) >= xMax {
										status = death
									} else {
										status = matrix[y-1][x+1]
									}
								case rightMiddle:
									if (x + 1) >= xMax {
										status = death
									} else {
										status = matrix[y][x+1]
									}
								case rightBottom:
									if (y+1) >= yMax || (x+1) >= xMax {
										status = death
									} else {
										status = matrix[y+1][x+1]
									}
								case centerBottom:
									if (y + 1) >= yMax {
										status = death
									} else {
										status = matrix[y+1][x]
									}
								case leftBottom:
									if (y+1) >= yMax || (x-1) < 0 {
										status = death
									} else {
										status = matrix[y+1][x-1]
									}
								case leftMiddle:
									if (x - 1) < 0 {
										status = death
									} else {
										status = matrix[y][x-1]
									}
								case leftTop:
									if (y-1) < 0 || (x-1) < 0 {
										status = death
									} else {
										status = matrix[y-1][x-1]
									}
								}
								rounds[p] = status

								// if (y - 1) < 0 {
								// 	// Matrixから上にはみ出している場合
								// 	rounds[p] = death
								// }
								// if (y + 1) >= yMax {
								// 	// Matrixから下にはみ出している場合
								// 	rounds[p] = death
								// }
								// if (x - 1) < 0 {
								// 	// Matrixから左にはみ出している場合
								// 	rounds[p] = death
								// }
								// if (x + 1) >= xMax {
								// 	// Matrixから右にはみ出している場合
								// 	rounds[p] = death
								// }

							}

							// if (y - 1) < 0 {
							// 	// Matrixから上にはみ出している場合
							// 	rounds[centerTop] = death
							// 	rounds[rightTop] = death
							// 	rounds[leftTop] = death
							// }
							// if (y + 1) >= yMax {
							// 	// Matrixから下にはみ出している場合
							// 	rounds[centerBottom] = death
							// 	rounds[rightBottom] = death
							// 	rounds[leftBottom] = death
							// }
							// if (x - 1) < 0 {
							// 	// Matrixから左にはみ出している場合
							// 	rounds[leftBottom] = death
							// 	rounds[leftMiddle] = death
							// 	rounds[leftTop] = death
							// }
							// if (x + 1) >= xMax {
							// 	// Matrixから右にはみ出している場合
							// 	rounds[rightTop] = death
							// 	rounds[rightMiddle] = death
							// 	rounds[rightBottom] = death
							// }

							alives := funk.Filter(rounds, func(s int) bool {
								return s == alive
							})
							aliveCount := len(alives.([]int))

							// Rule1
							if aliveCount == 3 {
								// Create a life.
								for {
									ry := rand.Intn(yMax) // Y座標を生成
									rx := rand.Intn(xMax) // X座標を生成
									if matrix[ry][rx] == death {
										matrix[ry][rx] = alive
										break
									}
								}
							}

							// Rule2
							if aliveCount == 2 || aliveCount == 3 {
								// Alive.
							}

							// Rule3
							if aliveCount <= 1 {
								matrix[y][x] = death
							}

							// Rule4
							if aliveCount >= 4 {
								matrix[y][x] = death
							}

						}
					}
				}

				// ゲーム終了判定
				isOver := true
				for y := 0; y < yMax; y++ {
					for x := 0; x < xMax; x++ {
						if matrix[y][x] == alive {
							isOver = false
						}
					}
				}
				if isOver {
					// ゲーム終了
					break
				}

				// 世代更新
				generation++

			}
		}
		t.Stop() // タイマを止める。

	}

}

package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"math"
	"os"
	"strings"
	"testing"
)

func TestDecodePic(t *testing.T) {
	// label_mask
	pics := `iVBORw0KGgoAAAANSUhEUgAAAQAAAADACAAAAADOhuK6AAABoElEQVR4AeXBgW3DQBDAMGn/oa+Jvyu8cYFIiZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4iZM4uUOGnyBXyMfwA+QGeQz7yQVyDPvJBXIM+8kFcgz7yQVyDPvJDfIY9pMr5GP4ARIncRIncRIncRIncRIncRIncRInmwgMr5JF5Gt4k+whx/Ai2UOO4UWyhxzDi2QPOYYXyR5yDC+SReRreJNsIjC8SuIkTuIkTuIkTuIkTuIkTuIkTuIkTuIkTuIkTuLkHmHYTq6Rr2E3uUUew25yixzDanKLHMNqcon8G1aTW+QYVpNb5DHsJtfIx7Cc3CMM20mcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxP0BhKscwa3upNMAAAAASUVORK5CYII=`
	//blob
	//pics := `iVBORw0KGgoAAAANSUhEUgAAAQAAAADACAAAAADOhuK6AAACGElEQVR4AeXByXGEMAAAwZkAyIwfcfIjMxKwfK0TWFFyTbfESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzESZzMcbHzL8gUF8POPyAzXHzZWZ9McPFjZ3nyfhcvO6uTt7sYdi6GndXJ210MOxfDzurk7S6GnYthZ3Xydhd/dlYn73fxsrM6meDix87yZIaLLzvrkykuhp1/QOIkTuIkTuIkTuIkTuJkJTew8ShZyM2njSfJOm6+bTxIlnHza+M5soqbPxuPkVXcDBs3w8ZjZBU3w8bNsPEYWcXNsHEzbDxGVnHzZ+MxsoybXxvPkXXcfNt4kCzk5tPGk2QlN7DxKImTOImTOImTOImTOImTOImTOImTOImTeU44WJ1Mc/LpYG0yy8m3g6XJJCe/DlYmc5y8HKxM5jgZDk6Gg4XJHCfDwclwsDCZ4mQ44GQ4WJjMcTIcnAwHC5M5Tl4OViaTnPw4WJrMcvLlYG0yzclwsDiZ54SD1UmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxEmcxH0AWJAyweJU8FcAAAAASUVORK5CYII=`
	bytes, err := base64.StdEncoding.DecodeString(pics)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("hysen_pic_label_mask.png", bytes, os.ModePerm)

	if err != nil {
		panic(err)
	}

	//rb := new(io.Buffer)
	reader := strings.NewReader(string(bytes))

	canvas, str, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
	fmt.Println(canvas.ColorModel())
	bounds := canvas.Bounds()
	fmt.Println(bounds.Min, bounds.Max)
	switch canvas.(type) {
	case *image.NRGBA:
		//img := canvas.(*image.NRGBA)
		fmt.Println("NRGBA")
	case *image.RGBA:
		//img := canvas.(*image.RGBA)
		fmt.Println("NRGBA")
	case *image.Gray:
		fmt.Println("Gray")
		img := canvas.(*image.Gray)

		var points [][]uint8
		var line []uint8
		first := -1
		for i, pix := range img.Pix {
			if pix != 0 {
				if first == -1 {
					first = i
				}
			}
			//fmt.Print(pix)

			//x := (i + 1) / img.Stride
			//y := (i+1)%img.Stride - 1
			line = append(line, pix)
			if (i+1)%img.Stride == 0 {
				//fmt.Println()
				points = append(points, line)
				line = nil
			}
		}
		fmt.Println("==========================================")
		fmt.Println("first", first)
		allMap := map[string]point{}
		var mpPoints [][]point
		allPics := []map[string]point{}
		for y := 0; y < len(points); y++ {
			for x := 0; x < len(points[y]); x++ {
				if points[y][x] == 0 {
					continue
				}
				start := point{
					X: x,
					Y: y,
				}
				kk := start.getK()
				if _, ok := allMap[kk]; ok {
					continue
				}
				mp := map[string]point{}
				dfs(start, points, mp, img.Bounds().Dx(), img.Bounds().Dy())
				if len(mp) > 0 {
					for k, v := range mp {
						allMap[k] = v
					}
					allPics = append(allPics, mp)
				}
				//fmt.Println("========================start", y, x)
			}

			//fmt.Println()
		}
		fmt.Println("pics.len", len(mpPoints))
		{
			maxX := -1
			maxY := -1
			for _, v := range allMap {
				if maxX < v.X {
					maxX = v.X
				}
				if maxY < v.Y {
					maxY = v.Y
				}
			}

			for y := 0; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					p := point{
						X: x,
						Y: y,
					}
					k := p.getK()
					if _, ok := allMap[k]; ok {
						fmt.Print(1)
					} else {
						fmt.Print(0)
					}
				}
				fmt.Println()
			}
		}
		//todo 深度搜索找出属于同一张图片的点位

		//for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		//	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		//		location := (y-bounds.Min.Y)*img.Stride + (x-bounds.Min.X)*1
		//		fmt.Print(img.Pix[location])
		//	}
		//	fmt.Println("")
		//}
	}

}

var left, right, top, boot = -1, 1, -1, 1

type point struct {
	X int
	Y int
}

func (p point) getK() string {
	return fmt.Sprintf("%d,%d", p.Y, p.X)
}

func (p point) Top() *point {
	return &point{
		X: p.X,
		Y: p.Y - 1,
	}
}

func (p point) TopLeft() *point {
	return &point{
		X: p.X - 1,
		Y: p.Y - 1,
	}
}
func (p point) Left() *point {
	return &point{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p point) LeftBoot() *point {
	return &point{
		X: p.X - 1,
		Y: p.Y + 1,
	}
}
func (p point) Boot() *point {
	return &point{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p point) BootRight() *point {
	return &point{
		X: p.X + 1,
		Y: p.Y + 1,
	}
}
func (p point) Right() *point {
	return &point{
		X: p.X + 1,
		Y: p.Y,
	}
}

func (p point) RightTop() *point {
	return &point{
		X: p.X + 1,
		Y: p.Y + 1,
	}
}
func (p point) Out(xMax, yMax int) bool {
	if p.X < 0 || p.X >= xMax {
		return true
	}
	if p.Y < 0 || p.Y >= yMax {
		return true
	}
	return false
}

// mp 表示一个图的集合 key= "x,y"   value=像素值
// pics 也表示一个图的集合
func dfs(point point, points [][]uint8, mp map[string]point, xMax, yMax int) {
	//节点越界
	if point.Out(xMax, yMax) {
		return
	}
	k := point.getK()
	x := point.X
	y := point.Y
	if points[y][x] == 0 {
		return
	}

	if _, ok := mp[k]; !ok {
		mp[k] = point
		Top := *point.Top()
		dfs(Top, points, mp, xMax, yMax)
		TopLeft := *point.TopLeft()
		dfs(TopLeft, points, mp, xMax, yMax)
		Left := *point.Left()
		dfs(Left, points, mp, xMax, yMax)
		LeftBoot := *point.LeftBoot()
		dfs(LeftBoot, points, mp, xMax, yMax)
		Boot := *point.Boot()
		dfs(Boot, points, mp, xMax, yMax)
		BootRight := *point.BootRight()
		dfs(BootRight, points, mp, xMax, yMax)
		Right := *point.Right()
		dfs(Right, points, mp, xMax, yMax)
		RightTop := *point.RightTop()
		dfs(RightTop, points, mp, xMax, yMax)
	}
	return
}

// todo 判断所有的点 是否在pointA pointB的同一侧 满足条件的话
// todo 也许要考虑横线 和 垂线等特殊情况
// y*(x2-x1)+x(y1-y2)+x1*y2-y1*x2 = 0
func isOutSide(a, b point, mp map[string]point) bool {
	z1 := 0
	for _, p := range mp {
		if z1 == 0 {
			z1 = p.line(a, b)
			if z1 == 0 {
				continue
			}
		}
		z := p.line(a, b)
		result := z * z1
		if result < 0 { //小于0的化表示不同侧
			return false
		}
	}
	return true
}

func (p point) onLine(a, b point) bool {
	return p.line(a, b) == 0
}

func (p point) line(a, b point) int {
	return p.Y*(b.X-a.X) + p.X*(a.Y-a.Y) + a.X*b.Y - a.Y*b.X
}

func findNext(head, node *Node, mp map[string]point) {
	for _, p := range mp {

		if node.Point.same(p) {
			continue
		}

		if node.Pre != nil && node.Pre.Point.same(p) {
			continue
		}

		out := isOutSide(node.Point, p, mp)
		if !out {
			continue

		}
		//如果新的点  和 head坐标一致则说明结束了,
		if head.Point.same(p) {
			return
		}
		//如果这个点在当前点与一个点组成的线上面，则忽略
		if p.onLine(node.Point, node.Pre.Point) {
			continue
		}
		next := &Node{
			Point: p,
			Next:  nil,
		}
		if node.Next == nil { //如果还不存在下一个节点，则加入
			node.Add(next)
		} else { //如果已经存在下一个点，则看谁的距离最远，选距离最远的
			if node.Point.length(node.Next.Point) < node.Point.length(p) {
				node.Add(next)
			}
		}

	}
	findNext(head, node.Next, mp)
}

func (p point) same(a point) bool {
	return p.X == a.X && p.Y == a.Y
}

func (p point) length(b point) float64 {
	return math.Sqrt(float64((p.X-b.X)*(p.X-b.X)) + float64((p.Y-b.Y)*(p.Y-b.Y)))
}

type Node struct {
	Point point
	Pre   *Node
	Next  *Node
}

func (node *Node) Add(next *Node) {
	node.Next = next
	next.Pre = node
}

package main

import ("fmt"
	"os"
	"math"
)
func distance(x1,x2,y1,y2 int) float64{
	return math.Sqrt(math.Pow(( float64(x2-x1)), 2) + math.Pow(( float64(y2-y1)), 2))
}


func main() {
	var catSpeed int
	fmt.Scan(&catSpeed)
	x,y:=0,0

	// radius := 500
	i:=10
	for {
		var mouseX, mouseY, catX, catY int
		fmt.Scan(&mouseX, &mouseY, &catX, &catY)
		f:= distance(mouseX, catX, mouseY,catY)
		//angle := getAngle(catX,catY,mouseX,mouseY)
		angle := getAngle2(mouseX,mouseY,catX,catY)
		// if angle < 45{
		if angle < 30 {
			x =  mouseX+(-1*i)
			y = mouseY+i
		}else if angle >30 && angle < 90{
			fmt.Fprintln(os.Stderr,"Here:",x,y)
			x =  mouseX+ (-1*i)
			y =  mouseY+ (-1*i)
			fmt.Fprintln(os.Stderr,"Here2:",x,y)
		}else{
			y =  mouseY+ (-1*i)
			x =  mouseX+ (1*i)
		}
		//}

		fmt.Fprintln(os.Stderr,"Angle:",angle)

		fmt.Fprintln(os.Stderr, fmt.Sprintf("Distance:%f, Cat(%d,%d), M(%d,%d)", f, catX,catY,mouseX,mouseY))
		fmt.Println(fmt.Sprintf("%d %d Escaping the cat",x,y ))

	}
}
func getAngle2(dx1, dy1, dx2, dy2 int) float64{
	x1:=float64(dx1)
	y1:=float64(dy1)
	x2 := float64(dx2)
	y2:= float64(dy2)

	n := x1 * y2 + x2*y1
	su := math.Sqrt(x1*x1+y1*y1)
	sv := math.Sqrt(x2*x2* + y2*y2)
	d  := su*sv
	aci:= math.Acos(n/d) * 180 / math.Pi

	if x1*y2> x2*y1{
		return aci
	}else{
		return 360-aci
	}

}
func getAngle(x2, y2, x1,  y1 int) float64{
	xDiff := float64(x2 - x1)
	yDiff := float64(y2 - y1)
	ang:= (math.Atan2(yDiff, xDiff) * (180 / math.Pi))
	if ang<0{
		return ang+360
	}
	return ang

}
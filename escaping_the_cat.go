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
	x,y:=0.0,0.0
	mouse_speed:= float64(10)
	cat_speed := float64(catSpeed)
	alpkan := 0.89000
	radius := float64(500)
	angle_iteration := float64(30)
	circle_distance := ((( float64(2) * math.Pi * radius) / cat_speed) * mouse_speed) / ( float64(2) * math.Pi)
	dash_distance := radius - ((math.Pi * radius) / cat_speed) * mouse_speed
	circle_r := ((circle_distance - dash_distance) * alpkan) + dash_distance
	dash := false
	circle:=false
	m:= float64(-100)
	for {
		var mouseX, mouseY, catX, catY int
		fmt.Scan(&mouseX, &mouseY, &catX, &catY)
		cat_mouse_angle := getAngle(catX,catY,0,0,mouseX,mouseY)
		mouse_0_1_angle := getAngle(1,0,0,0,mouseX,mouseY)
		distance_to_origin := distance(mouseX,0,mouseY,0)
		fmt.Fprintln(os.Stderr,fmt.Sprintf("cat speed%f\ndash:%b\n circle:%b\nangle_with_cat:%f\nangle_with_1_0:%f\ndistance_to_origin:%f\ncircle_r:%f ",
			cat_speed,dash,circle,cat_mouse_angle,mouse_0_1_angle,distance_to_origin,circle_r))

		if distance_to_origin <= circle_r && cat_mouse_angle <=180 && !dash{
			angle := getAngle(1,0,0,0,mouseX,mouseY) + angle_iteration
			x,y = findNextPosition(angle,circle_r)
			circle =true
			fmt.Fprintln(os.Stderr,"angle",angle)
			fmt.Fprintln(os.Stderr,"anglingggg")

		}else if 180 <= cat_mouse_angle && cat_mouse_angle <190 && !dash{
			x,y = m * float64(catX) ,m * float64(catY)
			fmt.Fprintln(os.Stderr,"dashhhinggg")
			dash=true
		}else{
			if !circle && !dash{
				x,y = m * float64(catX) ,m * float64(catY)
			}
		}

		fmt.Println(fmt.Sprintf("%.0f %.0f Escaping the cat",x,y ))

	}
}
func getAngle(a0,a1,b0,b1,c0,c1 int) float64{

	first := math.Atan2(float64(c1 - b1),float64(c0 - b0))
	second := math.Atan2(float64(a1 - b1),float64(a0 - b0))
	diff := first - second
	ang:=  convert("d",diff)
	if ang <0{
		return ang+360
	}
	return ang
}

func findNextPosition(a,r float64) (float64,float64){

	x:= r * math.Cos(convert("r",a))
	y:= r * math.Sin(convert("r",a))
	return x,y
}
func convert(s string,n float64)float64{
	if s=="r"{
		return n * (math.Pi/180)
	}else if s=="d"{
		return n * (180/math.Pi)
	}
	return 0
}

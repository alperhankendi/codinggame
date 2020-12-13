package main

import (
	"fmt"
	"math"
)
func main() {
	surface := NewSurface()
	rover := NewRover(surface)
	for {
		rover.Read()
		outputThrust,outputAngle := rover.Accelerate()
		fmt.Println(fmt.Sprintf("%d %d", outputAngle,outputThrust ))
	}
}
//Constants
const  farDistance float64 = 2000
const  initHorizontalSpeed int = 80
const  desiredVerticalLandingSpeed int = 40
type Point struct{
	Left    int
	Mid     int
	Right   int
}
type Surface struct {
	Points []Point
	TargetPoint Point
}
type Rover struct {
	X, Y, hSpeed, vSpeed, fuel, rotate, power,initHspeed int
	Surface		*Surface
}

func NewSurface() *Surface {

	var points []Point
	var flat Point
	var surfaceN int
	fmt.Scan(&surfaceN)
	for i := 0; i < surfaceN; i++ {
		var landX, landY int
		fmt.Scan(&landX, &landY)
		points = append(points, Point{
			Left:landX,
			Right:landY,
		})
		if i>0 && points[i].Right == points[i-1].Right{
			flat = Point{
				Left: points[i-1].Left,
				Right: points[i].Left,
				Mid: int( (points[i-1].Left + points[i].Left) / 2),
			}
		}
	}
	return &Surface{
		Points:      points,
		TargetPoint: flat,
	}
}

func NewRover(surface *Surface) *Rover {
	rover := &Rover{
		Surface: surface,
	}
	fmt.Scan(&rover.X, &rover.Y, &rover.hSpeed, &rover.vSpeed, &rover.fuel, &rover.rotate, &rover.power)
	fmt.Println("0 0")
	rover.initHspeed = initHorizontalSpeed
	if Single(surface.TargetPoint.Mid - rover.X) == Single(rover.hSpeed){
		rover.initHspeed= int(math.Abs(float64(rover.hSpeed)))
	}
	return rover
}
func (receiver Rover) isFar() bool {

	if math.Abs(float64(receiver.Surface.TargetPoint.Mid - receiver.X)) > farDistance{
		return true
	}
	return false
}
func (receiver Rover) isInTargetRange() bool {

	return receiver.X >= receiver.Surface.TargetPoint.Left &&
		receiver.X <= receiver.Surface.TargetPoint.Right
}
func (receiver *Rover) Read()  {

	fmt.Scan(&receiver.X, &receiver.Y, &receiver.hSpeed, &receiver.vSpeed, &receiver.fuel, &receiver.rotate, &receiver.power)

}
func (receiver Rover) landingVerticalSpeedIsOk() bool {
	return math.Abs(float64(receiver.vSpeed)) < float64(desiredVerticalLandingSpeed)
}
func (receiver Rover) calculateAngle(desiredSpeedX,maxAngle int) (outputAngle int) {

	// -1=left, +1=right, 0 = no move horizantially
	desiredDirection := Single(receiver.Surface.TargetPoint.Mid - receiver.X)
	desiredVelocity := desiredSpeedX * desiredDirection
	desiredAcceleration := desiredVelocity - receiver.hSpeed
	outputAngle = Clamp(desiredAcceleration * 2, -1* maxAngle, 1* maxAngle)

	return
}
func (receiver Rover) Accelerate() (outputThrust,outputAngle int) {

	outputThrust 	=   4
  	outputAngle 	=   0
	desiredSpeedX	:=  0
	maxAngle 		:= 33
	if  receiver.isInTargetRange() {

		if receiver.landingVerticalSpeedIsOk() && receiver.hSpeed ==0{
			outputThrust =3
		}
		maxAngle = 33
	}else{
		if receiver.vSpeed > 0{
			outputThrust=3
		}
		if receiver.isFar(){
			desiredSpeedX= receiver.initHspeed
			maxAngle = 60
		}else{
			desiredSpeedX=20
			maxAngle = 45
		}
	}
	outputAngle = receiver.calculateAngle(desiredSpeedX, maxAngle)
	return
}
func Single(n int) int{

	if n>0{
		return 1
	}else if n ==0{
		return 0
	}else if n <0{
		return -1
	}
	return 0
}
func Clamp(v, lo, hi int) int {
	val := Min(Max(v, lo), hi)
	return -1 * val
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
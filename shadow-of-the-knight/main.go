package main
import ("fmt"
	"os"
)
func main() {
	var W, H int
	fmt.Scan(&W, &H)

	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	var HighX,HighY int = W-1, H-1 // index-0
	var LowX,LowY int = 0,0
	fmt.Scan(&X0, &Y0)
	isCenter := false
	counter:=0
	for {
		var bombDir string
		fmt.Scan(&bombDir)

		if !isCenter{
			isCenter=true
			X0 = (W-1)/2 // index-0
			Y0 = (H-1)/2
		}else {
			for _,k := range bombDir{
				if k=='L'{
					HighX = X0 - 1
				}else if k=='R'{
					LowX = X0 + 1
				}else if k=='U'{
					HighY = Y0 -1
				}else if k=='D'{
					LowY = Y0 +1
				}
			}
		}
		X0 = (HighX + LowX) / 2
		Y0 = (HighY + LowY) / 2
		counter++
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Counter:%d,Command: %s, HighX,Y(%d,%d), LowX,Y(%d,%d), X0:%d,Y0:%d",counter, bombDir,HighX,HighY,LowX,LowY,X0,Y0 ))
		fmt.Println(fmt.Sprintf("%d %d",X0,Y0 ))
	}

}

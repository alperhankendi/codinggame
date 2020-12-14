package codinggame
import ("fmt"
"math"
"os"
)
func distance(x1,x2,y1,y2 int) float64{

	return math.Sqrt(math.Pow(( float64(x2-x1)), 2) + math.Pow(( float64(y2-y1)), 2))
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var A string
		var xA, yA int
		var B string
		var xB, yB int
		var C string
		var xC, yC int
		var D string
		var xD, yD int
		var parallelogram,rhombus,rectangle,square bool = false,false,false,false
		fmt.Scan(&A, &xA, &yA, &B, &xB, &yB, &C, &xC, &yC, &D, &xD, &yD)

		dist_a_b := distance(xA, xB, yA,yB )
		dist_b_c := distance(xB, xC, yB,yC )
		dist_c_d := distance(xC, xD, yC,yD )
		dist_d_a := distance(xD, xA, yD,yA )
		dist_a_c := distance(xC, xA, yC,yA )
		dist_b_d := distance(xD, xB, yD,yB )
		fmt.Fprintln(os.Stderr, fmt.Sprintf("A_B:%f\nB_C:%f\nC_D:%f\nD_A:%f\nA_C:%f\nB_D:%f",dist_a_b,dist_b_c,dist_c_d,dist_d_a,dist_a_c,dist_b_d))
		fmt.Print(A+B+C+D+" is a ")

		if dist_a_b == dist_c_d && dist_b_c == dist_d_a{
			parallelogram = true
		}
		if (dist_a_b == dist_c_d) && (dist_b_c == dist_d_a) && (dist_a_b==dist_b_c) {
			rhombus = true
		}
		if dist_a_c == dist_b_d{
			rectangle=true
		}

		fmt.Fprintln(os.Stderr ,parallelogram,rhombus,rectangle,square )

		if rectangle && rhombus{
			fmt.Println("square.")
		}else if rectangle && !rhombus{
			fmt.Println("rectangle.")
		}else if !rectangle && rhombus{
			fmt.Println("rhombus.")
		}else if parallelogram{
			fmt.Println("parallelogram.")
		}else{
			fmt.Println("quadrilateral.")
		}
	}

}

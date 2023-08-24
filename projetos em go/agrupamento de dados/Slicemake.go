
import "fmt"

func main() {
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
	slice = append(slice, 11) // capacidade do array aumentada para 20 agora

	fmt.Println(slice, len(slice), cap(slice))

}


//tela antes dos append
// [ 1, 2, 3, 4, 5] 5 10  


// tela depois
// [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11] 11 20 
// números fora do colchetes: 11 len, quantos tem
// números fora do colchetes; 20 cap, capacidade

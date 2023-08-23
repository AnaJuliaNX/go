///////////////////////////////////////////////////////////
//PACOTE SORT COM STRING

package main

import (
	"fmt"
	"sort"
)

func main() {

	slicestring := []string{"batata", "agua", "dadinho", "caramelo"}

	fmt.Println(slicestring)
	sort.Strings(slicestring) //precisa ser maiusculo e com "s" no final
	fmt.Println(slicestring)
}

//tela
//[batata agua dadinho caramelo]
//[agua batata caramelo dadinho]

///////////////////////////////////////////////////////////
                      //PACOTE SORT COM INT
package main

import (
	"fmt"
	"sort"
)

func main() {

	sliceint := []int{2, 1, 3, 4, 6, 5}

	fmt.Println(sliceint)
	sort.Ints(sliceint)    //precisa ser maiusculo e com "s" no final
	fmt.Println(sliceint)
}
//tela
//[2 1 3 4 6 5]
//[1 2 3 4 5 6]

///////////////////////////////////////////////////////////
                     //PACOTE SORT COM FLOAT64

package main

import (
	"fmt"
	"sort"
)

func main() {

	slicefloat := []float64{2.2, 1.1, 3.3, 3.4, 6.6, 5.5, 4.4}

	fmt.Println(slicefloat)
	sort.Float64s(slicefloat)  //precisa ser maiusculo e com "s" no final
	fmt.Println(slicefloat)
}

//tela
//[2.2 1.1 3.3 3.4 6.6 5.5 4.4]
//[1.1 2.2 3.3 3.4 4.4 5.5 6.6]

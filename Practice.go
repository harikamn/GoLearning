package main
import "fmt"
func main(){

	fmt.Println("Hello")



	x := [6]string{"a","b","c","d","e","f"}

	fmt.Println(x[2:5])

	x1 := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	var smallest = x1[0]
	//var largest = x1[0]
	for i:=0;i<len(x1); i++ {
		if x1[i]<smallest {
			smallest = x1[i]
		}

	}
	fmt.Println(smallest)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

		elements2 := map[string]map[string]string{
			"H": map[string]string{
				"name":"Hydrogen",
				"state":"gas",
			},
			"He": map[string]string{
				"name":"Helium",
				"state":"gas",
			},
			"Li": map[string]string{
				"name":"Lithium",
				"state":"solid",
			},
			"Be": map[string]string{
				"name":"Beryllium",
				"state":"solid",
			},
			"B":  map[string]string{
				"name":"Boron",
				"state":"solid",
			},
			"C":  map[string]string{
				"name":"Carbon",
				"state":"solid",
			},
			"N":  map[string]string{
				"name":"Nitrogen",
				"state":"gas",
			},
			"O":  map[string]string{
				"name":"Oxygen",
				"state":"gas",
			},
			"F":  map[string]string{
				"name":"Fluorine",
				"state":"gas",
			},
			"Ne":  map[string]string{
				"name":"Neon",
				"state":"gas",
			},
		}

		if el, ok := elements2["Li"]; ok {
			fmt.Println(el["name"], el["state"])
		}
	}
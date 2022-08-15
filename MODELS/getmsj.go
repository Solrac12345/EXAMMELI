package MODELS

import (
	"fmt"
	"strings"
)

func LeerMensaje(men1 []string, men2 []string, men3 []string) string {

	MensajeFinal := make([]string, len(men1))

	for i := 0; i < len(men1); i++ {
		if men1[i] != "''" {

			MensajeFinal[i] = men1[i]
		}
	}
	for k := 0; k < len(men1); k++ {
		if men2[k] != "''" {

			MensajeFinal[k] = men2[k]
		}
	}
	for j := 0; j < len(men1); j++ {
		if men3[j] != "''" {

			MensajeFinal[j] = men3[j]

		}
	}
	fmt.Println("--------------MENSAJE FINAL----------------")

	a := strings.Join(MensajeFinal, " ")
	fmt.Println(a)
	return a
}

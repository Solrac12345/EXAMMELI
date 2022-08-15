package main

import (
	"fmt"
	"paquetes/MODELS"
)

func main() {

	t1 := MODELS.Satelite{}
	t1.Constructor("Kenobi", 832, 5148, 87, ", es un mensaje")

	t2 := MODELS.Satelite{}
	t2.Constructor("Skywalker", 741, 5242, 69, "Este , un ,")

	t3 := MODELS.Satelite{}
	t3.Constructor("Sato", 864, 5245, 55, "Este , , mensaje")

	list := MODELS.TaskSat{}
	list.AppendTask(&t1)
	list.AppendTask(&t2)
	list.AppendTask(&t3)

	t1.ToPrint()
	t2.ToPrint()
	t3.ToPrint()

	fmt.Println("--------------MENSAJES SATELITES----------------")

	for _, task := range list.Tasks {
		fmt.Println(task.GetMjs())
	}

	men1 := list.ArreglarMensaje(&t1)
	men2 := list.ArreglarMensaje(&t2)
	men3 := list.ArreglarMensaje(&t2)

	MODELS.LeerMensaje(men1, men2, men3)

	a := t1.GetCoorx()
	b := t1.GetCoory()
	c := t1.GetDist()

	d := t2.GetCoorx()
	e := t2.GetCoory()
	f := t2.GetDist()

	g := t3.GetCoorx()
	h := t3.GetCoory()
	i := t3.GetDist()

	Coor := MODELS.GetLocation(a, b, c, d, e, f, g, h, i)

	fmt.Println("--------------COORDENADAS DEL SATELITE----------------")

	fmt.Println("La coordenada en x es:", Coor[0])
	fmt.Println("La coordenada en y es:", Coor[1])
}

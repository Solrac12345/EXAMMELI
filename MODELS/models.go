package MODELS

import (
	"fmt"
	"strings"
)

type TaskSat struct {
	Tasks []*Satelite
}

type Satelite struct {
	nomsat  string
	coorx   float64
	coory   float64
	dist    float64
	mensaje string
}

func (tl *TaskSat) AppendTask(t *Satelite) {
	tl.Tasks = append(tl.Tasks, t)

}

func (s *TaskSat) ArreglarMensaje(t *Satelite) []string {

	cmensaje := strings.ReplaceAll(t.mensaje, ",", "''")
	arraymensaje := strings.Split(cmensaje, " ")
	return arraymensaje
}

func (t *Satelite) ToPrint() {
	fmt.Printf("Nombre: %s\nCoodenanda_x: %g\nCoodenanda_y: %g\nmensaje: %s\n\n", t.nomsat, t.coorx, t.coory, t.mensaje)

}

func (s *Satelite) Constructor(nombre string, coorx float64, coory float64, dist float64, mensaje string) {
	s.nomsat = nombre
	s.coorx = coorx
	s.coory = coory
	s.dist = dist
	s.mensaje = mensaje
}

func (s *Satelite) GetNombsat() string {
	return s.nomsat
}

func (s *Satelite) SetNombsat(nombre string) {
	s.nomsat = nombre
}

func (s *Satelite) GetCoorx() float64 {
	return s.coorx
}

func (s *Satelite) SetCoorx(coorx float64) {
	s.coorx = coorx
}

func (s *Satelite) GetCoory() float64 {
	return s.coory
}

func (s *Satelite) SetCoory(coory float64) {
	s.coory = coory
}

func (s *Satelite) GetDist() float64 {
	return s.dist
}

func (s *Satelite) SetDist(dist float64) {
	s.dist = dist
}

func (s *Satelite) GetMjs() string {
	return s.mensaje
}

func (s *Satelite) SetMjs(mensaje string) {
	s.mensaje = mensaje
}

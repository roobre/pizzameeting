package main

import (
	"fmt"
	"roob.re/pizzameeting"
	"roob.re/pizzameeting/combinator"
	"roob.re/pizzameeting/pizzerias"
)

func main() {
	m := pizzameeting.Meeting{}
	m.CombinatorMaker = combinator.NewRecursiveCombinatorMaker()
	m.Pizzeria = pizzerias.AlTaglioBCN

	roobre := &pizzameeting.Person{}
	roobre.Score("Margherita", 10)
	roobre.Score("Diavola", 11)
	roobre.Score("Prosciutto", 11)

	matteo := &pizzameeting.Person{}
	matteo.Score("Diavola", 10)
	matteo.Score("Prosciutto", 11)
	matteo.Score("Margherita", 11)

	mikel := &pizzameeting.Person{}
	mikel.Score("Coppata", 10)
	//mikel.Score("Prosciutto", 10)
	//mikel.Score("4Fromatges", 11)

	xavi := &pizzameeting.Person{}
	xavi.Score("Parmesana", 10)
	xavi.Score("4Fromatges", 11)
	xavi.Score("Margherita", 11)
	xavi.Score("Prosciutto", 11)

	fran := &pizzameeting.Person{}
	fran.Score("Vegetariana", 10)
	fran.Score("Margherita", 10)
	fran.Score("Parmesana", 10)
	fran.Score("4Fromatges", 11)

	jaume := &pizzameeting.Person{}
	jaume.Score("Prosciutto", 10)
	jaume.Score("Parmesana", 10)
	jaume.Score("Vegetariana", 10)
	jaume.Score("4Fromatges", 11)

	leonidas := &pizzameeting.Person{}
	leonidas.Score("Parmesana", 10)
	leonidas.Score("Prosciutto", 10)
	leonidas.Score("Margherita", 10)
	leonidas.Score("4Fromatges", 11)

	sergi := &pizzameeting.Person{}
	sergi.Score("Parmesana", 10)
	sergi.Score("Vegetariana", 10)
	sergi.Score("4Fromatges", 11)

	jordi := &pizzameeting.Person{}
	jordi.Score("Parmesana", 10)
	jordi.Score("Vegetariana", 10)
	jordi.Score("Huevos", 11)

	hamid := &pizzameeting.Person{}
	hamid.Score("Parmesana", 10)
	hamid.Score("Prosciutto", 10)
	hamid.Score("Margherita", 10)
	hamid.Score("4Fromatges", 11)

	m.Invite(roobre, matteo, mikel, xavi, fran, jaume, leonidas, sergi, jordi, hamid)
	fmt.Println(m.Menu())
}

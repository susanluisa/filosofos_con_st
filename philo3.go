package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type palillo struct {sync.Mutex}
type Filosofo struct {
	cod 						int
	palillo_left, palillo_right *palillo 
}

func (f Filosofo) dine(){

for j:=0; j<3;j++{

	f.palillo_left.Lock()
	f.palillo_right.Lock()

	imp("comenzando a comer " , f.cod)
	randomPause(3)

	f.palillo_left.Unlock()
	f.palillo_right.Unlock()

	imp("terminando de comer " , f.cod)
	randomPause(5)
	
		}
		wg.Done()
}

func randomPause (max int){
	time.Sleep(time.Millisecond*time.Duration(rand.Intn(max*1000)))
}
func imp(estado string, cod int){
	fmt.Printf("nro: %d  %s\n", cod +1,estado)
}
//func init(){
	//rand.Seed(time.Now().UTC().UnixNano())
//}
var wg sync.WaitGroup
func main(){
	num := 5
	apalillos :=make([]*palillo, num)
		for i:=0; i<num; i++{
			apalillos[i]=new(palillo)
		}

	aFilosofos:= make([]*Filosofo,num)
	for i:=0; i<num; i++{
		aFilosofos[i]=&Filosofo{
			cod:i, palillo_left:apalillos[i], palillo_right:apalillos[(i+1)%num]}
		wg.Add(1)
		 go aFilosofos[i].dine()
	}
	wg.Wait()
}

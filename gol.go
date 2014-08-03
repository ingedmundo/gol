package main

import "math/rand"

type Cell struct{
  status bool
}

func (c Cell) IsAlive() bool{
  return c.status
}

func (c *Cell) Die(){
  c.status = false
}

func (c *Cell) Live(){
  c.status = true
}

type Universe struct{
  Space [10][10]Cell
}

type Statistics struct{
  Alive int
  Dead int
}

func (u Universe) Statistics() Statistics{
  alive_count := 0
  dead_count := 0

  for i:=0; i < 10; i++{
    for j:=0; j < 10; j++{
      if u.Space[i][j].IsAlive() == true{
        alive_count++
      }else{
        dead_count++
      }
    }
  }

  var stats = Statistics{ Alive: alive_count, Dead: dead_count }
  return stats
}

func (u *Universe) BigBang(amount int){
  for i := 1; i < amount; i++{
    x := rand.Intn(9)
    y := rand.Intn(9)
    u.Space[x][y].Live()
  }
}


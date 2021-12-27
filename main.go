package main

import (
	"fmt"
	"log"
)

func main() {
  var month, seasons string
	fmt.Println("Времена года.\nВведите месяц:")
fmt.Scan(&month)

switch month{
  case "январь", "февраль", "декабрь":{
seasons="зима"
  }
  case "март", "апрель", "май":{
seasons="весна"
  } 
  case "июнь", "июль", "август":{
seasons="лето"
  } 
  case "сентябрь", "октябрь", "ноябрь":{
seasons="осень"
  }
default:{
  fmt.Println("Введено неизвестное наименование")
  log.Fatal()
}
}
 fmt.Printf("Время года — %s.", seasons)
 fmt.Println()
}

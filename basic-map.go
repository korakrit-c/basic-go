package main

import (
   "fmt"
)
func main() {
   
   basicMap();
   fmt.Printf("#######################\n");
   basicMap2();

}

func basicMap(){
   setName := make(map[int]string); // (map[Key]value)
   
   setName[1] = "Korakrit";
   setName[2] = "Arthur";
   setName[3] = "Chariyasathian";

   fmt.Println(setName);
}

func basicMap2(){
   setProfile := make(map[string]int);
   
   setProfile["Korakrit"] = 1234;
   setProfile["Arthur"] = 6969;
   setProfile["Chariyasathian"] = 2525;

   for key, value := range setProfile {
      fmt.Println("Name:", key, "\tMoney:", value)
   }
}

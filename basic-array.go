package main

import (
   "fmt"
)
func main() {
   
   basicArray1();
   basicArray2();
   basicSlice();

}

func basicArray1(){
   setA := []int{1, 2, 3, 4, 5};
   setB := [5]string{"One", "Two", "Three", "Four", "Five"};
   fmt.Println(setA);
   fmt.Println(setB);
}

func basicArray2(){
   setA := []int{1, 2, 3, 4, 5}
   setB := [5]string{"One", "Two", "Three", "Four", "Five"};
   
   for _, x := range setA {
      fmt.Printf("- %d\n", x);
   }

   for _, x := range setB {
      fmt.Printf("* %s\n", x);
   }

}

func basicSlice(){
   setIndex := []int{};
   setName := []string{};

   setName = append(setName, "Korakrit");
   setName = append(setName, "Arthur");
   setName = append(setName, "Chariyasathian");

   setIndex = append(setIndex, 1);
   setIndex = append(setIndex, 2);
   setIndex = append(setIndex, 3);

   if len(setIndex) == len(setName){
      for i := range setIndex{
         fmt.Printf("[%d] %s\n", setIndex[i], setName[i])
      }
   }

}
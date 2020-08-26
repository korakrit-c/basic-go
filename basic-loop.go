package main

func main() {
   println("For Loop")
   basicForLoop()

   println("Do While Loop")
   basicDoWhileLoop()

   println("White Loop")
   basicWhiteLoop()
}

func basicForLoop() {
   for i:=1; i<10; i++ {
      println(i);
   }
}

func basicDoWhileLoop() {
   i := 1;
   for {
      if i==69 {
         println(i);
         break;
      }
      i++;
   }
}

func basicWhiteLoop(){
   i := 1;
   for i <= 6 {
      print(i);
      i = i+1;
   }

}

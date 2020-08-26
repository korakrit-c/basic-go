package main

func main() {
   
   println("If Else");
   basicIfElse();

   println("If Else If");
   basicIfElseIf();

   println("ABC == X");
   doSomething()
}

func basicIfElse(){
   x := 10;
   
   if x < 10 {
      println("X is less than 10");
   } else {
      println("X is greater than or equal 10");
   }
}

func basicIfElseIf(){
   x := 10;
   
   if x < 10 {
      println("X is less than 10");
   } else if x == 10 {
      println("X is equal to 10");
   } else {
      println("X is greater than 10");
   }
}

func doSomething(){
   a, b, c := 1, 2, 3;
   x := 0;

   if x = a+b+c ; x==6 {
      println("X is equal to 6");

   }
}

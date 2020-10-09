package main

import "fmt"    
import "math"

func swap(x,y,z string) (string,string,string,string){
    return y,x,z,y
}

func (v Vertex) Abs() float64{
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    fmt.Printf("Hari ini %d-%s-%d",8,"oktober", 2020)
    fmt.Println("halo")

    //panggil swap
    a,b,c,d := swap("hello","world","hello")
    fmt.Printf("var 1: %s,var 2: %s,var 3: %s,var 4: %s",a,b,c,d)

    //panggil vertwx
    v:=Vertex{3,4}
    fmt.Println(v.Abs())

}
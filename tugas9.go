package main

import "fmt"
import "strconv"

func main()  {
  defer fixing()
  var input string
  fmt.Println("Masukan Angka Undian :")
  fmt.Scanln(&input)

  var number int
  var err error
  number, err = strconv.Atoi(input)

  if err== nil{
    fmt.Println(number, "Selamat No anda sudah terinput")
    } else{
      fmt.Println(input, "Maaf yang anda masukan bukan Angka Undian")
      //fmt.Println(err.Error()) //ini comment error biasa kalau mau error panic di bawah ini.
      panic(err.Error())
      fmt.Println("Coba Lagi") // apabila error biasa, comment ini akan muncul di bawahnya. dan apabila panic error tidak akan muncul
    }
}

func fixing()  { //fungsi recover untuk menyembunyikan nilai dari pada panic error yang di timbulkan.
  if r:=recover(); r!=nil{
    fmt.Println("Mohon di input kembali dengan angka Undian ya")
  }else{
    fmt.Println("Silakan tunggu, apabila kamu beruntung kami akan segera menghubungi anda")
  }
}

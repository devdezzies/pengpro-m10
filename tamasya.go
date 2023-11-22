package main

import "fmt"

func main() {
  // var jumlah, total int

  // total = 0
  // fmt.Scan(&jumlah)
  // for jumlah != 0 {
  //   if jumlah <= 7 {
  //     total = total + 1
  //     jumlah = jumlah / 7
  //   } else if jumlah > 7 {
  //     total = total + (jumlah / 7)
  //     jumlah = jumlah % 7
  //   }
  // }
  // fmt.Println(total, "Mobil dibutuhkan")

  var mobil, jumlah int

  fmt.Scan(&jumlah)
  mobil = jumlah / 7
  if jumlah % 7 != 0 {
    mobil = mobil + 1
  }
  fmt.Println(mobil)
}

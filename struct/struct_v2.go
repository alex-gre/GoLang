package main

import (
        "fmt"
        "os"
        )  

type goods struct {
  name string
  price float32
  percent float32 // процент наценки
  vol int         // колличество товара значение  
  date string     // дата поступления товара
  
}
  

func main() {


  // товар пальто
  var coat goods
 // create file goods.txt import os must include
  file, err := os.Create("goods.txt")

  if err != nil{
      fmt.Println(err) 
      os.Exit(1) 
  }
  defer file.Close() //закрыть после записи


     // ввод наименования товара
     fmt.Println("Введите наименование товара: ")
     fmt.Scan(&coat.name)
   
     // ввод цены за штуку
     fmt.Println("цену за штуку: ")
     fmt.Scan(&coat.price)
   
     // ввод розничной наценки (моржа за товар для бизнесмена)
     fmt.Println("процент наценки: ")
     fmt.Scan(&coat.percent)
   
     // ввод колличество штук (vol это volue - значение)
     fmt.Println("колличество: ")
     fmt.Scan(&coat.vol)
   
   
      // ввод даты поступления товара
     fmt.Println("дата поступления: ")
     fmt.Scan(&coat.date)
    
  /*
    coat:= goods{
      name   : "Пальто кашемировое",
     // price  : 100.50,
     // percent: 15.5,
     // vol    : 10,
      //date   : "24.01.2003",

      } 
   */
  

  //fmt.Printf("Наименование товара -> %s\n дата поступления -> %s",coat.name,coat.date)
  // для переноса использовать + 
  fmt.Fprintf(file, "Наименование товара => %s \n цена за штуку => %4.2f \n процент => %2.2f\n"+
  " дата привоза => %s\n  ", coat.name,coat.price,coat.percent,coat.date)
  fmt.Println("Файл goods.txt создан")

        
}
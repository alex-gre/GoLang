package main

import ("fmt"
        "net/http"
        "html/template")


type User struct {
     Name   string
     Age    uint16
     Money  int16
     Avg_grades,Heppiness float64
     Hobbies []string
} 
func (u User) getAllInfo() string {

    return fmt.Sprintf("User name is: %s.He is %d he has is money %d ",u.Name,u.Age,u.Money)
}

func (u *User) setNewName(newName string) {
    u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request){
    bob := User {"Bob",25,-50,4.2,0.8,  []string {"Football","Skate","Dance"}}
    //bob.setNewName("Alex")
     // fmt.Fprintf(w, bob.getAllInfo())
     tmpl, _ := template.ParseFiles("templates/home_page.html")
     tmpl.Execute(w, bob)
   
} 

func contacts_page(w http.ResponseWriter, r *http.Request){

        fmt.Fprintf(w,"Contacts page")
    } 

func hanlde_request(){

     http.HandleFunc("/", home_page)
     //http.HandleFunc("/contacts/", contacts_page)
     http.ListenAndServe(":8080",nil)
  
}     

func main()  {
    //bob := User {name: "Bob",age:25, money: -50 avg_grades:4.2,heppiness:0.8}
    	
     hanlde_request()

}		
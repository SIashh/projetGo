package main

import(
  "fmt"
  "bufio"
  "net"
  "sync"
  "strconv"
  "unicode/utf8"

)

var wg sync.WaitGroup


func main(){
  //le serveur accepte 20 clients
  wg.Add(20)
  listener,err := net.Listen("tcp",":9999")
  if err != nil {
    fmt.Println("Problème d'écoute sur le port 9999")
  }else{
  //boucle dans laquelle le serveur attend qu'un client demande une connection
      for {
        connection, _ := listener.Accept()
        go handleConn(connection)
      }
  }
  wg.Wait()
}

//fonction traitant la connection d'un client
func handleConn(connection net.Conn){

  defer wg.Done()

  //récupération des arguments passés par le client
  reader := bufio.NewReader(connection)
  iteration, er := reader.ReadString('\n')
  if er != nil {
    fmt.Println("Problème à la réception de la requête")
  } else {
    //conversion des arguments passés par le client en int
    nbIteration := ""
    for cpt := 0; cpt < utf8.RuneCountInString(iteration)-1; cpt++{
      nbIteration = nbIteration + string(iteration[cpt])
    }
    it, erreur := strconv.Atoi(nbIteration)
    if erreur != nil {
      fmt.Println("Erreur lors de la conversion de l'argument")
    }else{
      for i := 0; i < it ; i ++ {
        fmt.Println(i)
      }
    }
  }

}

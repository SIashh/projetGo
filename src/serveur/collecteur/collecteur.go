package collecteur

import (
  "strconv"
  "net"
  "fmt"
  "bufio"
  "unicode/utf8"

)
//fonction du collecteur
//param : fromCollector le canal dans lequel le collecteur doit placer les tâches des clients (ces tâches sont simplement des int)
func Collecteur(fromCollector chan int){
  listener, err := net.Listen("tcp",":9999") // port a ecouter
  if err != nil {
    fmt.Println("Problème d'écoute sur le port 9999")
  }else{
    for{ // boucle principale du serveur
      connexion, err := listener.Accept() // connexion
      if err != nil {
        fmt.Println("Problème pour accepter une connexion entrante avec un client")
      }else{
        //récupération des arguments passés par le client
        reader := bufio.NewReader(connexion)
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
            fromCollector <- it; // ajout de la requete au channel
          }
        }
      }
    }
  }
}

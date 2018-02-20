/*
Info utiles :
  os.Args
  strconv.Atoi // convertir des string en int
  flag : pouvoir directement déclarer des option pour le programme
*/

package main

import(
  "net"
  "bufio"
  "fmt"
  "sync"
  "sticonv"
)

var wg sync.WaitGroup
var tablePseudo = make([]string, 99)

func main() {
  wg.Add(3)
  listener, err := net.Listen("tcp",":10000")
  fromCollector := make(chan int, 5)
  go collecteur(fromCollector)

// //***************************************************************************************************
//
//   if err != nil {
//     fmt.Println(err)
//   }
//   for{
//     connexion, err := listener.Accept()
//     if err != nil {
//       fmt.Println(err)
//     }
//     go traitementConnection(connexion, messageAAfficher)
//   }
//   wg.Wait()
}

func collecteur(fromCollector chan int){
  listener, err := net.Listen("tcp",":10000") // port a ecouter
  if err != nil {
    fmt.Println(err)
  }else{
    for{ // boucle principale du serveur
      connexion, err := listener.Accept() // connexion
      if err != nil {
        fmt.Println(err)
      }else{
        writer := bufio.NewWriter(connexion)
        _, _ = writer.WriteString("Quelle tâche voulez-vous effectuer ?\n")
        writer.Flush();
        reader := bufio.NewReader(connexion)
        tache, _ := reader.ReadString('\n') // reception de la requete du client
        valeur, err := strconv.Atoi(tache);
        if err != nil {
          fmt.Println(err)
        }else{
          fromCollector <- valeur; // ajout de la requete au channel
        }
      }
    }
  }
}

// func traitementConnection(connexion net.Conn, messageAAfficher chan string){
//   writer := bufio.NewWriter(connexion)
//   _, _ = writer.WriteString("tki?\n")
//   writer.Flush()
//   reader := bufio.NewReader(connexion)
//   pseudo, _ := reader.ReadString('\n')
//   var nouveauPseudo bool = false
//   for !nouveauPseudo {
//     nouveauPseudo = true
//     for _, pseudos := range tablePseudo {
//       nouveauPseudo = pseudo != pseudos
//       if ! nouveauPseudo {
//         break
//       }
//     }
//     if ! nouveauPseudo {
//       _, _ = writer.WriteString("tkn!tki?\n")
//     }
//   }
//   _, _ = writer.WriteString("ok!\n")
//   fmt.Print("***** Un utilisateur vient de se connecter : " + pseudo)
//   tablePseudo = append(tablePseudo, pseudo)
//   writer.Flush()
//   var erreur error = nil
//   var message string
//   for erreur == nil {
//     message, erreur = reader.ReadString('\n')
//     if erreur == nil{
//       messageAAfficher <- ("----- " + pseudo + "> " + message)
//     }else{
//       fmt.Print("***** Un utilisateur vient de se déconnecter : " + pseudo)
//     }
//   }
// }
//
// func afficherMessage(messageAAfficher chan string){
//   compteur := 1
//   for{
//     select {
//     case msg := <-messageAAfficher:
//       fmt.Print(compteur, " : ", msg)
//       compteur++
//     }
//   }
// }

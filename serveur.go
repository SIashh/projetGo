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
var tablePseudo = make([]string, 99)


func main(){

  //le serveur accepte 3 clients
  wg.Add(3)
  fromCollector := make(chan int, 5)
  go collecteur(fromCollector)
  availableWorkers := make(chan chan int, 5)
  go repartiteur(availableWorkers, fromCollector)
  chanWorker1 := make(chan int, 5)
  nom :="travailleur 1"
  go travailleur(chanWorker1, availableWorkers, nom)
  nom ="travailleur 2"
  chanWorker2 := make(chan int, 5)
  go travailleur(chanWorker2, availableWorkers, nom)
  nom ="travailleur 3"
  chanWorker3 := make(chan int, 5)
  go travailleur(chanWorker3, availableWorkers, nom)
  wg.Wait()
}

func collecteur(fromCollector chan int){
  defer wg.Done()
  listener, err := net.Listen("tcp",":9999") // port a ecouter
  if err != nil {
    fmt.Println("Problème d'écoute sur le port 9999")
  }else{
    for{ // boucle principale du serveur
      connexion, err := listener.Accept() // connexion
      if err != nil {
        fmt.Println("Problème pour accepter une connexion entrante avec un client")
      }else{
        fmt.Println("Connecté")
        // writer := bufio.NewWriter(connexion)
        // _, _ = writer.WriteString("Quelle tâche voulez-vous effectuer ?\n")
        // writer.Flush();
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

func travailleur(workerChan chan int, availableWorkers chan chan int, nom string){
  defer wg.Done()
  availableWorkers <- workerChan //Lors du démarrage le travailleur est disponible
  for{
    select {
      case tache := <-workerChan: // On récupère la tache à effectuer dans le canal du travailleur
        for tache > 0 {
            fmt.Println(strconv.Itoa(tache)+ " " + nom)
            tache--
        }
        availableWorkers <- workerChan // Une fois la tache terminée, on ajoute le canal du travailleur au canal des canaux disponibles.
    }
  }

}

func repartiteur(availableWorkers chan chan int, fromCollector chan int){
  for{
    select {
    case work := <-fromCollector: // On récupère la première tache à effectuer
      worker := <-availableWorkers
      if worker != nil {
        worker <- work
      }
    }
  }
}

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
  "strconv"
)

var wg sync.WaitGroup
var tablePseudo = make([]string, 99)

func main() {
  wg.Add(5)
  listener, err := net.Listen("tcp",":10000")
  fromCollector := make(chan int, 5)
  go collecteur(fromCollector)
  availlableWorkers := make(chan chan int, 5)
  go repartiteur(availlableWorkers, fromCollector)
  worker1 := make(chan int, 5)
  go travailleur(worker1, availlableWorkers)
  worker2 := make(chan int, 5)
  go travailleur(worker1, availlableWorkers)
  worker3 := make(chan int, 5)
  go travailleur(worker1, availlableWorkers)
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
        writer.Flush()
        reader := bufio.NewReader(connexion)
        tache, _ := reader.ReadString('\n') // reception de la requete du client
        valeur, err := strconv.Atoi(tache)
        if err != nil {
          fmt.Println(err)
        }else{
          fromCollector <- valeur // ajout de la requete au channel
        }
      }
    }
  }
}

func repartiteur(availlableWorkers chan chan int, fromCollector chan int){
  for{
    select {
    case work := <-fromCollector: // On récupère la première tache à effectuer
      assignerTravail:
      for{
        select {
        case worker := <-availlableWorkers: // On récupère le premier canal disponible
          worker <- work
          break assignerTravail //on sort de la boucle pour passer à la tache suivante
        }
      }
    }
  }
}

func travailleur(workerChan chan int, availlableWorkers chan chan int){
  for{
    select {
      case boucle := <-workerChan: // On récupère la tache à effectuer dans le canal du travailleur
      for boucle > 0 {
        boucle--
      }
      availlableWorkers <- workerChan // Une fois la tache terminée, on ajoute le canal du travailleur au canal des canaux disponibles.
    }
  }
}

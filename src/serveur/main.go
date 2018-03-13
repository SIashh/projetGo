package main

import(
  "sync"
  "serveur/collecteur"
  "serveur/travailleur"
  "serveur/repartiteur"
)

var wg sync.WaitGroup

func main(){
  wg.Add(3)
  //initialisation du canal de collecte
  fromCollector := make(chan int, 5)
  go collecteur.Collecteur(fromCollector)
  //initialisation du canal comportant les travailleurs disponibles
  availableWorkers := make(chan chan int, 5)
  go repartiteur.Repartiteur(availableWorkers, fromCollector)
  //initialisation de canal d'un travailleur
  chanWorker1 := make(chan int, 1)
  nom :="travailleur 1"
  go travailleur.Travailleur(chanWorker1, availableWorkers, nom)
  nom ="travailleur 2"
  chanWorker2 := make(chan int, 1)
  go travailleur.Travailleur(chanWorker2, availableWorkers, nom)
  nom ="travailleur 3"
  chanWorker3 := make(chan int, 1)
  go travailleur.Travailleur(chanWorker3, availableWorkers, nom)
  wg.Wait()
}

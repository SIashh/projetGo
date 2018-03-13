package travailleur

import (
  "fmt"
  "strconv"
)
//fonction du travailleur
//param : workerChan le canal du travailleur pouvant accueillir une tâche donnée par le serveur
//param : availableWorkers le canal dans lequel le travailleur met son canal après réalisation d'une tâche
//param : nom le nom du travailleur nous ayant permis de tester si plusieurs travailleurs effectuaient des tâches "en même temps"
func Travailleur(workerChan chan int, availableWorkers chan chan int, nom string){
  availableWorkers <- workerChan //Lors du démarrage, le travailleur est disponible
  for{
    select {
      case tache := <-workerChan: // On récupère la tache à effectuer dans le canal du travailleur
        for tache > 0 {
            fmt.Println(strconv.Itoa(tache)+ " " + nom) //affichage nous permettant de tester
            tache--
        }
        availableWorkers <- workerChan // Une fois la tache terminée, on ajoute le canal du travailleur au canal des canaux disponibles.
    }
  }

}

package repartiteur

//fonction du repartiteur
//param : availableWorkers le canal qu'utilise le répartiteur pour sélectionner le travailleur à qui assigner la prochaine tâche
//param : fromCollector le canal comportant les tâches demandées par les clients
func Repartiteur(availableWorkers chan chan int, fromCollector chan int){
  for{
    select {
    case work := <-fromCollector: // On récupère la première tache à effectuer
      worker := <-availableWorkers //On récupère le premier travailleur disponible
      if worker != nil {
        worker <- work //On donne au travailleur le tâche à réaliser
      }
    }
  }
}

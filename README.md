- Installer JProfiler [download](https://www.ej-technologies.com/download/jprofiler/files)
- Installer Docker [installation](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04)
  - Les commandes docker devraient être lancées en mode sudo!
- Télécharger le répertoire git [download](https://github.com/Gabrielbdry/weka-3.8)
- Construire l'image docker à partir du répertoire du fichier **Dockerfile**:
```
sudo docker image build . -t jguwekarestv1:local
```
- Lancer le conteneur docker à partir du répertoire du fichier **docker-compose.yml**:
```
sudo docker-compose up
```
- Connectez-vous au conteneur Docker:
```
sudo docker exec -it jguwekarest bash
```
- Lancer l'agent JProfiler et entrez "**1**" comme premier argument et "**8849**"comme deuxième argument, comme suit:
```
tomcat@fa21547de714:/usr/local/tomcat$ /usr/local/jprofiler10.1.5/bin/jpenable
Connecting to org.apache.catalina.startup.Bootstrap start [1] ...
Please select the profiling mode:
GUI mode (attach with JProfiler GUI) [1, Enter]
Offline mode (use config file to set profiling settings) [2]
1
Please enter a profiling port
[33179]
8849
You can now use the JProfiler GUI to connect on port 8849
```
- Vous pouvez quitter la connexion (CTRL+D)
- Ouvrez l'outil JProfiler, en haut à gauche cliquez sur Session et dans la liste déroulante cliquez sur l'option New Session.
- Utilisez la configuration suivant et cliquez sur OK:
![img](https://i.imgur.com/QO7SfU0.png "Image")
- Sélectionnez l'option d'évaluation -> Sampling (Recommended) -> OK
- Vous pouvez maintenant envoyer des requêtes au serveur et son exécution devrait être tracée.
- Pour faire une sauvegarde, cliquer sur Save Snapshot.
- Un script python est fourni pour faire les requêtes, les valeurs **concurrent** et **request_per_second** contrôlent la fréquence des requêtes. Dans la requête, le fichier **contact-lenses.arff** est passé au rest API weka. Vous pouvez le modifier pour avoir des fichiers variables. La requête s'effectue à la ligne 16 du script.

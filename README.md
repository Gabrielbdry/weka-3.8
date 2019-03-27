## Manuel d'usage JProfiler

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

## Manuel d’usage de JMeter

- Télécharger les fichier binaires de JMeter [ici](http://apache.forsale.plus//jmeter/binaries/apache-jmeter-5.1.1.tgz).
- Lancer JMeter ./path-to-binaries/bin/jmeter
- Ajouter un ThreadGroup au plan de test (Clic droit sur le plan de test : add -> Threads -> ThreadGroup)
- Dans le thread group, vous pouvez configurer le nombre de thread lancés, ce qui est l’équivalent du nombre d’utilisateurs qu’on veut simuler, par exemple 10 pour la charge moyenne. Aussi cocher Loop Count: x Forever.
- Toujours à partir du plan de test, ajouter un Constant Throughput Timer (add -> Timer -> Constant Throughput Timer). Vous pouvez définir le nombre de requêtes envoyées par minute. Sélectionner l’option de calcul this thread only.
- Ajouter un HTTP Request à partir du Thread Group (add -> Sampler -> HTTP Request). Définir la méthode comme étant POST, le path comme étant /algorithm/BayesNet. Ajouter les éléments nécessaires dans le Body Data. Ajouter les fichiers arff dans Files Upload. Cocher Use multipart/form-data. Dans Advanced sélectionner l’implémentation client Java.
- Ajouter un Simple Data Writer à partir du Thread Group (add -> listener -> Simple Data Writer). Entrer un fichier dans filename, ce fichier sera réutilisé plus tard.
- Ajouter un Summary Report à partir du Thread Group (add -> listener -> Summary Report). Entrer le même nom de fichier qu’à l’étape précédente.
- Ajouter la ligne “jmeter.save.saveservice.print_field_names=true” dans le fichier ./path-to-binaries/bin/saveservice.properties
- Ajouter le contenu du fichier ./path-to-binaries/bin/reportgenerator.properties dans le fichier ./path-to-binaries/bin/user.properties
- Lancer le plan de test pour un certain temps (~1 minute).
- Dans un terminal, lancer ./path-to-binaries/bin/jmeter  -g ./path-to-report-file/report-file.csv
- Ouvrir le fichier index.html dans ./path-to-binaries/bin/report-output/
- Vous pouvez visualiser les données recueillies par le test.

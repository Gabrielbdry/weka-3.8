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
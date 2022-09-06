# Test Technique
## Suppositions
Pour ce test j'ai fait les suppositions suivantes:
- Il n'y a pas de mécanique de retry/fallback côté client, une fois le message reçu c'est à ce service de le gérer dans tous les cas.
- Les requêtes envoyées par le client ne doivent pas tomber en timeout. 
- Il n'est pas nécessaire de traiter le contenu des requêtes immédiatement.
- Le fallback pour les payloads invalides est hors scope.
- Les ids des ordres et produits sont uniques.


## Solution
### Modèles
Deux modèles identifiés, `Order` et `Product`, liés par une relation many to many.

### Traitement
- On expose un endpoint en POST `/web_api/tickets`
- Dans le handler associé, le payload de la requête est lu et envoyé dans une channel "buffurisée" de taille arbitrairement grande. (Ceci n'est pas optimal et provoquera un bloquage si le buffer est plein)
- Un dispatcher, lancé en parallèle du serveur, consomme les payloads envoyés dans la channel précédente. Ce dispatcher a à disposition 10 consumers qui vont traiter les messages en parallèle.
- Chaque message est parsé, validé et enregistré en db, si aucune erreur ne se produit.
- Les messages invalides sont envoyés vers le fallback.

### Améliorations
- Ne pas utiliser de channel mais une queue type RMQ ou SQS afin de ne pas encombrer la mémoire du service et d'avoir une plus grande capacité de stockage.
- En cas d'erreur de traitement une mécanique de fallback est nécessaire. On peut imaginer l'envoi du payload dans une queue dédiée pour un traitement spécifique en fonction de l'erreur levée.
- Ajouter des validations sur les données reçues.
- Ajouter du rate limiting et/ou de l'autoscaling.

### Fonctionnement
- Créer un fichier `.env` à la racine du projet contenant les variables d'environnement décrites dans le fichier `env.conf`
- Lancer le service

```
go run main.go
```
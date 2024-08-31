# go-sendtext

![SendText Logo](https://github.com/abdotop/go-sendtext/blob/main/sendtext-logo.png)

Le package `go-sendtext` est une bibliothèque en Go conçue pour faciliter l'intégration de l'API Sendtext dans des applications Web ou mobiles. Ce SDK permet aux développeurs de communiquer facilement avec l'API Sendtext pour envoyer des SMS, gérer des campagnes SMS, consulter le solde SMS et l'historique des envois.

## Fonctionnalités

- **Envoi de SMS** : Permet d'envoyer un SMS à un numéro spécifique.
- **Gestion de Campagnes SMS** : Supporte l'envoi de SMS en masse.
- **Consultation du Solde SMS** : Permet de vérifier le solde SMS disponible.
- **Historique des Envois** : Permet de consulter l'historique des SMS envoyés.

## Installation

Pour installer le package, vous pouvez utiliser `go get` :

```bash
go get -u github.com/abdotop/go-sendtext
```

## Utilisation

Voici un exemple simple d'utilisation du package pour envoyer un SMS :

```go
package main

import (
    "log"
    "github.com/abdotop/go-sendtext"
)

func main() {
    client := sendtext.NewClient("your_api_key", "your_api_secret")

    smsRequest := sendtext.SMSRequest{
        SenderName: "YourSenderName",
        SMSType: "normal",
        Phone: "221763983535",
        Text: "Votre message ici",
    }

    response, err := client.SendSMS(smsRequest)

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Response: %+v\n", response)
}
```

## Contribution

Les contributions à ce projet sont les bienvenues. Pour contribuer, veuillez forker le dépôt, créer une branche, effectuer vos modifications, et soumettre une pull request.

## Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.
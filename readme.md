VERSIEBEHEER:
Versie 1 - web, geen DB : Bij versie 1 is er een webserver met een html pagina die nog niks doet. De webserver draait op port 8080.

Versie 2 connectie met DB : Bij versie 1 is er een webserver met een html pagina die nog niks doet. De webserver draait op port 8080. Dit keer is er wel een connectie opgezet met de database, maar verder kan er niks gedaan worden op de website zelf.

Versie 3 - interactie met DB : Bij versie 3 is er een webserver waar je nu wel "apparaten mee kunt toekennen". Wat je invult op de website wordt bij een POST-request naar de database gestuurd.

Versie 3.1 - interactie met DB : Is hetzelfde als versie 3, maar met wat andere comments.

Versie 4 - toevoeging Platform en Serienummer : Bij versie 4 kun je nu ook bij de toekenningen een platform en Serienummer kiezen voor het apparaat. Dit wordt ook bij een POST-request naar de database gestuurd. 

Versie 5 - Environment var, HTML pad, Logging : Bij versie 5 is de wachtwoord van de database beschermd door een environment variabel. De pad van de HTML file is aangepast zodat het gedeployd kan worden in Azure. Ook wordt er in een error file de fouten bijgehouden met logging.
# Generátor náhodných dat o nosnicích

Tento program generuje soubor ve formátu `.CSV` s údaji o různých nosnicích. Program vytvoří datový soubor se sloupci: `ID`, `Rok`, `Plemeno`, `Hala`, `Snáška (vajec / rok)`, `Doba snášky (týdnů)` a `Hmotnost (kg)`. Data jsou generována náhodně na základě předem definovaných parametrů pro každý sloupec.

## Způsob použití
    
- 	Stáhněte spustitelný soubor `nosnice.exe` ze sekce "Release"
- 	Otevřete příkazový řádek
-	Přesuňte se do úmístění, kde je spustitelný soubor stažen
-	Spusťte soubor s následujícími parametry

```bash

nosnice.exe -n <pocet_zaznamu> -o <cilovy_soubor>


```

Nahraďte `<pocet_zaznamu>` za číslo, které představuje počet záznamů (řádků), které chcete vygenerovat a `<cilovy_soubor>` za název cílového souboru ve formátu CSV

Například:

```bash

nosnice.exe -n 100 -o data.csv


```

Vygeneruje 100 rádků se záznamy a uloží je do souboru `data.csv`
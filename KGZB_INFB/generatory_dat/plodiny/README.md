# Generátor náhodných dat o plodinách

Tento program generuje soubor ve formátu `.CSV` s údaji o různých plodínách. Program vytvoří datový soubor se sloupci: `Rok`, `Region`, `Plodina`, `Výnos (t/ha)` a `Plocha (ha)`. Data jsou generována náhodně na základě předem definovaných parametrů pro každý sloupec.

## Způsob použití
    
- 	Stáhněte spustitelný soubor `plodiny.exe` ze sekce "Release"
- 	Otevřete příkazový řádek
-	Přesuňte se do úmístění, kde je spustitelný soubor stažen
-	Spusťte soubor s následujícími parametry

```bash

plodiny.exe -n <pocet_zaznamu> -o <cilovy_soubor>


```

Nahraďte `<pocet_zaznamu>` za číslo, které představuje počet záznamů (řádků), které chcete vygenerovat a `<cilovy_soubor>` za název cílového souboru ve formátu CSV

Například:

```bash

plodiny.exe -n 100 -o data.csv


```

Vygeneruje 100 rádků se záznamy a uloží je do souboru `data.csv`
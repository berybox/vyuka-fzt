# Generátor náhodných dat o skotu

Tento program generuje soubor ve formátu `.CSV` s údaji o různých plemenech skotu. Program vytvoří datový soubor se sloupci: `Region`, `Mléčná užitkovost`, `Masná užitkovost`, `Velikost populace` a `Plemeno`. Data jsou generována náhodně na základě předem definovaných parametrů pro každý sloupec.

## Způsob použití
    
- 	Stáhněte spustitelný soubor `skot.exe` ze sekce "Release"
- 	Otevřete příkazový řádek
-	Přesuňte se do úmístění, kde je spustitelný soubor stažen
-	Spusťte soubor s následujícími parametry

```bash

skot.exe -n <pocet_zaznamu> -o <cilovy_soubor>


```

Nahraďte `<pocet_zaznamu>` za číslo, které představuje počet záznamů (řádků), které chcete vygenerovat a `<cilovy_soubor>` za název cílového souboru ve formátu CSV

Například:

```bash

skot.exe -n 100 -o data.csv


```

Vygeneruje 100 rádků se záznamy a uloží je do souboru `data.csv`
@REM Získat název rodičovského adresáře
@FOR %%a IN ("%~dp0\.") DO @SET "PARENT=%%~nxa"

@REM Získat aktuální datum a čas ve formátu yyyy-mm-dd_hh-mm
@FOR /F "usebackq tokens=1,2 delims==" %%i IN (`wmic os get LocalDateTime /VALUE 2^>NUL`) DO @IF '.%%i.'=='.LocalDateTime.' SET DATE=%%j
@SET DATE=%DATE:~0,4%-%DATE:~4,2%-%DATE:~6,2%_%DATE:~8,2%:%DATE:~10,2%

@REM Nastavit implicitní komentář
@SET COMMENT=%PARENT% [%DATE%]
@IF [%1] NEQ [] SET COMMENT=%*

@REM Příkazy pro Git
git add -A .
git commit -m "%COMMENT%"
git push -f origin --all

pause
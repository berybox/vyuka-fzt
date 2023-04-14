# Příklady na regulářní výrazy

1.  **Výraz**: `^[a-z]+$`  
**Popis**: Hledá řetězce, které obsahují pouze malá písmena od a do z.  
**Příklad dat**:  
    
    -   abcd
    -   efghij
    -   klmnopqrstuvwxyz
    -   12345 (nebude shodovat)
    -   ABCD (nebude shodovat)
    -   zyxwvutsrqponmlkjihgfedcba
2.  **Výraz**: `^[A-Z]+$`  
**Popis**: Hledá řetězce, které obsahují pouze velká písmena od A do Z.   
**Příklad dat**:  
    
    -   ABCDEFGHIJKLMNOPQRSTUVWXYZ
    -   UVWXYZ
    -   QWERTY (nebude shodovat)
    -   12345 (nebude shodovat)
    -   AbCdEfGhIjKlMnOpQrStUvWxYz (nebude shodovat)
3.  **Výraz**: `^[0-9]+$`   
**Popis**: Hledá řetězce, které obsahují pouze číslice od 0 do 9.  
**Příklad dat**:  
    
    -   1234567890
    -   9876543210
    -   456789
    -   abcdefg (nebude shodovat)
    -   12345a6789 (nebude shodovat)
4.  **Výraz**: `^[a-zA-Z0-9]+$`  
**Popis**: Hledá řetězce, které obsahují pouze alfanumerické znaky (malá a velká písmena a číslice).  
**Příklad dat**:  
    
    -   Abc123
    -   DeF456
    -   789Ghi
    -   &*()_+ (nebude shodovat)
    -   !@#$%^&*()_+ (nebude shodovat)
5.  **Výraz**: `^[a-z]{3}$`  
**Popis**: Hledá řetězce, které mají právě tři malá písmena.  
**Příklad dat**:
    
    -   abc
    -   def
    -   ghi
    -   ab (nebude shodovat)
    -   ABC (nebude shodovat)
6.  **Výraz**: `^[a-z]{3,5}$`  
**Popis**: Hledá řetězce, které mají mezi třemi a pěti malými písmeny.  
**Příklad dat**:
    
    -   abc
    -   defg
    -   hijkl
    -   ab (nebude shodovat)
    -   abcdef (nebude shodovat)
7.  **Výraz**: `^[A-Z][a-z]+$`  
**Popis**: Hledá řetězce, které začínají velkým písmenem a obsahují následně alespoň jedno malé písmeno.   **Příklad dat**:
    
    -   HelloWorld
    -   OpenAI
    -   ArtificialIntelligence
    -   helloWorld (nebude shodovat)

8. **Výraz**: `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$`  
**Popis**: Hledá řetězce, které představují platné emailové adresy.  
**Příklad dat**:
    
    -   john.doe@example.com
    -   jane.doe@example.co.uk
    -   info+test123@example.net
    -   user@example (nebude se shodovat)
    -   user@.com (nebude se shodovat)
9. **Výraz**: `^+?[0-9]{9,15}$`   
**Popis**: Hledá řetězce, které představují platná telefonní čísla (včetně mezinárodního formátu s prefixem +).  
**Příklad dat**:
    
    -   +420123456789
    -   123456789
    -   +12125551212
    -   +123 (nebude se shodovat)
    -   abcdefghij (nebude se shodovat)

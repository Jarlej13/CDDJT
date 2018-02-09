<h1>Opg 1</h1>  
<table>
    <tr>
        <th>Binære tall</th>  
        <th>Hexadesimale tall</th>  
        <th>Desimale tall</th>
    </tr>
    <tr>
        <td>1101</td>
        <td>0xD</td>
        <td>13</td>
    </tr>
    <tr>
        <td>110111101010</td>
        <td>0xDEA</td>
        <td>3562</td>
    </tr>
    <tr>
        <td>1010111100100100</td>
        <td>0xAF34</td>
        <td>44852</td>
    </tr>
    <tr>
        <td>01111111111111111</td>
        <td>0xFFFF</td>
        <td>65535</td>
    </tr>
    <tr>
        <td>00010001011110001010</td>
        <td>0x1178A</td>
        <td>71562</td>
    </tr>
</table>  
<h2>a)</h2>  
<p>Binærtall til hexadesimaltall: Vi deler opp rekken med 1-ere og 0-ere i grupper på fire. Dersom det ikke går opp legger vi til 0-ere foran. Deretter bruker vi tabellen for å finne ut hva det tilsvarer i hexadesimaltall.</p>  
<p>Hexadesimaltall til binærtall: Her sjekker vi bare opp mot tabellen og setter tallene vi får ut etter hverandre.</p>  
<p>Binærtall til desimaltall: Vi tilegner verdien 1 til det tallet lengst til høyre og dobler verdien for hvert tall til venstre. Deretter sjekker vi hvert tall fra høyre og dersom det er en 1-er legger vi til verdien. Summen vi sitter igjen med på slutten er desimaltallet som tilsvarer det originale binærtallet.<br />
   Eksempel: 1101 = 1[8] 1[4] 0[2] 1[1]= 8 + 4 + 1 = 13</p>  
<p>Desimaltall til binærtall: Vi skriver først ned desimaltallet. Deretter deler vi det på 2 og skriver ned resultatet (rundet ned). Dersom resultatet får rest noterer vi 1. Om ikke, noterer vi 0. Vi fortsetter slik til resultatet blir 0. Vi skriver ned 1-erne og 0-ene ved siden av hverandre (siste først) og da har vi binærtallet.<br />
   Eksempel: 13 = 13[1] --> 6[0] --> 3[1] --> 1[1] --> 0 = 1101 </p>  
<h2>b)</h2>
<p>Hexadesimaltall til desimaltall: Her er matematikken litt mer avansert. Siden hexadesimaltall har base 16 må vi multiplisere hver verdi i hexadesimaltallet med 16 opphøyd i plasseringen minus 1. Altså i hexadesimaltallet 4A må A som står på plass 1 multipliseres med 16 opphøyd i 1-1=0 og 4 som står på plass 2 multipliseres med 16 opphøyd i 2-1=1. Deretter legger vi sammen alt så har vi desimaltallet.<br />
   Eksempel: 4A = 4 * 16^1 + A(10) * 16^0 = 74</p>
<p>Desimaltall til hexadesimaltall: Vi deler desimaltallet på 16 og noterer resten. Fortsett på samme måte med resultatet til vi bare sitter igjen med rest. Bruk tabellen til å finne hexadesimaltallet v.h.a. restene (siste først).<br />
   Eksempel: 74 = 74 / 16 --> 4 (10 i rest). 4 / 16 --> 0 (4 i rest) = 4A</p>  
<h1>Opg 2</h1>
<h2>c)</h2>   
<h3>Big-O</h3  
<p>Kort sagt er Big-O i hvilken grad tiden det tar å utføre operasjonen vokser i forhold til størrelsen på lista. Big-O for bubblesort og quicksort er vanligvis O(N^2). Det vil si at tiden det tar å utføre øker proporsjonalt med størrelsen på lista. Dette observerer vi også for de 3 algoritmene.<br />
   Eksempel: Vi tar utgangspunkt i den modifiserte bubblesort funksjonen. Etter å ha kjørt testen observerer vi at når størrelsen på lista øker fra 100 elementer til 1000 elementer, altså 10 ganger, øker ns/op 55 ganger. Deretter observerer vi at når lista øker videre fra 1000 elementer til 10000 elementer, altså 10 ganger til, øker ns/op 138 ganger. Under finner du alle 3 algoritmene illustrert grafisk og sammenlignet med O(N^2).</p>  
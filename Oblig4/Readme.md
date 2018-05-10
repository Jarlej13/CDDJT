<h1>SYSTEMSPESIFIKASJON</h1>

<p>Vi har laget ett et system for live værvarsling. Dette kan gi værinformasjon for alle byer i hele verden. Du vil få en beskrivelse av sted og oppgitt temperatur, nedbør, beskrivelse av været, vindhastighet, vindretning og opplevd temperatur basert på vind. I tillegg til værmeldingen vil brukeren få en beskrivelse av hvordan man burde kle deg. Når programmet kjøres vil du først få opp en nettside hvor du skriver inn hvor du ønsker værinformasjon fra. Applikasjonen vil så hente værdata direkte fra www.apixu.com. Vi bruker Json for å få ut dataen vi trenger og bruker html og css for å få et fornuftig og oversiktlig format i nettleseren. Du vil i tillegg for å kart over den aktuelle byen hvor du også har mulighet for å se rundt. Kartet hentet fra googlemaps. </p>


<h1>SYSTEMARKITEKTUR</h1> 

<p>Nodene består av klient, databaseserver og webserver. Applikasjonen åpner webserver på port :8080, får input fra bruker og presenterer dataen på en annen side. Klienten sender også en forespørsel til Apixu.com sine servere for å hente værdata.</p>

<img src="https://user-images.githubusercontent.com/35611995/39879725-e9eec64c-547b-11e8-934b-6f9251a2f0e0.png" width="90%"></img> 

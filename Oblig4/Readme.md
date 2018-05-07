<h1>SYSTEMSPESIFIKASJON</h1>

<p>Vi har laget et system for live værvarsling. Dette kan gi værinformasjon for alle byer i hele verden, og du vil i tillegg til værmeldingen få en beskrivelse av hvordan du burde kle deg. Når programmet kjøres vil du først få opp en nettside hvor du skriver inn hvor du ønsker værinformasjon fra. Applikasjonen vil så hente værdata direkte fra www.apixu.com. Vi bruker Json for å få ut dataen vi trenger og bruker html og css for å få et fornuftig og oversiktlig format i nettleseren.</p>


<h1>SYSTEMARKITEKTUR</h1> 

<p>Nodene består av klient, databaseserver og webserver. Applikasjonen åpner webserver på port :8080, får input fra bruker og presenterer dataen på en annen side. Klienten sender også en forespørsel til Apixu.com sine servere for å hente værdata.</p>

<img src="https://user-images.githubusercontent.com/35611995/39714350-12ee54da-522a-11e8-9179-2080cbd305f2.png" width="90%"></img> 

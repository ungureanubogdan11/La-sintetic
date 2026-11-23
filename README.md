# La Sintetic

Aplicație web propusă pentru organizarea meciurilor de fotbal pe teren sintetic. Organizatorul publică un meci pe un teren pe care l-a rezervat deja. Jucătorii văd unde și când se joacă, ce posturi sunt libere și cât îi revine fiecăruia din costul terenului.

## Funcționalități 

- **Organizarea meciurilor.** Un membru va putea crea, modifica și șterge propriile anunțuri. Pentru fiecare meci va preciza terenul, orașul, data, ora de început și de sfârșit, formatul de joc și tariful total.
- **Căutare și filtrare.** Pagina principală va afișa următoarele meciuri cu locuri libere. Catalogul va permite căutare după teren, filtrare după oraș, dată și post disponibil, sortare după ora de început sau cost estimat și afișare paginată.
- **Înscrieri pe posturi.** Jucătorul va alege portar sau jucător de câmp. La 5 contra 5 vor fi disponibile două locuri de portar și opt de câmp. Un jucător va putea avea o singură înscriere activă la același meci.
- **Listă de așteptare.** După ocuparea unui post, înscrierile pentru acel post vor intra în așteptare, în ordine cronologică. La retragerea unui participant înainte de ora de început, primul jucător în așteptare pe postul eliberat va fi confirmat automat.
- **Cost pe jucător.** Estimarea va fi tariful total împărțit la numărul de jucători confirmați și se va actualiza după fiecare schimbare. Până la prima confirmare se va afișa doar tariful total.
- **Profiluri publice sau private.** Membrii își vor putea edita numele, e-mailul, fotografia, descrierea și postul preferat. Proprietarul va avea acces la propriul profil. Pentru un profil privat, descrierea și postul preferat vor fi vizibile și administratorului și colegilor confirmați la un meci comun. E-mailul va fi folosit pentru autentificare.

## Roluri

| Rol | Drepturi |
| --- | --- |
| Vizitator | Consultă anunțurile publice și poate crea un cont. |
| Membru | Își gestionează profilul, propriile meciuri și propriile înscrieri. |
| Administrator | Are drepturi complete asupra resurselor și gestionează conturile și rolurile. |

Organizatorul este membrul care a creat meciul. Drepturile sale de organizare se aplică acelui meci.

## Implementare 

- **Go și Gin** pentru API. Handlerele vor procesa cererile HTTP, serviciile vor aplica regulile de înscriere, iar repository-urile vor gestiona accesul la date.
- **PostgreSQL și GORM** pentru persistență, cu migrații versionate prin golang-migrate.
- **JWT și bcrypt** pentru autentificare și stocarea parolelor. Middleware-ul va verifica tokenul și rolul, iar serviciile vor verifica proprietarul resursei.
- **Validare în interfață și pe server** pentru câmpuri obligatorii, tarif și interval orar. Serverul va verifica disponibilitatea în momentul confirmării pentru a preveni ocuparea simultană a ultimului loc.
- **Răspunsuri JSON standardizate** și coduri HTTP pentru succes, date invalide, autentificare necesară, acces interzis, resurse inexistente și erori interne.
- **Date demonstrative** cu meciuri în mai multe orașe, profiluri publice și private și înscrieri confirmate sau în așteptare.

## Exemplu pentru demonstrație

La un meci de 5 contra 5, ambele locuri de portar sunt ocupate și mai există un loc de câmp. Un al treilea portar intră în așteptare, în timp ce un jucător de câmp se poate înscrie imediat. Dacă unul dintre portarii confirmați se retrage, aplicația promovează automat portarul din așteptare. Pentru un teren de 300 de lei și zece participanți confirmați, contribuția estimată este de 30 de lei de persoană.

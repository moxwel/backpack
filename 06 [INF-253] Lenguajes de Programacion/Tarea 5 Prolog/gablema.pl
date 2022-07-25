% Resto.
% 
% resto(Elem, Lista, Resto).
% Revisa si: a partir de la primera aparicion de 'Elem' en la lista 'Lista',
% se obtiene el 'Resto' de la lista.

% Caso base:
% En una lista '[First|Rest]': si se toma el primer elemento,
% el resto de la lista sera efectivamente la lista 'Rest'.
resto(Elem, [Elem|Resto], Resto).

% Paso recursivo:
% En una lista '[First|Rest]': si se toma un elemento cualquiera (que no esta
% al inicio de la lista), hay que revisar el resto en la lista 'Rest'.
resto(Elem, [_|Cola], Resto) :- resto(Elem, Cola, Resto).

% Ejemplos:
% 
% Obtener el resto de la lista.
% resto(a, [b,c,a,d,e,f], Resto).
%	Resto = [d,e,f]
%
% Obtener las posibles combinaciones de restos.
% resto(Elem, [a,b,c,d], Resto).
% 	Elem = a,
% 	Resto = [b, c, d];
%	Elem = b,
%	Resto = [c, d];
%	Elem = c,
%	Resto = [d];
%	Elem = d,
%	Resto = []

% ====================

% Seleccionar.
% 
% selecc(Num, Lista, Result).
% Revisa si: al tomar un numero 'Num' de elementos de la lista 'Lista', se obtiene
% alguna combinacion de elementos de la lista en 'Result'

% Caso base:
% Seleccionar 0 elementos de cualquier lista, siempre resulta en una lista vacia.
% Corte de backtracking para evitar permutaciones.
selecc(0, _, []) :- !.

% Paso recursivo:
% Mientras 'Num' sea positivo, se van viendo los posibles restos de la lista 'Lista',
% y a medida que se va haciendo, ir disminuyendo 'Num'.
selecc(Num, Lista, [First|Rest]) :-
    Num > 0,
    resto(First, Lista, Result),
    Num1 is Num - 1,
    selecc(Num1, Result, Rest).

% Ejemplos:
% 
% Obtener las combinaciones.
% selecc(3, [a,b,c], Comb).
% 	Comb = [a, b, c]
% 	
% selecc(2, [a,b,c,d], Comb).
% 	Comb = [a, b];
%	Comb = [a, c];
%	Comb = [a, d];
%	Comb = [b, c];
%	Comb = [b, d];
%	Comb = [c, d]

% ====================

% Equipos.
% 
% equipos(Lista, Equipos, Result).
% Revisa si: de una lista 'Lista' de miembros, se arman equipos segun la lista 'Equipos'
% en 'Result'.

% Caso base:
% Equpos vacios siguen siendo vacios.
equipos([],[],[]).

% Paso recursivo:
% Hay que ir viendo la lista de numeros de 'Equipos', e ir seleccionando grupos de ese tamano,
% para luego restarlo con la lista de miembros, y volver a revisar otros equipos.
% 
% *subtract(Lista1, Lista2, Result)
% Revisa si: de la lista 'Lista1', se eliminan todos los elementos que hay en 'Lista2', y
% resulta en 'Result'.
equipos(Miemb,[Num1|NumN],[Equipo1|EquipoN]) :- 
   selecc(Num1,Miemb,Equipo1),
   subtract(Miemb,Equipo1,Result),
   equipos(Result,NumN,EquipoN).

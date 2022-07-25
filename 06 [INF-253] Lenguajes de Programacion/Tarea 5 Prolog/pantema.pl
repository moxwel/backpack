arco(a, b).
arco(a, f).
arco(f, g).
arco(f, i).
arco(h, f).
arco(d, h).
arco(d, f).

% Regla para considerar enlaces en grafos no dirigidos.
enlace(X, Y) :-
    arco(X, Y);
    arco(Y, X).

% Camino:
%
% Revisa si hay un camino entre el nodo 'X' y el nodo 'Y'

% Caso base:
% Hay un camino si es que hay un enlace entre esos dos nodos.
camino(X, Y) :- enlace(X, Y).

% Paso recursivo:
% Hay un camino si es que hay caminos entre los hijos de un nodo.
camino(X, Y) :- enlace(X, Z), camino(Z, Y).
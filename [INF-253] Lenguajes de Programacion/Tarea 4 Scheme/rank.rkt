#lang scheme

;; (valor nodo)
;; Retorna el valor del nodo dado
;; Retorno: un numero (o una lista vacia)
(define valor
  (lambda (nodo)
    (if (null? nodo)
      ; Si la lista esta vacia...
      '()

      ; Si no...
      (first nodo)
      ; (   >[ nodo ]<   hijo-izq   hijo-der  )
     )
   )
 )



;; (der nodo)
;; Retorna el hijo derecho del nodo dado
;; Retorno: una lista
(define der
  (lambda (nodo)
    (if (null? nodo)
      ; Si la lista esta vacia...
      '()

      ; Si no...
      (first (rest (rest nodo)))
      ; (  nodo [  hijo-izq  >[ hijo-der ]<  ]  )
     )
   )
 )



;; (izq nodo)
;; Retorna el hijo izquierdo del nodo dado
;; Retorno: una lista
(define izq
  (lambda (nodo)
    (if (null? nodo)
      ; Si la lista esta vacia...
      '()
       
      ; Si no...
      (first (rest nodo))
      ; (  nodo [  >[ hijo-izq ]<  hijo-der  ]  )
     )
   )
 )



;; (in-orden arbol)
;; Recorre el arbol en in-orden
;; Retorno: una lista
(define in-orden
  (lambda (arbol)
    (if (null? arbol)
      ; Caso base: si el arbol esta vacio, terminar
      '()

      ; Si aun nos queda por recorrer...
      (append
        ; Hay que concatenar las listas que se nos generaran.
        (in-orden (izq arbol)) ; Recorrer el subarbol izquierdo en in-orden
        (list (valor arbol))   ; Visitar nodo. Obtener valor
        (in-orden (der arbol)) ; Recorrer el subarbol derecho en in-orden
       )
     )
   )
 )



; Funcion principal
(define rank
  (lambda (numero arbol)
    (let recursion
      ; Parametros
      ((l (in-orden arbol)) ; l = Lista generada por 'in-orden' del arbol (esta en orden creciente)
       (c 0)                ; c = Contador
       )

      ; Funcion
      (if (or (null? l) (>= (first l) numero))
        ; Caso base: Si el primer elemento es mayor o igual que el numero a analizar,
        ; o la lista esta vacia
        ; ¿Que retornar?
        ; El contador
        c

        ; Si es menor, agregar contador y contar el resto de la lista
        (recursion
          (rest l) ; l
          (+ c 1)  ; c
         )
       )
     )
   )
 )



; Arboles de prueba
(define mi_arbol1 '(5 (3 (2 () ()) (4 () ())) (8 (6 () ()) ())) )
(define mi_arbol2 '(28 (16 (5 () (12 () ())) (19 (17 () ()) (21 () ()))) (30 (29 () ()) (32 () ()))) )

; Pruebas
(in-orden mi_arbol1)
(rank 7 mi_arbol1)
(in-orden mi_arbol2)
(rank 23 mi_arbol2)
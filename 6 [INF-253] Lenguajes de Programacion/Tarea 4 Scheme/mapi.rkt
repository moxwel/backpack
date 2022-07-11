#lang scheme

;; (sumar-lista numero lista_de_numeros)
;; Suma un numero a todos los numeros de una lista
;; Retorno: una lista
(define sumar-lista
  (lambda (n l)
    (if (null? l)
      ; Caso base: Si la lista esta vacia, terminar.
      '()

      ; Si no esta vacia, ir explorando la lista.
      (append
        (list (+ (first l) n))
        (sumar-lista n (rest l))
       )
     )
   )
 )



;; (resta-lista numero lista_de_numeros)
;; Resta un numero a todos los numeros de una lista
;; Retorno: una lista
(define resta-lista
  (lambda (n l)
    (if (null? l)
      ; Caso base: Si la lista esta vacia, terminar.
      '()

      ; Si no esta vacia, ir explorando la lista.
      (append
        (list (- (first l) n))
        (resta-lista n (rest l))
       )
     )
   )
 )



;; (multi-lista numero lista_de_numeros)
;; Multiplica todos los numeros de una lista por un numero
;; Retorno: una lista
(define multi-lista
  (lambda (n l)
    (if (null? l)
      ; Caso base: Si la lista esta vacia, terminar.
      '()

      ; Si no esta vacia, ir explorando la lista.
      (append
        (list (* (first l) n))
        (multi-lista n (rest l))
       )
     )
   )
 )



;; (div-lista numero lista_de_numeros)
;; Divide todos los numeros de una lista por un numero
;; Retorno: una lista
(define div-lista
  (lambda (n l)
    (if (not (= n 0))
      ; Hay que asegurarse que el numero no sea 0
      ; No se permite division por 0.
      (if (null? l)
        ; Caso base: Si la lista esta vacia, terminar.
        '()

        ; Si no esta vacia, ir explorando la lista.
        (append
          (list (/ (first l) n))
          (div-lista n (rest l))
         )
       )

      ; Si se divide por 0, no hacer nada.
      '()
     )
   )
 )



; Funcion principal
(define mapi
  (lambda (oper num lista)
    (cond
      ; Si el operador es de suma
      ((eqv? oper 'S)
        (sumar-lista num lista)
       )
      
      ; Si el operador es de resta
      ((eqv? oper 'R)
        (resta-lista num lista)
       )

      ; Si el operador es de multiplicacion
      ((eqv? oper 'M)
        (multi-lista num lista)
       )

      ; Si el operador es de division
      ((eqv? oper 'D)
        (div-lista num lista)
       )
     )
   )
 )



; Pruebas
(mapi 'S 1 '(1 2 3 4 5) )
(mapi 'M 2 '(1 2 3 4 5) )
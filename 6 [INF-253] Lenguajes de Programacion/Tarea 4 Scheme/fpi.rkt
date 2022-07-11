#lang scheme

; Funcion principal
(define fpi
  (lambda (funcion umbral i)
    (let recursion
      ; Parametros
      ((x0 (funcion i)) ; x0 = Funcion evaluada  |  f(x)
       (x1 i)           ; x1 = Numero            |   x
       (c 0)            ; c = Contador
       )

      ; Funcion
      (if (<= (abs (- x0 x1)) umbral)
        ; Caso base: Cuando se alcance el umbral
        ; ¿Que retornar?
        ; El contador de iteraciones.
        c

        ; Si no, hay que seguir...
        (recursion
          (funcion x0) ; x0  |  f(f(x))
          x0           ; x1  |   f(x)
          (+ c 1)      ; c
         )
       )
     )
   )
 )



; Pruebas
(fpi (lambda (x) (cos x)) 0.02 1)
#lang scheme

;; (dec-hex numero)
;; Transforma un numero decimal en [0, 15] a hexadecimal
;; Retorno: un string
(define dec-hex
  (lambda (num)
    (if (and (>= num 0) (< num 10))
      ; Si numero esta en [0, 9], simplemente hay que retornarlo (como string)
      (number->string num)

      ; Si es mayor o igual que 10, hay que transformarlo a hex
      (cond
        ; 10=A, 11=B, 12=C, 13=D, 14=E, 15=F
        ((= num 10)
         "A"
         )
        ((= num 11)
         "B"
         )
        ((= num 12)
         "C"
         )
        ((= num 13)
         "D"
         )
        ((= num 14)
         "E"
         )
        ((= num 15)
         "F"
         )
       )
     )
   )
 )



; Funcion principal
(define hex
  (lambda (num)
    (let recursion
      ; Parametros
      ((n num)               ; n = Numero actual
       (q (quotient num 16)) ; q = Cuociente actual
       (s "")                ; s = String que contene el hex
       )

      ; Funcion
      (if (= q 0)
        ; Caso base: Si el cuociente es menor que 16
        ; ¿Que hay que retornar?
        ; El string completado.
        (string-append (dec-hex (modulo n 16)) s)

        ; Si no es 0, entonces falta por ver...
        (recursion
          q                                         ; n
          (quotient q 16)                           ; q
          (string-append (dec-hex (modulo n 16)) s) ; s
         )
       )
     )
   )
 )



; Pruebas
(hex 160)
(hex 462)
(hex 14794430)
(hex 81985529216486895)
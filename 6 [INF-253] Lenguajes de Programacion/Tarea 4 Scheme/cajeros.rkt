#lang scheme

;; [TOMADO DE MAPI.RKT]
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



;; (min-lista lista_de_numeros)
;; Obtiene el menor numero de una lista dada
;; Retorno: un numero
(define min-lista
  (lambda (lista)
    (first (sort lista <))
   )
 )



;; Variable donde guardar la ultima caja donde un cliente ingreso
(define ultima_caja 0)



;; (avanzar-cola lista_cajas lista_cola)
;; Dado una lista con el estado de las cajas y una lista de la cola de clientes,
;; devuelve el estado despues de que un cliente de las cajas se desocupe e
;; ingrese el proximo cliente de la cola.
;; Retorno: una lista.
(define avanzar-cola
  (lambda (cajas cola)
    ; Suponiendo que se comienza con 3 cajas: (406 424 87)
    ; Y la cola esta asi: (888 871 915 516 ...)

    (list-set
      (resta-lista (min-lista cajas) cajas) ; Estado de las cajas cuando una se desocupe (319 337 0)

      (begin
        ; Primero hay que guardar el numero de la caja en donde el nuevo cliente va a ir (3)
        (set! ultima_caja (+ (index-where (resta-lista (min-lista cajas) cajas) zero?) 1))

        ; Retornar ese numero (3) para que 'list-set' pueda usarlo
        (index-where (resta-lista (min-lista cajas) cajas) zero?)
       )
      
      (first cola) ; Alli (en la posicion 3), poner el proximo cliente de la cola (319 337 888)
     )
   )
 )



; Funcion principal
(define caja
  (lambda (cajeros lista)
    (let recursion
      ; Parametros
      ((cajas (make-list cajeros 0)) ; cajas  = Lita de cajas inicial (0 0 0 ...)
       (cola lista)                  ; cola   = Lista de la cola actual
       (avance '())                  ; avance = Lista de los clientes usando las cajas 
       )

      ; Procedimiento
      (if (null? cola)
        ; Caso base: la cola esta vacia.
        ; Que hay que retornar?
        ; La lista de los clientes usando las cajas
        avance

        ; Si no, hay que seguir añadiendo clientes a las cajas
        (recursion
         (avanzar-cola cajas cola)          ; cajas con la cola de clientes avanzada
         (rest cola)                        ; cola
         (append avance (list ultima_caja)) ; avance
         )
       )
     )
   )
 )



; Pruebas
(caja 3 '(406 424 87 888 871 915 516 81 275 578))
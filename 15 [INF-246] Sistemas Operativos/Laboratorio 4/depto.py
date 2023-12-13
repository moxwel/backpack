from threading import Lock, Condition, Thread, Timer
from time import sleep
from datetime import datetime
import random
import os

class Departamento:
    def __init__(self, nombre: str, capacidad_fila: int, duracion_consulta: int, capacidad_depto: int, archivo: str):
        self.nombre: str = nombre
        self.fila_depto: list[Persona] = []
        self.fila_patio: list[Persona] = []

        self.capacidad_fila: int = capacidad_fila
        self.duracion_consulta: int = duracion_consulta
        self.capacidad_depto: int = capacidad_depto

        self.enUso: bool = False

        self.lock_fila_depto: Lock = Lock()  # Lock para la fila del departamento
        self.lock_fila_patio: Lock = Lock()  # Lock para el patio de las lámparas
        self.lock_depto: Lock = Lock()  # Lock para el departamento

        self.cv_espacioFilaDepto: Condition = Condition(self.lock_fila_depto)  # La fila del departamento tiene espacio
        self.cv_capacidadDepto: Condition = Condition(self.lock_depto)  # El departamento tiene capacidad para atender

        self.archivo: str = archivo

        # El temporizador se usa para saber cuando no sucede nada en el departamento
        # Cada vez que se atiende a una persona, se reinicia el temporizador
        self.tiempo_espera: int = 10
        self.temporizador: Timer = Timer(self.tiempo_espera, self.noSucedeNada)
        self.temporizador.start()

    def hayEspacioEnFilaDpto(self) -> bool:
        return len(self.fila_depto) < self.capacidad_fila

    def hayPersonasSuficientes(self) -> bool:
        return len(self.fila_depto) >= self.capacidad_depto

    def hayPersonasEnFilaPatio(self) -> bool:
        return len(self.fila_patio) > 0

    def incapazDeLlenarDepto(self) -> bool:
        return (len(self.fila_depto) + len(self.fila_patio)) < self.capacidad_depto

    def estaOcupado(self) -> bool:
        return self.enUso


    # Funcion que se ejecuta cuando no sucede nada en el departamento
    # Se ejecuta cada 'tiempo_espera' segundos.
    # Si no sucede nada en el departamento, entonces se atiende a las personas que estan en la fila del depto
    def noSucedeNada(self) -> None:
        print(f"\n\t\t---Departamento {self.nombre} NO SUCEDIO NADA DURANTE BASTANTE TIEMPO. ({personas_restantes})\n")

        if len(self.fila_depto) > 0:
            self.atender()

            with self.lock_depto:
                # Notifica a las hebras que estan esperando (por que no habia suficientes personas en la fila del depto)
                # para que despierten y verifiquen si pueden entrar al depto.
                self.cv_capacidadDepto.notify_all()
        else:
            if personas_restantes == 0:
                print(f"\n\n\t---Derpartamento {self.nombre} NO HAY mas personas EN LA UNIVERSIDAD. Terminando...\n\n")
            else:
                print(f"\n\t\t---Departamento {self.nombre} NO HAY NADIE EN LA FILA, Esperando... ({personas_restantes})\n")
                self.temporizador = Timer(self.tiempo_espera, self.noSucedeNada)
                self.temporizador.start()


    # Toma a las personas de la fila del depto y las atiende
    def atender(self):
        self.enUso = True

        # Hay que atender a tantas personas hasta llenar la capacidad del depto
        # o hasta que no haya mas personas en la fila del depto
        for i in (range(min(self.capacidad_depto, len(self.fila_depto)))):
            persona = self.fila_depto.pop(0)
            persona.atendido = True
            persona.hora_entrada_depto = datetime.now().strftime("%H:%M:%S.%f")
            print(f"\t\t=0=Persona {persona.id} ENTRO al DEPARTAMENTO en {self.nombre}. ({personas_restantes})")

        with self.lock_fila_depto:
            # Notifica a las hebras que estan esperando en la fila del patio del depto (por que no habia espacio en la fila del depto)
            # para que despierten y verifiquen si pueden entrar a la fila del depto.
            self.cv_espacioFilaDepto.notify_all()

        self.temporizador.cancel()

        print(f"\n\t\t---Departamento {self.nombre} ATENDIENDO... ({personas_restantes})\n")
        sleep(self.duracion_consulta)
        self.enUso = False
        print(f"\n\t\t@@@Departamento {self.nombre} TERMINA de ATENDER. ({personas_restantes})\n")

        # Reiniciar el temporizador
        self.temporizador = Timer(self.tiempo_espera, self.noSucedeNada)
        self.temporizador.start()

    # Registra la actividad de la persona en el archivo del depto
    def registrarActividad(self, persona_id: int, hora_entrada_fila: int, hora_entrada_depto: int, depto_n: int):
        with open(self.archivo, "a") as archivo:
            archivo.write(f"Persona{persona_id}, {hora_entrada_fila}, {hora_entrada_depto}, {depto_n}\n")


class Persona:
    def __init__(self, id: int, lista_deptos: list[Departamento]):
        self.id: int = id
        self.atendido: bool = False # Si ha sido atendido en el depto 0 y 1
        self.enPatio: bool = False  # Si está en el patio de las lámparas
        self.deptos: list[Departamento] = random.sample(lista_deptos, 2) # Deptos a los que quiere ir. Es una lista de 2 elementos
        self.hora_entrada_patio: str = ""
        self.hora_entrada_fila: str = ""
        self.hora_entrada_depto: str = ""

    def esPrimeroEnPatio(self, depto_n: int) -> bool:
        # Verifica si la persona es la primera en la fila del patio del depto.
        # 'depto_n' = 0 para el primero, 1 para el segundo

        return self.deptos[depto_n].fila_patio[0] == self

    def esPrimeroEnDepto(self, depto_n: int) -> bool:
        # Verifica si la persona es la primera en la fila del depto.
        # 'depto_n' = 0 para el primero, 1 para el segundo

        return self.deptos[depto_n].fila_depto[0] == self

    def entrarAPatio(self, depto_n: int):
        # La persona primero llega al patio de las lamparas y se pone en la fila del patio del depto.
        # 'depto_n' = 0 para el primero, 1 para el segundo

        with self.deptos[depto_n].lock_fila_patio:

            # Entrando a la fila del patio del depto
            self.deptos[depto_n].fila_patio.append(self)
            self.hora_entrada_patio = datetime.now().strftime("%H:%M:%S.%f")
            self.enPatio = True
            print(f"Persona {self.id} LLEGA al PATIO de las lámparas para {self.deptos[depto_n].nombre} ({personas_restantes})")

    def avanzarAFila(self, depto_n: int):
        # Ahora la persona en la fila del patio del depto debe intentar entrar a la fila del depto.
        # para eso, primero debe verificar que haya espacio en la fila del depto, y ademas
        # debe ser el primero en la fila del patio del depto.
        # 'depto_n' = 0 para el primero, 1 para el segundo

        with self.deptos[depto_n].lock_fila_depto:

            # Espera a que haya espacio en la fila del depto y que sea el primero en la fila del patio del depto
            # antes de intentar entrar a la fila del depto
            while not (self.deptos[depto_n].hayEspacioEnFilaDpto() and self.esPrimeroEnPatio(depto_n)):
                # print(f"Persona {self.id} ESPERA en el PATIO de las lámparas para {self.deptos[depto_n].nombre}... ({personas_restantes})")
                self.deptos[depto_n].cv_espacioFilaDepto.wait()
                # print(f"\n!!!Persona {self.id} DESPIERTA en el PATIO de las lámparas para {self.deptos[depto_n].nombre}. ({personas_restantes})\n")

            # Pasa a la fila del depto
            self.deptos[depto_n].fila_patio.pop(0)
            self.deptos[depto_n].fila_depto.append(self)
            self.hora_entrada_fila = datetime.now().strftime("%H:%M:%S.%f")
            self.enPatio = False

            print(f"Persona {self.id} ENTRA a FILA de {self.deptos[depto_n].nombre}. ({personas_restantes})")

            # Notifica a las hebras que estan esperando (por que no habia espacio en la fila del depto) para
            # que despierten y verifiquen si pueden entrar a la fila del depto.
            self.deptos[depto_n].cv_espacioFilaDepto.notify_all()

        with self.deptos[depto_n].lock_depto:
            # Notifica a las hebras que estan esperando (por que no habia suficientes personas en la fila del depto)
            # para que despierten y verifiquen si pueden entrar al depto.
            self.deptos[depto_n].cv_capacidadDepto.notify_all()

    def entrarADepto(self, depto_n):
        # La persona ahora debe intentar entrar al depto. Para eso, primero debe verificar que hay suficientes
        # personas en la fila del depto, y ademas debe ser el primero en la fila del depto.
        # 'depto_n' = 0 para el primero, 1 para el segundo

        with self.deptos[depto_n].lock_depto:

            # Para entrar al depto, la persona debe esperar a que haya suficientes personas en la fila del depto,
            # que sea el primero en la fila del depto, y que el depto no este ocupado.
            # O bien, que ya haya sido atendido en el depto.
            while not ((self.deptos[depto_n].hayPersonasSuficientes() and self.esPrimeroEnDepto(depto_n) and not self.deptos[depto_n].estaOcupado()) or self.atendido):
                # print(f"\tPersona {self.id} ESPERA en la FILA de {self.deptos[depto_n].nombre}... ({personas_restantes})\nHay peronas suficientes: {self.deptos[depto_n].hayPersonasSuficientes()}\nEs el primero en la fila: {self.esPrimeroEnDepto(depto_n)}\nEl depto esta ocupado: {self.deptos[depto_n].estaOcupado()}\nYa fue atendido: {self.atendido}\nHay gente en la fila del patio: {self.deptos[depto_n].hayPersonasEnFilaPatio()}\nIncapaz de llenar depto: {self.deptos[depto_n].incapazDeLlenarDepto()}\n")
                self.deptos[depto_n].cv_capacidadDepto.wait()
                # print(f"\n\t!!!Persona {self.id} DESPIERTA en la FILA de {self.deptos[depto_n].nombre}. ({personas_restantes})\n")

            if not self.atendido:
                print(f"\tPersona {self.id} NO ATENDIDA. Iniciando proceso en {self.deptos[depto_n].nombre}. ({personas_restantes})")
                # Si no ha sido atendido en el depto, entonces entra al depto y se atiende
                self.deptos[depto_n].atender()

            # Las personas que saltan el 'if' anterior son las que ya fueron atendidas en el depto.
            print(f"\t\t=X=Persona {self.id} YA ha sido ATENDIDO en {self.deptos[depto_n].nombre}. ({personas_restantes})")
            self.deptos[depto_n].registrarActividad(self.id, self.hora_entrada_fila, self.hora_entrada_depto, depto_n+1)

            # Notifica a las hebras que estan esperando (por que no habia suficientes personas en la fila del depto)
            # para que despierten y verifiquen si pueden entrar al depto.
            self.deptos[depto_n].cv_capacidadDepto.notify_all()


NUM_PERSONAS = 500
personas_restantes = NUM_PERSONAS

carpeta_arch = "./salida/"

if not os.path.exists(carpeta_arch):
    os.makedirs(carpeta_arch)

arch_mat = carpeta_arch + "Departamento_de_Matematicas.txt"
arch_inf = carpeta_arch + "Departamento_de_Informatica.txt"
arch_fis = carpeta_arch + "Departamento_de_Fisica.txt"
arch_qui = carpeta_arch + "Departamento_de_Quimica.txt"
arch_defider = carpeta_arch + "DEFIDER.txt"
arch_mec = carpeta_arch + "Departamento_de_Mecanica.txt"
arch_minas = carpeta_arch + "Departamento_de_Minas.txt"
arch_lamparas = carpeta_arch + "PdLamparas.txt"
arch_salidas = carpeta_arch + "Salidas.txt"

mat = Departamento("Matematicas", 20, 9, 10, arch_mat)
inf = Departamento("Informatica", 8, 5, 2, arch_inf)
fis = Departamento("Fisica", 15, 7, 5, arch_fis)
qui = Departamento("Quimica", 6, 4, 3, arch_qui)
defider = Departamento("DEFIDER", 6, 1, 5, arch_defider)
mec = Departamento("Mecanica", 9, 4, 4, arch_mec)
minas = Departamento("Minas", 7, 5, 2, arch_minas)

lista_deptos: list[Departamento] = [mat, inf, fis, qui, defider, mec, minas]

random.seed(1)

def simulacion(id: int):
    persona: Persona = Persona(id+1, lista_deptos)
    hora_llegada_u: str = datetime.now().strftime("%H:%M:%S.%f")

    print(f"\n==> Persona {persona.id} INGRESA a UNIVERSIDAD [ {persona.deptos[0].nombre}, {persona.deptos[1].nombre} ]\n")

    # La persona llega al patio de las lámparas
    persona.entrarAPatio(0)

    # La persona avanza a la fila del primer depto
    persona.avanzarAFila(0)

    # La persona entra al primer depto
    persona.entrarADepto(0)

    # Guardar tiempos y reiniciar
    hora_fila_depto1: str = persona.hora_entrada_fila
    persona.atendido = False

    # La persona llega al patio de las lámparas
    persona.entrarAPatio(1)

    # La persona avanza a la fila del segundo depto
    persona.avanzarAFila(1)

    # La persona entra al segundo depto
    persona.entrarADepto(1)

    # Guardar tiempos y reiniciar
    hora_fila_depto2: str = persona.hora_entrada_fila
    hora_salida_u: str = datetime.now().strftime("%H:%M:%S.%f")
    persona.atendido = False
    print(f"\n<== Persona {persona.id} SALE de UNIVERSIDAD [ {persona.deptos[0].nombre}, {persona.deptos[1].nombre} ]\n")

    global personas_restantes
    personas_restantes -= 1

    # Guardar tiempos en el archivo de las lámparas
    with open(arch_lamparas, "a") as archivo:
        archivo.write(f"Persona{persona.id}, {hora_llegada_u}, {persona.deptos[0].nombre}, {hora_fila_depto1}, {persona.deptos[1].nombre}, {hora_fila_depto2}\n")

    # Guardar tiempos en el archivo de las salidas
    with open(arch_salidas, "a") as archivo:
        archivo.write(f"Persona{persona.id}, {hora_salida_u}\n")



threads: list[Thread] = []

for i in range(NUM_PERSONAS):
    thread = Thread(target=simulacion, args=(i,))
    threads.append(thread)
    thread.start()
    sleep(0.01)

for thread in threads:
    thread.join()

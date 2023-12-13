JFLAGS = -g
JC = javac
JAVA = java
.SUFFIXES: .java .class
.java.class:
	$(JC) $(JFLAGS) $*.java

CLASSES = \
	MainFork.java \
	MainThread.java \
	MainClasico.java

default: classes

classes: $(CLASSES:.java=.class)

clean:
	$(RM) *.class

runClasico:
	for file in $$(find ./prueba -name sopa_de_letras.txt); do \
		$(JAVA) MainClasico $$file; \
	done

runFork:
	for file in $$(find ./prueba -name sopa_de_letras.txt); do \
		$(JAVA) MainFork $$file; \
	done

runThread:
	for file in $$(find ./prueba -name sopa_de_letras.txt); do \
		$(JAVA) MainThread $$file; \
	done

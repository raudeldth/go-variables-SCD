#
#   secuencia de finocci de N numero
#

def fibonacci(x):
	if x == 1:
		return 1
	elif x == 0:
		return 0

	return fibonacci(x - 1) + fibonacci(x - 2)

numero = (int)(input("Ingrese un numero: "))

print(fibonacci(numero))

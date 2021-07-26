ALUNO: GABRIEL SOARES DA SILVA SANTOS<br>
Processador Intel i5 - 8 núcleos</br>

### Como executar:
- Instale o [Go](https://golang.org/);
- Navegue até a pasta do projeto baixado;
- Execute o seguinte comando:
```
$ go run main.go
```

### Resultados
#### 1 THREAD
```
Digite a quantidade de termos: 1000000000
Digite a quantidade de threads: 1
Valores - Valores - ( Pi / Tempo de execução ):  [
    {3.1415926525880504 1m37.2685418s}
    {3.1415926525880504 1m38.7006092s}
    {3.1415926525880504 1m39.7156669s}
    {3.1415926525880504 1m39.4897706s}
    {3.1415926525880504 1m39.8962521s}
]
Duração média (ms):  99013
Desvio padrão: 963.54
Coeficiente de variação: 0.9731 %
```
#### 2 THREADS 
```
Digite a quantidade de termos: 1000000000
Digite a quantidade de threads: 2
Valores - Valores - ( Pi / Tempo de execução ):  [
     {3.141592652589258 57.649799s}
     {3.141592652589258 59.0475278s}
     {3.141592652589258 58.6973399s}
     {3.141592652589258 58.5226576s}
     {3.141592652589258 59.8028005s}
]
Duração média (ms):  58743
Desvio padrão: 701.67
Coeficiente de variação: 1.1945 %
```
#### 4 THREADS 
```
Digite a quantidade de termos: 1000000000
Digite a quantidade de threads: 4
Valores - Valores - ( Pi / Tempo de execução ):  [
     {3.1415926525892104 37.1753543s}
     {3.1415926525892104 40.4343631s}
     {3.1415926525892104 40.5554022s}
     {3.1415926525892104 40.0690481s}
     {3.1415926525892104 41.3294813s}
]
Duração média (ms):  39912
Desvio padrão: 1428.96
Coeficiente de variação: 3.5803 %
```
#### 8 THREADS 
```
Digite a quantidade de termos: 1000000000
Digite a quantidade de threads: 8
Valores - ( Pi / Tempo de execução ):  [
    {3.141592652589324 27.4800085s}
    {3.141592652589324 29.0397896s}
    {3.141592652589324 30.3070001s}
    {3.141592652589324 30.4137545s}
    {3.141592652589324 30.3784084s}
]
Duração média (ms):  29523
Desvio padrão: 1144.19
Coeficiente de variação: 3.8756 %
```
#### 16 THREADS 
```
Digite a quantidade de termos: 1000000000
Digite a quantidade de threads: 16
Valores - ( Pi / Tempo de execução ):  [
    {3.141592652590108 27.3846109s}
    {3.141592652590108 29.5881879s}
    {3.141592652590108 30.3071212s}
    {3.141592652590108 30.2456789s}
    {3.141592652590108 30.3531532s}
]
Duração média (ms):  29565
Desvio padrão: 1127.49
Coeficiente de variação: 3.8136 %
```
#### 32 THREADS
```
Digite a quantidade de termos: 1000000000
Digite a quantidade de threads: 32
Valores - ( Pi / Tempo de execução ):  [
    {3.141592652590108 27.3846009s}
    {3.141592652590108 29.5881903s}
    {3.141592652590108 30.3861253s}
    {3.141592652590108 30.1096725s}
    {3.141592652590108 30.3589542s}
]
Duração média (ms):  29565
Desvio padrão: 1127.49
Coeficiente de variação: 3.8136 %
```

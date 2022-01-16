#Golang Standard Project

O objetivo desse projeto é servir de base para futuras aplicações em Golang.

A ideia aqui é sugerir uma divisão de pastas que reflita a estrutura em camadas da aplicação,
sendo essas `domínio`, `aplicação` e `infraestrutura`. 

Além de definir tais camadas esse modelo também pretende definir os seus limites. 
- `infraestrutura`: A camada mais exterior, a ela cabe lidar com requisições http, consumo e produção de mensageria, 
  logs, integrações com serviços externos, bancos de dados, etc... Essa camada pode acessar recursos das camadas inferiores
  como funções da camada de aplicação ou estruturas de dados do domínio da aplicação.
- `aplicação`: Este é o nível em que os componentes de domínio interagem formando as funcionalidades da aplicação, essa camada pode acessar
  recursos da camada de domínio apenas.
- `domínio`: Essa camada é responsável por definir estruturas de dados que representam as entidades de negócio e também o seu comportamento,
  outro tipo de componente que pode existir nessa camada são interfaces que definem como elementos de domínio, como repositórios de entidade.

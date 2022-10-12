### DESIGN PATTERNS IN GOLANG

<details>
  <summary><strong>S.O.L.I.D</strong></summary>

**S** - **Single Responsability Principle**
- De acordo com esse principio, cada classe, funcão ou componente deve ter apenas uma responsabilidade.

- Ganhos: 
  - Reaproveitamento de código;
  - Facilita a refatoração;
  - Aplicar teste automatizados com mais facilidade;
  - Gerar menos bugs;

**O** - **Open-Closed Principle**
- Aberto para extensão, fechado para modificações; 
  - Não é correto adicionar um novo `if` dentro da clase para adicionar um novo comportamento.

**L** - **Liskov Substituion Principle** 
- Deve-se poder intercambiar implementações de uma determinada classe;
- Uma subclasse não pode quebrar as expectativas estabalecidas pelo superclasse;
- Faz pensar o que realmente a classe pai fornece de comum para toas as subclasse;

**I** - **Interface Segregation Principle**
- Dividir em varias interface que somadas definirão todo o objeto.
- É melhor criar interfaces mais especificas ao invés de temos uma unica interface genérica.

**D** - **Dependency Inversion Principle**
- Não dependa implementações, dependa de abstrações;
- Utilizar de contratos(interfaces) para abstrair as implementações;
</details>

<details>
  <summary><strong>BUILDER</strong></summary>
  - O padrão Builder sugere que você extraia o código de construção do objeto para fora de sua própria classe e mova ele para objetos separados chamados builders.
  - Trata da criação de objetos complexos.
  - O objeto final pode variar de acordo com a necessidade.

  - BOM:
    - Os mesmos objetos são simples e podem ser criados em uma única chamada ao construtor.
    - Permite a criação de um objeto em etapas.
    - Separa criação de utilização.
    - O cliente não precisa criar objetos diretament
    - O mesmo código pode construir objetos diferentes
    - Ajuda na aplicação dos princípios SRP e OCP
  - RUIM:
    - O código final pode se tornar muito complexo
  - [Ref](https://refactoring.guru/pt-br/design-patterns/builder)
  - Builder Facets
</details>
<details>
  <summary><strong>Factory</strong></summary>
  - Factory são simplesmente operacões que criam objetos.
  - O FactoryMethod é um padrão de projeto de criação (lida com a criação de objetos).
  - Oculta a lógica de instanciação de código cliente. O método fábrica será responsável 
  por instanciar as classes desejadas.
  - Dá flexibilidade ao código client permitindo a criação de novas factories sem a necessidade
   de alerar o código já escrito(OCP)
</details>
<details>
  <summary><strong>Prototype</strong></summary>
  
  - Especificar os tipos de objetos a serem criados usando uma instância-protótipo e
    criar novos objetos pelo cópia desse protótipo.
</details>

<details>
<summary><strong>Singleton</strong></summary>

- Garantir que uma classe tenha apenas uma única instância;

- Fornece um ponto de acesso global para aquela instância;

- Utilize o padrão Singleton quando você precisa de um controle mais estrito sobre as variáveis globais;
</details>

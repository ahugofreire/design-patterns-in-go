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
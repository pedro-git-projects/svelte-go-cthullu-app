# Cthullhu Online

## O que é

Cthullhu Online é uma aplicação web full stack constituída por uma RESTful API escrita em Go e um frontend escrito em Svelte.

O frontend é um site completo, mas sua comunicação com o backend se dá na rota /criar. 
Neste endereço há um formulário com checagem do lado do cliente que faz uma requisição HTTP do tipo POST para o backend e exibe a resposta dinâmicamente na forma de tabelas.

Todos os cálculos seguem as regras oficiais da sétima edição do RPG Call of Cthullhu da Chaosium.

# Backend

## Endpoints

O Backend é escrito em Go e expõe dois endpoints:


| Método | Padrão da URL | Ação|
|--------|---------------------|----------------------------|
| GET    | /v1/healthcheck     | Mosta as informações do servidor      |
| POST   | /v1/criar 			| Gera um novo investigador |


## Dependências

### julienschmidt/httprouter

Ao construir endpoints de API sem bibliotecas de terceiros  em Go temos de lidar com o fato de que o http.ServeMux não permite rotear para handlers distintos utilizando o mesmo método. Além disso também não oferece suporte para URLs limpos com parâmetros interpolados.

Assim, a escolha foi feita para usar o httprouter de julienschmidt, por não ter nenhuma dessas duas limitações além de ser extremamente eficiente pelo uso de radix sort para identificação das URLs. 

### rs/cors

É um handler net/http que implementa a especificação W3 de Compartilhamento de Recursos de Origem Cruzada em Golang.

# Frontend

## Rotas

Escrito em Svelte utilizando o framework Sveltekit, o frontend conta com 7 rotas:

- /home, que conta com uma imagem, texto e um botão que leva para a criação de personagens. 

- /criar, contendo um formulário com validação que exibe erros dinâmicamente e ao ser submetido faz uma requisição POST assíncrona para o backend e renderiza a resposta formatada em uma tabela.  

- /lovcraft, que conta com uma imagem e textos sobre a vida e obra do autor

- /instrucoes, que, além de texto, contém um carrossel que pode ser controlado com o teclado ou passado com botões e um componente do youtube que utiliza a biblioteca vime.

- /produtos, que conta com uma lista de 10 produtos renderizados dinamicamente em cards. Ao clicar em "adicionar ao carrinho", um novo card com a imagem do produto seleicionado, a quantidade e um botão para adicionar ou remover mais cópias do mesmo produto é renderizado, assim como o valor total da soma de todos os produtos selecionados.

-  /galeria consiste em um uma galeria responsiva de diversas imagens, as quais são organizadas automaticamente na melhor resolução.

- /codigo onde temos o link para o github, tabelas e blocos de código com destaque sintático utilizando a biblioteca Highlight js.

## Temas

O site conta com dois esquemas de cor que podem ser alterados a qualquer momento clicando em um botão na navbar. Eles são Gruvbox e Dracula. Todos os componentes trocam de cor imediatamente e de forma consiste, este resultado foi obtido através do uso de slectors e uma Svelte Store.

Além disso, por fazer uso consistente das classes do bootstrap 5, o site é completamente responsivo.

# Makefile

Para facilitar a utilização a aplicação conta com um Makefile.

Com o GNU Make instalado o comando:

- Make  install

Instala todas as bibliotecas necessárias, tanto para front quanto para o backend

- Make Dev

Inicia o servidor do frontend 3000

- Make run

Inicia o servidor do backend na porta 4000

# UX

O site conta com um wireframe desnvolvido no Figma. Ele pode ser encontrado na pasta UX.

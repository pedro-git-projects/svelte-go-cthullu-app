<script>
	import Highlight from "svelte-highlight"
	import gruvbox from "svelte-highlight/styles/gruvbox-dark-medium"
	import dracula from "svelte-highlight/styles/dracula"
	import json from "svelte-highlight/languages/json"

	import { darkMode } from "../store" 

	/* isDark recebe o valor de darkMode da Store */
	let isDark
	darkMode.subscribe(valor => {
		isDark = valor
	})


	let request = `{
	"nome":"Pedro",
	"idade":30,
	"residencia":"Boston",
	"nascimento":"new york",
	"ocupacao":"archeologist" 
}  `

	let response = `{
	"investigator": {
		"name": "Pedro",
		"age": 30,
		"residence": "Boston",
		"birthplace": "new york",
		"occupation": "archeologist",
		"str": 75,
		"con": 15,
		"pow": 15,
		"dex": 75,
		"app": 45,
		"siz": 80,
		"int": 50,
		"edu": 80,
		"luck": 30,
		"mp": 3,
		"db": 3,
		"build": 1,
		"hp": 9,
		"san": 15,
		"mv": 7,
		"description": {
			"str_description": "average, for a human.",
			"app_description": "ugly, possibly difigured due to injury or birth.",
			"con_description": "weak health, you're prone to bouts of ill health, great propensity for feeling pain.",
			"int_description": "average human intellect.",
			"siz_descrpition": "very tall, strongly built, or obese (240lbs/150 kg).",
			"pow_description": "weak-willed, easily dominated by those with a greater intellect or willpower.",
			"edu_description": "degree level graduate.",
			"dex_description": "fast, nimble and able to perform feats of fine manipulation."
		}
	}
}`
</script>


<svelte:head>
	<title>Cthullhu Online - Fonte</title>
	{#if isDark}
		{@html dracula}
	{:else}
		{@html gruvbox}
	{/if}
</svelte:head>


<h1>Código Fonte</h1>
<hr>

<div id="checkDiv" class="" style="display: none;"></div>

<p>
	Este site é Open Source, licenciado sob a <strong>GLP3</strong> e seu código fonte está disponível <a href="https://github.com/pedro-git-projects/cthullu-online" target="blank"> aqui</a>.
	Ele consiste em uma aplicação full stack, databaseless com backend escrito em <strong>Go</strong> e frontend escrito utilizando <strong>Svelte</strong>. Além disso <strong>Bootstrap</strong> foi utilizado visando simplificar o desenvolvimento dos estilos do css. Por fim, o <strong>GNU Make</strong> foi utilizado para automatizar o processo de levantar o site.
</p>


<h2>Go</h2>
<p>
	O backend da aplicação foi escrito em Go e consiste em uma <strong>RESTful API</strong>. Esta API expõe dois endpoints:	
</p>

<table id="tabela" class="{isDark === false ? 'table table-light' : 'table table-dark text-white'} ">
	<thead class="thead">
		<tr>
			<th scope="col">Método</th>
			<th scope="col">Padrão da URL</th>
			<th scope="col">Ação</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<th scope="row">GET</th>
			<td>/v1/healthcheck</td>
			<td>Exibe a disponibilidade do servidor</td>
		</tr>
		<tr>
			<th scope="row">POST</th>
			<td>/v1/criar</td>
			<td>Cria um novo investigator</td>
		</tr> </tbody>
</table>

<p>
	A aplicação backend cria investigadores recebendo um <strong>POST</strong> contendo o nome, idade, residência  e ocupação do investigator em JSON do seguinte tipo:
</p>


<div class="container">
	<div class="row">
		<div class="col">
			<Highlight language={json} code={request} />
		</div>
	</div>
</div>


E responde também com JSON uma ficha de personagem como:

<div class="container">
	<div class="row">
		<div class="col">
			<Highlight language={json} code={response} />
		</div>
	</div>
</div>




<h2>Go - Dependências</h2>

<p>
	As dependências de backend são apenas duas: <strong>julienschmidt/httprouter</strong> e <strong>rs/cors</strong>.
</p>

<p>
	A primeira foi utilizada porque ao construir endpoints de API sem bibliotecas de terceiros, enfrentamos a limitação de que o http.ServeMux não permite roteamento para diferentes manipuladores com base no método de solicitação. Ele também não oferece suporte para URLs limpos com parâmetros interpolados.
</p>

<p>
	Assim, a escolha foi feita para usar o httprouter de julienschmidt, pois ele resolve ambos os problemas e é extremamente eficiente porque usa um algoritmo de classificação radix para correspondência de URL.
</p>

<h2>Svelte</h2>

<p>
	O frontend foi escrito utilizando <strong>Svelte</strong> e <strong>SapperJS</strong>. Esta escolha foi feita dada a grande flexibilidade desse framework, assim como sua alta preformance dada a ausência de difernciação de DOM virutal.
</p>


<h2>Svelte - Dependências</h2>

Foram utilizadas as seguintes bibliotecas:

<ul>
	<li>@vime/core</li>
	<li>@vime/svelte</li>
	<li>svelte-image-gallery</li>
</ul>



<style>
	:global() h1 { 
		color:#98971a;
		font-family: 'JMH Cthulhumbus UGalt1'; 
	}
	:global() h2 { 
		color:#d79921;
		font-family: 'JMH Cthulhumbus UGalt1'; 
	}

	:global(.dark) h1 { 
		color:#50fa7b;
	}

	:global(.dark) h2 { 
		color:#ff79c6;
	}
</style>

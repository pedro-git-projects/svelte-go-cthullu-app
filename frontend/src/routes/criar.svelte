<svelte:head>
	<title>Cthullu Online - Criar</title>
</svelte:head>

<script>
	import { createForm } from "svelte-forms-lib"

	import { darkMode } from "../store" 

	let isDark
	darkMode.subscribe(valor => {
		isDark = valor
	})

	let promise = Promise.resolve([])

	const { form, errors, handleChange, handleSubmit } = createForm({
		initialValues: {
			nome: '',
			idade: '',
			residencia: '',
			nascimento: '',
			ocupacao: '',
		},
		validate: values => {
			let errs = {}
			if(values.nome === "") {
				errs["nome"] = "Nome é um campo obrigatório"
			}

			if(values.idade=== "") {
				errs["idade"] = "Insira um número entre 1 e 99"
			}

			if(values.idade < 1 || values.idade > 99) {
				errs["idade"] = "Insira um número entre 1 e 99"
			}

			if( isNaN(values.idade)) {
				errs["idade"] = "Insira um número entre 1 e 99"
			}

			if(values.residencia === "") {
				errs["residencia"] = "Residência é um campo obrigatório"
			}		

			if(values.nascimento === "") {
				errs["nascimento"] = "Nascimento é um campo obrigatório"
			}

			if(values.ocupacao === "") {
				errs["ocupacao"] = "Ocupação é um campo obrigatório"
			}		
			return errs
		},

		onSubmit: values => {
			submitForm(values)
		}
	})


	function submitForm(values) {
		promise = fetchInvestigador(values)
	}

	async function fetchInvestigador(values) {
		const response = await self.fetch('http://localhost:4000/v1/criar', {
			method: "POST",
			body: JSON.stringify({
				nome: values.nome,	
				idade: parseInt(values.idade),
				residencia: values.residencia,
				nascimento: values.nascimento,
				ocupacao: values.ocupacao,
			})
		})
		console.log(values)
		if (response.ok) {
			return response.json()

		} else {
			throw new Error()
		}
	}
</script>

<h1>Criar Investigador</h1>
<hr>


<form on:submit|preventDefault={handleSubmit}>

	<div class="form-group">

		<label for="nome">Nome</label>
		<input
	  class="form-control"
   			id="nome"
   			name="nome"
	  		placeholder="Nome do investigador"
   			on:change={handleChange}
   			bind:value={$form.nome}
   />
   {#if $errors.nome}
	   <p class="error-alert">{$errors.nome}</p>
   {/if}

   <label for="idade">Idade</label>
   <input
	   	class="form-control"
		id="idade"
 		name="idade"
	  	placeholder="1-99"
 		on:change={handleChange}
 		bind:value={$form.idade}
 />
 {#if $errors.idade}
	 <p class="error-alert">{$errors.idade}</p>
 {/if}

 <label for="residencia">Residência</label>
 <input
	 class="form-control"
  id="residencia"
  name="residencia"
  placeholder="Cidade de residência"
  on:change={handleChange}
  bind:value={$form.residencia}
  />
  {#if $errors.residencia}
	 <p class="error-alert">{$errors.residencia}</p>
  {/if}

  <label for="nascimento">Nascimento</label>
  <input
	  class="form-control"
	  id="nascimento"
   	  name="nascimento"
   	  placeholder="Local de nascimento"
   on:change={handleChange}
   bind:value={$form.nascimento}
   />
   {#if $errors.nascimento}
	 <p class="error-alert">{$errors.nascimento}</p>
   {/if}

   <label for="ocupacao">Ocupação</label>
   <input
	   class="form-control"
		id="ocupacao"
 		name="ocupacao"
   		placeholder="Ocupação, ex: arqueólogo"
 on:change={handleChange}
 bind:value={$form.ocupacao}
 />
 {#if $errors.ocupacao}
	 <p class="error-alert">{$errors.ocupacao}</p>
 {/if}
	</div>

	<div class="container">
		<div class="row">
			<div class="col text-center mt-4 mb-4">
				<button class="btn btn-primary" type="submit">Criar!</button>
			</div>
		</div>
	</div>
</form>


{#await promise}
	<div></div>
{:then caracteristicas}
	<h2>Características</h2>
	<table shadow class="{isDark === false ? 'table table-light' : 'table table-dark'}">
		<thead>
			<tr>
				<th scope="col">Característica</th>
				<th scope="col">Valor</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<th scope="row">Nome</th>
				<td>{caracteristicas.investigator.name}</td>
			</tr>
			<tr>
				<th scope="row">Idade</th>
				<td>{caracteristicas.investigator.age}</td>
			</tr>
			<tr>
				<th scope="row">Residência</th>
				<td>{caracteristicas.investigator.residence}</td>
			</tr>
			<tr>
				<th scope="row">Local de Nascimento</th>
				<td>{caracteristicas.investigator.birthplace}</td>
			</tr>
			<tr>
				<th scope="row">Ocupação</th>
				<td>{caracteristicas.investigator.occupation}</td>
			</tr>
			<tr>
				<th scope="row">Força</th>
				<td>{caracteristicas.investigator.str}</td>
			</tr>
			<tr>
				<th scope="row">Constituição</th>
				<td>{caracteristicas.investigator.con}</td>
			</tr>
			<tr>
				<th scope="row">Poder</th>
				<td>{caracteristicas.investigator.pow}</td>
			</tr>
			<tr>
				<th scope="row">Destreza</th>
				<td>{caracteristicas.investigator.dex}</td>
			</tr><tr>
				<th scope="row">Aparência</th>
				<td>{caracteristicas.investigator.app}</td>
			</tr><tr>
				<th scope="row">Poder</th>
				<td>{caracteristicas.investigator.pow}</td>
			</tr><tr>
				<th scope="row">Tamanho</th>
				<td>{caracteristicas.investigator.siz}</td>
			</tr><tr>
				<th scope="row">Inteligência</th>
				<td>{caracteristicas.investigator.int}</td>
			</tr><tr>
				<th scope="row">Educação</th>
				<td>{caracteristicas.investigator.edu}</td>
			</tr><tr>
				<th scope="row">Sorte</th>
				<td>{caracteristicas.investigator.luck}</td>
			</tr><tr>
				<th scope="row">MP</th>
				<td>{caracteristicas.investigator.mp}</td>
			</tr><tr>
				<th scope="row">Bônus de dano</th>
				<td>{caracteristicas.investigator.db}</td>
			</tr>
			<tr>
				<th scope="row">Constituição</th>
				<td>{caracteristicas.investigator.build}</td>
			</tr>
			<tr>
				<th scope="row">HP</th>
				<td>{caracteristicas.investigator.hp}</td>
			</tr>
			<tr>
				<th scope="row">Sanidade</th>
				<td>{caracteristicas.investigator.san}</td>
			</tr>  
			<tr>
				<th scope="row">Taxa de movimento</th>
				<td>{caracteristicas.investigator.mv}</td>
			</tr>  
		</tbody>
	</table>
	<br>
	<h2> Descrição </h2>
	<ul class="  list-group list-group-flush">
		<li class="shadow list-group-item 
		{isDark === false ? 'bg-light' : 'bg-dark'}	
		text-white">Força: {caracteristicas.investigator.description.str_description}</li>

		<li class="list-group-item {isDark === false ? 'bg-light' : 'bg-dark'} text-white">Aparência: {caracteristicas.investigator.description.app_description}</li>
		<li class="list-group-item {isDark === false ? 'bg-light' : 'bg-dark'} text-white">Constituição: {caracteristicas.investigator.description.con_description}</li>
		<li class="list-group-item {isDark === false ? 'bg-light' : 'bg-dark'} text-white">Inteligência: {caracteristicas.investigator.description.int_description}</li>
		<li class="list-group-item {isDark === false ? 'bg-light' : 'bg-dark'} text-white">Tamanho: {caracteristicas.investigator.description.siz_descrpition}</li>
		<li class="list-group-item {isDark === false ? 'bg-light' : 'bg-dark'} text-white">Poder: {caracteristicas.investigator.description.pow_description}</li>
		<li class="list-group-item {isDark === false ? 'bg-light' : 'bg-dark'} text-white">Educação: {caracteristicas.investigator.description.edu_description}</li>
		<li class="list-group-item {isDark === false ? 'bg-light' : 'bg-dark'} text-white">Destreza: {caracteristicas.investigator.description.dex_description}</li>
	</ul>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}

<style>
	/* Sobreescrevendo a cor do botao */
	:global(.dark) button {
		background-color: #bd93f9 !important;
		border-color:#bd93f9 !important;
	}

	:global() h1 { 
		font-family: 'JMH Cthulhumbus UGalt1'; 
		color:#98971a;
	}
	:global() h2 { 
		color:#d79921;
	}

	:global(.dark) h1 { 
		color:#50fa7b;
	}

	:global(.dark) h2 { 
		color:#ff79c6;
	}


	:global(.dark) .form-control:focus {
		border-color: #28a745;
		box-shadow: 0 0 0 0.2rem rgba(189, 147, 249, 0.5);
	} 

	:global(.dark) .form-control {
		background-color:#f8f8f2;
	}

	:global() .error-alert {
		border: 1px solid #ebdbb2 !important;
		padding: 6px;
		text-align: center;
		color: #ebdbb2;
		background-color: #cc241d;
		border-radius: 3px;
		margin-top: 1em;
	}

	:global(.dark) .error-alert {
		border: 1px solid #f8f8f2 !important;
		padding: 6px;
		text-align: center;
		color: #f8f8f2;
		background-color: #ff5555;
		border-radius: 3px;
		margin-top: 1em;
	}
	
	:global() button {
		font-family: 'JMH Cthulhumbus UGalt1'; 
	}

	:global(.dark) button:hover {
		box-shadow: inset 0 0 100px 100px rgba(255, 255, 255, 0.1);		
	}

	:global(.dark) button:focus{
		border:#bd93f9;
		box-shadow: 0 5px 15px rgba(255, 255, 255, 0.1);		
	}
	
</style>

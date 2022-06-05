<svelte:head>
	<title>Cthullu Online - Criar</title>
</svelte:head>

<script>
	let promise = Promise.resolve([])

	let nome = ''
	let idade = '' 
	let residencia = '' 
	let nascimento = '' 
	let ocupacao = ''  

	function submitForm() {
		promise = fetchInvestigador()
	}

	async function fetchInvestigador() {
		const response = await self.fetch('http://localhost:4000/v1/criar', {
			method: "POST",
			body: JSON.stringify({
				nome,	
				idade,
				residencia,
				nascimento,
				ocupacao,
			})
		})

		if (response.ok) {
			return response.json()

		} else {
			throw new Error()
		}
	}
</script>

<h1>Criar Investigador</h1>
<hr>
<form on:submit|preventDefault={submitForm}>

	<div class="form-group">

		<label for="nome"> <strong> Nome </strong> </label>
		<input class="form-control" bind:value="{nome}" id="nome" name="nome" type="text" />

		<label for="idade"> <strong> Idade </strong> </label>
		<input class="form-control" bind:value="{idade}" id="idade" name="idade" type="number" />

		<label for="residencia"> <strong> Residência </strong> </label>
		<input class="form-control" bind:value="{residencia}" id="residencia" name="residencia" type="text" />

		<label for="nascimento"> <strong> Nascimento </strong> </label>
		<input class="form-control" bind:value="{nascimento}" id="nascimento" name="nascimento" type="text" />

		<label for="ocupacao"> <strong> Ocupação </strong> </label>
		<input class="form-control" bind:value="{ocupacao}" id="ocupacao" name="ocupacao" type="text" />

	</div>
	<br />

	<div class="container">
		<div class="row">
			<div class="col text-center mb-4">
				<button class="btn btn-primary" type="submit">Criar!</button>
			</div>
		</div>
	</div>
</form>


{#await promise}
	<div></div>
{:then caracteristicas}
	<table class="table table-light">
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
	<ul class="list-group list-group-flush">
		<li class="list-group-item 
		bg-light 
		text-white">Força: {caracteristicas.investigator.description.str_description}</li>

		<li class="list-group-item bg-light text-white">Aparência: {caracteristicas.investigator.description.app_description}</li>
		<li class="list-group-item bg-light text-white">Constituição: {caracteristicas.investigator.description.con_description}</li>
		<li class="list-group-item bg-light text-white">Inteligência: {caracteristicas.investigator.description.int_description}</li>
		<li class="list-group-item bg-light text-white">Tamanho: {caracteristicas.investigator.description.siz_descrpition}</li>
		<li class="list-group-item bg-light text-white">Poder: {caracteristicas.investigator.description.pow_description}</li>
		<li class="list-group-item bg-light text-white">Educação: {caracteristicas.investigator.description.edu_description}</li>
		<li class="list-group-item bg-light text-white">Destreza: {caracteristicas.investigator.description.dex_description}</li>
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

</style>

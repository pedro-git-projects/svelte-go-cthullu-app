<svelte:head>
	<title>Cthullhu Online - Home</title>
</svelte:head>

<script>
	import cultist from  "../../static/home_image_3.png"
	import { goto } from '$app/navigation'

	const goToCriar = () => {
		goto("/criar")
	}
</script>

<div class="container">
	<div class="row">
		<div class="col text-center mb-3"> <h1> Cthullu Online</h1>
		</div>
	</div>
</div>

<div class="container">
	<div class="row">
		<div class="col text-center mb-3">
			Bem vindo ao <b>Cthullu Online</b>, um site onde você
			pode gerar investigadores para a sétima edição do jogo 
			<i>Call of Cthullu</i> da Chaosium.			

		</div>
	</div>
</div>

<div class="container">
	<div class="row">
		<div class="col text-center">
			<button 
	   on:click="{goToCriar}"
	   type="button" 
	class="btn btn-primary mt-3">Criar</button>
		</div>
	</div>
</div>

<div class="container mt-0">
	<div class="row">
		<div class="col text-center">
			<img alt="Cultist" class="img-fluid" src="{cultist}">
		</div>
	</div>
</div>

<style>

	:global() h1 { 
		font-family: 'JMH Cthulhumbus UGalt1'; 
		color:#98971a;
	}

	:global(.dark) h1 { 
		color:#50fa7b;
	}

	:global() button {
		font-family: 'JMH Cthulhumbus UGalt1'; 
	}

	:global(.dark) button {
		background-color: #bd93f9 !important;
		border-color: #bd93f9 !important;
	}

	:global(.dark) button:hover {
		box-shadow: inset 0 0 100px 100px rgba(255, 255, 255, 0.1);		
	}

</style>

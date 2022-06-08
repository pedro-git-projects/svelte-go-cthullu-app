<script>
	import { page } from '$app/stores'
	import Toggle from './Toggle.svelte'
</script>

<header>
	<nav id="nav" class="navbar shadow-sm navbar-expand-lg navbar-light bg-light">
		<div class="container-fluid">
			<div class="navbar-brand">
				<a sveltekit:prefetch href="/">Cthullu Online</a>
			</div>

			<button 
	   			class="navbar-toggler" 
	   			type="button" 
				data-bs-toggle="collapse" 
 				data-bs-target="#navbarNav" 
 				aria-controls="navbarNav" 
 				aria-expanded="false" 
 				aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>

			<div class="collapse navbar-collapse" id="navbarNav">
				<ul class="navbar-nav">

					<li class="nav-item" class:active={$page.url.pathname === '/criar'}>
						<div class="nav-link">
							<a sveltekit:prefetch href="/criar">Criar</a>
						</div>
					</li>
			
				<li class="nav-item">
					<div class="nav-link" class:active={$page.url.pathname === "/lovecraft"} to="/lovecraft">
						<a sveltekit:prefetch href="/lovecraft">Lovecraft</a>
					</div>
				</li>

				<li class="nav-item">
					<div class="nav-link" class:active={$page.url.pathname === "/instrucoes"} to="/instrucoes">
						<a sveltekit:prefetch href="/instrucoes">Instruções</a>
					</div>
				</li>

				<li class="nav-item">
					<div class="nav-link" class:active={$page.url.pathname === "/comprar"} to="/comprar">
						<a sveltekit:prefetch href="/comprar">Produtos</a>
					</div>
				</li>

				<li class="nav-item">
					<div class="nav-link" class:active={$page.url.pathname === "/galeria"} to="/galeria">
						<a sveltekit:prefetch href="/galeria">Galeria</a>
					</div>
				</li>				
				<li class="nav-item">
					<div class="nav-link" class:active={$page.url.pathname === "/fonte"} to="/fonte">
						<a sveltekit:prefetch href="/fonte">Código</a>
					</div>
				</li>

					<Toggle/>
				</ul>
			</div>
		</div>
	</nav>
</header>

<style>
	.navbar-brand {
		font-family: 'JMH Cthulhumbus UGalt1'; 
	}
	
	:global(a) {
		color: var(--link-color);
		text-decoration: none;
	}

	:global(a:hover) {
		color: #d79921;
		text-decoration: none;
	}

	:global(a.dark) {
		color: #f8f8f2;
		text-decoration: none;
	}

	:global(a.dark:hover) {
		color: #ff79c6;
		text-decoration: none;
	}

	:global() .navbar-toggler-icon {
		background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' width='30' height='30' viewBox='0 0 30 30'%3e%3cpath stroke='rgba(251, 241, 199, 0.5)' stroke-linecap='round' stroke-miterlimit='10' stroke-width='2' d='M4 7h22M4 15h22M4 23h22'/%3e%3c/svg%3e");	
	}

	:global(.dark) .navbar-toggler-icon {
		background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' width='30' height='30' viewBox='0 0 30 30'%3e%3cpath stroke='rgba(248, 248, 242, 0.5)' stroke-linecap='round' stroke-miterlimit='10' stroke-width='2' d='M4 7h22M4 15h22M4 23h22'/%3e%3c/svg%3e");	
	}

</style>

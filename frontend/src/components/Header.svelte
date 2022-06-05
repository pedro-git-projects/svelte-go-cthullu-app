<script>
	import { page } from '$app/stores'
	import Toggle from './Toggle.svelte'
</script>

<header>
	<nav id="nav" class="navbar navbar-expand-lg navbar-dark bg-dark">
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
					<li class="nav-item" class:active={$page.url.pathname === '/'}>
						<div class="nav-link">
							<a sveltekit:prefetch href="/">Home</a>
						</div>
					</li>
			
				<li class="nav-item">
					<div class="nav-link" class:active={$page.url.pathname === "/lovecraft"} to="/lovecraft">
						<a sveltekit:prefetch href="/lovecraft">Lovecraft</a>
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
		font-family: 'Macondo', cursive;
	}

	:root{
		--bg-color: #FFFFFF;
		--text-color: #000000;
		--link-color: #FFFFFF;
		--bootstrap-nav: #212529;
	}

	:global(body) {
		background: var(--bg-color);
		color: var(--text-color);
		transition: background-color 0.3s
	}

	:global(body.dark) {
		--bg-color: #000000;
		--text-color: #FFFFFF;
	}

	:global(a) {
		color: var(--link-color);
		text-decoration: none;
	}

	:global(a.dark) {
		color: #000000;
		text-decoration: none;
	}
</style>

<script>
	let darkMode = false;
	function toggle() {
		darkMode = !darkMode

		/* Muda o tema do corpo */
		window.document.body.classList.toggle('dark')
		let w = window.document.body.classList

		/* Muda o tema dos links */
		let a = document.getElementsByTagName('a') 
		for (let i = 0; i < a.length; i++) { 
			if (!(a[i].classList.contains("dark"))) { 
				a[i].classList.add("dark") 
			} else  { 
				a[i].classList.remove("dark") 
			} 
		}  

		/* Muda o tema da navbar */
		let nav = document.getElementById('nav') 
		if(nav.classList.contains("navbar-dark")) {
			nav.classList.remove('navbar-dark') 
			nav.classList.remove('bg-dark') 
			nav.classList.add('navbar-light') 
			nav.classList.add('bg-light') 
		} else { 
			nav.classList.remove('navbar-light') 
			nav.classList.remove('bg-light') 
			nav.classList.add('navbar-dark') 
			nav.classList.add('bg-dark') 
		} 

		/* Muda o tema do footer */
		let f = document.getElementById('meuFooter')
		if(f.classList.contains('bg-light')) {
			f.classList.remove('bg-light')
			f.classList.add('bg-dark')
		} else {
			f.classList.remove('bg-dark')
			f.classList.add('bg-light')
		}

		/* Muda o tema das tabelas */
		let t = document.getElementsByTagName('table') 
		for (let i = 0; i < t.length; i++) { 
			if ( w.length > 0 && t[i].classList.contains("table-light")) { 
				t[i].classList.remove("table-light") 
				t[i].classList.add("table-dark") 
			} else  {
				t[i].classList.remove("table-dark") 
				t[i].classList.add("table-light") 
			} 
		}

		/* Muda o tema de listas */
		let lg = document.getElementsByClassName('list-group-item')
		for (let i = 0; i < lg.length; i++) { 
			if (w.length > 0 && lg[i].classList.contains('bg-light')) { 
				lg[i].classList.remove("bg-light") 
				lg[i].classList.add("bg-dark") 
			} else  {
				lg[i].classList.remove("bg-dark") 
				lg[i].classList.add("bg-light") 
			} 
		}	

		/* Muda o tema dos pontos do carrossel */
		let d = document.getElementsByClassName('dot')
		for (let i = 0; i < d.length; i++) {
			if(d[i].classList.contains('bg-light')) {
				d[i].classList.remove('bg-light')
				d[i].classList.add('bg-dark')
			} else {
				d[i].classList.remove('bg-dark')
				d[i].classList.add('bg-light')
			}
		}

		/* Por serem objetos gerados dinâmicamente, 
		o tema dos cards é alterado no próprio componente de forma global */
	}
</script>

<button class="ml-2" on:click={toggle}>
	{#if darkMode }
		<!-- emoji do sol  -->		
		&#9728;&#65039; 	
	{:else}
		<!-- emoji do da lua -->		
		&#127769;	
	{/if}
</button>

<style>
	button {
		background: var(--botsrap-nav);
		border: 2px var(--text-color);
		border-radius: 5px;
		color: var(--text-color);
		padding: 5px 10px;
	}

	button:active {
		background: inherit;
	}
</style>

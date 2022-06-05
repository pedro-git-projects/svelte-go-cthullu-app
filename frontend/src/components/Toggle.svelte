<script>
	let darkMode = false;
	function toggle() {
		darkMode = !darkMode

		/* Muda o tema do corpo */
		window.document.body.classList.toggle('dark')

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


		/* Muda o tema da tabela */
		let t = document.getElementById('tabela')
		if(t.classList[0] === "table") {
			t.classList.remove('table')
			t.classList.add('table-dark')
		} else {
			t.classList.remove('table-dark')
			t.classList.add('table')
		}
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

<svelte:head>
	<title>Cthullu Online - Instruções</title>
</svelte:head>

<script>
	import { hslide } from '../components/Slider'	
	import { onMount } from 'svelte'
	import { darkMode } from "../store" 

	let isDark
	darkMode.subscribe(valor => {
		isDark = valor
	})

	let isMounted = false
	let video
	let Player
	let Youtube

	onMount(async () => {
		const vime = await import('@vime/svelte')
		Player = vime.Player
		Youtube = vime.Youtube
		isMounted = true

		setTimeout(() => {
			video.play();
		}, 100);
	})

	import img1 from '../../static/img1.png'
	import img2 from '../../static/img2.png'
	import img3 from '../../static/img3.png'

	let slides = [
		{ content: '1', bg: img1 },
		{ content: '2', bg: img3 },
		{ content: '3', bg: img2 },
	]

	let cur = 0;

	function changeSlide(slide) {
		cur = slide;
	}

	const clamp = (number, min, max) => Math.min(Math.max(number, min), max)
	const transition_args = {
		duration: 450,
	}

	function prev() {
		cur = clamp( --cur, 0, slides.length-1 )
	}

	function next() {
		cur = clamp( ++cur, 0, slides.length-1 )
	}

	const ARROW_LEFT = 37;
	const ARROW_RIGHT = 39;
	function handleShortcut(e) {
		if (e.keyCode === ARROW_LEFT ) {
			prev();
		}
		if (e.keyCode === ARROW_RIGHT ) {
			next();
		}
	}	

</script>

<h1>Sobre o Jogo</h1>
<hr>
<h2 class="mb-2"> O que é Call of Cthullhu</h2>
<p>
	Call of Cthulhu é um jogo sobre segredos, mistérios e horror. Desempenhando o papel de investigadores inabaláveis, equipes de personagens viajam para lugares estranhos e perigosos, descobrem tramas sujas e enfrentam os terrores da noite. No caminho, eles encontrarão entidades destruidoras de sanidade de além do espaço e do tempo, monstros hediondos e cultistas insanos. Dentro de tomos estranhos e esquecidos, eles descobrirão segredos que a humanidade não deveria conhecer. Essas pessoas comuns enfrentarão muitos desafios, mas sua posição heróica pode muito bem decidir o destino do mundo.
</p>

<p>
	Criado por Sandy Petersen e publicado pela primeira vez em 1981, Call of Cthulhu há mais de quarenta anos define o gênero e é consistentemente considerado um dos melhores jogos disponíveis para jogar. Para aqueles corajosos o suficiente para descobrir seus segredos, as recompensas estão além da compreensão!
</p>

<p>
	No jogo, cada jogador assume o papel de um personagem, enquanto um jogador é o árbitro – o Guardião do Conhecimento Arcano (“Guardião”) que modera o jogo e apresenta o enredo e o cenário para os outros jogadores. Usando dados e as regras do jogo, você determina o sucesso e o fracasso das ações dos personagens, enquanto eles são lançados em situações dramáticas e perigosas.
</p>


<svelte:window on:keyup={handleShortcut} />
<div class="extra-outer-wrapper">
	<div class="outer-wrapper">
		<div class="inner-wrapper">
			{#each slides as slide, id}
				{#if id === cur}
					<div
		 class="slide"
   in:hslide={transition_args}
   out:hslide={transition_args}
   >
   <!--- Esse estilo faz a imagem ficar contida na div --->
   <img style='height: 100%; width: 100%; object-fit: contain'  src="{slide.bg}" alt="carousel img">		
					</div>
				{/if}
			{/each}
		</div>
	</div>
</div>

<div class="dots">
	{#each slides as slide, i}
		<button on:click={()=>changeSlide(i)} class="{isDark === false ? 'dot' : 'dot-dark'}" class:selected={cur == i}>{i+1}</button>
	{/each}
</div>

<h2> Papel dos Investigadores </h2>
<hr>
<p>
	Em cada jogo, os jogadores assumem um dos dois papéis: um investigador ou o Guardião. A maioria assume o papel de investigadores (já que é exatamente isso que eles fazem) tentando resolver um mistério ou resolver alguma situação terrível. As tramas encontradas são projetadas para desafiar esses investigadores, que podem se machucar, sofrer experiências de quebrar a sanidade ou até ser comidos por um monstro! À medida que o jogo avança, os investigadores podem aprender sobre magias estranhas e monstros alienígenas horrendos, obter conhecimento especial de livros arcanos de conhecimento esquecido e avançar em suas habilidades à medida que se tornam mais experientes e adeptos.
</p>

<p>
	Um jogador assume o papel de Guardião. Eles escolhem um cenário para jogar ou podem criar um de sua própria criação. No jogo, o Guardião prepara o cenário, descreve as cenas e retrata as pessoas que os investigadores encontram (chamados de “Personagens Não-Jogadores” ou NPCs). O Guardião também ajuda a resolver a ação e arbitra as regras do jogo. Como o Guardião deve fazer uma preparação extra, os jogadores geralmente alternam o dever do Guardião entre diferentes cenários. Pense no papel do Guardião como o de um diretor fazendo um filme em que os atores (os investigadores) não sabem como a história se desenvolverá.
</p>

<p>
	O jogo é uma interação em evolução entre os jogadores – sob o disfarce de seus personagens desvendando um mistério – e o Guardião, que apresenta e julga o mundo em que o mistério ocorre.
</p>

<p>
	Não há tabuleiro para jogar. O jogo é principalmente falar: uma situação é apresentada e delineada pelo Guardião, e então os jogadores dizem o que eles, como seus investigadores, pretendem fazer. Usando as regras para manter as coisas consistentes e justas, o Guardião diz a eles se eles podem fazer o que propuseram e os passos que devem seguir, o que geralmente significa rolar alguns dados para determinar o sucesso. Os dados ajudam a resolver encontros e situações e mantêm todos honestos, além de adicionar drama e suspense - o resultado de uma jogada pode significar uma surpresa imprevista, uma derrota sombria ou uma fuga da morte por um fio de cabelo! Uma vez que um resultado é determinado, o Guardião narra o que acontece, pedindo aos jogadores suas reações e assim por diante.
</p>

<p>
	O objetivo do roleplaying é se divertir. Até o coração batendo e as sobrancelhas suadas, faz parte da natureza humana encontrar prazer em sentir medo, desde que esse medo não seja real. Para alguns, o relaxamento após o susto é o resultado mais importante. Para outros, é o próprio susto. Call of Cthulhu é um veículo para assustar e tranquilizar os jogadores alternadamente. Diversão agradável para todos os interessados!
</p>

<h2>Tutorial</h2>
<hr>
{#if isMounted === true}
	<Player this={Player} bind:this={video} controls>
		<Youtube this={Youtube} videoId="wouSEjZHj9U" />
	</Player>
{/if}

<style>
	button {
		background: transparent;
		color: #FFF;
		border-color: transparent;
		width: 3.2rem;
		height: 3.2rem;
	}

	button:hover,
	button:focus{
		background: rgba(0,0,0,0.5);
	}

	.dots {
		display: flex;
		align-items: center;
		justify-content: center;
		margin-top: 8px;
	}

	.dot {
		width: 1rem;
		height: 1rem;
		background: #928374;
		border-radius: 100%;
		font-size: 0;
		margin: 0.3rem;
		opacity: 0.3;
	}

	.dot-dark {
		width: 1rem;
		height: 1rem;
		background: #bd93f9;
		border-radius: 100%;
		font-size: 0;
		margin: 0.3rem;
		opacity: 0.3;
	}

	.dot.selected {
		opacity: 1;
	}

	.dot-dark.selected {
		opacity: 1;
	}
	.extra-outer-wrapper {
		width: 80%;
		margin: 0 auto;
	}

	.outer-wrapper {
		width: 100%;
		padding: 0 0 56.25%;
		position: relative;
	}

	.inner-wrapper {
		height: 100%;
		width: 100%;
		display: flex;
		position: absolute;
	}

	.slide {
		flex: 1 0 auto;
		width: 100%;
		height: 100%;
		align-items: center;
		justify-content: center;
		display: flex;
		text-align: center;
		font-weight: bold;
		font-size: 2rem;
		color: white;
	}
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

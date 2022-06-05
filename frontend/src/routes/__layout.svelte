<svelte:head>
	<title>Cthullu Online - Home</title>
</svelte:head>

<script>
	import "../app.scss";
	import Header from '../components/Header.svelte'
	import Footer from '../components/Footer.svelte'
</script>

<Header/>
<div class="wrapper d-flex flex-column min-vh-100 px-5 mt-3">
	<slot></slot>
</div>
<Footer/>


<style>
	/* :global(.dark) { */
	/* 	background-color: #282828; */	
	/* } */
	/* :global(body) { */
	/* 	background-color: #FFFFFF; */
	/* } */
</style>

<script>
	import Header from '../components/Header.svelte'
	import Footer from '../components/Footer.svelte'
</script>

<Header/>
<div class="wrapper d-flex flex-column min-vh-100 px-5 mt-3">
	<slot></slot>
</div>
<Footer/>


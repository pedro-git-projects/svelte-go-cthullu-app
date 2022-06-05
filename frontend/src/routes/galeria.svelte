
<svelte:head>
	<title>Cthullhu Online - Galeria</title>
</svelte:head>

<script>
	import Gallery from 'svelte-image-gallery'

	function HandleClick(e) {
		console.log(
			`src: ${e.detail.src}, alt: ${e.detail.alt}, loading: ${e.detail.loading}, class: ${e.detail.class}`
		);
	}

</script>

<h1 class="text-center mb-3">Galeria</h1>
<Gallery
	gap={15}
	maxColumnWidth={250}
	on:click={HandleClick}
	hover={true}
	loading="lazy"
 >
 <img   src="https://cdna.artstation.com/p/assets/images/images/015/745/616/large/bram-sels-cthulhu-rises.jpg?1549474030" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/018/163/051/large/daniel-zrom-danielzrom-dust1947-cthulhucustists1.jpg?1558520931" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/010/721/737/large/daniel-zrom-danielzrom-cthulhuspawnartstation.jpg?1525882711" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/042/617/303/large/eren-arik-cthulhu3.jpg?1634998017" alt="">
 <img src="https://cdna.artstation.com/p/assets/images/images/016/504/068/large/andy-walsh-cthulhu3-2000px-andywalshart.jpg?1552409478" alt="">
 <img src="https://www.artstation.com/artwork/xBG6E" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/008/636/751/large/andy-walsh-offering-to-cthulhu-by-andy-walsh-1700px-sig.jpg?1514138060" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/012/729/229/large/richard-wright-cthulhu-the-office.jpg?1536219278" alt="">
 <img src="https://cdna.artstation.com/p/assets/images/images/012/708/650/large/richard-wright-cthulhu-catacombs.jpg?1536123412" alt="">
 <img src="https://cdna.artstation.com/p/assets/images/images/036/974/692/large/andy-walsh-house-on-the-rocks-1500px-by-andy-walsh.jpg?1619119935" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/012/675/039/large/richard-wright-cthulhu-ghoul.jpg?1535978919" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/042/321/845/large/pascal-blanche-hastur-low.jpg?1634179026" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/042/321/855/large/pascal-blanche-dagon-low.jpg?1634179056" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/003/927/989/large/juan-pablo-corredor-cthulhu.jpg?1483648462" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/000/421/203/large/tom-isaksen-cthulhu-by-tomisaksen-d6xvdt5.jpg?1421830343" alt="">
 <img src="https://cdna.artstation.com/p/assets/images/images/001/487/802/large/_-dofresh-_-hpl-was-right-03-def.jpg?1447276437" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/035/338/199/large/vladimir-somov-highresscreenshot00006.jpg?1614709680" alt="">
 <img src="https://cdna.artstation.com/p/assets/images/images/022/182/574/large/tj-foo-squidboat-small.jpg?1574415228" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/022/208/123/large/tj-foo-squidattack-small.jpg?1574523984" alt="">
 <img src="https://cdna.artstation.com/p/assets/images/images/022/182/572/large/tj-foo-beachedsquid-small.jpg?1574415222" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/010/912/727/large/aurore-folny-olivemcbride-aurorefolny.jpg?1526913891" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/041/147/903/large/maxlm-salty-rain-done2.jpg?1630908207" alt="">
 <img src="https://cdna.artstation.com/p/assets/images/images/000/237/344/large/james-daly-oldone-unnamedguardian.jpg?1412605641" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/008/513/371/large/guillem-h-pongiluppi-61-x-86-27-cm-cthulhumaster.jpg?1513244580" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/007/343/597/large/guillem-h-pongiluppi-sketch-7-guille.jpg?1505468960" alt="">
 <img src="https://cdnb.artstation.com/p/assets/images/images/007/641/497/large/jake-murray-jakemurray-masksofnyarlathotep-mark.jpg?1507559038" alt="">
 <img src="https://cdna.artstation.com/p/assets/images/images/025/612/188/large/roman-gaydakov-midnight-room-eevee.jpg?1586416915" alt="">
</Gallery>

<div id="myModal" class="modal">

	<!-- The Close Button -->
	<span class="close">&times;</span>

	<!-- Modal Content (The Image) -->
	<img class="modal-content" alt="" id="img01">

	<!-- Modal Caption (Image Text) -->
	<div id="caption"></div>
</div>

<style>
	img { opacity: .9; transition: all .2s }
	img:hover { opacity: 1; transform: scale(1.04) }
</style>

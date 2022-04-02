<script lang="ts">
	import { onMount } from 'svelte';

	let mapContainer: HTMLElement;

	onMount(() => {
		const options = {
			center: new window.kakao.maps.LatLng(33.450701, 126.570667),
			level: 4
		};
		new window.kakao.maps.Map(mapContainer, options);
	});
</script>

<svelte:head>
	<script
		type="text/javascript"
		src="//dapi.kakao.com/v2/maps/sdk.js?appkey={import.meta.env.VITE_KAKAO_MAP_KEY}"></script>
</svelte:head>

<div class="map-container" bind:this={mapContainer} />

<style>
	.map-container {
		box-sizing: border-box;
		width: 100vw;
		height: 100vh;
	}
</style>

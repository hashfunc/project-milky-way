<script lang="ts">
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { store as mapStore } from '$lib/store/map';

	let mapContainer: HTMLElement;

	onMount(() => {
		const { latitude, longitude } = get(mapStore);

		const position = new window.kakao.maps.LatLng(latitude, longitude);

		const options = { center: position, level: 4 };
		const map = new window.kakao.maps.Map(mapContainer, options);

		const marker = new window.kakao.maps.Marker({ map, position });

		mapStore.subscribe(({ latitude, longitude }) => {
			const position = new window.kakao.maps.LatLng(latitude, longitude);
			map.setCenter(position);
			marker.setPosition(position);
		});
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

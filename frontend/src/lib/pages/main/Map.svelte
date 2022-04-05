<script lang="ts">
	import { onMount } from 'svelte';
	import { actions as mapActions } from '$lib/store/map';

	let ref: HTMLElement;

	onMount(() => {
		mapActions.initialize(ref);
	});
</script>

<svelte:head>
	<script
		src="//dapi.kakao.com/v2/maps/sdk.js?appkey={import.meta.env.VITE_KAKAO_MAP_KEY}"></script>
</svelte:head>

<div class="map-container" bind:this={ref} />

<style>
	.map-container {
		box-sizing: border-box;
		width: 100vw;
		height: 100vh;
	}
</style>

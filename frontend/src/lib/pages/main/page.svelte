<script lang="ts">
	import Header from '$lib/components/Header/Header.svelte';
	import Map from './Map.svelte';
</script>

<main>
	<Header />
	<Map />
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		gap: 20px;
		box-sizing: border-box;
		width: 100vw;
		height: 100vh;
	}
</style>

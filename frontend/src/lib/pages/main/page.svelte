<script lang="ts">
	import MapButton from './MapButton.svelte';
	import ProjectTitle from './ProjectTitle.svelte';
	import SearchForm from './SearchForm.svelte';
</script>

<main>
	<ProjectTitle />
	<SearchForm placeholder="장소를 검색해주세요." />
	<MapButton />
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
